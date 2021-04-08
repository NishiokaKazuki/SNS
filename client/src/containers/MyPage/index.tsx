import React from 'react'
import styled from 'styled-components'
import { useDispatch, useSelector } from 'react-redux'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

import { signIn } from '../../actions'

const MyPage: React.FC = () => {
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

  return (
    <Root>

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

export default MyPage