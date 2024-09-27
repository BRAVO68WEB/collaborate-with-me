import { ApolloClient, InMemoryCache } from '@apollo/client'

const createApolloClient = () => {
    const token = localStorage.getItem('token')

    return new ApolloClient({
        uri: process.env.NEXT_PUBLIC_GRAPHQL_URL,
        cache: new InMemoryCache(),
        headers: {
            authorization: token ? `Bearer ${token}` : '',
        },
    })
}

export default createApolloClient

export const apolloClient = createApolloClient()
