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
import React from "react";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";

const Product = [
  {
    id: 1,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 2,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 3,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 4,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 5,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 6,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 7,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 8,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 9,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 10,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
];

export default function   Basket() {
  return (
    <Box
      sx={{
        width: { xs: "100%", md: "64.5%" },
        mb:4
      }}
    >
      <Typography variant="h4">Корзина</Typography>
      <Grid
        container
        spacing={{ xs: 2, md: 5 }}
        columns={{ xs: 4, sm: 4, md: 4 }}
        sx={{ mt: 4 }}
      >
        {Product.map((e) => {
          return (
            <Grid item key={e.id} xs={1} sm={1} md={1}>
              <Card
                sx={{
                  width: { xs: "100%", lg: "100%" },
                  background: "#F5FCFF",
                  display: "flex",
                  p: { xs: 0, md: 3 },
                }}
              >
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "center",
                    pt: 3,
                    pl: 3,
                  }}
                >
                  <CardMedia
                    component="img"
                    image={e.image}
                    alt={"wheelchair"}
                    sx={{
                      width: { xs: "125px", md: "200px" },
                      height: { xs: "125px", sm: "200px", md: "200px" },
                      objectFit: "cover",
                    }}
                  />
                </Box>

                <CardContent
                  sx={{
                    display: "flex",
                    flexDirection: { xs: "column", md: "unset" },
                  }}
                >
                  <Box>
                    <CardHeader title={e.title} sx={{ p: { xs: 0, md: 2 } }} />
                    <Typography
                      variant="body2"
                      color="text.secondary"
                      sx={{ ml: { xs: 0, md: 2 } }}
                    >
                      {e.country}
                    </Typography>
                  </Box>
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "space-between",
                      alignItems: "center",
                      flexDirection: { xs: "unset", md: "column-reverse" },
                      mt: "20px",
                    }}
                  >
                    <Typography variant="h6" sx={{ color: "black" }}>
                      {e.price}
                    </Typography>
                    <IconButton>
                      <DeleteOutlineIcon color="error" fontSize="large" />
                    </IconButton>
                  </Box>
                </CardContent>
              </Card>
            </Grid>
          );
        })}
      </Grid>
    </Box>
  );
}
