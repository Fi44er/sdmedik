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
} from "@mui/material";
import useProductStore from "../../../../store/productStore";

const AdminProductTable = () => {
  const { fetchProducts, products, deleteProduct } = useProductStore();
  const [filteredProducts, setFilteredProducts] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 20;

  useEffect(() => {
    fetchProducts();
  }, []);

  useEffect(() => {
    if (Array.isArray(products.data)) {
      setFilteredProducts(products.data);
    }
  }, [products]);

  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = filteredProducts.slice(
    indexOfFirstItem,
    indexOfLastItem
  );

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
  };

  const handleDeleteProduct = async (id) => {
    await deleteProduct(id);
    fetchProducts();
  };

  return (
    <Box sx={{ padding: 2 }}>
      <Paper sx={{ width: "100%" }}>
        {/* Таблица для больших экранов */}
        <TableContainer
          sx={{ overflowX: "auto", display: { xs: "none", sm: "block" } }}
        >
          <Table>
            <TableHead>
              <TableRow>
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
              {currentItems.map((product) => (
                <TableRow key={product.id}>
                  <TableCell>
                    <Box sx={{ display: "flex", gap: 1 }}>
                      {product.images.map((image) => (
                        <img
                          key={image.id}
                          src={`http://127.0.0.1:8080/api/v1/image/${image.name}`}
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
                          window.location.href = `/update_product/${product.id}`;
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

        {/* Карточки для мобильных устройств */}
        <Box sx={{ display: { xs: "block", sm: "none" } }}>
          {currentItems.map((product) => (
            <Paper key={product.id} sx={{ mb: 2, p: 2 }}>
              <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
                <Box sx={{ display: "flex", gap: 1 }}>
                  {product.images.map((image) => (
                    <img
                      key={image.id}
                      src={`http://127.0.0.1:8080/api/v1/image/${image.name}`}
                      alt="product"
                      style={{ width: 50, height: 50, borderRadius: "4px" }}
                    />
                  ))}
                </Box>
                <Box>Название: {product.name}</Box>
                <Box>Цена: {product.price}</Box>
                <Box>
                  Категория:{" "}
                  {product.categories
                    .map((category) => category.name)
                    .join(", ")}
                </Box>
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
                      window.location.href = `/update_product/${product.id}`;
                    }}
                  >
                    редактировать
                  </Button>
                </Box>
              </Box>
            </Paper>
          ))}
        </Box>

        {/* Пагинация */}
        <Box sx={{ display: "flex", justifyContent: "center", mt: 2, mb: 2 }}>
          <Pagination
            count={Math.ceil(filteredProducts.length / itemsPerPage)}
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
