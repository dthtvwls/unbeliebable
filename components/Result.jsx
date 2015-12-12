import $ from 'jquery';
import React from 'react';

export default class Result extends React.Component {

  handleClick(event) {
    event.preventDefault();
    $.post("/playlist", { name: event.target.dataset.name, artist: event.target.dataset.artist });
  }

  render() {
    return <li><a href="#" onClick={this.handleClick} data-name={this.props.name} data-artist={this.props.artist}>
      {this.props.name + " - " + this.props.artist}
    </a></li>;
  }
}
