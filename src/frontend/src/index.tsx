import React from 'react';
import ReactDOM from 'react-dom/client';
import { HashRouter } from "react-router-dom";
import './index.css';
import App from './App';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  // hash router has been used instead of browser router
  // because I didn't want to change the nginx config 
  // to redirect to home page in case on 404 and 
  // also there is no requirement for the tool box to be SEO friendly
  // location / {
  //   try_files $uri /index.html;
  //  }
  <HashRouter>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </HashRouter>
);

