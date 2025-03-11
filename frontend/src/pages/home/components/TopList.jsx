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
import useProductStore from "../../../store/productStore";
import useBascketStore from "../../../store/bascketStore";
import { urlPictures } from "../../../constants/constants";
export default function TopList() {
  const { fetchTopList, products } = useProductStore();
  const [quantity, setQuantity] = useState(1);
  const { addProductThisBascket } = useBascketStore();
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

  useEffect(() => {
    fetchTopList();
    console.log(products);
  }, []);
  const hendleAddProductThithBascket = async (id) => {
    setQuantity(quantity);
    const product_id = id;
    console.log(id, quantity);

    await addProductThisBascket(product_id, quantity);
  };

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
        <img style={{ width: "100%" }} src="/Line 1.png" alt="Линия" />
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
            spacing={{ xs: 1, md: 4, lg: 2 }}
            columns={{ xs: 4, sm: 4, md: 4 }}
          >
            {products.data &&
              products?.data.map((item, index) => (
                <Grid item="true" xs={1} sm={1} md={1} key={index}>
                  <Card
                    sx={{
                      maxWidth: { xs: "167px", md: "261px" },
                      height: { xs: "385px", md: "514px" },
                      background: "#F5FCFF",
                      boxShadow: "0 4px 20px rgba(0, 0, 0, 0.1)",
                      borderRadius: "8px",
                      transition: "transform 0.2s, box-shadow 0.2s",
                      "&:hover": {
                        transform: "scale(1.05)",
                        boxShadow: " 0 8px 30px rgba(0, 0, 0, 0.2)",
                      },
                      display: "flex",
                      flexDirection: "column",
                      cursor: "pointer",
                    }}
                  >
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",
                        height: "300px",
                        borderBottom: "1px solid #E0E0E0",
                      }}
                      onClick={() => {
                        window.location.href = `/product/${item.id}`;
                      }}
                    >
                      <CardMedia
                        component="img"
                        image={`${urlPictures}/${item.image}`}
                        alt={item.name}
                        sx={{
                          width: "100%",
                          height: { xs: "200px", md: "300px" },
                          objectFit: "cover",
                        }}
                        loading="lazy"
                      />
                    </Box>
                    <CardContent>
                      <Typography
                        sx={{
                          fontSize: { xs: "0.9rem", md: "1.2rem" },
                          fontWeight: "bold",
                          mb: 1,
                          width: { xs: "150px", md: "235px" },
                          overflow: "hidden",
                          textOverflow: "ellipsis",
                          whiteSpace: "nowrap",
                        }}
                        onClick={() => {
                          window.location.href = `/product/${item.id}`;
                        }}
                      >
                        {item.name}
                      </Typography>
                      <Typography
                        variant="body"
                        sx={{
                          color: "text.secondary",
                          // textDecoration: "line-through",
                        }}
                      >
                        Всего заказов {item.order_count} шт.
                      </Typography>
                      <Box
                        sx={{
                          display: "flex",
                          justifyContent: "space-between",
                          alignItems: "center",
                          mt: 1,
                        }}
                      >
                        <Typography
                          variant="h6"
                          sx={{ color: "#00B3A4", fontWeight: "bold" }}
                          onClick={() => {
                            window.location.href = `/product/${item.id}`;
                          }}
                        >
                          {item.price} ₽
                        </Typography>
                        {/* {e.oldPrice && (
                              <Typography
                                variant="body2"
                                sx={{ color: "text.secondary", textDecoration: "line-through" }}
                              >
                                {e.oldPrice} ₽
                              </Typography>
                            )} */}
                      </Box>
                      <Button
                        variant="contained"
                        sx={{
                          width: "100%",
                          mt: { xs: 1, md: 5 },
                          backgroundColor: "#00B3A4",
                          color: "#FFFFFF",
                          borderRadius: "8px",
                          "&:hover": {
                            backgroundColor: "#00B3A4",
                          },
                        }}
                        onClick={() => {
                          hendleAddProductThithBascket(item.id);
                        }}
                      >
                        В корзину
                      </Button>
                    </CardContent>
                    {/* <Box
                          sx={{
                            display: "flex",
                            justifyContent: "space-between",
                            alignItems: "center",
                            p: 2,
                            borderTop: "1px solid #E0E0E0",
                          }}
                        >
                          <Button
                            sx={{
                              width: "100%",
                              height: "40px",
                              border: `2px solid #00B3A4`,
                              borderRadius: "12px",
                              color: "#00B3A4",
                            }}
                            variant="outlined"
                            onClick={() => {
                              hendleAddProductThithBascket(e.id);
                            }}
                          >
                            В корзину
                          </Button>
                        </Box> */}
                  </Card>
                </Grid>
              ))}
          </Grid>
          {/* <Box sx={{ mt: 4, mb: 4, display: "flex", justifyContent: "right" }}>
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
          </Box> */}
        </Box>
      </motion.div>
    </Box>
  );
}
