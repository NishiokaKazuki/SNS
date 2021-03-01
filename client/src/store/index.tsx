import { combineReducers } from "redux";
import { configureStore } from "@reduxjs/toolkit";

// それぞれ slice.reducer を default export している前提
import authReducer from "./AuthReducer";
import userReducer from "./UserReducer";

const reducer = combineReducers({
  auth: authReducer,
  user: userReducer
});

const store = configureStore({ reducer });

export default store;