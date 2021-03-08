import { createSlice } from '@reduxjs/toolkit';
// import { iSignin } from '../serviceClient'

interface Auth {
    isAuthenticated: boolean;
    token: string;
}

interface AuthState {
    nextAuthId: number;
    list: Auth[];
};

// Stateの初期状態
const initialAuthState:AuthState = {
    nextAuthId: 0,
    list: [],
};

// Sliceを生成する
const slice = createSlice({
    name: 'Auth',
    initialState:initialAuthState,
    reducers: {
        setAuth: (state, action)=>{
            state.list.push(action.payload)
        },
        setToken:(state, action)=>{
            return Object.assign({}, state, { token: action.payload })
        },
        setIsAuthenticated: (state, action)=>{
            return Object.assign({}, state, { isAuthenticated: action.payload })
        },
        clearToken: state=>{
            return Object.assign({}, state, { token: '' })
        },
    }
});

export default slice.reducer;
export const { setAuth, setIsAuthenticated, setToken, clearToken } = slice.actions;