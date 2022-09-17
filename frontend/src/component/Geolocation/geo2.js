import {useEffect, useState} from "react";

export default function Weather() {
    const [observationalData, setObservationalData] = useState({
        stationId: "",
        stationName: "",
        timestamp: 0,
        time: "",
        observationDate: "",
        observationTime: "",
        temperature: 0,
        weather: "",
        title: "",
        humidity: 0,
    });

    const options = {
        enableHighAccuracy: true,
        timeout: 5000,
        maximumAge: 0,
    };

    const getPosition = (data) => {
        //use data to destructure
        //setObservationalData(custom-data);
    };

    function error(err) {
        console.warn(`ERROR(${err.code}): ${err.message}`);
    }

    useEffect(() => {
        const timer = setInterval(getPosition, 5000);
        console.log(new Date());

        return () => {
            clearInterval(timer);
        };
    });

    useEffect(() => {
        navigator.geolocation.getCurrentPosition(getPosition, error, options);
    }, []);
}