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
  Table,
  TableContainer,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
} from "@mui/material";
import useProductStore from "../../../../store/productStore";

const AdminProductTable = () => {
  const { fetchProducts, products, deleteProduct } = useProductStore();
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

  return (
    <Box sx={{ padding: 2 }}>
      <Paper sx={{ width: "100%" }}>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Фото</TableCell>
                <TableCell>Название</TableCell>
                <TableCell>Цена</TableCell>
                <TableCell>Категория</TableCell>
                <TableCell>Управление</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {filteredProducts.map((product) => (
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
                  <TableCell>
                    {product.categories.map((category) => category.name).join(", ")}
                  </TableCell>
                  <TableCell>
                    <Box>
                      <Button
                        variant="contained"
                        color="error"
                        onClick={() => deleteProduct(product.id)}
                      >
                        удалить
                      </Button>
                      <Button
                        variant="contained"
                        color="info"
                        onClick={(e) => {
                          e.preventDefault();
                          window.location.href = "/";
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
      </Paper>
    </Box>
  );
};

export default AdminProductTable;