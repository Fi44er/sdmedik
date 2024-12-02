import {
  Box,
  Button,
  Container,
  Link,
  Paper,
  TextField,
  Typography,
} from "@mui/material";
import React from "react";
import { motion } from "framer-motion"; // Импортируем motion

const scaleVariants = {
  hidden: {
    opacity: 0,
    scale: 0, // Начальное состояние (уменьшенный)
  },
  visible: {
    opacity: 1,
    scale: 1, // Конечное состояние (нормальный размер)
    transition: {
      type: "spring",
      stiffness: 100,
      damping: 25,
    },
  },
};

export default function Auth() {
  return (
    <Box sx={{ display: "flex", justifyContent: "center" }}>
      <motion.div
        initial="hidden" // Начальное состояние
        animate="visible" // Конечное состояние
        variants={scaleVariants} // Используем определенные варианты анимации
        style={{ transformOrigin: "center" }} // Устанавливаем точку трансформации в центр
      >
        <Paper sx={{ p: 2, mt: 5, mb: 5, width: { xs: 320, md: 500 } }}>
          <Container>
            <Box
              sx={{ display: "flex", alignItems: "center", gridGap: 15, mb: 4 }}
            >
              <img src="/previwLogo.svg" alt="" />
              <Typography variant="h6" sx={{ color: "#2CC0B3" }}>
                Sdmedik
              </Typography>
            </Box>
            <Box sx={{ display: "flex", flexDirection: "column", gridGap: 30 }}>
              <Typography variant="h4">Вход</Typography>
              <TextField
                variant="outlined"
                label="Email"
                placeholder="your@email.com"
                sx={{
                  "& .MuiOutlinedInput-root": {
                    "&.Mui-focused fieldset": {
                      borderColor: "#2CC0B3", // Изменение цвета рамки при фокусе
                    },
                  },
                  "& .MuiInputLabel-root": {
                    "&.Mui-focused": {
                      color: "#2CC0B3", // Изменение цвета метки при фокусе
                    },
                  },
                }}
              />
              <TextField
                variant="outlined"
                label="Пороль"
                placeholder="Пороль"
                type="password"
                sx={{
                  "& .MuiOutlinedInput-root": {
                    "&.Mui-focused fieldset": {
                      borderColor: "#2CC0B3", // Изменение цвета рамки при фокусе
                    },
                  },
                  "& .MuiInputLabel-root": {
                    "&.Mui-focused": {
                      color: "#2CC0B3", // Изменение цвета метки при фокусе
                    },
                  },
                }}
              />
              <Button variant="contained" sx={{ background: "#2CC0B3" }}>
                Войти
              </Button>
            </Box>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                gridGap: 10,
                mt: 3,
                mb: 3,
                flexDirection: { xs: "column", md: "unset" },
              }}
            >
              <Typography>У вас нету аккаунта?</Typography>
              <Link href="/register" sx={{ color: "#2CC0B3" }}>
                Зарегистрироваться
              </Link>
            </Box>
          </Container>
        </Paper>
      </motion.div>
    </Box>
  );
}
