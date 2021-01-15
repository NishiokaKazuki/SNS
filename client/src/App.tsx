import React from 'react';
import logo from './logo.svg';
import './App.css';

import Footer from './containers/Footer'
import SignIn from './containers/SignIn'
import Header from './containers/Header'

const App: React.FC = () => {
  return (
    <>
      <Header />
      <Footer />
    </>
  );
}

export default App;
