import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

const SignIn: React.FC = () => {
  const classes = useStyles();

  return (
    <Root>
      <form className={classes.root} noValidate autoComplete="off">
        <div>
          <TextField
            id="outlined-password-input"
            label="User"
            variant="outlined"
          />
          <TextField
            id="outlined-password-input"
            label="Password"
            type="password"
            autoComplete="current-password"
            variant="outlined"
          />
        </div>
      </form>
      <Button variant="contained" color="secondary">
        サインイン
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

const StyledLink = styled(Link)`
    margin: auto;
    text-decoration: none;
    color: white;
`

export default SignIn