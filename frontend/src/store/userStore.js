import { create } from "zustand";
import axios from "axios";
import { useState } from "react";
import { url } from "../constants/constants";

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
      const response = await axiosInstance.get(`${url}/user/me`);
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
  users: [],
  fetchUsers: async () => {
    try {
      const response = await axios.get(`${url}/user`);
      set({ users: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },
  refreshToken: async () => {
    if (get().logoutCalled) {
      return;
    }
    try {
      const response = await axiosInstance.post(`${url}/auth/refresh`);
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
      const response = await axiosInstance.post(`${url}/auth/logout`);
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
