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
import React, { useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/swiper-bundle.css"; // Импорт стилей Swiper

export default function ProductDetailPage() {
  const [mainImage, setMainImage] = useState(0);
  const images = ["/wheelchair.png", "/wheelchair.png", "/slideGalary1.png"];

  const productDetails = {
    title: "Инвалидная коляска Trend 40",
    brand: "Ortonica",
    price: "25000 ₽",
    description:
      "Кресло-коляска для инвалидов обеспечивает комфортное передвижение людям с нарушениями опорно-двигательного аппарата. Она выполнена из облегченного алюминия, оснащена регулировками, которые позволяют адаптировать коляску под особенности пользователя.",
    features: [
      "Грузоподъемность: 130кг",
      "Регулируемая спинка",
      "Легкий алюминиевый каркас",
      "Подходит для самостоятельного передвижения и с помощью сопровождающего лица",
    ],
    reviews: [
      { text: "Отличная коляска, очень удобная!", author: "Ирина", rating: 5 },
      { text: "Качество на высоте, рекомендую!", author: "Сергей", rating: 4 },
    ],
  };

  const handlePrevImage = () => {
    setMainImage((prev) => (prev === 0 ? images.length - 1 : prev - 1));
  };

  const handleNextImage = () => {
    setMainImage((prev) => (prev === images.length - 1 ? 0 : prev + 1));
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
                    image={images[mainImage]} // Отображаем основное изображение
                    alt={`Product Image ${mainImage + 1}`}
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
            {images.map((image, index) => (
              <Button
                variant="outlined"
                key={index}
                onClick={() => setMainImage(index)}
                sx={{
                  border: mainImage === index ? "2px solid #00B3A4" : "none",
                  color: "#00B3A4",
                  margin: 1,
                }}
              >
                <img
                  src={image}
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
          <Typography variant="h5">{productDetails.title}</Typography>
          <Typography variant="subtitle1">{productDetails.brand}</Typography>
          <Typography variant="h5">{productDetails.price}</Typography>
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
              {productDetails.features.map((feature, index) => (
                <ListItem key={index}>
                  <Typography>{feature}</Typography>
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
        <Typography>{productDetails.description}</Typography>
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
