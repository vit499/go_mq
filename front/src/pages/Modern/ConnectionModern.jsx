import { observer } from "mobx-react-lite";
import React from "react";
import authStore from "../../store/AuthStore";
import mq from "../../store/Mq";

const ConnectionModern = observer(() => {
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

export default ConnectionModern;
