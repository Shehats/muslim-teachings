import React, { Component } from 'react';
import { withTracker } from 'meteor/react-meteor-data';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import InputBase from '@material-ui/core/InputBase';
import { fade, makeStyles } from '@material-ui/core/styles';
import MenuIcon from '@material-ui/icons/Menu';
import SearchIcon from '@material-ui/icons/Search';
import logo from 'assets/logo.png';
import TextField from '@material-ui/core/TextField';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  logo: {
    maxWidth: 40,
    maxHeight: 40
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
    display: 'none',
    [theme.breakpoints.up('sm')]: {
        display: 'block',
      },
    },
    search: {
      position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.white, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.white, 0.25),
      },
      marginLeft: 0,
      width: '65%',
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
    },
    searchIcon: {
      width: theme.spacing(7),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'right',
      justifyContent: 'right',
    },
    inputRoot: {
      color: 'inherit',
    },
    appBar: {
      background: 'white',
      borderStyle: 'solid bottom',
      borderColor: '96ceb4',
      color: '#96ceb4',
      borderRadius: theme.shape.borderRadius,
    },
    nputInput: {
      padding: theme.spacing(1, 1, 1, 7),
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        width: 120,
        '&:focus': {
          width: 200,
        },
      },
    },
  }
));

let NavBar = () => {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <AppBar position="static" variant="outlined">
        <Toolbar className={classes.appBar}>
          <IconButton
            edge="start"
            className={classes.menuButton}
            color="inherit"
            aria-label="open drawer">
            <MenuIcon />
          </IconButton>
          <img src={logo} className={classes.logo}/>
          <Typography className={classes.title} variant="h6" noWrap>
            Islamic Teachings
          </Typography>
          <div className={classes.search}>
            <TextField
              placeholder="Search teachings........"
              fullWidth/>
          </div>
          <SearchIcon />
        </Toolbar>
      </AppBar>
    </div>
  )
}
export default NavBar;
// export default NavBar = withTracker(() => {
//   return {
//     links: Links.find().fetch(),
//   };
// })(NavBar);
