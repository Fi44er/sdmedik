import { create } from "zustand";
import axios from "axios";
import { useState } from "react";

const axiosInstance = axios.create({
  timeout: 5000, // таймаут в миллисекундах (5 секунд)
  withCredentials: true,
});

const useUserStore = create((set, get) => ({
  user: null,
  isLoggedOut: false,
  isRefreshingToken: false,
  logoutCalled: false,
  getUserInfo: async () => {
    try {
      if (get().logoutCalled) {
        return;
      }
      const response = await axiosInstance.get(
        "http://localhost:8080/api/v1/user/me"
      );
      set({ user: response.data, isLoggedOut: false });
    } catch (error) {
      if (error.response.status === 401 && !get().isLoggedOut) {
        if (!get().isRefreshingToken && !get().isLoggingOut) {
          set({ isRefreshingToken: true });
          await get().refreshToken();
          set({ isRefreshingToken: false });
        }
      } else if (error.code === "ECONNABORTED") {
        console.error("Таймаут запроса истек");
      } else {
        console.error("Ошибка при получении данных:", error);
      }
    }
  },
  refreshToken: async () => {
    if (get().logoutCalled) {
      return;
    }
    try {
      const response = await axiosInstance.post(
        "http://localhost:8080/api/v1/auth/refresh"
      );
      // После обновления токена повторно вызываем функцию getUserInfo
      await get().getUserInfo();
    } catch (error) {
      if (error.code === "ECONNABORTED") {
        console.error("Таймаут запроса истек");
      } else {
        console.error("Error:", error);
      }
    }
  },
  Logout: async () => {
    try {
      set({ isLoggingOut: true, logoutCalled: true });
      const response = await axiosInstance.post(
        `http://localhost:8080/api/v1/auth/logout`
      );
      set({ isLoggedOut: true, isLoggingOut: false });
    } catch (error) {
      if (error.code === "ECONNABORTED") {
        console.error("Таймаут запроса истек");
      } else {
        console.error("Ошибка при выходе:", error); // Рекомендуется обработать ошибку
      }
    }
  },
}));

export default useUserStore;
