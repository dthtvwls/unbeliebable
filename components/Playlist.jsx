import $ from 'jquery';
import React from 'react';
import Song from './Song.jsx';

export default class Playlist extends React.Component {

  constructor() {
    super();
    this.state = { playlist: [] };
  }

  loadPlaylistFromServer() {
    $.getJSON("/playlist", function (playlist) {
      if (!playlist) playlist = [];
      this.setState({ playlist: playlist });
    }.bind(this));
    $.getJSON("/nowplaying", function (nowplaying) {
      this.setState({ nowplaying: nowplaying });
    }.bind(this));
    $.getJSON("/elapsedtime", function (elapsedtime) {
      this.setState({ elapsedtime: elapsedtime });
    }.bind(this));
  }

  componentDidMount() {
    this.loadPlaylistFromServer();
    this.state.interval = setInterval(this.loadPlaylistFromServer.bind(this), 1000);
  }

  componentWillUnmount() {
    clearInterval(this.state.interval);
  }

  render() {
    return <div className="container">
      <ul className="list-group">
        {this.state.nowplaying ?
          <li className="list-group-item active">
            <div className="progress">
              <div className="progress-bar progress-bar-success progress-bar-striped active" style={{ width: this.state.elapsedtime + '%' }}></div>
            </div>
            <span className="badge">
              <a href="#">👎</a>
              &nbsp; {this.state.nowplaying.Votes ? this.state.nowplaying.Votes.reduce(function (score, vote) { return score + vote.Value; }, 0) : 0} &nbsp;
              <a href="#">👍</a>
            </span>
            {this.state.nowplaying.Name + " - " + this.state.nowplaying.Artist}
          </li>
        : ''}
        {this.state.playlist.map(function (song) {
          return <Song key={song.ID} id={song.ID} name={song.Name} artist={song.Artist} score={song.Votes ? song.Votes.reduce(function (score, vote) { return score + vote.Value; }, 0) : 0} />;
        })}
      </ul>
    </div>;
  }
}
