import { create } from "zustand";
import axios from "axios";
import { useState } from "react";

const useCategoryStore = create((set, get) => ({
  category: [],
  createCategory: async (data) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/api/v1/category",
        data,
        {
          withCredentials: true, // Если нужно отправлять куки
        }
      );
      console.log("Response:", response.data);
      // Обработка успешного ответа
      if (response.data.status === "success") {
        alert("Категория успешно сохранена!");
      } else {
        alert("Ошибка: " + response.data.message);
      }
    } catch (error) {
      console.error("Error:", error);
      alert(
        "Ошибка при сохранении категории: " +
          (error.response?.data?.message || error.message)
      );
    }
  },
  getAllCategory: async () => {
    try {
      const response = await axios.get(
        "http://localhost:8080/api/v1/category",
        {},
        {
          withCredentials: true,
        }
      );
      set({ category: response.data });
      console.log("Response:", response.data);
    } catch (error) {
      console.error("Error:", error);
      alert(
        "Ошибка при получении категорий: " +
          (error.response?.data?.message || error.message)
      );
    }
  },
}));
export default useCategoryStore;
