import { create } from "zustand";
import axios from "axios";

const useBascketStore = create((set, get) => ({
  product: {
    article: "",
    category_ids: [],
    characteristic_values: [],
    description: "",
    name: "",
  },
  products: [],
  basket: {},
  addProductThisBascket: async (product_id, quantity) => {
    try {
      const response = await axios.post(
        "http://localhost:8080/api/v1/basket",
        { product_id, quantity },
        {
          withCredentials: true,
        }
      );
      // Обработка успешного ответа
      console.log("Продукт создан:", response.data);
    } catch (error) {
      // Обработка ошибки
      console.error("Ошибка при создании продукта:", error);
      if (error.response.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        // await get().refreshToken();
        // await get().addProductThisBascket(product_id,
        //     quantity,); // Повторяем запрос
      } else {
        alert(
          "Ошибка при сохранении продукта: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },
  fetchUserBasket: async () => {
    try {
      const response = await axios.get(`http://localhost:8080/api/v1/basket`, {
        withCredentials: true,
      });

      if (response && response.data) {
        set({ basket: response.data });
        get().fetchProductsByIds(response.data.items);
      } else {
        throw new Error("No data in response");
      }
    } catch (error) {
      console.error("Error fetching basket:", error);
    }
  },

  fetchProductsByIds: async (items) => {
    if (Array.isArray(items) && items.length > 0) {
      let products = [];
      for (let item of items) {
        try {
          const response = await axios.get(
            "http://localhost:8080/api/v1/product",
            {
              params: { id: item.toString() },
            }
          );

          products.push(response.data);
        } catch (error) {
          console.error(`Error fetching product with ID ${item}:`, error);
        }
      }
      set({ products });
    } else {
      console.warn("Items is not an array or is empty.");
    }
  },
  deleteProductThithBasket: async (id) => {
    try {
      const response = await axios.delete(
        `http://localhost:8080/api/v1/basket/${id}`,
        {},
        {
          withCredentials: true,
        }
      );
    } catch (error) {
      console.error("Error deleting basket:", error);
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

export default useBascketStore;
