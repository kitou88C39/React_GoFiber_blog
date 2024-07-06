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
          if (response?.data.statusText === 'OK') {
            setApiData(response?.data?.blog_records);
          }
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
        <Col xs='12' className='py-2'>
          <h1 className='text-center'>
            React Application with Go fiber Backend
          </h1>
        </Col>
        {apiData &&
          apiData.map((record, index) => (
            <Col xs='4'>
              <div>{record.title}</div>
              <div>{record.post}</div>
            </Col>
          ))}
      </Row>
    </Container>
  );
}

export default App;
