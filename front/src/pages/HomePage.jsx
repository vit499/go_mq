import { observer } from "mobx-react-lite";
import React, { useEffect } from "react";
import { Link } from "react-router-dom";
import Connection from "../components/mqtt/Connection";
import SetOut from "../components/outs/SetOut";
import TemperComp from "../components/outs/TemperComp";
import { LOGIN_ROUTE } from "../components/router/constRouter";
import authStore from "../store/AuthStore";
import descrStore from "../store/DescrStore";

const HomePage = observer(() => {
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
          {descrStore.getDescrOut(0, 0) !== "" && (
            <SetOut indObj={0} indOut={0} />
          )}
          {descrStore.getDescrOut(0, 1) !== "" && (
            <SetOut indObj={0} indOut={1} />
          )}
          {descrStore.getDescrOut(0, 2) !== "" && (
            <SetOut indObj={0} indOut={2} />
          )}
          {descrStore.getDescrOut(1, 1) !== "" && (
            <SetOut indObj={1} indOut={1} />
          )}
          {descrStore.getDescrOut(1, 2) !== "" && (
            <SetOut indObj={1} indOut={2} />
          )}
          {descrStore.getDescrOut(2, 0) !== "" && (
            <SetOut indObj={2} indOut={0} />
          )}
          <TemperComp />
        </div>
      )}
    </div>
  );
});

export default HomePage;
