import React from 'react';
import { Theme, createStyles, makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';

interface Props {
  isOther: boolean,
  message: string
}

const Bubble: React.FC<Props> = (props) => {
  const classes = useStyles(props);
  return (
    <div className={classes.root}>
      <Paper elevation={3}>{props.message}</Paper>
    </div>
  );
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      justifyContent: (props:Props) =>
        props.isOther
          ? 'flex-start'
          : 'flex-end',
      display: 'flex',
      flexWrap: 'wrap',
      '& > *': {
        background: (props:Props) =>
          props.isOther
            ? '#d3d3d3'
            : '#98fb98',
        margin: theme.spacing(1),
        padding: theme.spacing(1),
      },
    },
  }),
);

export default Bubble