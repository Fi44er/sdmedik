import {
  Box,
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import React, { useEffect } from "react";
import useCategoryStore from "../../../../store/categoryStore";

export default function AdminCategoriesTable() {
  const { fetchCategory, category } = useCategoryStore();

  useEffect(() => {
    fetchCategory();
  }, []);

  return (
    <Box>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell align="center">Название Категори</TableCell>
              <TableCell align="center">Действия</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {Array.isArray(category.data) &&
              category.data.map((item) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.name}</TableCell>
                  <TableCell align="center">
                    <Button>Удалить</Button>
                  </TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
}
