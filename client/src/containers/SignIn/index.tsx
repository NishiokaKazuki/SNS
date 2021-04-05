import React from 'react'
import styled from 'styled-components'
import { useDispatch, useSelector } from 'react-redux'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

import { signIn, user } from '../../actions'

const SignIn: React.FC = () => {
  const [handle, setHandle] = React.useState("")
  const [pw, setPw] = React.useState("")
  const auth = useSelector((state: any) => state.auth);
  const dispatch = useDispatch()
  const classes = useStyles()

  const onChangeHandle = (e: any) => {
    setHandle(e.target.value)
  }
  const onChangePassword = (e: any) => {
    setPw(e.target.value)
  }
  const handleSignin = () => dispatch(
    // setAuth({isAuthenticated: true, token:'aaa'})
    signIn({handle, pw})
  )
  const handleUser = () => dispatch(
    // setAuth({isAuthenticated: true, token:'aaa'})
    user({token:auth.list[0].token})
  )

  return (
    <Root>
      <form className={classes.root} noValidate autoComplete="off">
        <div>
          <TextField
            id="outlined-password-input"
            label="User"
            variant="outlined"
            onChange={onChangeHandle}
          />
          <TextField
            id="outlined-password-input"
            label="Password"
            type="password"
            autoComplete="current-password"
            variant="outlined"
            onChange={onChangePassword}
          />
        </div>
      </form>
      <Button variant="contained" color="secondary" onClick={handleSignin}>
        サインイン
      </Button>
      <Button variant="contained" color="secondary" onClick={handleUser}>
        ユーザ
      </Button>
    </Root>
  )
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      '& .MuiTextField-root': {
        margin: theme.spacing(1),
        width: '25ch',
      },
    },
  }),
);


const Root = styled.div`
  margin: auto;
  padding-top: 80px;
  padding-bottom: 50px;
`

export default SignIn