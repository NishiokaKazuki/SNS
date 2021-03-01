import { createSlice } from '@reduxjs/toolkit';
import serviceClient from '../serviceClient'
// import { iSignin } from '../serviceClient'

// Stateの初期状態
const initialState = {
    isAuthenticated: false,
    token: '',
};

// Sliceを生成する
const slice = createSlice({
    name: 'Auth',
    initialState,
    reducers: {
        signIn: (state, action)=>{
            const {handle, pw} = action.payload
            const res: any = serviceClient.signinRequest({handle:handle, pw:pw})
            console.log('aaa')
            // setToken(res.getToken())
            return Object.assign({}, state, { isAuthenticated: action.payload })
        },
        setIsAuthenticated: (state, action)=>{
            return Object.assign({}, state, { isAuthenticated: action.payload })
        },
        setToken: (state, action)=>{
            console.log(action.payload)
            return Object.assign({}, state, { token: action.payload })
        },
        clearToken: state=>{
            return Object.assign({}, state, { token: '' })
        },
    }
});

export default slice.reducer;
export const { signIn, setIsAuthenticated, setToken, clearToken } = slice.actions;