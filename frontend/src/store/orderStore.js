import { create } from "zustand";
import axios from "axios";
import { useState } from "react";
import Cookies from "js-cookie";
import { toast } from "react-toastify";

const useOrderStore = create((set, get) => ({
  email: "",
  fio: "",
  phone_number: "",
  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  order: {},

  payOrder: async () => {
    const { email, fio, phone_number } = useOrderStore.getState();

    try {
      const response = await axios.post(
        `http://localhost:8080/api/v1/order`,
        {
          email,
          fio,
          phone_number,
        },
        {
          withCredentials: true,
        }
      );
      console.log("Response:", response);
      //   set({ order: response.data });
      // Исправлено: проверка статуса ответа
      if (response.data.status === "success") {
        window.location.href = response.data.data.id;
      }
    } catch (error) {
      toast.error(
        "Ошибка оплаты: " + (error.response?.data?.message || error.message)
      );
      console.error("Error Registrations:", error);
    }
  },
}));
export default useOrderStore;
