import { gql } from '../../../../../shared/src/graphql/graphql'
import { ActorQuery } from '../../../actor/graphql'

export const ThreadFragment = gql`
    fragment ThreadFragment on Thread {
        __typename
        id
        number
        title
        state
        kind
        url
        createdAt
        externalURLs {
            url
            serviceType
        }
        repository {
            name
            url
        }
        author {
            ${ActorQuery}
        }
        comments {
            totalCount
        }
    }
`