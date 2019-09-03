import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';


class ContextViewer extends Component {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
    this.handleRefreshClick = this.handleRefreshClick.bind(this);
  }

  componentDidMount() {
    const { dispatch, selectedTopic } = this.props
    dispatch(fetchPostsIfNeeded(selectedSubreddit))
  }
}