import React from 'react';

export default class Song extends React.Component {

  render() {
    return <li className="list-group-item active">
      <div className="progress">
        <div className="progress-bar progress-bar-success progress-bar-striped active" style={{ width: '0%' }}></div>
      </div>
      <span className="badge"><a href="#">👎</a> 0 <a href="#">👍</a></span>
      {this.props.name + " - " + this.props.artist}
    </li>;
  }
}
