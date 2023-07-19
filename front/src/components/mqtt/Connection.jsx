import React from "react";
import mq from "../../store/Mq";
import { observer } from "mobx-react-lite";
import authStore from "../../store/AuthStore";

const Connection = observer(() => {
  return (
    <div className="mb-2">
      {authStore.isAuth ? (
        <button className="me-2" onClick={() => mq.mqTryConnect()}>
          {mq._connectStatus}
        </button>
      ) : (
        <p>No auth</p>
      )}
      {mq._connectStatus === "Connected" && (
        <button onClick={() => mq.mqttDisconnect()}>Disconnect</button>
      )}
    </div>
  );
});

export default Connection;
