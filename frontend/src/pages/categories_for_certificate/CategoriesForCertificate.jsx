import {
  Box,
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  Container,
  Typography,
} from "@mui/material";

import Grid from "@mui/material/Grid2";

import React, { useEffect } from "react";
import useCategoryStore from "../../store/categoryStore";
import { urlPictures } from "../../constants/constants";
import { Helmet } from "react-helmet";

export default function CategoriesForCertificate() {
  const { fetchCategory, category } = useCategoryStore();

  useEffect(() => {
    fetchCategory();
    console.log(category.data);
  }, []);

  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Helmet>
        <title>Категории товаров | Sdmedik</title>
        <meta
          name="description"
          content="Ознакомьтесь с нашими категориями товаров. Мы предлагаем широкий ассортимент продукции для ваших нужд."
        />
        <meta
          name="keywords"
          content="категории, товары, ассортимент, продукция"
        />
      </Helmet>
      <Container>
        <Grid
          container
          spacing={{ xs: 2, md: 2 }}
          columns={{ xs: 4, sm: 4, md: 4 }}
          sx={{ display: "flex", justifyContent: "center" }}
        >
          {Array.isArray(category.data) && category.data.length > 0 ? (
            category.data.map((item) => (
              <Grid item={"true"} xs={1} sm={1} md={1} key={item.id}>
                <Card
                  sx={{
                    width: { xs: "340px", md: "276px" },
                    background: "#fff",
                    borderRadius: "20px",
                    height: "360px",
                    display: "flex",
                    flexDirection: "column",
                    justifyContent: "space-around",
                    textAlign: "center",
                    cursor: "pointer",
                  }}
                  onClick={(e) => {
                    e.preventDefault();
                    window.location.href = `/products/certificate/${item.id}`;
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
                      image={`${urlPictures}/${item.images[0].name}`}
                      alt={`Изображение категории ${item.name}`}
                      sx={{
                        width: { xs: "270px", md: "180px", lg: "200px" },
                        height: {
                          xs: "270px",
                          sm: "200px",
                          md: "200px",
                          lg: "200px",
                        },
                        objectFit: "cover",
                      }}
                    />
                  </Box>
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <CardContent>
                      <Typography variant="h6">{item.name}</Typography>
                    </CardContent>
                  </Box>
                </Card>
              </Grid>
            ))
          ) : (
            <Typography variant="h6">Нет данных</Typography>
          )}
        </Grid>
      </Container>
    </Box>
  );
}
