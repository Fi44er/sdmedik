import { Box, Button, CardMedia, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

export default function PromotionalSlider() {
  const [slides, setSlides] = useState([]);

  useEffect(() => {
    // Замените на статические данные для тестирования
    const testSlides = [
      {
        title: "Слайд 1",
        description: "Описание для слайда 1",
        link: "https://example.com/slide1",
        image: "/public/wheelchair.png",
        altText: "Слайд 1",
      },
      {
        title: "Слайд 2",
        description: "Описание для слайда 2",
        link: "https://example.com/slide2",
        image: "https://via.placeholder.com/600x400?text=Slide+2",
        altText: "Слайд 2",
      },
      {
        title: "Слайд 3",
        description: "Описание для слайда 3",
        link: "https://example.com/slide3",
        image: "https://via.placeholder.com/600x400?text=Slide+3",
        altText: "Слайд 3",
      },
    ];

    setSlides(testSlides);
  }, []);

  const settings = {
    infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
  };

  return (
    <Box sx={{ mb: 2 }}>
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
                  {slide.title}
                </Typography>
                <Typography variant="h6" color="white" component="p">
                  {slide.description}
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
                    window.location.href = slide.link;
                  }}
                >
                  Подробнее
                </Button>
              </Box>
              <Box sx={{ width: { xs: "100%", lg: "50%" } }}>
                <CardMedia
                  component="img"
                  image={slide.image}
                  alt={slide.altText}
                  sx={{
                    width: "400px",
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
