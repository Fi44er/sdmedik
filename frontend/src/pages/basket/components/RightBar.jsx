import {
  Box,
  Button,
  Container,
  InputBase,
  Paper,
  Typography,
} from "@mui/material";
import React, { useEffect } from "react";
import useBascketStore from "../../../store/bascketStore";

export default function RightBar() {
  const {
    fetchUserBasket,
    basket,
    products,
    fetchProductsByIds,
    deleteProductThithBasket,
  } = useBascketStore();

  return (
    <Box
      sx={{
        mt: { xs: 0, md: "75px" },
        width: { xs: "100%", md: "29.5%" },
      }}
    >
      <Paper sx={{}}>
        <Container
          sx={{
            pt: 5,
            pb: 5,
          }}
        >
          <Typography variant="h5" sx={{ textAlign: "center" }}>
            Оформление заказа
          </Typography>
          <Box
            sx={{
              display: "flex",
              flexDirection: "column",
              mt: "20px",
            }}
          >
            <Box>
              <Typography>
                Итого товаров в корзине : {basket.data && basket.data.quantity}
              </Typography>
            </Box>
            <Box>
              <Typography>
                Общая сумма заказа : {basket.data && basket.data.total_price} ₽
              </Typography>
            </Box>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignContent: "center",
              mt: 5,
            }}
          >
            <Button
              variant="contained"
              sx={{
                background: "#00B3A4",
              }}
            >
              Перейти к оплате
            </Button>
          </Box>
        </Container>
      </Paper>
    </Box>
  );
}
