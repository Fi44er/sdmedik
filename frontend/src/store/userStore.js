import { create } from "zustand";
import axios from "axios";
import { url } from "../constants/constants";

const axiosInstance = axios.create({
  timeout: 5000, // таймаут в миллисекундах (5 секунд)
  withCredentials: true, // Важно для работы с куками
});

const useUserStore = create((set, get) => ({
  user: null,
  isLoggedOut: false,
  isRefreshingToken: false,
  logoutCalled: false,
  isLoggingOut: false,

  // Функция для получения информации о пользователе
  getUserInfo: async () => {
    try {
      if (get().logoutCalled) {
        return;
      }
      const response = await axiosInstance.get(`${url}/user/me`);
      set({ user: response.data, isLoggedOut: false });
    } catch (error) {
      if (error.code === "ECONNABORTED") {
        console.error("Таймаут запроса истек");
      } else {
        console.error("Ошибка при получении данных:", error);
      }
    }
  },

  // Функция для обновления токена
  refreshToken: async () => {
    if (get().logoutCalled) {
      return;
    }
    try {
      // Отправляем запрос на обновление токена
      await axiosInstance.post(
        `${url}/auth/refresh`,
        {},
        { withCredentials: true }
      );
      // После обновления токена повторно вызываем функцию getUserInfo
      await get().getUserInfo();
    } catch (error) {
      if (error.code === "ECONNABORTED") {
        console.error("Таймаут запроса истек");
      } else {
        console.error("Ошибка при обновлении токена:", error);
      }
      throw error; // Пробрасываем ошибку, чтобы обработать её в интерцепторе
    }
  },

  // Функция для выхода из системы
  Logout: async () => {
    try {
      set({ isLoggingOut: true, logoutCalled: true });
      await axiosInstance.post(
        `${url}/auth/logout`,
        {},
        {
          withCredentials: true,
          skipAuthRefresh: true, // Флаг для пропуска обновления токена
        }
      );
      // Очищаем состояние пользователя и сбрасываем флаги
      set({
        user: null,
        isLoggedOut: true,
        isLoggingOut: false,
        logoutCalled: false,
      });
    } catch (error) {
      if (error.code === "ECONNABORTED") {
        console.error("Таймаут запроса истек");
      } else {
        console.error("Ошибка при выходе:", error);
      }
      // В случае ошибки сбрасываем флаги
      set({ isLoggingOut: false, logoutCalled: false });
    }
  },
}));

// Интерцептор для обработки ошибок 401
axiosInstance.interceptors.response.use(
  (response) => response, // Если ответ успешный, просто возвращаем его
  async (error) => {
    const originalRequest = error.config;

    // Если ошибка 401 и это не запрос на обновление токена
    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.skipAuthRefresh
    ) {
      originalRequest._retry = true; // Помечаем запрос как повторный

      try {
        await useUserStore.getState().refreshToken(); // Обновляем токен
        return axiosInstance(originalRequest); // Повторяем оригинальный запрос
      } catch (refreshError) {
        // Если не удалось обновить токен, выходим из системы
        await useUserStore.getState().Logout();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export default useUserStore;
