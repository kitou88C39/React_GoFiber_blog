import { useEffect, useState } from 'react';
import { Col, Container, Row } from 'react-bootstrap';
import axios from 'axios';

function App() {
  const [apiData, setApiData] = useState(false);
  useEffect(() => {
    const fetchData = async () => {
      try {
        const apiUrl = 'http://localhost:8000';
        const response = await axios.get(apiUrl);

        if (response.status === 200) {
          setApiData(response?.data);
        }
      } catch (error) {
        console.log(error.response);
      }
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
