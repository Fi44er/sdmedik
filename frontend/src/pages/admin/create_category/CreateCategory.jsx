import React, { useState } from "react";
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
import useCategoryStore from "../../../store/categoryStore";

export default function CreateCategory() {
  const [name, setName] = useState("");
  const [characteristics, setCharacteristics] = useState([
    { data_type: "string", name: "" },
  ]);
  const { createCategory, fetchCategory, category } = useCategoryStore();

  const handleCharacteristicChange = (index, value, type) => {
    const newCharacteristics = [...characteristics];
    if (type === "name") {
      newCharacteristics[index].name = value;
    } else if (type === "data_type") {
      newCharacteristics[index].data_type = value;
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

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      sx={{
        maxWidth: 400,
        margin: "auto",
        padding: 3,
        borderRadius: 2,
        boxShadow: 3,
        backgroundColor: "#f5f5f5",
        mt: 4,
        mb: 4,
      }}
    >
      <Typography variant="h5" sx={{ mb: 2, textAlign: "center" }}>
        Создать категорию
      </Typography>
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
          sx={{ display: "flex", flexDirection: "column", mb: 1 }}
        >
          <TextField
            label={`Характеристика ${index + 1}`}
            variant="outlined"
            fullWidth
            value={characteristic.name}
            onChange={(e) =>
              handleCharacteristicChange(index, e.target.value, "name")
            }
            sx={{ mb: 1 }}
          />
          <FormControl fullWidth sx={{ mb: 1 }}>
            <InputLabel>Тип данных</InputLabel>
            <Select
              value={characteristic.data_type}
              onChange={(e) =>
                handleCharacteristicChange(index, e.target.value, "data_type")
              }
              required
            >
              <MenuItem value="string">Строковое значение</MenuItem>
              <MenuItem value="int">Целочисленое значение</MenuItem>
              <MenuItem value="float">Дробное значение</MenuItem>
              <MenuItem value="bool">Есть/нету</MenuItem>
            </Select>
          </FormControl>

          <IconButton
            onClick={() => removeCharacteristic(index)}
            color="error"
            sx={{ alignSelf: "flex-end" }}
          >
            <DeleteIcon />
          </IconButton>
        </Box>
      ))}
      <Button variant="outlined" onClick={addCharacteristic} sx={{ mb: 2 }}>
        Добавить характеристику
      </Button>
      <Button
        type="submit"
        variant="contained"
        sx={{ backgroundColor: "#3f51b5", color: "#fff" }}
      >
        Сохранить категорию
      </Button>
      <Box sx={{ mt: 3 }}>
        {Array.isArray(category.data) && category.data.length > 0 ? (
          category.data.map((item, index) => (
            <Box
              key={index}
              sx={{
                border: "1px solid #ccc",
                borderRadius: 1,
                padding: 2,
                mb: 2,
              }}
            >
              <Typography variant="h6">{item.name}</Typography>
              <Typography variant="body2">{item.description}</Typography>
            </Box>
          ))
        ) : (
          <Typography variant="body2" sx={{ textAlign: "center" }}>
            Данных нет
          </Typography>
        )}
      </Box>
    </Box>
  );
}
