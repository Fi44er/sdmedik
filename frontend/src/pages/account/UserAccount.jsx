import React, { useEffect, useState } from "react";
import {
  Box,
  Typography,
  CircularProgress,
  Container,
  Card,
  CardContent,
  CardHeader,
  Button,
  Grid,
} from "@mui/material";
import useUserStore from "../../store/userStore";
import useOrderStore from "../../store/orderStore";

export default function UserAccount() {
  const { getUserInfo, user, Logout } = useUserStore();
  const { fetchUserOrders, userOrders } = useOrderStore();
  const [currentProducts, setCurrentProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      await getUserInfo();
      setLoading(false);
    };

    fetchData();
  }, [getUserInfo]);

  useEffect(() => {
    const fetchOrders = async () => {
      await fetchUserOrders();
    };
    fetchOrders();
  }, [fetchUserOrders]);

  useEffect(() => {
    if (userOrders.data?.length > 0) {
      const allItems = userOrders.data.flatMap((order) => order.items);
      setCurrentProducts(allItems);
    }
  }, [userOrders]);

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

  const statusStyles = {
    pending: { color: "orange", backgroundColor: "#fff3e0" },
    processing: { color: "blue", backgroundColor: "#e3f2fd" },
    completed: { color: "green", backgroundColor: "#e8f5e9" },
    canceled: { color: "red", backgroundColor: "#ffebee" },
  };

  const statusTranslations = {
    pending: "В ожидании",
    processing: "Рассмотрен",
    completed: "Завершен",
    canceled: "Отменен",
  };

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
                onClick={() => Logout()}
                sx={{ mt: 2 }}
              >
                Выйти
              </Button>
            </Box>
          ) : (
            <Typography>Нет данных о пользователе</Typography>
          )}
        </Box>
        <Box sx={{ mt: 3 }}>
          <Typography variant="h4" sx={{ fontWeight: "bold", mb: 2 }}>
            Мои заказы
          </Typography>
          <Grid container spacing={2}>
            {currentProducts.map((e) => (
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
                  <CardContent>
                    <CardHeader title={e.title} sx={{ p: 0 }} />
                    <Typography variant="h6" color="text.secondary">
                      Название товара: {e.name}
                    </Typography>
                    <Typography variant="h6" sx={{ color: "black" }}>
                      Сумма: {e.price} руб
                    </Typography>
                    <Box sx={{ mt: 2 }}>
                      <Typography variant="h6" sx={{ fontWeight: "bold" }}>
                        Статус заказа
                      </Typography>
                      <Box
                        sx={{
                          ...statusStyles[
                            userOrders.data.find(
                              (order) => order.id === e.order_id
                            )?.status
                          ],
                          padding: "5px",
                          borderRadius: "4px",
                          display: "inline-block",
                        }}
                      >
                        <Typography variant="body2">
                          {statusTranslations[
                            userOrders.data.find(
                              (order) => order.id === e.order_id
                            )?.status
                          ] || "Неизвестный статус"}
                        </Typography>
                      </Box>
                    </Box>
                    <Box sx={{ mt: 2 }}>
                      <Typography variant="h6" sx={{ fontWeight: "bold" }}>
                        Количество приобретенного товара:
                      </Typography>
                      <Typography variant="body2">{e.quantity}</Typography>
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
