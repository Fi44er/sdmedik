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

export default function CatalogDynamicPage() {
  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Grid
        container
        spacing={{ xs: 2, md: 3 }}
        columns={{ xs: 4, sm: 4, md: 4 }}
      >
        {Product.map((e) => (
          <Grid item key={e.id} xs={1} sm={1} md={1}>
            <Card
              sx={{
                width: { xs: "100%", lg: "261px" },
                background: "#F5FCFF",
              }}
            >
              <Box
                sx={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <CardMedia
                  component="img"
                  image={e.image}
                  alt={"wheelchair"}
                  sx={{
                    width: "200px",
                    height: { xs: "200px", sm: "200px", md: "200px" },
                    objectFit: "cover",
                  }}
                />
              </Box>

              <CardContent>
                <CardHeader title={e.title} />
                <Typography variant="body2" color="text.secondary">
                  {e.country}
                </Typography>
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                  }}
                >
                  <Typography variant="h6" sx={{ color: "#004B8D" }}>
                    {e.price}
                  </Typography>
                  <Typography
                    variant="body2"
                    sx={{
                      color: "text.secondary",
                      textDecoration: "line-through",
                    }}
                  >
                    {e.price} 
                  </Typography>
                </Box>

                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                    mt: "20px",
                  }}
                >
                  <Button
                    sx={{
                      width: "157px",
                      height: "50px",
                      border: "2px solid #1E90FF",
                      borderRadius: "20px",
                    }}
                    variant="outlined"
                  >
                    В 1 клик
                  </Button>
                  <IconButton>
                    <img
                      style={{ width: "50px", height: "50px" }}
                      src="/public/basket.png"
                      alt=""
                    />
                  </IconButton>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
}
