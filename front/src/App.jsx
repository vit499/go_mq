//import "./App.css";
import "./bootstrap.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import TopBar from "./components/router/TopBar";
import HomePage from "./pages/HomePage";
import AboutPage from "./pages/AboutPage";
import ErrorPage from "./pages/ErrorPage";
import LoginPage from "./pages/LoginPage";
import OutPage from "./pages/OutPage";
import TemperComp from "./components/outs/TemperComp";
import HostPage from "./pages/HostPage";
import Description from "./components/mqtt/Description";
import { useEffect } from "react";
import wsStore from "./store/WsStore";
//import router from "./components/router/routes";

const router = createBrowserRouter([
  {
    path: "/",
    element: <TopBar />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <HomePage /> },
      {
        path: "/about",
        element: <AboutPage />,
      },
      {
        path: "/login",
        element: <LoginPage />,
      },
      {
        path: "/host",
        element: <HostPage />,
      },
      {
        path: "/out/:indObj/:indOut",
        element: <OutPage />,
      },
      {
        path: "/temper",
        element: <TemperComp />,
      },
      {
        path: "/descr",
        element: <Description />,
      },
    ],
  },
]);

function App() {
  useEffect(() => {
    console.log("app init");
    wsStore.Init();
    return () => {
      wsStore.Disconnect();
    };
  }, []);
  return <RouterProvider router={router} />;
}

export default App;
