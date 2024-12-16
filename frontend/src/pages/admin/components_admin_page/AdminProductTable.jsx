import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  Paper,
  Typography,
} from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";
import useProductStore from "../../../store/productStore";

const AdminProductTable = () => {
  const { fetchProducts, products } = useProductStore();
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [categoryFilter, setCategoryFilter] = useState("");
  const [priceFilter, setPriceFilter] = useState("");

  useEffect(() => {
    fetchProducts();
  }, []);

  useEffect(() => {
    console.log(products.data); // Логирование данных
    if (Array.isArray(products.data)) {
      setFilteredProducts(products.data);
    }
  }, [products]);

  const handleFilterChange = () => {
    let filtered = products.data || []; // Убедитесь, что это массив

    if (categoryFilter) {
      filtered = filtered.filter(
        (product) => product.category === categoryFilter
      );
    }

    if (priceFilter) {
      filtered = filtered.filter((product) => product.price <= priceFilter);
    }

    setFilteredProducts(filtered);
  };

  const columns = [
    {
      field: "image",
      headerName: "Фото",
      width: 100,
      renderCell: (params) => (
        <img
          src={params.value}
          alt="product"
          style={{ width: 50, height: 50, borderRadius: "4px" }}
        />
      ),
    },
    { field: "name", headerName: "Название", width: 200 },
    {
      field: "price",
      headerName: "Цена",
      width: 150,
      //   valueFormatter: (params) => `${params.value} ₽`,
    },
    {
      field: "categories",
      headerName: "Категория",
      width: 200,
      renderCell: (params) => (
        <Box>
          {params.value.map((category) => category.name).join(", ")}
        </Box>
      ),
    },
  ];

  return (
    <Box sx={{ padding: 2 }}>
      <Box sx={{ display: "flex", justifyContent: "space-between", mb: 2 }}>
        <FormControl sx={{ minWidth: 120 }}>
          <InputLabel>Категория</InputLabel>
          <Select
            value={categoryFilter}
            onChange={(e) => setCategoryFilter(e.target.value)}
          >
            <MenuItem value="">
              <em>Все</em>
            </MenuItem>
            <MenuItem value="category1">Категория 1</MenuItem>
            <MenuItem value="category2">Категория 2</MenuItem>
          </Select>
        </FormControl>
        <TextField
          label="Максимальная цена"
          type="number"
          value={priceFilter}
          onChange={(e) => setPriceFilter(e.target.value)}
          sx={{ width: 200 }}
        />
        <Button variant="contained" onClick={handleFilterChange}>
          Применить фильтры
        </Button>
      </Box>

      <Paper sx={{ width: "100%", height: 400 }}>
        <DataGrid
          rows={filteredProducts}
          columns={columns}
          pageSize={5}
          rowsPerPageOptions={[5, 10, 20]}
          checkboxSelection
          disableSelectionOnClick
        />
      </Paper>
    </Box>
  );
};

export default AdminProductTable;
