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
  products: [],
  fetchProducts: async () => {
    try {
      const response = await axios.get(`http://localhost:8080/api/v1/product`);
      set({ products: response.data });
    } catch (error) {
      console.error("Error fetching category:", error);
    }
  },
}));

export default useProductStore;
