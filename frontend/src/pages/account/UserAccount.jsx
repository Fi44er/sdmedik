import {
  Box,
  Typography,
  CircularProgress,
  Container,
  Card,
  CardMedia,
  CardContent,
  CardHeader,
  IconButton,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import Grid from "@mui/material/Grid2";
import useUserStore from "../../store/userStore";
const Product = [
  {
    id: 1,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
  {
    id: 2,
    title: `Инвалидная коляскаTrend 40`,
    country: `Бельгия`,
    price: `25000,00  ₽ `,
    image: `/public/wheelchair.png`,
  },
];

export default function UserAccount() {
  const { getUserInfo, user } = useUserStore();
  const [loading, setLoading] = useState(true); // Состояние загрузки

  useEffect(() => {
    const fetchData = async () => {
      await getUserInfo(); // Ждем, пока данные будут загружены
      setLoading(false); // Устанавливаем состояние загрузки в false
    };

    fetchData();
  }, [getUserInfo]);

  // Если данные еще загружаются, показываем индикатор загрузки
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

  // Если данные загружены, отображаем их
  return (
    <Box>
      <Container>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gridGap: 40,
            flexDirection: { xs: "column", md: "unset" },
          }}
        >
          <Box>
            <img src="/user_Profile.png" alt="" />
          </Box>
          {user && user.data ? ( // Проверяем, есть ли данные
            <Box>
              <Typography variant="h5">{user.data.fio}</Typography>
              <Typography variant="h6">{user.data.email}</Typography>
              <Typography variant="h6">{user.data.phone_number}</Typography>
              
            </Box>
          ) : (
            <Typography>No user data available</Typography>
          )}
        </Box>
        <Box>
          <Typography>Заказы</Typography>

          <Box>
            <Grid
              container
              spacing={{ xs: 2, md: 5 }}
              columns={{ xs: 4, sm: 4, md: 4 }}
              sx={{ mt: 4 }}
            >
              {Product.map((e) => {
                return (
                  <Grid
                    item
                    key={e.id}
                    xs={1}
                    sm={1}
                    md={1}
                    sx={{
                      width: "100%",
                    }}
                  >
                    <Card
                      sx={{
                        width: { xs: "100%", lg: "100%" },
                        background: "#F5FCFF",
                        display: "flex",
                        flexDirection: { xs: "column", md: "unset" },
                        p: { xs: 0, md: 3 },
                      }}
                    >
                      <Box
                        sx={{
                          display: "flex",
                          justifyContent: "center",
                          pt: 3,
                          pl: 3,
                        }}
                      >
                        <CardMedia
                          component="img"
                          image={e.image}
                          alt={"wheelchair"}
                          sx={{
                            width: { xs: "125px", md: "200px" },
                            height: { xs: "125px", sm: "200px", md: "200px" },
                            objectFit: "cover",
                          }}
                        />
                      </Box>
                      <CardContent
                        sx={{
                          display: "flex",
                          flexDirection: { xs: "column", md: "unset" },
                        }}
                      >
                        <Box>
                          <CardHeader
                            title={e.title}
                            sx={{ p: { xs: 0, md: 2 } }}
                          />
                          <Typography
                            variant="body2"
                            color="text.secondary"
                            sx={{ ml: { xs: 0, md: 2 } }}
                          >
                            {e.country}
                          </Typography>
                          <Typography
                            variant="h6"
                            sx={{ color: "black", ml: { xs: 0, md: 2 } }}
                          >
                            {e.price}
                          </Typography>
                        </Box>
                      </CardContent>
                      <CardContent
                        sx={{
                          display: "flex",
                          flexDirection: { xs: "column", md: "column" },
                        }}
                      >
                        <Box
                          sx={{
                            display: "flex",
                            flexDirection: "column",
                          }}
                        >
                          <Typography
                            variant="h5"
                            sx={{ color: "black", mt: { xs: 0, md: 2 } }}
                          >
                            Статус заказа
                          </Typography>
                          <Typography
                            variant="h6"
                            component="p"
                            sx={{ color: "black", mt: { xs: 0, md: 2 } }}
                          >
                            Арсений сделай нормальную карточку
                          </Typography>
                        </Box>
                        <Box
                          sx={{
                            display: "flex",
                            flexDirection: "column",
                          }}
                        >
                          <Typography
                            variant="h5"
                            sx={{ color: "black", mt: { xs: 0, md: 2 } }}
                          >
                            Адрес
                          </Typography>
                          <Typography
                            variant="h6"
                            sx={{ color: "black", mt: { xs: 0, md: 2 } }}
                          >
                            Арсений сделай нормальную карточку
                          </Typography>
                        </Box>
                      </CardContent>
                    </Card>
                  </Grid>
                );
              })}
            </Grid>
          </Box>
        </Box>
      </Container>
    </Box>
  );
}
