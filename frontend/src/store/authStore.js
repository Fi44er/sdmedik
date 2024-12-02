import { create } from "zustand";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const useAuthStore = create((set, get) => ({
  email: "",
  fio: "",
  password: "",
  phone_number: "",
  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  setPassword: (password) => set({ password }),

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

      // Логируем весь ответ
      console.log("Response:", response);
      // Проверяем только статус ответа
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
