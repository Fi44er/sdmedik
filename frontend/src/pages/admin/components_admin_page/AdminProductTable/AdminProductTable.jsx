import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  Paper,
  Table,
  TableContainer,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Pagination,
  Typography,
  Select,
  MenuItem,
} from "@mui/material";
import useProductStore from "../../../../store/productStore";
import { urlPictures } from "../../../../constants/constants";
import useCategoryStore from "../../../../store/categoryStore";

const AdminProductTable = () => {
  const { fetchProducts, products, deleteProduct } = useProductStore();
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 20;
  const [selectedCategory, setSelectedCategory] = useState(""); // Состояние для выбранной категории
  const { fetchCategory, category } = useCategoryStore();

  useEffect(() => {
    const offset = (currentPage - 1) * itemsPerPage; // Рассчитываем offset
    fetchProducts(selectedCategory, null, offset, itemsPerPage); // Передаем offset и limit в fetchProducts
  }, [currentPage, selectedCategory]); // Добавляем currentPage и selectedCategory в зависимости

  useEffect(() => {
    if (Array.isArray(products.data)) {
      setFilteredProducts(products.data);
    }
  }, [products]);

  const handleCategoryChange = (event) => {
    const category = event.target.value;
    setSelectedCategory(category);
    setCurrentPage(1); // Сброс страницы при изменении категории
  };

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
  };

  const handleDeleteProduct = async (id) => {
    await deleteProduct(id);
    // После удаления, обновляем список товаров
    const offset = (currentPage - 1) * itemsPerPage; // Рассчитываем новый offset
    fetchProducts(selectedCategory, null, offset, itemsPerPage);
  };

  useEffect(() => {
    fetchCategory();
  }, []);

  return (
    <Box sx={{ padding: 2 }}>
      <Typography sx={{ fontSize: "30px", mb: 2, mt: 2 }}>
        Таблица с Продуктами
      </Typography>

      {/* Фильтрация по категориям */}
      <Select
        value={selectedCategory}
        onChange={handleCategoryChange}
        displayEmpty
        sx={{ mb: 2, minWidth: 200 }}
      >
        <MenuItem value="">
          <em>Все категории</em>
        </MenuItem>
        {category.data &&
          category.data.map((cat) => (
            <MenuItem key={cat.id} value={cat.id}>
              {cat.name}
            </MenuItem>
          ))}
      </Select>

      <Paper sx={{ width: "100%" }}>
        <TableContainer sx={{ overflowX: "auto", height: "600px" }}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Id товара</TableCell>
                <TableCell>Фото</TableCell>
                <TableCell>Название</TableCell>
                <TableCell>Цена</TableCell>
                <TableCell sx={{ display: { xs: "none", sm: "table-cell" } }}>
                  Категория
                </TableCell>
                <TableCell>Управление</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {filteredProducts.map((product) => (
                <TableRow key={product.id}>
                  <TableCell>
                    <Box sx={{ display: "flex", gap: 1 }}>{product.id}</Box>
                  </TableCell>
                  <TableCell>
                    <Box sx={{ display: "flex", gap: 1 }}>
                      {product.images.map((image) => (
                        <img
                          key={image.id}
                          src={`${urlPictures}/${image.name}`}
                          alt="product"
                          style={{ width: 50, height: 50, borderRadius: "4px" }}
                        />
                      ))}
                    </Box>
                  </TableCell>
                  <TableCell>{product.name}</TableCell>
                  <TableCell>{product.price}</TableCell>
                  <TableCell sx={{ display: { xs: "none", sm: "table-cell" } }}>
                    {product.categories
                      .map((category) => category.name)
                      .join(", ")}
                  </TableCell>
                  <TableCell>
                    <Box sx={{ display: "flex", gap: 1 }}>
                      <Button
                        variant="contained"
                        color="error"
                        onClick={() => handleDeleteProduct(product.id)}
                      >
                        удалить
                      </Button>
                      <Button
                        variant="contained"
                        color="info"
                        onClick={(e) => {
                          e.preventDefault();
                          window.location.href = `/admin/update_product/${product.id}`;
                        }}
                      >
                        редактировать
                      </Button>
                    </Box>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>

        {/* Пагинация */}
        <Box sx={{ display: "flex", justifyContent: "center", mt: 2, mb: 2 }}>
          <Pagination
            count={Math.ceil(products.total / itemsPerPage)} // Обновите общее количество страниц
            page={currentPage}
            onChange={handlePageChange}
            color="primary"
            sx={{ mt: 2, mb: 2 }}
          />
        </Box>
      </Paper>
    </Box>
  );
};

export default AdminProductTable;
