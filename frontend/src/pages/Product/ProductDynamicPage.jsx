import {
  Box,
  Button,
  Container,
  IconButton,
  Typography,
  Paper,
  Rating,
  Divider,
  List,
  ListItem,
  CardMedia,
} from "@mui/material";
import { ArrowBack, ArrowForward } from "@mui/icons-material";
import React, { useEffect, useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/swiper-bundle.css"; // Импорт стилей Swiper
import useProductStore from "../../store/productStore";
import { useParams } from "react-router-dom";

export default function ProductDetailPage() {
  const [mainImageIndex, setMainImageIndex] = useState(0);
  const [images, setImages] = useState([]);
  const { fetchProductById, products } = useProductStore();
  const { id } = useParams();

  const productDetails = {
    reviews: [
      { text: "Отличная коляска, очень удобная!", author: "Ирина", rating: 5 },
      { text: "Качество на высоте, рекомендую!", author: "Сергей", rating: 4 },
    ],
  };

  useEffect(() => {
    fetchProductById(id);
    console.log(products);
  }, [id]);
  useEffect(() => {
    if (products.data && products.data.images) {
      const fetchedImages = products.data.images.map(
        (image) => `http://127.0.0.1:8080/api/v1/image/${image.name}`
      );
      setImages(fetchedImages); // Сохраните изображения в состоянии
    }
  }, [products.data]);

  const handleNextImage = () => {
    setMainImageIndex((prevIndex) => (prevIndex + 1) % images.length);
  };

  const handlePrevImage = () => {
    setMainImageIndex(
      (prevIndex) => (prevIndex - 1 + images.length) % images.length
    );
  };

  return (
    <Container sx={{ mt: 5, mb: 5 }}>
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          flexDirection: { xs: "column", sm: "column", md: "unset" },
        }}
      >
        {/* Основная картинка с слайдером */}

        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            width: { xs: "100%", sm: "100%", md: "50%" },
          }}
        >
          {/* Слайдер с изображениями */}
          <Container>
            <Box
              sx={{ display: "flex", justifyContent: "center", marginTop: 2 }}
            >
              <Swiper spaceBetween={10} slidesPerView={1} navigation={false}>
                <SwiperSlide
                  style={{
                    display: "flex",
                    justifyContent: "center",
                  }}
                >
                  <CardMedia
                    component="img"
                    image={images[mainImageIndex]} // Отображаем основное изображение
                    alt={`Product Image ${mainImageIndex + 1}`}
                    style={{
                      width: { xs: 300, sm: 350, md: 400 },
                      height: { xs: 300, sm: 350, md: 400 },
                      objectFit: "cover",
                    }}
                  />
                </SwiperSlide>
              </Swiper>
            </Box>
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                marginTop: 2,
              }}
            >
              <IconButton onClick={handlePrevImage}>
                <ArrowBack />
              </IconButton>
              <IconButton onClick={handleNextImage}>
                <ArrowForward />
              </IconButton>
            </Box>
          </Container>

          {/* Превью лента изображений */}
          <Box sx={{ display: "flex", justifyContent: "center", marginTop: 2 }}>
            {products.data &&
              products.data.images &&
              products.data.images.map((image, index) => (
                <Button
                  variant="outlined"
                  key={index}
                  onClick={() => setMainImageIndex(index)}
                  sx={{
                    border:
                      mainImageIndex === index ? "2px solid #00B3A4" : "none",
                    color: "#00B3A4",
                    margin: 1,
                  }}
                >
                  <img
                    src={`http://127.0.0.1:8080/api/v1/image/${image.name}`}
                    alt={`Thumbnail ${index + 1}`}
                    style={{
                      width: 100,
                      height: 100,
                      objectFit: "cover",
                      borderRadius: "10px",
                    }}
                  />
                </Button>
              ))}
          </Box>
        </Box>

        {/* Информация о товаре */}
        <Box sx={{ width: { xs: "100%", sm: "100%", md: "50%" } }}>
          {products.data ? (
            <Box>
              <Typography variant="h5">
                Название товара : {products.data.name}
              </Typography>
              <Typography variant="subtitle1">
                Артикул: {products.data.article}
              </Typography>
              <Typography variant="h5">{products.data.price} р </Typography>
              {/* Other components */}
            </Box>
          ) : (
            <Typography>Loading...</Typography>
          )}
          <Box sx={{ marginTop: 2 }}>
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
              <img src="/basket_Cards.png" alt="Добавить в корзину" />
            </IconButton>
          </Box>
          <Divider sx={{ marginY: 2 }} />
          <Box sx={{ marginTop: 2 }}>
            <Typography variant="h6">Характеристики:</Typography>
            <List>
              {products.data?.characteristic_values?.map((feature, index) => (
                <ListItem key={index}>
                  <Typography>{feature.value}</Typography>
                </ListItem>
              ))}
            </List>
          </Box>
          <Divider sx={{ marginY: 2 }} />
        </Box>
      </Box>
      {/*Описание*/}

      <Box sx={{ marginTop: 2 }}>
        <Typography variant="h6">Описание:</Typography>
        <Typography>{products.data?.description}</Typography>
      </Box>
      {/* Отзывы */}
      <Box sx={{ marginTop: 4 }}>
        <Typography variant="h6">Отзывы:</Typography>
        {productDetails.reviews.map((review, index) => (
          <Paper
            key={index}
            elevation={2}
            sx={{
              padding: 2,
              marginTop: 1,
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Box>
              <Typography>{review.text}</Typography>
              <Typography variant="caption">
                - {review.author}, {review.rating} звезды
              </Typography>
            </Box>
            <Rating name="half-rating" defaultValue={5} precision={1} />
          </Paper>
        ))}
      </Box>
    </Container>
  );
}
