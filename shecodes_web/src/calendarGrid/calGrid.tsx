import * as React from 'react'
import FullCalendar, { EventApi, DateSelectArg, EventClickArg, EventContentArg, formatDate } from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import { INITIAL_EVENTS, createEventId } from './event-utils'
import { CalendarServiceClient } from '../api/CalendarServiceClientPb'
import { CreateEventRequest, GetRemindRequest, Location, RangeTime } from '../api/calendar_pb'
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
import { triggerAsyncId } from 'async_hooks'
// import { FormControl, InputLabel, Input, FormHelperText } from '@mui/material';


// import Box from '@mui/material/Box';
// import Button from '@mui/material/Button';
// import Typography from '@mui/material/Typography';
// import Dialog from '@mui/material/Dialog';
// import Modal from '@mui/material/Modal';

interface CalGridState {
  weekendsVisible: boolean
  currentEvents: EventApi[]
  currentPos: string
}

let map: google.maps.Map;


const configURL = "https://0aa8-115-73-142-201.ap.ngrok.io"
const getRemoteHost = () => {
 try {
  const request = new XMLHttpRequest();
  request.open('GET', configURL, false);  // `false` makes the request synchronous
  request.send(null);

  console.log(request.responseText);
  return JSON.parse(request.responseText).host as string;

 } catch (error) {
  console.error('get host err',error)
  return '';
  
 }
}
console.log('initial with host', getRemoteHost());
const client = new CalendarServiceClient(getRemoteHost(), null, null)

