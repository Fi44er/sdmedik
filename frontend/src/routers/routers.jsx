import React from "react"; // Импортируйте React
import { createBrowserRouter } from "react-router-dom"; // Убедитесь, что импортируете правильно
import HomePage from "../pages/home/HomePage";
import СategoriesPage from "../pages/categories/СategoriesPage";
import CatalogsLayout from "../pages/catalog/CatalogsLayout";
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
  //   {
  //     path: "/Admin",
  //     element: <Admin />,
  //   },
]);
