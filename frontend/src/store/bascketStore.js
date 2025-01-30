import { create } from "zustand";
import axios from "axios";
import { toast } from "react-toastify";
import { url } from "../constants/constants";

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
        `${url}/basket`,
        { product_id, quantity },
        {
          withCredentials: true,
        }
      );
      // Обработка успешного ответа
      toast.success("Продукт добавлен в корзину");
      console.log("Продукт добавлен в корзину:", response.data);
    } catch (error) {
      // Обработка ошибки
      toast.error("Ошибка при создании продукта:", error);
    }
  },
  editCountProductBascket: async (product_id, quantity) => {
    try {
      const response = await axios.post(
        `${url}/basket`,
        { product_id, quantity },
        {
          withCredentials: true,
        }
      );
      console.log("Продукт добавлен в корзину:", response.data);
    } catch (error) {
      // Обработка ошибки
      if (error.response.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().addProductThisBascket(product_id, quantity); // Повторяем запрос
      } else {
        toast.error(
          "Ошибка при сохранении продукта: " +
            (error.response?.data?.message || error.message)
        );
      }
    }
  },
  fetchUserBasket: async () => {
    try {
      const response = await axios.get(`${url}/basket`, {
        withCredentials: true,
      });

      set({ basket: response.data });

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

  //   fetchProductsByIds: async (items) => {
  //     if (Array.isArray(items) && items.length > 0) {
  //       let products = [];
  //       for (let item of items) {
  //         try {
  //           const response = await axios.get(
  //             "http://localhost:8080/api/v1/product",
  //             {
  //               params: { id: item.toString() },
  //             }
  //           );

  //           products.push(response.data);
  //         } catch (error) {
  //           console.error(`Error fetching product with ID ${item}:`, error);
  //         }
  //       }
  //       set({ products });
  //     } else {
  //       console.warn("Items is not an array or is empty.");
  //     }
  //   },
  deleteProductThithBasket: async (id) => {
    try {
      const response = await axios.delete(`${url}/basket/${id}`, {
        withCredentials: true,
      });
      toast.success("Продукт удален из корзины");
      if (error.response.status === 401) {
        // Если статус 401, обновляем токены и повторяем запрос
        await get().refreshToken();
        await get().deleteProductThithBasket(id); // Повторяем запрос
      } else {
        throw new Error("No data in response");
      }
    } catch (error) {
      console.error("Ошибка при удалении продукта:");
      // console.error("Error deleting basket:", error);
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

export default useBascketStore;
