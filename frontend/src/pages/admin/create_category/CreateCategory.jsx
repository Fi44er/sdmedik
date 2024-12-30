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
  const [category, setCategory] = useState({
    name: "",
    characteristics: [{ data_type: "string", name: "" }],
    images: [], // Добавлено состояние для изображений
  });
  const { createCategory, fetchCategory } = useCategoryStore();

  const handleCharacteristicChange = (index, value, type) => {
    const newCharacteristics = [...category.characteristics];
    if (type === "name") {
      newCharacteristics[index].name = value;
    } else if (type === "data_type") {
      newCharacteristics[index].data_type = value;
    }
    // Correctly update the category state
    setCategory({ ...category, characteristics: newCharacteristics });
  };

  const addCharacteristic = () => {
    setCategory((prevCategory) => ({
      ...prevCategory,
      characteristics: [
        ...prevCategory.characteristics,
        { data_type: "string", name: "" },
      ],
    }));
  };

  const removeCharacteristic = (index) => {
    const newCharacteristics = category.characteristics.filter(
      (_, i) => i !== index
    );
    setCategory({ ...category, characteristics: newCharacteristics });
  };

  const handleFileChange = (e) => {
    const files = Array.from(e.target.files);
    setCategory((prevCategory) => ({
      ...prevCategory,
      images: files,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    // Create an object with category data
    const categoryData = {
      name: category.name,
      characteristics: category.characteristics.map((characteristic) => {
        return {
          data_type: characteristic.data_type,
          name: characteristic.name,
        };
      }),
    };

    // Convert object to JSON string
    const jsonData = JSON.stringify(categoryData);

    // Create FormData
    const formData = new FormData();

    // Add JSON string to FormData
    formData.append("json", jsonData);

    // Add array of images to FormData
    category.images.forEach((file) => {
      formData.append("file", file);
    });

    // Log JSON for verification
    console.log(jsonData);

    // Send data
    createCategory(formData); // Ensure createCategory can handle FormData
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
        value={category.name}
        onChange={(e) => setCategory({ ...category, name: e.target.value })}
        sx={{ mb: 2 }}
      />
      <input
        type="file"
        multiple
        onChange={handleFileChange}
        accept="image/*"
      />
      {category.characteristics.map((characteristic, index) => (
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
      <Box sx={{ mt: 3 }}></Box>
    </Box>
  );
}
