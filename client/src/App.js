import { useEffect } from 'react';
import { Col, Container, Row } from 'react-bootstrap';
import axios from 'axios';

function App() {
  useEffect(() => {
    const fetchData = async () => {
      const apiUrl = 'http://localhost:8000';
    };
    fetchData();
    return () => {};
  }, []);

  return (
    <Container>
      <Row>
        <Col xs='12 py-2'>
          <h1 className='text-center'>
            React Application with Go fiber Backend
          </h1>
        </Col>
      </Row>
    </Container>
  );
}

export default App;
