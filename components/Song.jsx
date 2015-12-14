import React from 'react';

export default class Song extends React.Component {

  render() {
    return <li className="list-group-item">
      <span className="badge"><a href="#">👎</a> 0 <a href="#">👍</a></span>
      {this.props.name + " - " + this.props.artist}
    </li>;
  }
}
