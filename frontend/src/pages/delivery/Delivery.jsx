import React from "react";
import { Helmet } from "react-helmet";
import {
  Box,
  Container,
  List,
  ListItem,
  ListItemText,
  Paper,
  Typography,
} from "@mui/material";

export default function Delivery() {
  return (
    <Box>
      <Helmet>
        <title>Доставка - СД-МЕД</title>
        <meta
          name="description"
          content="Узнайте о доставке по Оренбургу и другим городам России. Бесплатная доставка по Оренбургу и информация о стоимости доставки в другие регионы."
        />
        <meta
          name="keywords"
          content="доставка, Оренбург, бесплатная доставка, доставка по России, курьерская доставка"
        />
      </Helmet>
      <Container>
        <Box sx={{ display: "flex", justifyContent: "center" }}>
          <Typography component="h1" variant="h2">
            Доставка
          </Typography>
        </Box>
        <Box sx={{ width: "100%" }}>
          <img style={{ width: "100%" }} src="/Line 1.png" alt="line" />
        </Box>

        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gridGap: "40px",
            mt: 5,
            mb: 5,
            flexDirection: { xs: "column", sm: "column", md: "unset" },
          }}
        >
          <Box sx={{ width: "100%" }}>
            <img style={{ width: "100%" }} src="/delivery.png" alt="" />
          </Box>
          <Box>
            <Paper
              elevation={3}
              style={{ padding: "16px", marginBottom: "16px" }}
            >
              <Typography variant="h5">
                Доставка по Оренбургу от 1 дня
              </Typography>
              <List>
                <ListItem>
                  <Typography variant="h6">
                    По Оренбургу – доставка бесплатная.
                  </Typography>
                </ListItem>
              </List>
              <Typography variant="h5">
                Доставка в другие города России
              </Typography>
              <List>
                <ListItem>
                  <Typography variant="h6">
                    Стоимость заказа включает в себя стоимость заказанных
                    товаров и стоимость почтовой/курьерской доставки до региона
                    получателя.
                  </Typography>
                </ListItem>
                <ListItem>
                  <Typography variant="h6">
                    Стоимость доставки зависит от региона получателя (при
                    доставке компанией СДЭК на стоимость доставки влияет также
                    общий вес заказа). Стоимость доставки видно на странице
                    оформления заказа после выбора региона проживания.
                  </Typography>
                </ListItem>
              </List>
            </Paper>
          </Box>
        </Box>
      </Container>
    </Box>
  );
}
