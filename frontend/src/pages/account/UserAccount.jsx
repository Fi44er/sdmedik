import { Box, Typography, CircularProgress } from "@mui/material";
import React, { useEffect, useState } from "react";
import useUserStore from "../../store/userStore";

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
      {user && user.data ? ( // Проверяем, есть ли данные
        <>
          <Typography>Email: {user.data.email}</Typography>
          <Typography>ФИО: {user.data.fio}</Typography>
          <Typography>Телефон: {user.data.phone_number}</Typography>
        </>
      ) : (
        <Typography>No user data available</Typography>
      )}
    </Box>
  );
}