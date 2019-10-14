/* eslint-disable @typescript-eslint/no-non-null-assertion */
import { FormattingOptions, Segment } from '@sqs/jsonc-parser'
import { setProperty } from '@sqs/jsonc-parser/lib/edit'
import path from 'path'
import * as sourcegraph from 'sourcegraph'
import { parseRepoURI } from '../../../../../shared/src/util/url'
import { ExecServerClient } from '../execServer/client'
import { ResolvedDependencyInPackage } from './packageManager'
import { PackageJsonDependencyCampaignContext } from './packageJsonDependency'
import { Observable, combineLatest, from } from 'rxjs'
import { switchMap, map } from 'rxjs/operators'

const MINIMAL_WORKTREE = true

export const editForCommands2 = (
    files: URL[],
    commands: string[][],
    execServerClient: ExecServerClient
): Observable<sourcegraph.WorkspaceEdit> =>
    combineLatest(files.map(url => from(sourcegraph.workspace.openTextDocument(url)))).pipe(
        switchMap(files => {
            const dir = path.dirname(parseRepoURI(files[0].uri).filePath!)

            const filesToText: { [path: string]: string } = {}
            for (const file of files) {
                const name = path.basename(parseRepoURI(file.uri).filePath!)
                filesToText[name] = file.text!
            }
            return from(
                execServerClient({
                    commands,
                    dir,
                    files: filesToText,
                    label: `editForCommands(${JSON.stringify({ files: files.map(f => f.uri), commands })})`,
                })
            ).pipe(
                map(result => {
                    const edit = new sourcegraph.WorkspaceEdit()
                    for (const file of files) {
                        const name = path.basename(parseRepoURI(file.uri).filePath!)
                        edit.set(new URL(file.uri), [sourcegraph.TextEdit.patch(result.fileDiffs![name])])
                    }
                    return edit
                })
            )
        })
    )

export const editForCommands = async (
    {
        packageJson,
        lockfile,
    }: {
        packageJson: Pick<sourcegraph.TextDocument, 'uri' | 'text'>
        lockfile: Pick<sourcegraph.TextDocument, 'uri' | 'text'>
    },
    commands: string[][],
    execServerClient: ExecServerClient
): Promise<sourcegraph.WorkspaceEdit> => {
    throw new Error('only MINIMAL_WORKTREE is supported')
}

export const editForDependencyAction = (
    dep: Pick<ResolvedDependencyInPackage, 'dependency'> & {
        packageJson: Pick<sourcegraph.TextDocument, 'uri' | 'text'>
        lockfile: Pick<sourcegraph.TextDocument, 'uri' | 'text'>
    },
    action: PackageJsonDependencyCampaignContext['action'],
    commands: Record<'upgradeCommands' | 'removeCommands', string[][]>,
    execServerClient: ExecServerClient
): Promise<sourcegraph.WorkspaceEdit> =>
    editForCommands(dep, action === 'ban' ? commands.removeCommands : commands.upgradeCommands, execServerClient)

// function computeDiffs(files: { old: sourcegraph.TextDocument; newText?: string }[]): sourcegraph.WorkspaceEdit {
//     const edit = new sourcegraph.WorkspaceEdit()
//     for (const { old, newText } of files) {
//         // TODO!(sqs): handle creation/removal
//         if (old.text !== undefined && newText !== undefined && old.text !== newText) {
//             edit.replace(
//                 new URL(old.uri),
//                 new sourcegraph.Range(new sourcegraph.Position(0, 0), old.positionAt(old.text.length)),
//                 newText
//             )
//         }
//     }
//     return edit
// }

const PACKAGE_JSON_FORMATTING_OPTIONS: FormattingOptions = {
    eol: '\n',
    insertSpaces: true,
    tabSize: 2,
}

export const editPackageJson = (
    doc: sourcegraph.TextDocument,
    operations: { path: Segment[]; value: any }[],
    workspaceEdit = new sourcegraph.WorkspaceEdit()
): sourcegraph.WorkspaceEdit => {
    for (const op of operations) {
        const edits = setProperty(doc.text!, op.path, op.value, PACKAGE_JSON_FORMATTING_OPTIONS)
        for (const edit of edits) {
            workspaceEdit.replace(
                new URL(doc.uri),
                new sourcegraph.Range(doc.positionAt(edit.offset), doc.positionAt(edit.offset + edit.length)),
                edit.content
            )
        }
    }
    return workspaceEdit
}