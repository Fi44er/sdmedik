import {
  Box,
  Button,
  Card,
  CardContent,
  CardMedia,
  IconButton,
  Pagination,
  Typography,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useProductStore from "../../../store/productStore";
import SidebarFilter from "./SidebarFilter";
import useBascketStore from "../../../store/bascketStore";
import { urlPictures } from "../../../constants/constants";

export default function CatalogDynamicPage() {
  const { id } = useParams();
  const { fetchProducts, products } = useProductStore();
  const { addProductThisBascket } = useBascketStore();
  const [currentPage, setCurrentPage] = useState(1);
  const [filters, setFilters] = useState(null); // Состояние для хранения фильтров
  const [currentProducts, setCurrentProducts] = useState([]); // Переменная для хранения текущих продуктов
  const [ProductsPerPage] = useState(20); // Количество продуктов на странице
  const [quantity, setQuantity] = useState(1);

  const category_id = id;

  useEffect(() => {
    const offset = (currentPage - 1) * ProductsPerPage; // Рассчитываем offset
    fetchProducts(category_id, filters, offset, ProductsPerPage); // Передаем offset и limit в fetchProducts
  }, [category_id, fetchProducts, filters, currentPage]); // Добавляем currentPage в зависимости

  useEffect(() => {
    if (products?.data) {
      let normalizedProducts = [];
      if (!Array.isArray(products.data)) {
        normalizedProducts = [products.data]; // Приводим объект к массиву
      } else {
        normalizedProducts = products.data;
      }
      setCurrentProducts(normalizedProducts);
    }
  }, [products]);

  const handleChangePage = (event, value) => {
    setCurrentPage(value);
  };

  const hendleAddProductThithBascket = async (id) => {
    setQuantity(quantity);
    const product_id = id;
    console.log(id, quantity);

    await addProductThisBascket(product_id, quantity);
  };

  return (
    <Box sx={{ mt: 1, mb: 5 }}>
      <Box sx={{ mb: 5 }}>
        <SidebarFilter setFilters={setFilters} />
      </Box>
      <Grid
        container
        spacing={{ xs: 2, md: 3 }}
        columns={{ xs: 4, sm: 4, md: 4 }}
        sx={{ height: 800, overflowX: "auto", pt: 2, pb: 2 }}
      >
        {currentProducts.length > 0 ? (
          currentProducts.map((e) => (
            <Grid item={"true"} key={e.id} xs={6} sm={4} md={3}>
              <Card
                sx={{
                  width: { xs: "100%", lg: "261px" },
                  height: "520px",
                  background: "#F5FCFF",
                  boxShadow: "0 4px 20px rgba(0, 0, 0, 0.1)",
                  borderRadius: "8px",
                  transition: "transform 0.2s, box-shadow 0.2s",
                  "&:hover": {
                    transform: "scale(1.05)",
                    boxShadow: "0 8px 30px rgba(0, 0, 0, 0.2)",
                  },
                  display: "flex",
                  flexDirection: "column",
                }}
              >
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                    height: "300px",
                    borderBottom: "1px solid #E0E0E0",
                  }}
                >
                  <CardMedia
                    component="img"
                    image={`${urlPictures}/${e.images[0].name}`}
                    alt={e.name}
                    sx={{
                      width: "100%",
                      height: "300px",
                      objectFit: "cover",
                    }}
                  />
                </Box>

                <CardContent sx={{ flexGrow: 1 }}>
                  <Typography
                    sx={{
                      fontSize: "1.2rem",
                      fontWeight: "bold",
                      mb: 1,
                      width: "235px",
                      overflow: "hidden",
                      textOverflow: "ellipsis",
                      whiteSpace: "nowrap",
                    }}
                  >
                    {e.name}
                  </Typography>
                  <Typography
                    variant="body2"
                    color="text.secondary"
                    sx={{ mb: 1 }}
                  >
                    Артикул: {e.article}
                  </Typography>
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "space-between",
                      alignItems: "center",
                      mt: 1,
                    }}
                  >
                    <Typography variant="h6" sx={{ color: "#00B3A4" }}>
                      {e.price} ₽
                    </Typography>
                    {e.oldPrice && (
                      <Typography
                        variant="body2"
                        sx={{
                          color: "text.secondary",
                          textDecoration: "line-through",
                        }}
                      >
                        {e.oldPrice} ₽
                      </Typography>
                    )}
                  </Box>
                </CardContent>

                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    p: 2,
                    borderTop: "1px solid #E0E0E0",
                  }}
                >
                  <Button
                    sx={{
                      width: "100%",
                      height: "40px",
                      border: `2px solid #00B3A4`,
                      borderRadius: "20px",
                      color: "#00B3A4",
                    }}
                    variant="outlined"
                    onClick={() => {
                      window.location.href = `/product/${e.id}`;
                    }}
                  >
                    Подробнее
                  </Button>
                  <IconButton
                    onClick={() => {
                      hendleAddProductThithBascket(e.id);
                    }}
                  >
                    <img
                      style={{ width: "50px", height: "50px" }}
                      src="/public/basket_cards.png"
                      alt="Добавить в корзину"
                    />
                  </IconButton>
                </Box>
              </Card>
            </Grid>
          ))
        ) : (
          <Typography>Нет данных для отображения</Typography>
        )}
      </Grid>
      {currentProducts.length > 0 && (
        <Pagination
          count={Math.ceil((products.count || 0) / ProductsPerPage)} // Обновите общее количество страниц
          page={currentPage}
          onChange={handleChangePage}
          sx={{
            mt: 4,
            mb: 4,
            display: "flex",
            justifyContent: "center",
          }}
        />
      )}
    </Box>
  );
}
