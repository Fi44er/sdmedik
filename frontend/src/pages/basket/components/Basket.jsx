import {
  Box,
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  IconButton,
  Typography,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import React, { useEffect, useState } from "react";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";
import useBascketStore from "../../../store/bascketStore";
import { urlPictures } from "../../../constants/constants";

export default function Basket() {
  const {
    fetchUserBasket,
    basket,
    deleteProductThithBasket,
    editCountProductBascket,
  } = useBascketStore();

  const [currentProducts, setCurrentProducts] = useState([]);

  useEffect(() => {
    fetchUserBasket();
  }, []);

  useEffect(() => {
    if (basket.data?.items) {
      let normalizedProducts = Array.isArray(basket.data.items)
        ? basket.data.items
        : [basket.data.items];
      setCurrentProducts(normalizedProducts);
    }
  }, [basket]);

  const handleDeleteProductBasket = async (id) => {
    await deleteProductThithBasket(id);
    setCurrentProducts(currentProducts.filter((product) => product.id !== id));
    fetchUserBasket();
  };

  const handleClick = async (product_id, action) => {
    try {
      // Изменяем количество товара
      await editCountProductBascket(product_id, action === "plus" ? 1 : -1);

      // Обновляем локальное состояние
      setCurrentProducts((prevProducts) =>
        prevProducts.map((product) =>
          product.product_id === product_id
            ? {
                ...product,
                quantity:
                  action === "plus"
                    ? product.quantity + 1
                    : Math.max(product.quantity - 1, 1), // Убедитесь, что количество не меньше 1
              }
            : product
        )
      );
      await fetchUserBasket();
    } catch (error) {
      console.error("Ошибка при изменении количества товара:", error);
    }
  };

  return (
    <Box sx={{ width: { xs: "100%", md: "64.5%" }, mb: 4 }}>
      <Typography variant="h4" sx={{ mb: 2 }}>
        Корзина
      </Typography>
      <Grid
        container
        spacing={2}
        columns={{ xs: 4, sm: 4, md: 4 }}
        sx={{ mt: 2 }}
      >
        {currentProducts.length > 0 &&
          currentProducts.map((product) => (
            <Grid item key={product.product_id} xs={12} sm={6} md={4}>
              <Card
                sx={{
                  display: "flex",
                  flexDirection: { xs: "column", md: "row" },
                  padding: 2,
                  borderRadius: 2,
                  boxShadow: 3,
                  width: { xs: "100%", md: 600 },
                }}
              >
                <CardMedia
                  component="img"
                  image={`${urlPictures}/${product.image}`}
                  alt={product.title}
                  sx={{
                    width: 120,
                    height: 120,
                    objectFit: "contain",
                    borderRadius: 1,
                  }}
                />
                <Box sx={{ flexGrow: 1, paddingLeft: 2 }}>
                  <CardHeader
                    title={product.name}
                    subheader={product.brand}
                    sx={{ paddingBottom: 0 }}
                  />
                  <CardContent sx={{ paddingTop: 0 }}>
                    <Typography variant="h6">
                      Цена: {product.price} ₽
                    </Typography>
                    <Typography variant="subtitle1">
                      Количество: {product.quantity}
                    </Typography>
                    <Box
                      sx={{
                        display: "flex",
                        alignItems: "center",
                        marginTop: 1,
                      }}
                    >
                      <IconButton
                        onClick={() => handleClick(product.product_id, "minus")}
                        disabled={product.quantity <= 1}
                      >
                        -
                      </IconButton>
                      <Typography variant="body1" sx={{ mx: 1 }}>
                        {product.quantity}
                      </Typography>
                      <IconButton
                        onClick={() => handleClick(product.product_id, "plus")}
                      >
                        +
                      </IconButton>
                      <IconButton
                        onClick={() => handleDeleteProductBasket(product.id)}
                        color="error"
                        sx={{ ml: 2 }}
                      >
                        <DeleteOutlineIcon />
                      </IconButton>
                    </Box>
                  </CardContent>
                </Box>
              </Card>
            </Grid>
          ))}
      </Grid>
    </Box>
  );
}
