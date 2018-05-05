import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Navigation from './Navigation';
import QueryForm from './QueryForm';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Navigation />
        <br />
        <QueryForm />
      </div>
    );
  }
}

export default App;
