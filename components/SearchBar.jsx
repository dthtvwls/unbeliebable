import $ from 'jquery';
import React from 'react';
import Autosuggest from 'react-autosuggest';

export default class SearchBar extends React.Component {

  constructor() {
    super();
    this.getSuggestions = this.getSuggestions.bind(this);
  }

  getSuggestions(input, callback) {
    $.getJSON("/search?q=" + input, function (results) {
      if (results) callback(null, results.slice(0, 5));
    }.bind(this));
  }

  renderSuggestion(suggestion, input) {
    return suggestion.value;
  }

  onSuggestionSelected(suggestion, event) {
    event.preventDefault();
    $.post("/playlist", { name: suggestion.name, artist: suggestion.artist });
  }

  render() {
    return <nav className="navbar navbar-default navbar-fixed-top">
      <div className="container">
        <form className="navbar-form">
          <Autosuggest
            suggestions={this.getSuggestions}
            suggestionRenderer={this.renderSuggestion}
            onSuggestionSelected={this.onSuggestionSelected}
            suggestionValue={() => ''}
            inputAttributes={{placeholder: 'search for a song...'}}
          />
        </form>
      </div>
    </nav>;
  }
}
