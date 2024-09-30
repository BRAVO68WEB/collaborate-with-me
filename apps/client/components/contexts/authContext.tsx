'use client'

import { createContext, useContext, useEffect, useState } from 'react'
import Cookies from 'js-cookie'
import { useLazyQuery, useQuery } from '@apollo/client'
import { GET_ME } from '@/lib/queries/user'
import createApolloClient, { apolloClient } from '@/lib/apolloClient'
import { headers } from 'next/headers'
// import api from '@/lib/api'

interface AuthContextType {
    authenticated: boolean
    // checkAuth: () => Promise<void>
    user: IUser | null
    loading: boolean
}

export const AuthContext = createContext<AuthContextType>({
    authenticated: false,
    loading: true,
    user: null,
    // checkAuth: async () => {},
})

export const useAuth = (): AuthContextType => {
    const context = useContext(AuthContext)
    return context
}

interface AuthProviderProps {
    children: React.ReactNode
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
    const [authenticated, setAuthenticated] = useState(false)
    const [user, setUser] = useState<IUser | null>(null)
    const [getMe,{ data, loading, refetch }] = useLazyQuery(GET_ME, {
        client: apolloClient,
        // context:{
        //     headers:{
        //         // authorization: Cookies.get('token') ? `Bearer ${Cookies.get('token')}` : '',
        //         authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0ODMzMDksImlhdCI6MTcyNzQ2ODkwOSwidXNlcl9pZCI6IjY2ZWIzOGQ5YTc4NWZjOGUxMDA0N2E0NyJ9.bujgTzQRdAMlFLzzPvqupx4S6cyJTxqWGQAVk-G9JNQ"
        //     }
        // }
    })
    console.log(data?.me, "data")
    console.log(Cookies.get('token'))
    useEffect(() => {
        // console.log('refetching')
        // console.log(Cookies.get('token'),"token")   
        // refetch({
        //     headers:{
        //         authorization: Cookies.get('token') ? `Bearer ${Cookies.get('token')}` : '',
        //     }
        // })
        getMe({
            context:{
                headers:{
                    authorization: Cookies.get('token') ? `Bearer ${Cookies.get('token')}` : '',
                }
            }
        })
    },[])

    return (
        <AuthContext.Provider value={{ authenticated, user, loading:false }}>
            {children}
        </AuthContext.Provider>
    )
}
