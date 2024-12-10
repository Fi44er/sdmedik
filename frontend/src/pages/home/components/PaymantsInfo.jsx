import { Box, Button, CardMedia, Typography } from "@mui/material";
import { motion } from "framer-motion";
import React from "react";

export default function PaymantsInfo() {
  return (
    <motion.div
      initial={{ y: -1000 }} // Начальная позиция (сверху)
      animate={{ y: 0 }} // Конечная позиция (по центру)
      transition={{ duration: 1.2 }} // Длительность анимации
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
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
              Оплата электронным сертификатом
            </Typography>
            <Typography variant="h6" color="white" component="p">
              Теперь оплачивать покупки на нашем сайте вы можете и электронным
              сертификатом
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
                window.location.href = "/certificate";
              }}
            >
              Подробнее
            </Button>
          </Box>
          <Box sx={{ width: { xs: "100%", lg: "50%" } }}>
            <CardMedia
              component="img"
              image="/public/Group 31.png"
              alt="title"
              sx={{
                width: "100%",
                height: { xs: "300px", sm: "300px", md: "400px" },
                objectFit: "cover",
              }}
            />
          </Box>
        </Box>
      </Box>
    </motion.div>
  );
}
