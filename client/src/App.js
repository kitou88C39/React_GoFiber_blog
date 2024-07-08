import { useEffect, useState } from 'react';
import { Col, Container, Row } from 'react-bootstrap';
import axios from 'axios';
import './App.css';

function App() {
  return (
    <Container>
      <Row>
        <Col xs='12' className='py-2'>
          <h1 className='text-center'>
            React Application with Go fiber Backend
          </h1>
        </Col>
        {apiData &&
          apiData.map((record, index) => (
            <Col key={index} xs='4' className='py-5 box'>
              <div className='title'>{record.title}</div>
              <div>{record.post}</div>
            </Col>
          ))}
      </Row>
    </Container>
  );
}

export default App;
