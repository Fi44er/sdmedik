import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  Paper,
  Table,
  TableContainer,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Pagination,
  Typography,
} from "@mui/material";
import useUserStore from "../../../../store/userStore";

const AdminUserTable = () => {
  const { fetchUsers, users } = useUserStore();
  const [filteredUsers, setFilteredUsers] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 20;

  useEffect(() => {
    fetchUsers();
    console.log(users);
  }, []);

  useEffect(() => {
    if (Array.isArray(users.data)) {
      setFilteredUsers(users.data);
    }
  }, [users]);

  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = filteredUsers.slice(indexOfFirstItem, indexOfLastItem);

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
  };

  //   const handleDeleteProduct = async (id) => {
  //     await deleteProduct(id);
  //     fetchProducts();
  //   };

  return (
    <Box sx={{ padding: 2 }}>
      <Typography sx={{ fontSize: "30px", mb: 2, mt: 2 }}>
        Таблица с пользователями
      </Typography>
      <Paper sx={{ width: "100%" }}>
        {/* Таблица для больших экранов */}
        <TableContainer
          sx={{ overflowX: "auto", display: { xs: "none", sm: "block" } }}
        >
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Email</TableCell>
                <TableCell>ФИО</TableCell>
                <TableCell>Номер телефона</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {currentItems.map((user) => (
                <TableRow key={user.id}>
                  <TableCell>{user.email}</TableCell>
                  <TableCell>{user.fio}</TableCell>
                  <TableCell>{user.phone_number}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>

        {/* Карточки для мобильных устройств */}
        <Box sx={{ display: { xs: "block", sm: "none" } }}>
          {currentItems.map((user) => (
            <Paper key={user.id} sx={{ mb: 2, p: 2 }}>
              <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
                <Box>Email: {user.email}</Box>
                <Box>ФИО: {user.fio}</Box>
                <Box>Номер телефона: {user.phone_number}</Box>
                <Box sx={{ display: "flex", gap: 1 }}>
                  <Button
                    variant="contained"
                    color="error"
                    onClick={() => handleDeleteProduct(user.id)}
                  >
                    удалить
                  </Button>
                  <Button
                    variant="contained"
                    color="info"
                    onClick={(e) => {
                      e.preventDefault();
                      window.location.href = `/update_product/${user.id}`;
                    }}
                  >
                    редактировать
                  </Button>
                </Box>
              </Box>
            </Paper>
          ))}
        </Box>

        {/* Пагинация */}
        <Box sx={{ display: "flex", justifyContent: "center", mt: 2, mb: 2 }}>
          <Pagination
            count={Math.ceil(filteredUsers.length / itemsPerPage)}
            page={currentPage}
            onChange={handlePageChange}
            color="primary"
            sx={{ mt: 2, mb: 2 }}
          />
        </Box>
      </Paper>
    </Box>
  );
};

export default AdminUserTable;
