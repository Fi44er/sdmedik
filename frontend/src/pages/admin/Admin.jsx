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
} from "@mui/material";
import { Delete as DeleteIcon } from "@mui/icons-material";
import useCategoryStore from "../../store/categoryStore";

export default function Admin() {
  const [name, setName] = useState("");
  const [characteristics, setCharacteristics] = useState([
    { data_type: "string", name: "" },
  ]);
  const { createCategory, getAllCategory, category } = useCategoryStore();

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
    getAllCategory();
  }, []);

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      sx={{ maxWidth: 400, margin: "auto" }}
    >
      <TextField
        label="Название категории"
        variant="outlined"
        fullWidth
        value={name}
        onChange={(e) => setName(e.target.value)}
        sx={{ mb: 2 }}
      />
      {characteristics.map((characteristic, index) => (
        <Box
          key={index}
          sx={{ display: "flex", flexDirection: "column-reverse", mb: 1 }}
        >
          <FormControl fullWidth sx={{ mr: 1 }}>
            <label sx={{}}>Тип данных</label>
            <Select
              value={characteristic.data_type}
              onChange={(e) =>
                handleCharacteristicChange(index, e.target.value, "data_type")
              }
              required
            >
              <MenuItem value="string">Строковое значение</MenuItem>
              <MenuItem value="int">Целочисленое значние</MenuItem>
              <MenuItem value="float">Дробное значение</MenuItem>
              <MenuItem value="bool">Есть\нету</MenuItem>
            </Select>
          </FormControl>
          <TextField
            label={`Характеристика ${index + 1}`}
            variant="outlined"
            fullWidth
            value={characteristic.name}
            onChange={(e) =>
              handleCharacteristicChange(index, e.target.value, "name")
            }
          />
          <IconButton
            onClick={() => removeCharacteristic(index)}
            color="error"
            sx={{ ml: 1 }}
          >
            <DeleteIcon />
          </IconButton>
        </Box>
      ))}
      <Button variant="outlined" onClick={addCharacteristic} sx={{ mb: 2 }}>
        Добавить характеристику
      </Button>
      <Button type="submit" variant="contained">
        Сохранить категорию
      </Button>
      <Box>
        {Array.isArray(category) && category.length > 0 ? (
          category.map((e, index) => (
            <Typography key={index}>{e.name}</Typography>
          ))
        ) : (
          <Typography>No categories available</Typography>
        )}
      </Box>
    </Box>
  );
}
