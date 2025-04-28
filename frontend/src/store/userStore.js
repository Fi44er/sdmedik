// import { create } from "zustand";
// import axios from "axios";
// import { url } from "../constants/constants";
// import Cookies from "js-cookie";
// import { toast } from "react-toastify";
// import CryptoJS from "crypto-js";

// const axiosInstance = axios.create({
//   timeout: 5000,
//   withCredentials: true,
// });

// // IndexedDB функции (без изменений)
// const openDB = () => {
//   return new Promise((resolve, reject) => {
//     const request = indexedDB.open("userDataDB", 1);
//     request.onupgradeneeded = (event) => {
//       const db = event.target.result;
//       db.createObjectStore("user", { keyPath: "id" });
//     };
//     request.onsuccess = (event) => {
//       resolve(event.target.result);
//     };
//     request.onerror = (event) => {
//       reject(event.target.error);
//     };
//   });
// };

// const encryptData = (data) => {
//   return CryptoJS.AES.encrypt(JSON.stringify(data), "secret_key").toString();
// };

// const decryptData = (encryptedData) => {
//   const bytes = CryptoJS.AES.decrypt(encryptedData, "secret_key");
//   return JSON.parse(bytes.toString(CryptoJS.enc.Utf8));
// };

// const generateHash = (data) => {
//   return CryptoJS.SHA256(JSON.stringify(data)).toString();
// };

// const saveUserToDB = async (user) => {
//   const db = await openDB();
//   const transaction = db.transaction("user", "readwrite");
//   const store = transaction.objectStore("user");
//   const hash = generateHash(user);
//   const encryptedUser = encryptData(user);
//   store.put({ id: 1, data: encryptedUser, hash });
// };

// const getUserFromDB = async () => {
//   const db = await openDB();
//   const transaction = db.transaction("user", "readonly");
//   const store = transaction.objectStore("user");
//   return new Promise((resolve, reject) => {
//     const request = store.get(1);
//     request.onsuccess = (event) => {
//       const result = event.target.result;
//       if (!result) {
//         resolve(null);
//         return;
//       }
//       try {
//         const user = decryptData(result.data);
//         const currentHash = generateHash(user);
//         if (currentHash !== result.hash) {
//           throw new Error("Данные были подменены!");
//         }
//         resolve(user);
//       } catch (error) {
//         reject(error);
//       }
//     };
//     request.onerror = (event) => {
//       reject(event.target.error);
//     };
//   });
// };

// const deleteUserFromDB = async () => {
//   const db = await openDB();
//   const transaction = db.transaction("user", "readwrite");
//   const store = transaction.objectStore("user");
//   store.delete(1);
// };

// const useUserStore = create((set, get) => ({
//   user: null,
//   allUsers: [],
//   isLoggedOut: false,
//   isRefreshingToken: false,
//   isLoggingOut: false,
//   email: "",
//   fio: "",
//   password: "",
//   phone_number: "",
//   showConfirmation: false,
//   code: "",
//   isAuthenticated: false,
//   isLoadingUser: true,
//   refreshTokenPromise: null,
//   setEmail: (email) => set({ email }),
//   setFio: (fio) => set({ fio }),
//   setPhone_number: (phone_number) => set({ phone_number }),
//   setPassword: (password) => set({ password }),
//   setShowConfirmation: (showConfirmation) => set({ showConfirmation }),
//   setCode: (code) => set({ code }),
//   setIsAuthenticated: (status) => set({ isAuthenticated: status }),

//   checkAuthAndExecute: async (callback) => {
//     if (!get().isAuthenticated) {
//       return;
//     }
//     try {
//       await callback();
//     } catch (error) {
//       if (error.response?.status === 401 && !get().isLoggingOut) {
//         try {
//           await get().refreshToken();
//           await callback();
//         } catch (refreshError) {
//           await get().logout();
//         }
//       } else {
//         throw error;
//       }
//     }
//   },

//   isAdmin: () => {
//     const user = get().user;
//     return user?.data?.role_id === 1; // Единая проверка по role_id
//   },

//   checkAuthStatus: async () => {
//     const loggedIn = Cookies.get("logged_in");
//     if (!loggedIn) {
//       set({ isAuthenticated: false, isLoadingUser: false });
//       return false;
//     }
//     try {
//       const response = await axiosInstance.get(`${url}/user/me`);
//       set({
//         isAuthenticated: true,
//         user: response.data, // Предполагаем что сервер возвращает { data: { role_id: 1, ... } }
//         isLoadingUser: false,
//       });
//       await saveUserToDB(response.data);
//       return true;
//     } catch (error) {
//       set({ isAuthenticated: false, isLoadingUser: false });
//       return false;
//     }
//   },

