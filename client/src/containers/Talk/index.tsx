import React from 'react'
import { Link, useHistory } from 'react-router-dom'
import styled from 'styled-components'
import Resizer from 'react-image-file-resizer'

import { makeStyles } from '@material-ui/core/styles'
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableRow from '@material-ui/core/TableRow';

const Talk: React.FC = () => {
  const classes = useStyles();
  const history = useHistory();
  const handleRowClick = (row:any) => {
    history.push('/talk/room/');//+row.link
  }
  return(
    <Root>
      <TableContainer className={classes.table}>
        <Table className={classes.table}>
          <colgroup>
            <col style={{width:'10%'}}/>
            <col style={{width:'20%'}}/>
            <col style={{width:'70%'}}/>
          </colgroup>
          <TableBody>
          {rows.map((row) => (
            <TableRow key={row.name} hover onClick={()=> handleRowClick(row)}>
              <TableCell component="th" scope="row">
                image
              </TableCell>
              <TableCell align="left">{row.name}</TableCell>
              <TableCell align="right">{row.message}</TableCell>
            </TableRow>
          ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Root>
  )
}

const useStyles = makeStyles({
  table: {
    width:"95%",
    marginLeft: "auto",
    marginRight: "auto",
  },
});

function createData(image: string, name: string, message: string, link: string) {
  return { image, name, message, link };
}

const rows = [
  createData('https://1.bp.blogspot.com/-_zO7P5V2zaI/X7zMIVkiQKI/AAAAAAABcYo/bZ2fpZQ_3i0okGLGRt_u6axbDsVtRIfaACNcBGAsYHQ/s1315/job_chocolatiere_woman.png','Tester_1', 'hi', '1'),
  createData('../../../public/job_chocolatiere_woman.png','Tester_2', 'bye', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_3', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_4', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_5', 'hi', '1'),
  createData('https://1.bp.blogspot.com/-_zO7P5V2zaI/X7zMIVkiQKI/AAAAAAABcYo/bZ2fpZQ_3i0okGLGRt_u6axbDsVtRIfaACNcBGAsYHQ/s1315/job_chocolatiere_woman.png','Tester_1', 'hi', '1'),
  createData('../../../public/job_chocolatiere_woman.png','Tester_2', 'bye', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_3', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_4', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_5', 'hi', '1'),
  createData('https://1.bp.blogspot.com/-_zO7P5V2zaI/X7zMIVkiQKI/AAAAAAABcYo/bZ2fpZQ_3i0okGLGRt_u6axbDsVtRIfaACNcBGAsYHQ/s1315/job_chocolatiere_woman.png','Tester_1', 'hi', '1'),
  createData('../../../public/job_chocolatiere_woman.png','Tester_2', 'bye', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_3', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_4', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_5', 'hi', '1'),
  createData('https://1.bp.blogspot.com/-_zO7P5V2zaI/X7zMIVkiQKI/AAAAAAABcYo/bZ2fpZQ_3i0okGLGRt_u6axbDsVtRIfaACNcBGAsYHQ/s1315/job_chocolatiere_woman.png','Tester_1', 'hi', '1'),
  createData('../../../public/job_chocolatiere_woman.png','Tester_2', 'bye', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_3', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_4', 'hi', '1'),
  createData('../../../public/pose_pien_uruuru_man.png','Tester_5', 'hi', '1'),
];

const StyledLink = styled(Link)`
  text-decoration: none;
`

const Root = styled.div`
  margin: auto;
  padding-top: 65px;
  padding-bottom: 50px;
`

export default Talk