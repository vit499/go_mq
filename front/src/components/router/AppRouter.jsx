import { observer } from "mobx-react-lite";
import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import AboutPage from "../../pages/AboutPage";
import AdminPage from "../../pages/AdminPage";
import ErrorPage from "../../pages/ErrorPage";
import HomePage from "../../pages/HomePage";
import LoginPage from "../../pages/LoginPage";
import {
  ABOUT_ROUTE,
  HOME_ROUTE,
  SETT_ROUTE,
  LOGIN_ROUTE,
} from "./constRouter";

const AppRouter = () => {
  // const userStore = { isAuth: true };
  //const userStore = { isAuth: false };
  return (
    <Routes>
      <Route path={LOGIN_ROUTE} element={<LoginPage />} />
      <Route path={SETT_ROUTE} element={<AdminPage />} />
      <Route
        path={HOME_ROUTE}
        element={<HomePage />}
        errorElement={<ErrorPage />}
      />
      <Route path="*" element={<Navigate to={<HomePage />} />} />
      <Route path={ABOUT_ROUTE} element={<AboutPage />} />
    </Routes>
  );
};

export default AppRouter;
