import React, { Fragment } from 'react';
import NavBar from './NavBar.jsx';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { ThemeProvider } from '@material-ui/styles';
import { createMuiTheme } from '@material-ui/core/styles';
import materialTheme from 'assets/theme.json';
import { createStore } from 'redux';

const _theme = createMuiTheme(materialTheme);

const App = () => (
  <ThemeProvider theme={_theme}>
    <NavBar/>
  </ThemeProvider>
);

export default App;
