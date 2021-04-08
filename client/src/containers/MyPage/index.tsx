import React from 'react'
import styled from 'styled-components'
import { useDispatch, useSelector } from 'react-redux'
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles'

const MyPage: React.FC = () => {
  const user = useSelector((state: any) => state.user);
  var name = "";
  var handle = "";
  var birthday = "";
  var profile = "";
  var isPrivate = false;

  if (user.list.length>0) {
    name = user.list[user.list.length-1].name
    handle = user.list[user.list.length-1].handle
    birthday = user.list[user.list.length-1].birthday
    profile = user.list[user.list.length-1].profile
    isPrivate = user.list[user.list.length-1].isPrivate
  }

  return (
    <Root>
      <div>名前: {name}</div>
      <div>ID: {handle}</div>
      <div>誕生日: {birthday}</div>
      <div>プロフィール: {profile}</div>
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

export default MyPage