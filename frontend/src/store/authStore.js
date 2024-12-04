import { create } from "zustand";
import axios from "axios";
import { useState } from "react";

const useAuthStore = create((set, get) => ({
  email: "",
  fio: "",
  password: "",
  phone_number: "",
  showConfirmation: false,
  code: "",
  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  setPassword: (password) => set({ password }),
  setShowConfirmation: (showConfirmation) =>
    set({ showConfirmation: showConfirmation }),
  setCode: (code) => set({ code }),

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
      } else {
        alert(
          "Ошибка регистрации: " +
            (response.data.message || "неизвестная ошибка")
        );
      }
    } catch (error) {
      alert(
        "Ошибка регистрации: " +
          (error.response?.data?.message || error.message)
      );
      console.error("Error Registrations:", error);
    }
  },
  loginFunc: async (navigate) => {
    const { email, password } = useAuthStore.getState();
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
        navigate("/profile");
      }
    } catch (error) {
      alert(
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
      }
    } catch (error) {
      // Если произошла ошибка, очищаем статус аутентификации
      alert("Ошибка: не правильный код верефикации" + error.message);
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
