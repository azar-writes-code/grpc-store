import { useEffect, useState } from "react";
import { SensorClient } from "./proto/sensor_grpc_web_pb";
import { SensorRequest } from "./proto/sensor_pb";

var client = new SensorClient("http://localhost:8000");

function App() {
  const [temp, setTemp] = useState(-9999);
  const [humidity, setHumidity] = useState(-99999);

  const getTemp = () => {
    console.log("Temp Called!");
    var sensorReq = new SensorRequest();
    var stream = client.tempSensor(sensorReq, {});

    stream.on("data", function (response) {
      setTemp(response.getValue());
    });
  };

  const getHumidity = () => {
    console.log("Humidity Called!");
    var sensorReq = new SensorRequest();
    var stream = client.humiditySensor(sensorReq, {});

    stream.on("data", function (response) {
      setHumidity(response.getValue());
    });
  };

  useEffect(() => {
    getTemp();
  }, []);

  useEffect(() => {
    getHumidity();
  }, []);

  return (
    <div className="bg-cyan-400 min-h-full max-w-screen">
      <h1>
        Temperature : {temp} F
        <br />
        Humidity : {humidity} %
      </h1>
    </div>
  );
}

export default App;
