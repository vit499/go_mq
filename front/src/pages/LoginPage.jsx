import { observer } from "mobx-react-lite";
import React from "react";
import { Navigate } from "react-router-dom";

import { HOME_ROUTE } from "../components/router/constRouter";
import authStore from "../store/AuthStore";
import hostStore from "../store/HostStore";
import mq from "../store/Mq";

const LoginPage = observer(() => {
  const singin = () => {
    hostStore.saveLoginToStorage();
    mq.mqTryConnect();
  };

  if (authStore.isAuth) {
    return <Navigate to={HOME_ROUTE} />;
  }
  return (
    <div className="container">
      <div className="mb-2 row">
        <label className="col-sm-2">Login</label>
        <div className="col-sm-2">
          <input
            onChange={(v) => hostStore.setLogin(v)}
            value={hostStore.login}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2">Password</label>
        <div className="col-sm-2">
          <input
            type="password"
            onChange={(v) => hostStore.setPassword(v)}
            value={hostStore.password}
          />
        </div>
      </div>
      <div className="mb-2 row">
        <label className="col-sm-2"></label>
        <div className="col-sm-2">
          <button onClick={singin}>SingIn</button>
        </div>
      </div>
    </div>
  );
});

export default LoginPage;
