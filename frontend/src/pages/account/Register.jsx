import {
  Box,
  Button,
  Container,
  Link,
  Paper,
  TextField,
  Typography,
  Snackbar,
  Modal,
} from "@mui/material";
import React, { useState } from "react";
import { motion } from "framer-motion";
import useAuthStore from "../../store/authStore";
import { useNavigate } from "react-router-dom";

const scaleVariants = {
  hidden: {
    opacity: 0,
    scale: 0,
  },
  visible: {
    opacity: 1,
    scale: 1,
    transition: {
      type: "spring",
      stiffness: 100,
      damping: 25,
    },
  },
};

export default function Register() {
  const {
    email,
    setEmail,
    fio,
    setFio,
    phone_number,
    setPhone_number,
    password,
    setPassword,
    registerFunc,
    showConfirmation,
    setShowConfirmation,
    code,
    setCode,

    verifyFunc,
  } = useAuthStore();

  const [error, setError] = useState(null);
  const navigate = useNavigate()

  const handleRegister = async () => {
    await registerFunc();
  };

  const handleConfirmationClose = () => {
    setShowConfirmation(false);
  };

  const handleVerify = async () => {
    await verifyFunc(navigate);
  };

  return (
    <Box sx={{ display: "flex", justifyContent: "center" }}>
      <motion.div
        initial="hidden"
        animate="visible"
        variants={scaleVariants}
        style={{ transformOrigin: "center" }}
      >
        <Paper sx={{ p: 2, mt: 5, mb: 5, width: { xs: 320, md: 500 } }}>
          <Container>
            <Box
              sx={{
                display: "flex",
                alignItems: "center",
                gridGap: 15,
                mb: 4,
              }}
            >
              <img src="/previwLogo.svg" alt="" />
              <Typography variant="h6" sx={{ color: "#2CC0B3" }}>
                Sdmedik
              </Typography>
            </Box>
            <Box sx={{ display: "flex", flexDirection: "column", gridGap: 30 }}>
              <Typography variant="h4">Регистрация</Typography>
              <TextField
                variant="outlined"
                label="Email"
                placeholder="your@email.com"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    "&.Mui-focused fieldset": {
                      borderColor: "#2CC0B3",
                    },
                  },
                  "& .MuiInputLabel-root": {
                    "&.Mui-focused": {
                      color: "#2CC0B3",
                    },
                  },
                }}
              />
              <TextField
                variant="outlined"
                label="Телефон"
                placeholder="+79228442121"
                value={phone_number}
                onChange={(e) => setPhone_number(e.target.value)}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    "&.Mui-focused fieldset": {
                      borderColor: "#2CC0B3",
                    },
                  },
                  "& .MuiInputLabel-root": {
                    "&.Mui-focused": {
                      color: "#2CC0B3",
                    },
                  },
                }}
              />
              <TextField
                variant="outlined"
                label="ФИО"
                placeholder="Иванов Дмитрий Сергеевич"
                value={fio}
                onChange={(e) => setFio(e.target.value)}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    "&.Mui-focused fieldset": {
                      borderColor: "#2CC0B3",
                    },
                  },
                  "& .MuiInputLabel-root": {
                    "&.Mui-focused": {
                      color: "#2CC0B3",
                    },
                  },
                }}
              />
              <TextField
                variant="outlined"
                label="Пороль"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                sx={{
                  "& .MuiOutlinedInput-root": {
                    "&.Mui-focused fieldset": {
                      borderColor: "#2CC0B3",
                    },
                  },
                  "& .MuiInputLabel-root": {
                    "&.Mui-focused": {
                      color: "#2CC0B3",
                    },
                  },
                }}
              />
              <Button
                type="button"
                variant="contained"
                sx={{ background: "#2CC0B3" }}
                onClick={handleRegister}
              >
                Зарегистрироваться
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
              <Typography>У вас есть аккаунт?</Typography>
              <Link href="/auth" sx={{ color: "#2CC0B3" }}>
                Войти
              </Link>
            </Box>
          </Container>
        </Paper>
      </motion.div>

      <Modal
        open={showConfirmation}
        onClose={handleConfirmationClose}
        aria-labelledby="confirmation-modal-title"
        aria-describedby="confirmation-modal-description"
        sx={{
          width: { xs: "350px", md: "500px" },
          position: "absolute",
          top: 400,
          left: { xs: 13, md: "37%" },
        }}
      >
        <Box sx={{ p: 4, bgcolor: "white", borderRadius: 2, boxShadow: 3 }}>
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              gridGap: 15,
              mb: 2,
            }}
          >
            <img src="/previwLogo.svg" alt="" />
            <Typography variant="h6" sx={{ color: "#2CC0B3" }}>
              Sdmedik
            </Typography>
          </Box>
          <Typography id="confirmation-modal-title" variant="h6">
            Подтверждение почты
          </Typography>
          <TextField
            variant="outlined"
            label="Введите код подтверждения"
            placeholder="Код"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            sx={{
              mt: 2,
              width: "100%",
              "& .MuiOutlinedInput-root": {
                "&.Mui-focused fieldset": {
                  borderColor: "#2CC0B3",
                },
              },
              "& .MuiInputLabel-root": {
                "&.Mui-focused": {
                  color: "#2CC0B3",
                },
              },
            }}
          />
          <Button
            variant="contained"
            sx={{ mt: 2, background: "#2CC0B3" }}
            onClick={handleVerify}
          >
            Подтвердить
          </Button>
        </Box>
      </Modal>
    </Box>
  );
}
