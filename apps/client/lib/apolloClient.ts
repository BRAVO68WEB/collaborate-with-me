import { ApolloClient, InMemoryCache } from '@apollo/client'

const createApolloClient = () => {
    return new ApolloClient({
        uri: process.env.NEXT_PUBLIC_GRAPHQL_URL,
        cache: new InMemoryCache(),
        headers: {
            authorization: `Bearer ${localStorage.getItem('token') ?? ''}`,
        },
    })
}

export default createApolloClient

export const apolloClient = createApolloClient()
