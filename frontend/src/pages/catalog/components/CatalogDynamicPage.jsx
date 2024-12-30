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
import { useParams } from "react-router-dom";
import useProductStore from "../../../store/productStore";

export default function CatalogDynamicPage() {
  const { category_id } = useParams();

  const { fetchProducts, products } = useProductStore();

  useEffect(() => {
    fetchProducts(category_id);
    console.log(products);
  }, []);

  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Grid
        container
        spacing={{ xs: 2, md: 3 }}
        columns={{ xs: 4, sm: 4, md: 4 }}
      >
        {Array.isArray(products.data) && products.data.length > 0 ? (
          products.data.map((e) => (
            <Grid item key={e.id} xs={1} sm={1} md={1}>
              <Card
                sx={{
                  width: { xs: "100%", lg: "261px" },
                  background: "#F5FCFF",
                  cursor: "pointer",
                }}
                onClick={(item) => {
                  window.location.href = `/product/${e.id}`;
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
                    image={`http://127.0.0.1:8080/api/v1/image/${e.images[0].name}`}
                    alt={"wheelchair"}
                    sx={{
                      width: "200px",
                      height: { xs: "200px", sm: "200px", md: "200px" },
                      objectFit: "cover",
                    }}
                  />
                </Box>

                <CardContent>
                  <CardHeader title={e.name} />
                  <Typography variant="body2" color="text.secondary">
                    {e.article}
                  </Typography>
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "space-between",
                      alignItems: "center",
                    }}
                  >
                    <Typography variant="h6" sx={{ color: "black" }}>
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
                        border: `2px solid #00B3A4`,
                        borderRadius: "20px",
                        color: "#00B3A4",
                      }}
                      variant="outlined"
                    >
                      В 1 клик
                    </Button>
                    <IconButton>
                      <img
                        style={{ width: "50px", height: "50px" }}
                        src="/public/basket_cards.png"
                        alt=""
                      />
                    </IconButton>
                  </Box>
                </CardContent>
              </Card>
            </Grid>
          ))
        ) : (
          <Typography variant="h6">Нет данных</Typography>
        )}
      </Grid>
    </Box>
  );
}
