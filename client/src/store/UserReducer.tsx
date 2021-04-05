import { createSlice } from '@reduxjs/toolkit';
// Stateの初期状態
const initialState = {
    id: 0,
    handle: '',
    name: '',
    birthday: '',
    profile: '',
    isPrivate: true,
};

interface User {
    id: number
    handle: string
    name: string
    birthday: string
    profile: string
    isPrivate: boolean
}

interface UserState {
    nextUserId: number;
    list: User[];
};

// Stateの初期状態
const initialUserState:UserState = {
    nextUserId: 0,
    list: [],
};


// Sliceを生成する
const slice = createSlice({
    name: 'User',
    initialState: initialUserState,
    reducers: {
        setUser: (state, action)=>{
            state.list.push(action.payload)
        },
        clearUser: state=>{
            return Object.assign({}, state, {
                id: 0,
                handle: '',
                name: '',
                birthday: '',
                profile: '',
                isPrivate: true,
            })
        },
    }
});

export default slice.reducer;
export const { setUser, clearUser } = slice.actions;