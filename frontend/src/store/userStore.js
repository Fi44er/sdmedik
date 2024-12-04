import { create } from "zustand";
import axios from "axios";
import { useState } from "react";

const useUserStore = create((set, get) => ({
  user: null,
  getUserInfo: async () => {
    try {
      const response = await axios.get("http://localhost:8080/api/v1/user/me", {
        withCredentials: true,
      });
      set({ user: response.data });
    } catch (error) {
      console.error("Ошибка при получении данных:", error);
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

export default useUserStore;
