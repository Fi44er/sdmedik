import { create } from "zustand";
import axios from "axios";

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
      const response = await axios.post(
        "http://localhost:8080/api/v1/product",
        formData,
        {
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data", // Убедитесь, что заголовок установлен правильно
          },
        }
      );
      // Обработка успешного ответа
      console.log("Продукт создан:", response.data);
    } catch (error) {
      // Обработка ошибки
      console.error("Ошибка при создании продукта:", error);
      if (error.response.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().createProduct(formData); // Повторяем запрос
      } else {
        alert(
          "Ошибка при сохранении продукта: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },

  updateProduct: async (id, formData) => {
    try {
      const response = await axios.put(
        `http://localhost:8080/api/v1/product/${id}`,
        formData,
        {
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data", // Убедитесь, что заголовок установлен правильно
          },
        }
      );
      // Обработка успешного ответа
      console.log("Продукт обновлен:", response.data);
    } catch (error) {
      // Обработка ошибки
      console.error("Ошибка при обновлении продукта:", error);
      if (error.response?.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().updateProduct(id, formData); // Повторяем запрос
      } else {
        alert(
          "Ошибка при обновлении продукта: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },

  deleteProduct: async (id) => {
    try {
      const response = await axios.delete(
        `http://localhost:8080/api/v1/product/${id}`
      );
    } catch (error) {
      console.error("Error deleting product:", error);
    }
  },

  products: [],
  fetchProducts: async (category_id, jsonData) => {
    try {
      const response = await axios.get(`http://localhost:8080/api/v1/product`, {
        params: {
          category_id: category_id,
          filters: jsonData,
        },
      });
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },
  fetchFiltersProducts: async (jsonData) => {
    try {
      const response = await axios.get(`http://localhost:8080/api/v1/product`, {
        params: {
          filters: jsonData,
        },
      });
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },
  fetchProductById: async (id, name) => {
    try {
      const response = await axios.get(`http://localhost:8080/api/v1/product`, {
        params: { id: id },
      });
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching product:", error);
    }
  },

  refreshToken: async () => {
    try {
      const response = await axios.post(
        "http://localhost:8080/api/v1/auth/refresh",
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
