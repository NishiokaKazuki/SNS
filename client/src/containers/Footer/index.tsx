import React from 'react'
import DrawerItems from './DrawerItems'

import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import { makeStyles } from '@material-ui/core/styles'

const Footer: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div>
                <AppBar position="fixed" color="secondary" className={classes.footer}>
                    <Toolbar>
                        <DrawerItems/>
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