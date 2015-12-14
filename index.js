window.jQuery = require('jquery');
require('bootstrap');

import React from 'react';
import ReactDOM from 'react-dom';

import SearchBar from './components/SearchBar.jsx';
import Playlist from './components/Playlist.jsx';

setInterval(function () {
  var swap = true;
  return function () {
    document.title = swap ? '♪♬' : '♬♪';
    swap = !swap;
  };
}(), 1000);

ReactDOM.render(<div><SearchBar /><Playlist /></div>, document.getElementById('root'));
