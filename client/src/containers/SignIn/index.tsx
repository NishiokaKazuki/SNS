import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import GitHubIcon from '@material-ui/icons/GitHub'

const SignIn: React.FC = () => {
    return (
        <>
            <div>
                <AppBar position="fixed">
                    <Toolbar>
                        <StyledLink to="/">
                            <GitHubIcon />チャット
                        </StyledLink>
                    </Toolbar>
                </AppBar>
            </div>
        </>
    )
}

const StyledLink = styled(Link)`
    margin: auto;
    text-decoration: none;
    color: white;
`

export default SignIn