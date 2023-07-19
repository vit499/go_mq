import { observer } from "mobx-react-lite";
import React, { useEffect } from "react";
import { Link, useParams } from "react-router-dom";
import { LOGIN_ROUTE } from "../../components/router/constRouter";
import authStore from "../../store/AuthStore";
import ConnectionModern from "./ConnectionModern";
import SetOutModern from "./SetOutModern";
//import "./modern.css";

const OutModernPage = observer(() => {
  const { indObj, indOut } = useParams();
  // console.log(`indObj=${indObj} indOut=${indOut}`);

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
          <ConnectionModern />
          <SetOutModern indObj={indObj} indOut={indOut} />
        </div>
      )}
    </div>
  );
});

export default OutModernPage;
