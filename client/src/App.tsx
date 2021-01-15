import React from 'react';
import logo from './logo.svg';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom"

import Footer from './containers/Footer'
import SignIn from './containers/SignIn'
import Header from './containers/Header'
import Top    from './containers/Top'

const App: React.FC = () => {
  return (
    <>
      <Router>
        <Header />
          <Switch>
            <Route exact path="/"><Top /></Route>
            <Route exact path="/signin"><SignIn /></Route>
          </Switch>
        <Footer />
      </Router>
    </>
  );
}

export default App;
