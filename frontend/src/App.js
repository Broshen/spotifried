import React from 'react';
import { BrowserRouter as Router, Route} from "react-router-dom";


import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'


import Profile from './Profile';
import Compare from './Compare';
import Home from './Home';
import {api_url} from "./variables"

import './App.css';

library.add(fas)

function ShareRedirect ({ match }) {
  window.location.href = api_url + "share/" + match.params.id
  return null 
}

function App() {
  return (
    <Router>

      <Route path="/" exact component={Home} />
      <Route path="/profile/:id" component={Profile} />
      <Route path="/compare/:id1/:id2" component={Compare} />
      <Route path="/share/:id" component={ShareRedirect} />

    </Router>
  );
}

export default App;
