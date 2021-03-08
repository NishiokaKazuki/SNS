import { Dispatch } from "react"
import { Action } from "redux"

import { setAuth } from '../store/AuthReducer'
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
            const { handle, pw } = arg
            const res:any  = await serviceClient.signinRequest({handle:handle, pw:pw})
            const auth = {
                isAuthenticated: res.getStatus(),
                token:res.getToken(),
            }
            dispatch(setAuth(auth))
        } catch (e) {
            const auth = {
                isAuthenticated: false,
                token:'',
            }
            dispatch(setAuth(auth))
        }
    }
}