import { create } from "zustand";
import axios from "axios";
import { url } from "../constants/constants";
import { toast } from "react-toastify";

const useProductStore = create((set, get) => ({
  product: {
    article: "",
    category_ids: [],
    characteristic_values: [],
    description: "",
    name: "",
  },
  createProduct: async (formData) => {
    try {
      const response = await axios.post(`${url}/product`, formData, {
        withCredentials: true,
        headers: {
          "Content-Type": "multipart/form-data", // Убедитесь, что заголовок установлен правильно
        },
      });
      // Обработка успешного ответа
      toast.success("Продукт успешно создан");
    } catch (error) {
      // Обработка ошибки
      console.error("Ошибка при создании продукта:", error);  
      if (error.response.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().createProduct(formData); // Повторяем запрос
      } else {
        toast.error(
          "Ошибка при сохранении продукта: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },

  updateProduct: async (id, formData) => {
    try {
      const response = await axios.put(`${url}/product/${id}`, formData, {
        withCredentials: true,
        headers: {
          "Content-Type": "multipart/form-data", // Убедитесь, что заголовок установлен правильно
        },
      });
      // Обработка успешного ответа
      console.log("Продукт обновлен:", response.data);
      toast.success("Продукт успешно обновлен");
    } catch (error) {
      // Обработка ошибки
      console.error("Ошибка при обновлении продукта:", error);
      if (error.response?.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().updateProduct(id, formData); // Повторяем запрос
      } else {
        toast.error(
          "Ошибка при обновлении продукта: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },

  deleteProduct: async (id) => {
    try {
      const response = await axios.delete(`${url}/product/${id}`);
    } catch (error) {
      console.error("Error deleting product:", error);
    }
  },

  products: [],
  fetchProducts: async (category_id, jsonData, offset, limit) => {
    try {
      const response = await axios.get(`${url}/product`, {
        params: {
          category_id: category_id,
          filters: jsonData,
          offset: offset, // Добавляем offset
          limit: limit, // Добавляем limit
        },
      });
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },
  fetchFiltersProducts: async (jsonData) => {
    try {
      const response = await axios.get(`${url}/product`, {
        params: {
          filters: jsonData,
        },
      });
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },
  fetchProductById: async (id, iso) => {
    try {
      const response = await axios.get(`${url}/product`, {
        params: { id: id, iso: iso },
      });
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },

  fetchTopList: async (category_id, jsonData) => {
    try {
      const response = await axios.get(`${url}/product/top/${4}`);
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },

  refreshToken: async () => {
    try {
      const response = await axios.post(
        `${url}/auth/refresh`,
        {},
        {
          withCredentials: true,
        }
      );
    } catch (error) {
      console.error("Error:", error);
    }
  },
}));

export default useProductStore;
