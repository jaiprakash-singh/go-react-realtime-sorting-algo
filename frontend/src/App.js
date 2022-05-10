import './App.css';
import React, { useState, useEffect } from 'react';
import { connect, sendMsg } from './components/api';
import Header from './components/Header';
import Chart from './components/Chart/Chart';
//import InteractiveChart from './components/InteractiveChart';


function App() {
  const [data, setData] = useState([]);
  const [status, setStatus] = useState("init")

  const sort = (event) => {
    sendMsg("sort");
    setStatus("started");
  }

  const stop = (event) => {
    sendMsg("stop");
    setStatus("stopped");
  }

  const resume = (event) => {
    sendMsg("resume");
    setStatus("started");
  }

  const newchart = (event) => {
    sendMsg("newchart");
    setStatus("init");
  }

  useEffect(() => {
    connect(msg => {
      console.log("New message",msg.data);
      if (msg.data.includes("solved:")) {
        setStatus("finished");
      }
      setData(() => msg);
      //console.log("Data:", data);
    });
  }, []);

  //<InteractiveChart />
  
  return (
    <div className="App">
      <Header />
      <Chart data={data} sort={sort} stop={stop} resume={resume} newchart={newchart} status={status} />
    </div>
  );
}

export default App;
