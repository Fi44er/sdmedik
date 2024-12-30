import { Box, Button, CardMedia, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import Slider from "react-slick";

export default function PromotionalSlider() {
  const [slides, setSlides] = useState([]);

  useEffect(() => {
    // Функция для получения данных с сервера
    const fetchSlides = async () => {
      try {
        const response = await fetch("/api/slides"); // Замените на ваш API
        const data = await response.json();
        setSlides(data);
      } catch (error) {
        console.error("Ошибка при загрузке слайдов:", error);
      }
    };

    fetchSlides();
  }, []);

  const settings = {
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
  };

  return (
    <Box>
      <Slider {...settings}>
        {slides.map((slide, index) => (
          <Box
            key={index}
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              background: `linear-gradient(280.17deg, #00B3A4 -56.17%, #66D1C6 100%)`,
              borderRadius: "10px",
              padding: { xs: "20px", lg: "70px" },
            }}
          >
            <Box
              sx={{
                display: "flex",
                flexDirection: { xs: "column", lg: "unset" },
                gridGap: { xs: "40px", lg: 0 },
              }}
            >
              <Box
                sx={{
                  width: "50%",
                  display: "flex",
                  flexDirection: "column",
                  gridGap: 20,
                }}
              >
                <Typography
                  variant="h2"
                  color="white"
                  sx={{ fontSize: { xs: "40px", lg: "60px" } }}
                >
                  {slide.title} {/* Динамическое заголовок */}
                </Typography>
                <Typography variant="h6" color="white" component="p">
                  {slide.description} {/* Динамическое описание */}
                </Typography>
                <Button
                  sx={{
                    display: "flex",
                    justifyContent: "left",
                    background: `linear-gradient(95.61deg, #A5DED1 4.71%, #00B3A4 97.25%)`,
                    width: "max-content",
                    padding: "13px 39px",
                    color: "white",
                    fontSize: "18px",
                  }}
                  onClick={(e) => {
                    e.preventDefault();
                    window.location.href = slide.link; // Динамическая ссылка
                  }}
                >
                  Подробнее
                </Button>
              </Box>
              <Box sx={{ width: { xs: "100%", lg: "50%" } }}>
                <CardMedia
                  component="img"
                  image={slide.image} // Динамическое изображение
                  alt={slide.altText} // Динамический alt текст
                  sx={{
                    width: "100%",
                    height: { xs: "300px", sm: "300px", md: "400px" },
                    objectFit: "cover",
                  }}
                />
              </Box>
            </Box>
          </Box>
        ))}
      </Slider>
    </Box>
  );
}
