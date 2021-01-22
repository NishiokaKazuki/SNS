import React from 'react'
import { Link } from 'react-router-dom'
import styled from 'styled-components'

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'
import Button from '@material-ui/core/Button';

const DrawerItems: React.FC = () => {
    const classes = useStyles();

    return (
      <>
        <StyledLink to="/">
            <Button className={classes.button} size="large">
                MyPage
            </Button>
        </StyledLink>
        <StyledLink to="/talk">
            <Button className={classes.button} size="large">
                Talk
            </Button>
        </StyledLink>
        <StyledLink to="/">
            <Button className={classes.button} size="large">
                TimeLine
            </Button>
        </StyledLink>
        <StyledLink to="/">
            <Button className={classes.button} size="large">
                Info
            </Button>
        </StyledLink>
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

export default DrawerItems