export default class CalGrid extends React.Component<{}, CalGridState> {
  state: CalGridState = {
    weekendsVisible: true,
    currentEvents: [],
    currentPos: ""
  }
  async create(title: string,x:number,y:number) {
    const req = new CreateEventRequest()
    req.setTitle(title)
    req.setLocation(new Location().setX(x).setY(y))
    req.setRangeTime(new RangeTime().setStart(new Timestamp().setSeconds(20)).setEnd(new Timestamp().setSeconds(20)))
    const res = await client.createEvent(req, {}, (err, response) => {
      if (err) {
        console.log("Error:", err);
      } else {
        console.log("Response: ", response)
        
      }
    })
  }
  async GetRemindRequest(x:number, y:number){
    const req = new GetRemindRequest();
    req.setLocation(new Location().setX(x).setY(y))
    const res = await client.getRemind(req, {}, (err, response) => {
      if (err) {
        console.log("Error:", err);
      } else {
        console.log("Response: ", response)
        prompt(""+(response.getDuration()), ""+response.getOkay())
      }
    })
  }
  render() {
    setTimeout(() => {      
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(
          (position: GeolocationPosition) => {
            this.GetRemindRequest(position.coords.latitude,position.coords.longitude) 
            const pos = {
              lat: position.coords.latitude,
              lng: position.coords.longitude,
            };
            console.log("current position", pos);
          },
          () => {
            console.log("Error", "The Geolocation service failed.")
          }
        );
      }      
    }, 1000*5*60);
    return (
      <div className='demo-app'>
        {this.renderSidebar()}
        <div className='demo-app-main'>
          <FullCalendar
            plugins={[dayGridPlugin, timeGridPlugin, interactionPlugin]}
            headerToolbar={{
              left: 'prev,next today',
              center: 'title',
              right: 'dayGridMonth,timeGridWeek,timeGridDay'
            }}
            initialView='dayGridMonth'
            editable={true}
            selectable={true}
            selectMirror={true}
            dayMaxEvents={true}
            weekends={this.state.weekendsVisible}
            initialEvents={INITIAL_EVENTS} // alternatively, use the `events` setting to fetch from a feed
            select={this.handleDateSelect}
            eventContent={renderEventContent} // custom render function
            eventClick={this.handleEventClick}
            eventsSet={this.handleEvents} // called after events are initialized/added/changed/removed
          /* you can update a remote database when these fire:
          eventAdd={function(){}}
          eventChange={function(){}}
          eventRemove={function(){}}
          */
          />
          <select 
          value={this.state.currentPos} 
          onChange={this.handleSelectLocation}>
            <option value={"10.765098712735305 106.65557228171438"} >184 Đ. Lê Đại Hành, Phường 15, Quận 11, Thành phố Hồ Chí Minh, Việt Nam</option>
            <option value={"10.77335143652137, 106.66063629226817"}>Trường Đại học Bách khoa - Đại học Quốc gia TP.HCM</option>
            <option value={"10.880832627089323, 106.80533378171523"}>Trường Đại học Bách Khoa - Đại học Quốc gia Thành phố Hồ Chí Minh (cơ sở 2)</option>
            <option value={"10.75373625368853, 106.65456075110404"}>Phuc Long Cofee  Tea Express</option>
            <option value={"10.760705483704767, 106.66339228803797"}>Sân vận động Thống Nhất</option>
          </select>
          <p>{`You selected ${this.state.currentPos}`}</p>

        </div>
      </div>
    )
  }

  renderSidebar() {
    return (
      <div className='demo-app-sidebar'>
        <div className='demo-app-sidebar-section'>
          <h2>Instructions</h2>
          <ul>
            <li>Select dates and you will be prompted to create a new event</li>
            <li>Drag, drop, and resize events</li>
            <li>Click an event to delete it</li>
          </ul>
        </div>
        <div className='demo-app-sidebar-section'>
          <label>
            <input
              type='checkbox'
              checked={this.state.weekendsVisible}
              onChange={this.handleWeekendsToggle}
            ></input>
            toggle weekends
          </label>
        </div>
        <div className='demo-app-sidebar-section'>
          <h2>All Events ({this.state.currentEvents.length})</h2>
          <ul>
            {this.state.currentEvents.map(renderSidebarEvent)}
          </ul>
        </div>
      </div>
    )
  }
  handleSelectLocation = (e: any) => {
    this.setState({
      currentPos: e.target.value
    })
  }
  handleWeekendsToggle = () => {
    this.setState({
      weekendsVisible: !this.state.weekendsVisible
    })
  }

  getLocation = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position: GeolocationPosition) => {
          const pos = {
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          };
          console.log("position", pos);
        },
        () => {
          console.log("Error", "The Geolocation service failed.")
        }
      );
    } else {
      // Browser doesn't support Geolocation
      console.log("Error", "The Geolocation service failed.")
    }
  }

  handleDateSelect = (selectInfo: DateSelectArg) => {
    let title = prompt('Please enter a new title for your event');

    let calendarApi = selectInfo.view.calendar;
    // this.create()
    let display = this.state.currentPos
    

    calendarApi.unselect() // clear date selection
    if (title) {
      calendarApi.addEvent({
        id: createEventId(),
        title,
        start: selectInfo.startStr,
        end: selectInfo.endStr,
        allDay: selectInfo.allDay,
        display:display//location
      })
      let xy = display.split(" ")
      this.create(title,Number(xy[0]),Number(xy[1]))
      // req.setTitle(title)
      // req.setLocation(new Location().setX(Number(xy[0])).setY(Number(xy[1])))
      // req.setRangeTime(new RangeTime().setStart(new Timestamp().setSeconds(20)).setEnd(new Timestamp().setSeconds(20)))
      // const res = client.createEvent(req, {}, (err, response) => {
      //   if (err) {
      //     console.log("Error:", err);
      //   } else {
      //     console.log("Response: ", response)
      //   }
      // })
      
      this.getLocation();
    }
  }//lf crlf

  handleEventClick = (clickInfo: EventClickArg) => {
    if (window.confirm(`Are you sure you want to delete the event '${clickInfo.event.title}'`)) {
      clickInfo.event.remove()
    }
  }

  handleEvents = (events: EventApi[]) => {
    this.setState({
      currentEvents: events
    })
  }

}

function renderEventContent(eventContent: EventContentArg) {
  return (
    <>
      <b>{eventContent.timeText}</b>
      <i>{eventContent.event.title}</i>
    </>
  )
}

function renderSidebarEvent(event: EventApi) {
  return (
    <li key={event.id}>
      <b>{formatDate(event.start!, { year: 'numeric', month: 'short', day: 'numeric' })}</b>
      <i>{event.title}</i>
      <p>{event.display}</p>
    </li>
  )
}