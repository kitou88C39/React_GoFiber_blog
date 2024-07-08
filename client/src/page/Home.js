import React from 'react';

const Home = () => {
  const [apiData, setApiData] = useState(false);
  useEffect(() => {
    return;
    const fetchData = async () => {
      try {
        const apiUrl = process.env.REACT_APP_API_ROOT;
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
};

export default Home;
