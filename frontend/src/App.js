import './App.css';
import Calendar from 'react-calendar'
import 'react-calendar/dist/Calendar.css';

import * as React from 'react';
import AddButton from "./Event/AddButton";
// import Button from '@mui/material/Button';
// import AddButton from "./Event/AddButton";
// import { useNavigate } from "react-router-dom";
//
// const style = {
//     position: 'absolute',
//     top: '50%',
//     left: '50%',
//     transform: 'translate(-50%, -50%)',
//     width: 400,
//     bgcolor: 'background.paper',
//     border: '2px solid #000',
//     boxShadow: 24,
//     p: 4,
// };

function App() {

    // let navigate = useNavigate()
    // const handleOpen = () =>{
    //     let path = <AddButton/>;
    //     navigate(path);
    // }

    return (
    <div className="App">
      <>
        <Calendar/>
          {/*<Button onClick={handleOpen}>Add events</Button>*/}
          <AddButton/>
      </>
    </div>
  );
}

export default App;
