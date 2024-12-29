import React, { useEffect, useState } from "react";
import {
  Box,
  TextField,
  Button,
  IconButton,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Typography,
  Container,
} from "@mui/material";
import { Delete as DeleteIcon } from "@mui/icons-material";
import useCategoryStore from "../../store/categoryStore";
import AdminProductTable from "./components_admin_page/AdminProductTable";

export default function Admin() {
  const [name, setName] = useState("");
  const [characteristics, setCharacteristics] = useState([
    { data_type: "", name: "" },
  ]);
  const { createCategory, fetchCategory, category } = useCategoryStore();

  const handleCharacteristicChange = (index, value, type) => {
    const newCharacteristics = [...characteristics];
    if (type === "name") {
      newCharacteristics[index].name = value; // Изменяем только имя характеристики
    } else if (type === "data_type") {
      newCharacteristics[index].data_type = value; // Изменяем тип данных
    }
    setCharacteristics(newCharacteristics);
  };

  const addCharacteristic = () => {
    setCharacteristics([...characteristics, { data_type: "string", name: "" }]);
  };

  const removeCharacteristic = (index) => {
    const newCharacteristics = characteristics.filter((_, i) => i !== index);
    setCharacteristics(newCharacteristics);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const data = {
      name: name,
      characteristics: characteristics,
    };
    createCategory(data);
    console.log("data:", data);
  };

  useEffect(() => {
    fetchCategory();
    console.log(category);
  }, []);

  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Container>
        <Box>
          <AdminProductTable />
        </Box>
      </Container>
    </Box>
  );
}
