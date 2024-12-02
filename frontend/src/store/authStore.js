import { create } from "zustand";
import axios from "axios";
import { url } from "../constants/contants";
import { useNavigate } from "react-router-dom";

const useAuthStore = create((set, get) => ({
  username: "",
  password: "",
  setUsername: (username) => set({ username }),
  setPassword: (password) => set({ password }),

  AuthFunc: async () => {
    const { username, password } = useAuthStore.getState();

    try {
      const response = await axios.post(
        `${url}/api/v1/auth/login`,
        {
          username,
          password,
        },
        {
          withCredentials: true,
        }
      );

      // Логируем весь ответ
      console.log("Response:", response);

      // Проверяем только статус ответа
      if (response.status === 200) {
      } else {
        alert("Неправильный логин или пароль");
      }
      checkUserAuth();
    } catch (error) {
      // Если произошла ошибка, очищаем статус аутентификации
      alert(
        "Ошибка авторизации: не правильный логин или пороль  " + error.message
      );
      console.error("Error Auth:", error);
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
