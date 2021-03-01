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

// Sliceを生成する
const slice = createSlice({
    name: 'User',
    initialState,
    reducers: {
        setUser: (state, action)=>{
            const { id, handle, name, birthday, profile, isPrivate} = action.payload
            return Object.assign({}, state, { id, handle, name, birthday, profile, isPrivate })
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