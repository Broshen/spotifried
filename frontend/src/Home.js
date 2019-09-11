import React from 'react';
import Container from 'react-bootstrap/Container';
import Button from 'react-bootstrap/Button';
import {api_url} from "./variables"


function Home () {
  return (
    
        <Container className="center-page text-center">
        <h1 className="pad-vertical">SPOTIFRIED</h1>
        <h3>Analyze your spotify music library and compare your tastes with your friends</h3>
        <Button variant="success" href={api_url + "authenticate?nocache=" + Math.random()}>Let's do it!</Button>
        </Container>
  )
}


export default Home;
