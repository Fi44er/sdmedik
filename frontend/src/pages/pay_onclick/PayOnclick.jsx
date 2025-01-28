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
import { useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import CloseIcon from "@mui/icons-material/Close";
import useOrderStore from "../../store/orderStore";
import { useParams } from "react-router-dom";

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

export default function PayOnclick() {
  const {
    email,
    setEmail,
    fio,
    setFio,
    phone_number,
    setPhone_number,

    payOrderById,
  } = useOrderStore();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const { id } = useParams();

  const [error, setError] = useState(null);
  //   const navigate = useNavigate();

  const handlePay = async () => {
    await payOrderById(id);
    // window.location.href = response.data.data.id;
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
              <Typography variant="h4">Укажите контактные данные</Typography>
              <form
                onSubmit={handleSubmit(handlePay)}
                style={{
                  display: "flex",
                  flexDirection: "column",
                  gridGap: 30,
                  marginTop: "10px",
                }}
              >
                <TextField
                  variant="outlined"
                  label="Email"
                  placeholder="your@email.com"
                  {...register("email", {
                    required: "Это поле обязательно для заполнения",
                    pattern: {
                      value: /^\S+@\S+$/i,
                      message: "Неправильный формат email",
                    },
                  })}
                  error={!!errors.email}
                  helperText={errors.email ? errors.email.message : ""}
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
                  {...register("phone_number", {
                    required: "Это поле обязательно для заполнения",
                    minLength: {
                      value: 11,
                      message:
                        "Неправильный формат номера телефона,номер телефона должен быть 11 цифр",
                    },
                  })}
                  error={!!errors.phone_number}
                  helperText={
                    errors.phone_number ? errors.phone_number.message : ""
                  }
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
                  {...register("fio", {
                    required: "Это поле обязательно для заполнения",
                  })}
                  error={!!errors.fio}
                  helperText={errors.fio ? errors.fio.message : ""}
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
                <Button
                  type="submit"
                  variant="contained"
                  sx={{ background: "#2CC0B3" }}
                >
                  Перейти к оплате
                </Button>
              </form>
            </Box>
          </Container>
        </Paper>
      </motion.div>
    </Box>
  );
}
