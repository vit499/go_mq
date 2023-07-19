import { ABOUT_ROUTE, SETT_ROUTE, HOME_ROUTE } from "./constRouter";
import HomePage from "../../pages/HomePage";
import AdminPage from "../../pages/AdminPage";
import AboutPage from "../../pages/AboutPage";
import { createBrowserRouter, RouterProvider, Route } from "react-router-dom";

// export const authRoutes = [
//   {
//     path: SETT_ROUTE,
//     element: AdminPage,
//   },
// ];

// export const publicRoutes = [
//   {
//     path: HOME_ROUTE,
//     element: HomePage,
//   },
//   {
//     path: ABOUT_ROUTE,
//     element: AboutPage,
//   },
// ];

const router = createBrowserRouter([
  {
    path: "/",
    element: <div>Hello world!</div>,
  },
]);

export default router;
