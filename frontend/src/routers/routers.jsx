import React, { useEffect, useState } from "react"; // Импортируйте React
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
import Electronic_certificate from "../pages/electronic_certificate/Electronic_certificate";
import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";
import CreateProduct from "../pages/admin/create_product/CreateProduct";
import Contacts from "../pages/contacts/Contacts";
import CreateCategory from "../pages/admin/create_category/CreateCategory";
import AdminDashboard from "../pages/admin/AdminLayout";
import UpdateProduct from "../pages/admin/update_product/UpdateProduct";
import AdminCategoriesTable from "../pages/admin/components_admin_page/AdminCategoriesTable/AdminCategoriesTable";
import axios from "axios";
import useUserStore from "../store/userStore";
import MainContent from "../pages/admin/components_admin_page/MainContent/MainContent";
import AdminUserTable from "../pages/admin/components_admin_page/AdminUserTable/AdminUserTable";
import AdminProductTable from "../pages/admin/components_admin_page/AdminProductTable/AdminProductTable";
import Paymants from "../pages/paymants/Paymants";
import PayOnclick from "../pages/pay_onclick/PayOnclick";

const UsersRoute = ({ children }) => {
  const isLoggedIn = Cookies.get("logged_in") === "true";

  if (!isLoggedIn) {
    return <Navigate to="/" replace />;
  }
  return children;
};

export default UsersRoute;

export const AdminRoute = ({ children }) => {
  const { getUserInfo, user } = useUserStore();
  const [loading, setLoading] = useState(true);
  const [isAdmin, setIsAdmin] = useState(false);

  useEffect(() => {
    const fetchUserInfo = async () => {
      await getUserInfo();
      setLoading(false);
    };
    fetchUserInfo();
  }, [getUserInfo]);

  useEffect(() => {
    if (user?.data) {
      setIsAdmin(user.data.role_id === 1);
    }
  }, [user]);

  if (loading) {
    return <div>Loading...</div>; // Можно добавить индикатор загрузки
  }

  if (!isAdmin) {
    return <Navigate to="/" replace />;
  }
  return children;
};

export const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/catalog",
    element: <СategoriesPage />,
  },
  {
    path: "/products/:id", // динамический маршрут
    element: <CatalogsLayout />, // Исправлено имя компонента
  },
  {
    path: "/product/:id", // динамический маршрут
    element: <ProductDynamicPage />, // Исправлено имя компонента
  },
  {
    path: "/basket", // динамический маршрут
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
    path: "/certificate", // динамический маршрут
    element: <Electronic_certificate />, // Исправлено имя компонента
  },
  {
    path: "/contacts", // динамический маршрут
    element: <Contacts />, // Исправлено имя компонента
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
    path: "/paymants",
    element: <Paymants />,
  },
  {
    path: "/paymants/:id",
    element: <PayOnclick />,
  },
  {
    path: "/profile",
    element: (
      <UsersRoute>
        <UserAccount />
      </UsersRoute>
    ),
  },
  {
    path: `/admin/*`,
    element: (
      <AdminRoute>
        <AdminDashboard />
      </AdminRoute>
    ),
  },
]);
