import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import GitHubIcon from '@material-ui/icons/GitHub'
import { makeStyles } from '@material-ui/core/styles'

const Footer: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div>
                <AppBar position="fixed" color="secondary" className={classes.footer}>
                    <Toolbar>
                        {/* <StyledLink to="/"> */}
                            <GitHubIcon />フッター
                        {/* </StyledLink> */}
                    </Toolbar>
                </AppBar>
            </div>
        </>
    )
}

const useStyles = makeStyles(() => ({
    footer: {
        top: 'auto',
        bottom: 0,
    }
}))


export default Footer