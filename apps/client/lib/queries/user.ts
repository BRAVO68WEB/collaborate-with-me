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
