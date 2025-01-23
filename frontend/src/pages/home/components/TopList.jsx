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
import { motion } from "framer-motion";
import React, { useEffect, useState } from "react";
import { Helmet } from "react-helmet";

export default function TopList() {
  const [isVisible, setIsVisible] = useState(false);

  const handleScroll = () => {
    const componentPosition = document
      .getElementById("top-list")
      .getBoundingClientRect().top;
    const windowHeight = window.innerHeight;

    if (componentPosition < windowHeight) {
      setIsVisible(true);
      window.removeEventListener("scroll", handleScroll);
    }
  };

  useEffect(() => {
    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <Box component="article" id="top-list">
      <Helmet>
        <title>Лучшие товары - Магазина СД-МЕД</title>
        <meta
          name="description"
          content="Посмотрите лучшие товары нашего магазина."
        />
        <meta name="keywords" content="товары, магазин, кресло-коляска" />
      </Helmet>
      <motion.div
        initial={{ y: "100%", opacity: 0 }}
        animate={{ y: isVisible ? 0 : "100%", opacity: isVisible ? 1 : 0 }}
        transition={{ duration: 1 }}
      >
        <img style={{ width: "100%" }} src="/public/Line 1.png" alt="Линия" />
        <Box sx={{ mt: 3 }}>
          {/* <header> */}
          <Typography
            variant="h5"
            color="Black"
            sx={{
              mb: 4,
            }}
          >
            Лучшие товары
          </Typography>
          {/* </header> */}
          <Grid
            container
            spacing={{ xs: 2, md: 4, lg: 2 }}
            columns={{ xs: 4, sm: 4, md: 4 }}
          >
            {Array.from({ length: 4 }).map((_, index) => (
              <Grid item xs={1} sm={1} md={1} key={index}>
                <Card
                  sx={{
                    width: { xs: "100%", sm: "100%", md: "100%", lg: "276px" },
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
                      image={"/public/wheelchair.png"}
                      alt={"Кресло-коляска"}
                      sx={{
                        width: "200px",
                        height: { xs: "200px", sm: "200px", md: "200px" },
                        objectFit: "cover",
                      }}
                    />
                  </Box>

                  <CardContent>
                    <CardHeader
                      title={
                        "Кресло-коляска облегчённая механическая MEYRA Eurochair 2.750"
                      }
                    />
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "space-between",
                        alignItems: "center",
                      }}
                    >
                      <Typography variant="h6" sx={{ color: "#39C8B8" }}>
                        124456 руб.
                      </Typography>
                      <Typography
                        variant="body2"
                        sx={{
                          color: "text.secondary",
                          textDecoration: "line-through",
                        }}
                      >
                        124456 руб.
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
                          background: `linear-gradient(95.61deg, #A5DED1 4.71%, #00B3A4 97.25%)`,
                          borderRadius: "10px",
                          color: "#fff",
                        }}
                        variant=" contained"
                      >
                        Подробнее
                      </Button>
                      <IconButton>
                        <img
                          style={{ width: "50px", height: "50px" }}
                          src="/public/basket_cards.png"
                          alt="Корзина"
                        />
                      </IconButton>
                    </Box>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
          <Box sx={{ mt: 4, mb: 4, display: "flex", justifyContent: "right" }}>
            <Button
              sx={{
                width: "260px",
                height: "50px",
                fontSize: "18px",
                border: "2px solid #2CC0B3",
                color: "#2CC0B3",
              }}
              variant="outlined"
            >
              Посмотреть все
            </Button>
          </Box>
        </Box>
      </motion.div>
    </Box>
  );
}
