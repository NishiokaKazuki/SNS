import React from 'react'
import Bubble from './Bubble'
import KeyBoard from './KeyBoard'
import styled from 'styled-components'

import { makeStyles } from '@material-ui/core/styles'
import Container from '@material-ui/core/Container';

const Room: React.FC = () => {
  const classes = useStyles();
  return(
    <Root>
        <Container className = {classes.color}>
            <Bubble isOther={true} message={"あああああああああああああああああああああああああああああああ"}/>
            <Bubble isOther={false} message={"あああああああああああああああああああああ"}/>
            <Bubble isOther={true} message={"aaaaa"}/>
            <Bubble isOther={false} message={"bbbbbb"}/>
            <Bubble isOther={true} message={"aaaaa"}/>
            <Bubble isOther={false} message={"bbbbbb"}/>
            <Bubble isOther={true} message={"aaaaa"}/>
            <Bubble isOther={true} message={"aaaaa"}/>
            <Bubble isOther={true} message={"aaaaa"}/>
            <Bubble isOther={false} message={"bbbbbb"}/>
            <Bubble isOther={false} message={"bbbbbb"}/>
            <Bubble isOther={false} message={"bbbbbb"}/>
        </Container>
        <KeyBoard/>
    </Root>
  )
}

const useStyles = makeStyles({
    color: {
      background:'linear-gradient(45deg, #2196F3 30%, #21CBF3 180%)',
      color:'white',
      height:'100%',
      paddingBottom:'100px',
    },
});

const Root = styled.div`
  margin: auto;
  padding-top: 60px;
  padding-bottom: 50px;
`

export default Room