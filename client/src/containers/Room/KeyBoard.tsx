import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

import { makeStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Grid from '@material-ui/core/Grid';
import Toolbar from '@material-ui/core/Toolbar'
import TextField from '@material-ui/core/TextField';
import PhotoIcon from '@material-ui/icons/Photo';
import SendIcon from '@material-ui/icons/Send';

const KeyBoard: React.FC = () => {
  const classes = useStyles();
  return(
    <>
        <AppBar position="fixed" color="secondary" className={classes.keyBoard}>
            <Toolbar>
                <StyledLink to="/">
                    <PhotoIcon color="primary"/>
                </StyledLink>
                <TextField
                    id="outlined-multiline-flexible"
                    className={classes.textField}
                    size="small"
                    fullWidth
                    multiline
                    rowsMax={2}
                    variant="outlined"
                />
                <SendIcon color="primary"/>
            </Toolbar>
        </AppBar>
    </>
  )
}

const useStyles = makeStyles({
    keyBoard:{
        top: 'auto',
        background: '#FFFFFF',
        bottom: 0,
        paddingBottom: '60px'
    },
    textField:{
        paddingLeft:'10px',
        paddingRight:'10px',
    }
});

const StyledLink = styled(Link)`
    margin: auto;
    text-decoration: none;
    color: white;
`

export default KeyBoard