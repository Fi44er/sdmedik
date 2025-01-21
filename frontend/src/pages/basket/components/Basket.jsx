import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  IconButton,
  Typography,
} from "@mui/material";
import Grid from "@mui/material/Grid2";
import React, { useEffect } from "react";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";
import useBascketStore from "../../../store/bascketStore";

export default function Basket() {
  const {
    fetchUserBasket,
    basket,
    products,
    fetchProductsByIds,
    deleteProductThithBasket,
  } = useBascketStore();

  const fetchUserBasketCards = () => {
    if (
      basket.data &&
      Array.isArray(basket.data.items) &&
      basket.data.items.length > 0
    ) {
      const items = basket.data.items.map((item) => item.product_id);
      console.log(items);

      fetchProductsByIds(items);
    } else {
      console.log("Data not available yet.");
    }
  };

  useEffect(() => {
    fetchUserBasket();
  }, []);

  useEffect(() => {
    if (
      basket.data &&
      Array.isArray(basket.data.items) &&
      basket.data.items.length > 0
    ) {
      fetchUserBasketCards();
      console.log(products);
    }
  }, [basket]);

  const hendleDeleteProductBascasket = () => {
        
    deleteProductThithBasket(id);
  };

  return (
    <Box
      sx={{
        width: { xs: "100%", md: "64.5%" },
        mb: 4,
      }}
    >
      <Typography variant="h4">Корзина</Typography>
      <Grid
        container
        spacing={{ xs: 2, md: 5 }}
        columns={{ xs: 4, sm: 4, md: 4 }}
        sx={{ mt: 4 }}
      >
        {products.length > 0 &&
          products.map((product) => (
            <Grid item key={product.id} xs={1} sm={1} md={1}>
              <Card
                sx={{
                  width: { xs: "100%", lg: "100%" },
                  backgroundColor: "#F5FCFF",
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "space-between",
                  padding: { xs: 0, md: 3 },
                  borderRadius: 4,
                }}
              >
                {/* Изображение продукта */}
                <Box
                  sx={{
                    display: "flex",
                    flexDirection: "column",
                    marginRight: 2,
                  }}
                >
                  <CardHeader
                    title={product.data.name}
                    subheader={product.data.brand}
                  />
                  <CardMedia
                    component="img"
                    image={`http://127.0.0.1:8080/api/v1/image/${product.data.images[0].name}`}
                    alt={product.data.title}
                    sx={{ width: "150px", height: "auto" }}
                  />
                </Box>

                {/* Описание продукта */}
                <CardContent
                  sx={{
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "space-between",
                    textAlign: "left",
                    width: "50%",
                  }}
                >
                  <Box sx={{ mt: 2 }}>
                    <Typography variant="h6">
                      Цена: {product.data.price} ₽
                    </Typography>
                    <Typography variant="subtitle1">
                      Артикул: {product.data.article}
                    </Typography>
                  </Box>
                </CardContent>

                {/* Кнопка удаления */}
                <IconButton
                  aria-label="удалить товар"
                  color="error"
                  size="large"
                  onClick={() =>
                    hendleDeleteProductBascasket(product.data.product_id)
                  }
                >
                  <DeleteOutlineIcon fontSize="inherit" />
                </IconButton>
              </Card>
            </Grid>
          ))}
      </Grid>
    </Box>
  );
}