//   initializeUser: async () => {
//     set({ isLoadingUser: true });
//     try {
//       const userFromDB = await getUserFromDB();
//       if (userFromDB && (await get().checkAuthStatus())) {
//         set({ user: userFromDB, isLoadingUser: false });
//       } else {
//         await get().getUserInfo(); // Попробуем загрузить с сервера
//         set({ isLoadingUser: false });
//       }
//     } catch (error) {
//       console.error("Ошибка загрузки данных из IndexedDB:", error);
//       set({ isLoadingUser: false });
//     }
//   },

//   registerFunc: async () => {
//     const { email, fio, phone_number, password } = get();
//     try {
//       const response = await axios.post(
//         `${url}/auth/register`,
//         { email, fio, phone_number, password },
//         { withCredentials: true }
//       );
//       if (response.data.status === "success") {
//         set({ showConfirmation: true });
//         toast.info("Пожалуйста, проверьте ваш email для подтверждения.");
//       }
//     } catch (error) {
//       toast.error(
//         "Ошибка регистрации: " +
//           (error.response?.data?.message || error.message)
//       );
//       console.error("Error Registrations:", error);
//     }
//   },

//   loginFunc: async (navigate) => {
//     const { email, password } = get();
//     try {
//       const response = await axios.post(
//         `${url}/auth/login`,
//         { email, password },
//         { withCredentials: true }
//       );
//       if (response.data.status === "success") {
//         await get().checkAuthStatus();
//         await get().getUserInfo();
//         navigate("/profile");
//         toast.success("Успешный вход!");
//       }
//     } catch (error) {
//       toast.error(
//         "Ошибка авторизации: " +
//           (error.response?.data?.message || error.message)
//       );
//       console.error("Error Auth:", error);
//     }
//   },

//   verifyFunc: async (navigate) => {
//     const { email, code } = get();
//     try {
//       const response = await axios.post(
//         `${url}/auth/verify-code`,
//         { email, code },
//         { withCredentials: true }
//       );
//       if (response.data.status === "success") {
//         navigate("/auth");
//         toast.success("Код подтвержден!");
//       }
//     } catch (error) {
//       toast.error("Ошибка: неправильный код верификации " + error.message);
//       console.error("Error Verify:", error);
//     }
//   },

//   getUserInfo: async () => {
//     if (get().user) return;
//     set({ isLoadingUser: true });
//     try {
//       const response = await axiosInstance.get(`${url}/user/me`);
//       set({
//         user: response.data,
//         isAuthenticated: true, // Добавляем синхронизацию статуса
//         isLoggedOut: false,
//         isLoadingUser: false,
//       });
//       await saveUserToDB(response.data);
//     } catch (error) {
//       console.error("Ошибка при получении данных:", error);
//       set({ isLoadingUser: false });
//     }
//   },

//   refreshToken: async () => {
//     if (get().isLoggingOut || get().refreshTokenPromise) {
//       return get().refreshTokenPromise;
//     }

//     set({ isRefreshingToken: true });
//     try {
//       const refreshPromise = axiosInstance.post(`${url}/auth/refresh`);
//       set({ refreshTokenPromise: refreshPromise });

//       await refreshPromise;
//       set({ isRefreshingToken: false, refreshTokenPromise: null });
//       return get().checkAuthStatus();
//     } catch (error) {
//       set({ isRefreshingToken: false, refreshTokenPromise: null });
//       if (error.response?.status === 401 || error.response?.status === 403) {
//         await get().logout(true); // Передаем флаг для принудительного логаута
//       }
//       throw error;
//     }
//   },

//   logout: async (force = false) => {
//     if (get().isLoggingOut && !force) return;

//     try {
//       set({ isLoggingOut: true });
//       await axiosInstance.post(`${url}/auth/logout`);
//     } catch (e) {
//       // Игнорируем ошибки при логауте
//     } finally {
//       // Полная очистка состояния
//       Cookies.remove("logged_in");
//       set({
//         user: null,
//         isLoggedOut: true,
//         isLoggingOut: false,
//         isAuthenticated: false,
//         isLoadingUser: false,
//       });
//       await deleteUserFromDB();
//       window.location.href = "/";
//     }
//   },

//   fetchUsers: async () => {
//     await get().checkAuthAndExecute(async () => {
//       try {
//         const response = await axios.get(`${url}/user`);
//         set({ allUsers: response.data });
//       } catch (error) {
//         toast.error(error.message);
//       }
//     });
//   },
// }));

// useUserStore.getState().initializeUser();

// axiosInstance.interceptors.response.use(
//   (response) => response,
//   async (error) => {
//     const originalRequest = error.config;
//     const { isLoggingOut, isAuthenticated } = useUserStore.getState();

//     // Пропускаем обработку если уже выходим или пользователь не авторизован
//     if (isLoggingOut || !isAuthenticated) {
//       return Promise.reject(error);
//     }

