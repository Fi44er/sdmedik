import { Box, Button, Container, Paper, Typography } from "@mui/material";
import React from "react";
import useBascketStore from "../../../store/bascketStore";

export default function RightBar() {
  const { basket } = useBascketStore();

  return (
    <Box
      sx={{
        mt: { xs: 0, md: "75px" },
        width: { xs: "100%", md: "29.5%" },
      }}
    >
      <Paper sx={{ padding: 3, borderRadius: 2, boxShadow: 3 }}>
        <Container>
          <Typography
            variant="h5"
            sx={{ textAlign: "center", fontWeight: "bold" }}
          >
            Оформление заказа
          </Typography>
          <Box sx={{ mt: 2 }}>
            <Typography>
              Итого товаров в корзине: {basket.data && basket.data.quantity}
            </Typography>
            <Typography variant="h6" sx={{ color: "#00B3A4", mt: 1 }}>
              Общая сумма заказа: {basket.data && basket.data.total_price} ₽
            </Typography>
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
                "&:hover": {
                  background: "#009B8A",
                },
              }}
              onClick={(e) => {
                e.preventDefault();
                window.location.href = "/paymants";
              }}
            >
              Продолжить
            </Button>
          </Box>
        </Container>
      </Paper>
    </Box>
  );
}
