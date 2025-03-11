import { create } from "zustand";
import axios from "axios";
import { useState } from "react";
import Cookies from "js-cookie";
import { toast } from "react-toastify";
import { url } from "../constants/constants";

const useOrderStore = create((set, get) => ({
  email: "",
  fio: "",
  phone_number: "",
  delivery_address: "",
  setEmail: (email) => set({ email }),
  setFio: (fio) => set({ fio }),
  setPhone_number: (phone_number) => set({ phone_number }),
  setDelivery_address: (delivery_address) => set({ delivery_address }),
  order: {},
  orders: {},
  userOrders: [],

  payOrder: async () => {
    const { email, fio, phone_number, delivery_address } =
      useOrderStore.getState();

    try {
      const response = await axios.post(
        `${url}/order`,
        {
          email,
          fio,
          phone_number,
          delivery_address,
        },
        {
          withCredentials: true,
        }
      );
      console.log("Response:", response);
      //   set({ order: response.data });
      // Исправлено: проверка статуса ответа
      if (response.data.status === "success") {
        window.location.href = response.data.data.url;
      }
    } catch (error) {
      toast.error(
        "Ошибка оплаты: " + (error.response?.data?.message || error.message)
      );
      console.error("Error Registrations:", error);
    }
  },
  payOrderById: async (id) => {
    const { email, fio, phone_number, delivery_address } =
      useOrderStore.getState();

    try {
      const response = await axios.post(
        `${url}/order/${id}`,
        {
          email,
          fio,
          phone_number,
          delivery_address,
        },
        {
          withCredentials: true,
        }
      );
      console.log("Response:", response);
      //   set({ order: response.data });
      // Исправлено: проверка статуса ответа
      if (response.data.status === "success") {
        window.location.href = response.data.data.url;
      }
    } catch (error) {
      toast.error(
        "Ошибка оплаты: " + (error.response?.data?.message || error.message)
      );
      console.error("Error Registrations:", error);
    }
  },
  changeStatus: async (order_id, status) => {
    try {
      const response = await axios.put(
        `${url}/order/status`,
        {
          order_id,
          status,
        },
        {
          withCredentials: true,
        }
      );
      console.log("Response:", response);
    } catch (error) {
      toast.error(
        "Ошибка оплаты: " + (error.response?.data?.message || error.message)
      );
      console.error("Error Registrations:", error);
    }
  },
  fetchOrders: async () => {
    try {
      const response = await axios.get(`${url}/order`, {
        withCredentials: true,
      });

      set({ orders: response.data });

      if (response.status === 401) {
        // {{ edit_1 }}
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().fetchUserBasket();
      } else {
        throw new Error("No data in response");
      }
    } catch (error) {
      console.error("Error fetching basket:", error);
    }
  },
  fetchUserOrders: async () => {
    try {
      const response = await axios.get(`${url}/order/my`, {
        withCredentials: true,
      });

      set({ userOrders: response.data });

      if (response.status === 401) {
        // {{ edit_1 }}
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
      } else {
        throw new Error("No data in response");
      }
    } catch (error) {
      console.error("Error fetching basket:", error);
    }
  },
}));
export default useOrderStore;
