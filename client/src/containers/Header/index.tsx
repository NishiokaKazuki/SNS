import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'
import { useDispatch, useSelector } from 'react-redux'

import { makeStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar'
import GitHubIcon from '@material-ui/icons/GitHub'

import { signOut } from '../../actions'


const Header: React.FC = () => {
    const classes = useStyles()
    const auth = useSelector((state: any) => state.auth);
    var isAuth = false
    if (auth.list.length>0) {
        isAuth = auth.list[auth.list.length-1].isAuthenticated
    }


    return (
        <>
            <div>
                <AppBar position="fixed" color="secondary">
                    <Toolbar>
                        <StyledLink to="/">
                            <GitHubIcon />チャット
                        </StyledLink>
                        <ButtonSwitch isAuth={isAuth}/>
                        {/* <ButtonLink to="/signin">
                            <Button className={classes.button}>
                                SignIn
                            </Button>
                        </ButtonLink> */}
                    </Toolbar>
                </AppBar>
            </div>
        </>
    )
}

interface Props {
    isAuth: boolean,
}

const ButtonSwitch: React.FC<Props> = (props) => {
    const classes = useStyles()
    const dispatch = useDispatch()
    const handleSignOut = () => dispatch(
        signOut()
    )
    if (props.isAuth){
        return (
            <ButtonLink to="/mypage">
                <Button className={classes.button} onClick={handleSignOut}>
                    SignOut
                </Button>
            </ButtonLink>
        );
    }
    return (
    <ButtonLink to="/signin">
        <Button className={classes.button}>
            SignIn
        </Button>
    </ButtonLink>
    );
}

const useStyles = makeStyles({
    button: {
      color: 'white',
    },
});


const StyledLink = styled(Link)`
    margin: auto;
    text-decoration: none;
    color: white;
`

const ButtonLink = styled(Link)`
    text-decoration: none;
`

export default Header