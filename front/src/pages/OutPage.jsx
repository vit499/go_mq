import { observer } from "mobx-react-lite";
import React, { useEffect } from "react";
import { Link, useParams } from "react-router-dom";
import Connection from "../components/mqtt/Connection";
import SetOut from "../components/outs/SetOut";
import { LOGIN_ROUTE } from "../components/router/constRouter";
import authStore from "../store/AuthStore";

const OutPage = observer(() => {
  const { indObj, indOut } = useParams();
  console.log(`indObj=${indObj} indOut=${indOut}`); //
  // useEffect(() => {
  //   // mq.mqttConnect();
  //   return () => mq.mqttDisconnect();
  // }, []);

  return (
    <div className="container">
      {!authStore.isAuth && (
        <div>
          <p>Need login</p>
          <Link className="me-2" to={LOGIN_ROUTE}>
            Login
          </Link>
        </div>
      )}
      {authStore.isAuth && (
        <div>
          <br />
          <Connection />
          <SetOut indObj={indObj} indOut={indOut} />
        </div>
      )}
    </div>
  );
});

export default OutPage;
