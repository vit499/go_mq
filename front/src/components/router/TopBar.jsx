import { observer } from "mobx-react-lite";
import React from "react";
import { Link, Outlet } from "react-router-dom";
import descrStore from "../../store/DescrStore";
import { ABOUT_ROUTE, HOME_ROUTE, SETT_ROUTE } from "./constRouter";

const TopBar = observer(() => {
  return (
    <>
      <div className="container mb-3">
        <Link className="me-2" to={HOME_ROUTE}>
          Home
        </Link>
        <Link className="me-2" to="/host">
          Host
        </Link>
        <Link className="me-2" to="/descr">
          Descr
        </Link>
        <Link className="me-2" to={ABOUT_ROUTE}>
          About
        </Link>
      </div>
      <div className="container mb-3">
        {descrStore.getDescrOut(0, 0) !== "" && (
          <Link className="me-2" to="/out/0/0">
            {descrStore.getDescrOut(0, 0)}
          </Link>
        )}
        {descrStore.getDescrOut(0, 1) !== "" && (
          <Link className="me-2" to="/out/0/1">
            {descrStore.getDescrOut(0, 1)}
          </Link>
        )}
        {descrStore.getDescrOut(0, 2) !== "" && (
          <Link className="me-2" to="/out/0/2">
            {descrStore.getDescrOut(0, 2)}
          </Link>
        )}
        {descrStore.getDescrOut(1, 0) !== "" && (
          <Link className="me-2" to="/out/1/0">
            {descrStore.getDescrOut(1, 0)}
          </Link>
        )}
        {descrStore.getDescrOut(1, 1) !== "" && (
          <Link className="me-2" to="/out/1/1">
            {descrStore.getDescrOut(1, 1)}
          </Link>
        )}
        {descrStore.getDescrOut(1, 2) !== "" && (
          <Link className="me-2" to="/out/1/2">
            {descrStore.getDescrOut(1, 2)}
          </Link>
        )}
        {descrStore.getDescrOut(2, 0) !== "" && (
          <Link className="me-2" to="/out/2/0">
            {descrStore.getDescrOut(2, 0)}
          </Link>
        )}
        <Link className="me-2" to="/temper">
          T_all
        </Link>
      </div>
      <div>
        <Outlet />
      </div>
    </>
  );
});

export default TopBar;
