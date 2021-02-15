import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

import { makeStyles } from '@material-ui/core/styles'
import AppBar from '@material-ui/core/AppBar'
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar'
import GitHubIcon from '@material-ui/icons/GitHub'

const Header: React.FC = () => {
    const classes = useStyles()
    return (
        <>
            <div>
                <AppBar position="fixed" color="secondary">
                    <Toolbar>
                        <StyledLink to="/">
                            <GitHubIcon />チャット
                        </StyledLink>
                        <ButtonLink to="/signin">
                            <Button className={classes.button}>
                                SignIn
                            </Button>
                        </ButtonLink>
                    </Toolbar>
                </AppBar>
            </div>
        </>
    )
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