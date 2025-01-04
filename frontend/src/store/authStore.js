import { create } from "zustand";
import axios from "axios";
import { useState } from "react";
import Cookies from "js-cookie";
import { toast } from "react-toastify";

const useAuthStore = create((set, get) => ({
  email: "",
  fio: "",
  password: "",
  phone_number: "",
  showConfirmation: false,
  code: "",
  isAuthenticated: false,
  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  setPassword: (password) => set({ password }),
  setShowConfirmation: (showConfirmation) =>
    set({ showConfirmation: showConfirmation }),
  setCode: (code) => set({ code }),
  setIsAuthenticated: (status) => set({ isAuthenticated: status }),

  checkAuthStatus: () => {
    const loggedIn = Cookies.get("logged_in");
    get().setIsAuthenticated(!!loggedIn);
  },

  registerFunc: async () => {
    const { email, fio, phone_number, password } = useAuthStore.getState();

    try {
      const response = await axios.post(
        `http://localhost:8080/api/v1/auth/register`,
        {
          email,
          fio,
          phone_number,
          password,
        },
        {
          withCredentials: true,
        }
      );
      console.log("Response:", response);

      // Исправлено: проверка статуса ответа
      if (response.data.status === "success") {
        set({ showConfirmation: true });
        toast.info("Пожалуйста, проверьте ваш email для подтверждения.");
      }
    } catch (error) {
      toast.error(
        "Ошибка регистрации: " +
          (error.response?.data?.message || error.message)
      );
      console.error("Error Registrations:", error);
    }
  },

  loginFunc: async (navigate, register) => {
    const { email, password, showLogin } = useAuthStore.getState();
    try {
      const response = await axios.post(
        `http://localhost:8080/api/v1/auth/login`,
        {
          email,
          password,
        },
        {
          withCredentials: true,
        }
      );
      console.log("login", response);
      if (response.data.status === "success") {
        get().checkAuthStatus();
        navigate("/profile");
        toast.success("Успешный вход!");
      }
    } catch (error) {
      toast.error(
        "Ошибка авторизации: " +
          (error.response?.data?.message || error.message)
      );
      console.error("Error Auth:", error);
    }
  },

  verifyFunc: async (navigate) => {
    const { email, code } = useAuthStore.getState();

    try {
      const response = await axios.post(
        `http://localhost:8080/api/v1/auth/verify-code`,
        {
          email,
          code,
        },
        {
          withCredentials: true,
        }
      );
      console.log("Response:", response);
      if (response.data.status === "success") {
        navigate("/auth");
        toast.success("Код подтвержден!");
      }
    } catch (error) {
      // Если произошла ошибка, очищаем статус аутентификации
      toast.error("Ошибка: не правильный код верификации " + error.message);
      console.error("Error Verify:", error);
    }
  },

  Logout: () => {
    try {
      const response = axios.post(`${url}/api/v1/auth/logout`, {
        withCredentials: true,
      });
    } catch (error) {
      console.error("Ошибка при выходе:", error); // Рекомендуется обработать ошибку
    }
  },
}));
export default useAuthStore;
