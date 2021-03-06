import { Dropdown, DropdownItem, DropdownMenu, DropdownToggle } from 'reactstrap'
import KeyboardIcon from 'mdi-react/KeyboardIcon'
import ViewQuiltIcon from 'mdi-react/ViewQuiltIcon'
import React, { useState, useCallback } from 'react'
import { Link } from 'react-router-dom'

interface Props {
    interactiveSearchMode: boolean
    toggleSearchMode: (event: React.MouseEvent<HTMLAnchorElement>) => void
}

export const SearchModeToggle: React.FunctionComponent<Props> = props => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const toggleIsOpen = useCallback(() => setIsOpen(!isOpen), [isOpen, setIsOpen])

    return (
        <Dropdown isOpen={isOpen} toggle={toggleIsOpen} className="search-mode-toggle">
            <DropdownToggle
                caret={true}
                className="search-mode-toggle__button e2e-search-mode-toggle"
                aria-label="Toggle search mode"
            >
                {props.interactiveSearchMode ? (
                    <ViewQuiltIcon className="icon-inline" size={8}></ViewQuiltIcon>
                ) : (
                    <KeyboardIcon className="icon-inline" size={8}></KeyboardIcon>
                )}
            </DropdownToggle>
            <DropdownMenu>
                <DropdownItem
                    active={props.interactiveSearchMode}
                    onClick={!props.interactiveSearchMode ? props.toggleSearchMode : undefined}
                    className="e2e-search-mode-toggle__interactive-mode"
                >
                    <ViewQuiltIcon className="icon-inline" size={8}></ViewQuiltIcon>
                    <span className="ml-1">Interactive mode</span>
                </DropdownItem>
                <DropdownItem
                    active={!props.interactiveSearchMode}
                    onClick={props.interactiveSearchMode ? props.toggleSearchMode : undefined}
                    className="e2e-search-mode-text-mode"
                >
                    <KeyboardIcon className="icon-inline"></KeyboardIcon>
                    <span className="ml-1">Plain text mode</span>
                </DropdownItem>
                <DropdownItem tag={Link} to="/search/query-builder">
                    Query builder&hellip;
                </DropdownItem>
            </DropdownMenu>
        </Dropdown>
    )
}
