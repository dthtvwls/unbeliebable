import $ from 'jquery';
import React from 'react';
import Result from './Result.jsx';

export default class SearchBar extends React.Component {

  constructor() {
    super();
    this.state = { results: [] };
    this.onChange = this.onChange.bind(this);
  }

  onChange(event) {
    // keep track of the most recently fired ajax request so that we don't update if an old request is late
    var nonce = Math.random();
    this.setState({ results: this.state.results, nonce: nonce });
    $.getJSON("/search?q=" + event.target.value, function (results) {
      if (this.state.nonce === nonce) {
        if (!results) results = [];
        this.setState({ results: results.slice(0, 10) });
      }
    }.bind(this));
  }

  render() {
    var results = this.state.results.map(function (result) {
      return <Result key={Math.random().toString(36)} {...result} />;
    });
    return <nav className="navbar navbar-default navbar-fixed-top">
      <div className="container">
        <form className="navbar-form" onSubmit={function (event) { event.preventDefault(); }}>
          <div className="input-group input-group-lg">
            <span className="input-group-addon">
              <span className="glyphicon glyphicon-search"></span>
            </span>
            <input ref="q" className="form-control" onChange={this.onChange} placeholder="search for a song..." data-toggle="dropdown" />
            {results.length > 0 ? <ul className="dropdown-menu">{results}</ul> : null}
          </div>
        </form>
      </div>
    </nav>;
  }
}
