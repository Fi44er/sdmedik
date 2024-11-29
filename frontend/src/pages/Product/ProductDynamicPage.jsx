import { Box, Button, Container, IconButton, Typography } from "@mui/material";
import React, { useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/swiper-bundle.css";

export default function ProductDynamicPage() {
  const [mainImage, setMainImage] = useState(0);
  const images = [
    "/wheelchair.png", // Замените на ваши изображения
    "/wheelchair.png",
    "/slideGalary1.png",
  ];

  return (
    <Box>
      <Container>
        <Box sx={{ display: "flex", justifyContent: "space-between" }}>
          {/* Основная картинка */}
          <Box sx={{ display: "flex", flexDirection: "column", width: "50%" }}>
            <Box>
              <img
                src={images[mainImage]}
                alt="Main"
                style={{ width: "400px", height: "400px", objectFit: "cover" }}
              />
            </Box>
            {/* Слайдер с превью картинками */}
            <Box sx={{ width: "50%" }}>
              <Swiper
                spaceBetween={10}
                slidesPerView={3}
                onSlideChange={(swiper) => setMainImage(swiper.activeIndex)}
              >
                {images.map((image, index) => (
                  <SwiperSlide key={index}>
                    <img
                      src={image}
                      alt={`Preview ${index + 1}`}
                      style={{
                        width: "100%",
                        cursor: "pointer",
                        objectFit: "cover",
                      }}
                      onClick={() => setMainImage(index)}
                    />
                  </SwiperSlide>
                ))}
              </Swiper>
            </Box>
          </Box>
          <Box>
            <Typography>Инвалидная коляска Trend 40</Typography>
            <Typography>Ortonica Trend 40</Typography>
            <Typography>25000 ₽</Typography>
            <Box>
              <IconButton></IconButton>
              <Button>В 1 клик</Button>
            </Box>
          </Box>
        </Box>
      </Container>
    </Box>
  );
}
