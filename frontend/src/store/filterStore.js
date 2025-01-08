import { create } from "zustand";
import axios from "axios";
import { useState } from "react";

const useFilterStore = create((set, get) => ({
  filters: [],
  fetchFilter: async (category_id) => {
    try {
      const response = await axios.get(
        `http://localhost:8080/api/v1/product/filter/${category_id}`
      );
      set({ filters: response.data });
    } catch (error) {
      console.error("Error fetching filters:", error);
    }
  },
}));
export default useFilterStore;
