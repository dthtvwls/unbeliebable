import $ from 'jquery';
import React from 'react';

export default class Song extends React.Component {

  constructor() {
    super();
    this.voteDown = this.voteDown.bind(this);
    this.voteUp = this.voteUp.bind(this);
  }

  voteDown(event) {
    event.preventDefault();
    $.post("/votes", { id: this.props.id, against: true });
  }

  voteUp(event) {
    event.preventDefault();
    $.post("/votes", { id: this.props.id, against: false });
  }

  render() {
    return <li className="list-group-item">
      <span className="badge">
        <a href="#" onClick={this.voteDown}>👎</a>
        &nbsp; {this.props.score} &nbsp;
        <a href="#" onClick={this.voteUp}>👍</a>
      </span>
      {this.props.name + " - " + this.props.artist}
    </li>;
  }
}