//     if (
//       error.response?.status === 401 &&
//       !originalRequest._retry &&
//       !originalRequest.url.includes("/auth/refresh") &&
//       !originalRequest.url.includes("/auth/logout")
//     ) {
//       originalRequest._retry = true;

//       try {
//         await useUserStore.getState().refreshToken();
//         return axiosInstance(originalRequest);
//       } catch (refreshError) {
//         await useUserStore.getState().logout(true);
//         return Promise.reject(refreshError);
//       }
//     }

//     return Promise.reject(error);
//   }
// );

// export default useUserStore;

import { create } from "zustand";
import axios from "axios";
import { url } from "../constants/constants";
import Cookies from "js-cookie";
import { toast } from "react-toastify";

const axiosInstance = axios.create({
  timeout: 5000,
  withCredentials: true,
});

const useUserStore = create((set, get) => ({
  user: localStorage.getItem("user")
    ? JSON.parse(localStorage.getItem("user"))
    : null,
  allUsers: [],
  isLoggedOut: false,
  isRefreshingToken: false,
  isLoggingOut: false,
  email: "",
  fio: "",
  password: "",
  phone_number: "",
  showConfirmation: false,
  code: "",
  isAuthenticated: !!Cookies.get("logged_in"),

  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  setPassword: (password) => set({ password }),
  setShowConfirmation: (showConfirmation) => set({ showConfirmation }),
  setCode: (code) => set({ code }),
  setIsAuthenticated: (status) => set({ isAuthenticated: status }),

  checkAuthAndExecute: async (callback) => {
    if (!get().isAuthenticated) {
      return;
    }

    try {
      await callback();
    } catch (error) {
      if (error.response?.status === 401 && !get().isLoggingOut) {
        try {
          await get().refreshToken();
          await callback();
          localStorage.setItem("user", JSON.stringify(get().user));
        } catch (refreshError) {
          await get().logout();
        }
      } else {
        throw error;
      }
    }
  },

  checkAuthStatus: () => {
    const loggedIn = Cookies.get("logged_in");
    get().setIsAuthenticated(!!loggedIn);
  },

  registerFunc: async () => {
    const { email, fio, phone_number, password } = get();
    try {
      const response = await axios.post(
        `${url}/auth/register`,
        { email, fio, phone_number, password },
        { withCredentials: true }
      );
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
    const { email, password } = get();
    try {
      const response = await axios.post(
        `${url}/auth/login`,
        { email, password },
        { withCredentials: true }
      );
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
    const { email, code } = get();
    try {
      const response = await axios.post(
        `${url}/auth/verify-code`,
        { email, code },
        { withCredentials: true }
      );
      if (response.data.status === "success") {
        navigate("/auth");
        toast.success("Код подтвержден!");
      }
    } catch (error) {
      toast.error("Ошибка: не правильный код верификации " + error.message);
      console.error("Error Verify:", error);
    }
  },

  getUserInfo: async () => {
    await get().checkAuthAndExecute(async () => {
      try {
        const response = await axiosInstance.get(`${url}/user/me`);
        set({ user: response.data, isLoggedOut: false });
        localStorage.setItem("user", JSON.stringify(response.data));
      } catch (error) {
        console.error("Ошибка при получении данных:", error);
      }
    });
  },

  refreshToken: async () => {
    if (get().isRefreshingToken || get().isLoggingOut) {
      return;
    }
    try {
      set({ isRefreshingToken: true });
      await axiosInstance.post(
        `${url}/auth/refresh`,
        {},
        { withCredentials: true }
      );
      set({ isRefreshingToken: false });
    } catch (error) {
      set({ isRefreshingToken: false });
      if (error.response?.status === 401) {
        await get().logout();
      }
      throw error;
    }
  },

  logout: async () => {
    try {
      set({ isLoggingOut: true });
      await axiosInstance.post(
        `${url}/auth/logout`,
        {},
        { withCredentials: true }
      );
      set({
        user: null,
        isLoggedOut: true,
        isLoggingOut: false,
        isAuthenticated: false,
      });
      localStorage.removeItem("user");
      window.location.href = "/";
    } catch (error) {
      set({ isLoggingOut: false });
      console.error("Ошибка при выходе:", error);
    }
  },

  fetchUsers: async () => {
    await get().checkAuthAndExecute(async () => {
      try {
        const response = await axios.get(`${url}/user`);
        set({ allUsers: response.data });
      } catch (error) {
        toast.error(error.message);
      }
    });
  },
}));

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.url.includes("/auth/refresh") &&
      !originalRequest.url.includes("/auth/logout")
    ) {
      originalRequest._retry = true;

      try {
        await useUserStore.getState().refreshToken();
        return axiosInstance(originalRequest);
      } catch (refreshError) {
        await useUserStore.getState().logout();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export default useUserStore;
