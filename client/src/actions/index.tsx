import { Dispatch } from "react"
import { Action } from "redux"

import { setAuth } from '../store/AuthReducer'
import { setUser } from '../store/UserReducer'
import serviceClient from '../serviceClient'

export interface Actions {
    type: String,
    payload: any
}


// export const signIn = createAsyncThunk(
//     'auth/signin',
//     async (arg: { handle: string, pw: string }) => {
//         const { handle, pw } = arg
//         const res:any  = await serviceClient.signinRequest({handle:handle, pw:pw})
//         const auth = {
//             isAuthenticated: res.getStatus(),
//             token:res.getToken(),
//         }
//         console.log(auth)
//         setAuth(auth)
//         return res.getStatus()
//     }
// )

export const signIn = (arg: { handle: string, pw: string }) => {
    return async (dispatch: Dispatch<Action>) => {
        try {
            console.log(arg)
            const { handle, pw } = arg
            const res:any  = await serviceClient.signinRequest({handle:handle, pw:pw})
            const auth = {
                isAuthenticated: res.getStatus(),
                token:res.getToken(),
            }
            dispatch(setAuth(auth))
            console.log(auth)
            const userRes:any  = await serviceClient.userRequest(res.getToken())
            const user = userRes.getUser()
            const iUser = {
                id: user.getId(),
                handle: user.getHandle(),
                name: user.getName(),
                birthday: user.getBirthday(),
                profile: user.getProfile(),
                isPrivate: user.getIsPrivate(),
            }
            dispatch(setUser(iUser))
            console.log(iUser)
        } catch (e) {
            const auth = {
                isAuthenticated: false,
                token:'',
            }
            dispatch(setAuth(auth))
            console.log(e)
        }
    }
}

export const user = (arg: { token: string }) => {
    return async (dispatch: Dispatch<Action>) => {
        try {
            const res:any  = await serviceClient.userRequest(arg.token)
            const user = {
                id: res.getId(),
                handle: res.getHandle(),
                name: res.getName(9),
                birthday: res.getBirthday(),
                profile: res.getProfile,
                isPrivate: res.getIsPrivate,
            }
            dispatch(setUser(user))
        } catch (e) {

        }
    }
}