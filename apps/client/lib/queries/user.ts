import { gql } from '@apollo/client'
import { graphql } from 'gql.tada'

export const LOGIN_USER = graphql(`
    mutation Login($email: String!, $password: String!) {
        login(email: $email, password: $password) {
            is_success
            access_token
        }
    }
`)

export const GET_ME = graphql(`
    query GetUser {
        me {
            id
            username
            email
            role
            is_active
            created_at
            updated_at
        }
    }
`)
