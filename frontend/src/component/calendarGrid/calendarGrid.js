import React from 'react'
import FullCalendar from '@fullcalendar/react' // must go before plugins
import dayGridPlugin from '@fullcalendar/daygrid' // a plugin!

function CalGrid(){
  const client = new CalendarServiceClient("http://localhost:50052", null, null)
  
  async function create(){
    const req = new CreateEventRequest()
    req.setTitle("1")
    req.setLocation(new Location().setX(1).setY(1))
    req.setRangeTime(new RangeTime().setStart(new Timestamp().setSeconds(20)).setEnd(new Timestamp().setSeconds(20)))
    const res = client.createEvent(req, {}, (err, response) => {
      if (err) {
        console.log("Error:", err);
      } else {
        console.log("Response: ", response)
      }
    })
  }
    return (
        <FullCalendar
        plugins={[ dayGridPlugin ]}
        initialView="dayGridMonth"
        weekends={true}
        events={[
          { title: 'event 1', date: '2022-09-17' }
        ]}
        />
    );
}

export default CalGrid;