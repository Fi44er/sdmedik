import { create } from "zustand";
import axios from "axios";
import { url } from "../constants/constants";
import Cookies from "js-cookie";
import { toast } from "react-toastify";

const axiosInstance = axios.create({
  timeout: 5000, // Таймаут в миллисекундах (5 секунд)
  withCredentials: true, // Важно для работы с куками
});

const useUserStore = create((set, get) => ({
  user: localStorage.getItem("user")
    ? JSON.parse(localStorage.getItem("user"))
    : null, // Получение данных из localStorage
  allUsers: [],
  isLoggedOut: false,
  isRefreshingToken: false,
  logoutCalled: false,
  isLoggingOut: false,
  email: "",
  fio: "",
  password: "",
  phone_number: "",
  showConfirmation: false,
  code: "",
  isAuthenticated: !!Cookies.get("logged_in"), // Проверка куки на авторизацию

  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  setPassword: (password) => set({ password }),
  setShowConfirmation: (showConfirmation) =>
    set({ showConfirmation: showConfirmation }),
  setCode: (code) => set({ code }),
  setIsAuthenticated: (status) => set({ isAuthenticated: status }),

  checkAuthAndExecute: async (callback) => {
    if (!get().isAuthenticated) {
      return; // Если пользователь не авторизован, ничего не делаем
    }

    try {
      await callback();

      // Сохранить данные пользователя в localStorage после успешного выполнения callback
      localStorage.setItem("user", JSON.stringify(get().user));
    } catch (error) {
      if (error.response && error.response.status === 401) {
        // Если произошла ошибка 401, пробуем обновить токен
        await get().refreshToken();

        // Повторная попытка вызова callback
        await callback();

        // Сохранить данные пользователя в localStorage после успешного выполнения callback
        localStorage.setItem("user", JSON.stringify(get().user));
      } else {
        throw error; // Пробрасываем остальные ошибки дальше
      }
    }
  },

  checkAuthStatus: () => {
    const loggedIn = Cookies.get("logged_in");
    get().setIsAuthenticated(!!loggedIn);
  },

  registerFunc: async () => {
    const { email, fio, phone_number, password } = useUserStore.getState();

    try {
      const response = await axios.post(
        `${url}/auth/register`,
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

  loginFunc: async (navigate) => {
    const { email, password, } = useUserStore.getState();
    try {
      const response = await axios.post(
        `${url}/auth/login`,
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
    const { email, code } = useUserStore.getState();

    try {
      const response = await axios.post(
        `${url}/auth/verify-code`,
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
  // Функция для получения информации о пользователе
  getUserInfo: async () => {
    await get().checkAuthAndExecute(async () => {
      try {
        const response = await axiosInstance.get(`${url}/user/me`);
        set({ user: response.data, isLoggedOut: false });
      } catch (error) {
        if (error.code === "ECONNABORTED") {
          console.error("Таймаут запроса истек");
        } else {
          console.error("Ошибка при получении данных:", error);
        }
      }
    });
  },

  // Функция для обновления токена
  refreshToken: async () => {
    if (!get().isAuthenticated || get().logoutCalled) {
      return;
    }
    try {
      set({ isRefreshingToken: true });
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
    } finally {
      set({ isRefreshingToken: false });
    }
  },

  // Функция для выхода из системы
  logout: async () => {
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
        isAuthenticated: false, // Сбрасываем статус аутентификации
      });
      localStorage.removeItem("user"); // Удаляем данные пользователя из localStorage
      window.location.href = "/"; // Перенаправляем на главную страницу
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

  fetchUsers: async () => {
    await get().checkAuthAndExecute(async () => {
      try {
        const response = await axios.get(`${url}/user`);
        set({ allUsers: response.data }); // Update allUsers state, not user state
      } catch (error) {
        toast.error(error.message);
      }
    });
  },

  // Метод для установки статуса аутентифицированности
  setIsAuthenticated: (value) => {
    set({ isAuthenticated: value });
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
