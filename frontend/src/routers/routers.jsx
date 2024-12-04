import React from "react"; // Импортируйте React
import { createBrowserRouter } from "react-router-dom"; // Убедитесь, что импортируете правильно
import HomePage from "../pages/home/HomePage";
import СategoriesPage from "../pages/categories/СategoriesPage";
import CatalogsLayout from "../pages/catalog/CatalogsLayout";
import BasketLayout from "../pages/basket/BasketLayout";
import Delivery from "../pages/delivery/Delivery";
import About from "../pages/about/About";
import Return_policy from "../pages/return_policy/Return_policy";
import Deteils from "../pages/deteils/Deteils";
import ProductDynamicPage from "../pages/Product/ProductDynamicPage";
import Auth from "../pages/account/Auth";
import Register from "../pages/account/Register";
import UserAccount from "../pages/account/UserAccount";
export const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/catalog",
    element: <СategoriesPage />,
  },

  //   {
  //     path: "/news",
  //     element: <News />,
  //   },
  {
    path: "/products/:id", // динамический маршрут
    element: <CatalogsLayout />, // Исправлено имя компонента
  },
  {
    path: "/product/:id", // динамический маршрут
    element: <ProductDynamicPage />, // Исправлено имя компонента
  },
  {
    path: "/basket/:id", // динамический маршрут
    element: <BasketLayout />, // Исправлено имя компонента
  },
  {
    path: "/delivery", // динамический маршрут
    element: <Delivery />, // Исправлено имя компонента
  },
  {
    path: "/about", // динамический маршрут
    element: <About />, // Исправлено имя компонента
  },
  {
    path: "/returnpolicy", // динамический маршрут
    element: <Return_policy />, // Исправлено имя компонента
  },
  {
    path: "/deteils", // динамический маршрут
    element: <Deteils />, // Исправлено имя компонента
  },

  {
    path: "/auth",
    element: <Auth />,
  },
  {
    path: "/register",
    element: <Register />,
  },
  {
    path: "/profile",
    element: <UserAccount />,
  },
  //   {
  //     path: "/Admin",
  //     element: <Admin />,
  //   },
]);
