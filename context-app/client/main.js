import React from 'react';
import { Meteor } from 'meteor/meteor';
import { render } from 'react-dom';
import App from 'imports/ui/App';
import { Provider } from 'react-redux';
import configureStore from 'store';

let store = configureStore();

Meteor.startup(() => {
  render(
    <Provider store={store}>
      <App />
    </Provider>, document.getElementById('react-target'));
});
