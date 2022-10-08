import React from 'react';
import './App.css';
import Dashboard from './components/Dashboard';

function App() {
  return (
    <div className="App">
      <div>{process.env.API_URL}</div>
      <Dashboard/>

    </div>
  );
}

export default App;
