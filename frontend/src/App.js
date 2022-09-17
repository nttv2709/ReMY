import Calendar from 'react-calendar'
import 'react-calendar/dist/Calendar.css';
import CalGrid from './component/calendarGrid/calendarGrid';

import * as React from 'react';
import Geolocation from "./component/Geolocation/geographic";
function App() {

    return (
    <div className="App">
      <>
        <Calendar/>
        <CalGrid/>
        <Geolocation/>
      </>
    </div>
  );
}

export default App;
