import {
  Box,
  Typography,
  CircularProgress,
  Container,
  Card,
  CardMedia,
  CardContent,
  CardHeader,
  Button,
  Grid,
  AppBar,
  Toolbar,
  IconButton,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import useUserStore from "../../store/userStore";
import { useNavigate } from "react-router-dom";
import MenuIcon from "@mui/icons-material/Menu";

const Product = [
  {
    id: 1,
    title: `Инвалидная коляска Trend 40`,
    country: `Бельгия`,
    price: `25,000.00 ₽`,
    image: `/public/wheelchair.png`,
  },
  {
    id: 2,
    title: `Инвалидная коляска Trend 40`,
    country: `Бельгия`,
    price: `25,000.00 ₽`,
    image: `/public/wheelchair.png`,
  },
];

export default function UserAccount() {
  const { getUserInfo, user, Logout } = useUserStore();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      await getUserInfo();
      setLoading(false);
    };

    fetchData();
  }, [getUserInfo]);

  if (loading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        height="100vh"
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box sx={{ backgroundColor: "#f5f5f5", minHeight: "100vh" }}>
      <Container sx={{ mt: 3 }}>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gridGap: 40,
            flexDirection: { xs: "column", md: "row" },
            mb: 4,
            padding: 2,
            backgroundColor: "#ffffff",
            borderRadius: 2,
            boxShadow: 3,
          }}
        >
          <Box>
            <img
              src="/user_Profile.png"
              alt="User  Profile"
              style={{ borderRadius: "50%", width: "100px", height: "100px" }}
            />
          </Box>
          {user && user.data ? (
            <Box>
              <Typography variant="h5" sx={{ fontWeight: "bold" }}>
                {user.data.fio}
              </Typography>
              <Typography variant="body1" sx={{ color: "gray" }}>
                {user.data.email}
              </Typography>
              <Typography variant="body1" sx={{ color: "gray" }}>
                {user.data.phone_number}
              </Typography>
              <Button
                variant="contained"
                color="error"
                onClick={(e) => {
                  e.preventDefault();
                  window.location.href = "/";
                  Logout();
                }}
                sx={{ mt: 2 }}
              >
                Выйти
              </Button>
            </Box>
          ) : (
            <Typography>No user data available</Typography>
          )}
        </Box>
        <Box sx={{ mt: 3 }}>
          <Typography variant="h4" sx={{ fontWeight: "bold", mb: 2 }}>
            Мои заказы
          </Typography>
          <Grid container spacing={2}>
            {Product.map((e) => (
              <Grid item key={e.id} xs={12} sm={6} md={4}>
                <Card
                  sx={{
                    background: "#ffffff",
                    borderRadius: 2,
                    boxShadow: 2,
                    transition: "0.3s",
                    "&:hover": {
                      boxShadow: 6,
                    },
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <CardMedia
                      component="img"
                      image={e.image}
                      alt={e.title}
                      sx={{ 
                        width: "250px",
                        height: "250px",
                        objectFit: "cover",
                        borderTopLeftRadius: 2,
                        borderTopRightRadius: 2,
                      }}
                    />
                  </Box>
                  <CardContent>
                    <CardHeader title={e.title} sx={{ p: 0 }} />
                    <Typography variant="body2" color="text.secondary">
                      {e.country}
                    </Typography>
                    <Typography variant="h6" sx={{ color: "black" }}>
                      {e.price}
                    </Typography>
                    <Box sx={{ mt: 2 }}>
                      <Typography variant="h6" sx={{ fontWeight: "bold" }}>
                        Статус заказа
                      </Typography>
                      <Typography variant="body2">В процессе</Typography>
                    </Box>
                    <Box sx={{ mt: 1 }}>
                      <Typography variant="h6" sx={{ fontWeight: "bold" }}>
                        Адрес
                      </Typography>
                      <Typography variant="body2">
                        Улица, Город, Страна
                      </Typography>
                    </Box>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
        </Box>
      </Container>
    </Box>
  );
}
