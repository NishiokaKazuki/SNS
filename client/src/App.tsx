import React from 'react';
import logo from './logo.svg';
import styled from 'styled-components'
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
import ScrollToTop from './containers/Wrapper/ScrollToTop'

const App: React.FC = () => {
  return (
    <>
      <Router>
        <ScrollToTop/>
        <Header />
          <Main>
            <Switch>
              <Route exact path="/"><Top /></Route>
              <Route exact path="/signin"><SignIn /></Route>
            </Switch>
          </Main>
        <Footer />
      </Router>
    </>
  );
}

const Main = styled.main`
    width: 500px;
    max-width: 100vw;
    min-height: 100vh;
    margin: 0 auto;
`

export default App;
