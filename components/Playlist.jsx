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
  }

  componentDidMount() {
    this.loadPlaylistFromServer();
    this.state.interval = setInterval(this.loadPlaylistFromServer.bind(this), 5000);
  }

  componentWillUnmount() {
    clearInterval(this.state.interval);
  }

  render() {
    return <div className="container">
      <ul className="list-group">
        {this.state.playlist.map(function (song) {
          return <Song key={song.ID} name={song.Name} artist={song.Artist} />;
        })}
      </ul>
    </div>;
  }
}
