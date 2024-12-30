import { create } from "zustand";
import axios from "axios";

const useCategoryStore = create((set, get) => ({
  category: [],
  categoryId: {},
  fetchCategory: async () => {
    try {
      const response = await axios.get(`http://localhost:8080/api/v1/category`);
      set({ category: response.data });
    } catch (error) {
      console.error("Error fetching category:", error);
    }
  },
  fetchCategoryId: async (id) => {
    try {
      const response = await axios.get(
        `http://localhost:8080/api/v1/category/${id}`
      );
      set({ categoryId: response.data });
    } catch (error) {
      console.error("Error fetching category:", error);
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

  createCategory: async (formData) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/api/v1/category",
        formData,
        {
          withCredentials: true,
          headers: {
            "Content-Type": "multipart/form-data", // Убедитесь, что заголовок установлен правильно
          },
        }
      );
      console.log("Response:", response.data);
      // Обработка успешного ответа
      if (response.data.status === "success") {
        alert("Категория успешно сохранена!");
        get().fetchCategory(); // Обновляем список категорий после создания новой
      } else {
        alert("Ошибка: " + response.data.message);
      }
    } catch (error) {
      console.error("Error:", error);
      if (error.response.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().createCategory(data); // Повторяем запрос
      } else {
        alert(
          "Ошибка при сохранении категории: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },
}));

export default useCategoryStore;
