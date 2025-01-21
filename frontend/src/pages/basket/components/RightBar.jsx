import {
  Box,
  Button,
  Container,
  InputBase,
  Paper,
  Typography,
} from "@mui/material";
import React from "react";

export default function RightBar() {
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
              gridGap: 15,
              mt:4
            }}
          >
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
              mt: "60px",
            }}
          >
            <Typography>
              Итог:<br></br>5 товаров
            </Typography>
            <Typography>125000 ₽</Typography>
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignContent: "center",
              mt: 5,
            }}
          >
            <Button variant="contained" sx={{
              background: "#00B3A4",
            }}>Перейти к оплате</Button>
          </Box>
        </Container>
      </Paper>
    </Box>
  );
}
