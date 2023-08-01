import React from "react";
import { observer } from "mobx-react-lite";
import authStore from "../../store/AuthStore";

const Connection = observer(() => {
  return (
    <div className="mb-2">
      {authStore.isAuth ? <div></div> : <p>No auth</p>}
    </div>
  );
});

export default Connection;
