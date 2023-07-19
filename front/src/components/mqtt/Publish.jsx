import React, { useState } from "react";
import mq from "../../store/Mq";

const Publish = () => {
  const [mes, setMes] = useState("setout2=10");
  const onMes = (e) => {
    setMes(e.target.value);
  };
  const send = () => {
    mq.mqttPublish({ payload: mes });
  };
  return (
    <div className="mb-2">
      <input className="me-2" onChange={onMes} value={mes} />
      <button onClick={send}>Publish</button>
    </div>
  );
};

export default Publish;
