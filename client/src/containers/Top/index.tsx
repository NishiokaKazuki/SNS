import React from 'react'
import styled from 'styled-components'

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

const useStyles = makeStyles((theme: Theme) => createStyles({
  toolbar: {
    ...theme.mixins.toolbar
  },
}))

const TopPage: React.FC = () => {
  const classes = useStyles()

  return (
    <Root>
      <div className={classes.toolbar} />
      <H1>~ トップページ ~</H1>
      <H1>　チャットアプリ　</H1>
    </Root>
  )
}

const Root = styled.div`
  padding-bottom: 50px;
`

const H1 = styled.h1`
  text-align: center;
`

export default TopPage