import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import CalGrid from './component/calendarGrid/calendarGrid';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <CalGrid />
  </React.StrictMode>
);

