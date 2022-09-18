import Calendar from 'react-calendar'
import 'react-calendar/dist/Calendar.css';
import CalGrid from './component/calendarGrid/calendarGrid';

import * as React from 'react';
import Area from "./component/Geolocation/geo2";
function App() {

    return (
    <div className="App">
      <>
        <Calendar/>
        <CalGrid/>
        <Area/>
      </>
    </div>
  );
}

export default App;
