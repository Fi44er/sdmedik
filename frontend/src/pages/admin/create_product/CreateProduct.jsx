import {
  Box,
  Typography,
  TextField,
  Checkbox,
  FormControlLabel,
  Button,
  Container,
  Paper,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import useCategoryStore from "../../../store/categoryStore";
import axios from "axios"; // Не забудьте импортировать axios
import useProductStore from "../../../store/productStore";

export default function CreateProduct() {
  const { createCategory, fetchCategory, category } = useCategoryStore();
  const { createProduct } = useProductStore();

  const [product, setProduct] = useState({
    article: "",
    category_ids: [],
    characteristic_values: [],
    description: "",
    name: "",
  });

  const [characteristics, setCharacteristics] = useState([]);
  const [selectedCategories, setSelectedCategories] = useState([]); // Изменено на массив
  const [characteristicValues, setCharacteristicValues] = useState({}); // Добавлено состояние для значений характеристик

  useEffect(() => {
    fetchCategory();
    console.log(category.data);
  }, []);

  const handleCheckboxChange = (id) => {
    setSelectedCategories((prevSelected) => {
      const isSelected = prevSelected.includes(id);
      const newSelected = isSelected
        ? prevSelected.filter((categoryId) => categoryId !== id) // Удаляем ID, если он уже выбран
        : [...prevSelected, id]; // Добавляем ID, если он не выбран

      // Обновляем category_ids в product
      setProduct((prevProduct) => ({
        ...prevProduct,
        category_ids: newSelected,
      }));

      // Обновляем характеристики в зависимости от выбранных категорий
      const selected = category.data.find((item) => item.id === id);
      if (selected) {
        setCharacteristics((prevCharacteristics) => {
          const newCharacteristics = selected.characteristic || [];
          // Если категория была снята, удаляем характеристики
          if (isSelected) {
            return prevCharacteristics.filter(
              (char) =>
                !newCharacteristics.some((newChar) => newChar.id === char.id)
            );
          } else {
            return [
              ...new Set([...prevCharacteristics, ...newCharacteristics]),
            ]; // Уникальные характеристики
          }
        });
      }

      return newSelected;
    });
  };

  const handleValueChange = (id, value) => {
    // Обновляем состояние при изменении значения инпута
    setCharacteristicValues((prevValues) => ({
      ...prevValues,
      [id]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Обновляем product с новыми значениями характеристик
    const productData = {
      ...product,
      characteristic_values: Object.entries(characteristicValues).map(
        ([id, value]) => ({
          characteristic_id: Number(id), // Преобразуем id в число
          value: String(value),
        })
      ),
    };
    console.log(productData);
    createProduct(productData);
    // Здесь можно добавить логику для отправки данных на сервер
  };

  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Container>
        <Paper>
          <Container sx={{ p: 2 }}>
            <Typography
              variant="h4"
              sx={{ display: "flex", justifyContent: "center" }}
            >
              Создать продукт
            </Typography>
            <Box component="form" onSubmit={handleSubmit}>
              <TextField
                label="Название"
                value={product.name}
                onChange={(e) =>
                  setProduct({ ...product, name: e.target.value })
                }
                fullWidth
                margin="normal"
              />
              <TextField
                label="Артикул"
                value={product.article}
                onChange={(e) =>
                  setProduct({ ...product, article: e.target.value })
                }
                fullWidth
                margin="normal"
              />
              <TextField
                label="Описание"
                value={product.description}
                onChange={(e) =>
                  setProduct({ ...product, description: e.target.value })
                }
                fullWidth
                margin="normal"
                multiline
                rows={4}
              />
              <Box>
                <Box>
                  <Typography variant="h5">Категории</Typography>
                </Box>
                <Box sx={{ display: "flex", flexWrap: "wrap" }}>
                  {Array.isArray(category.data) && category.data.length > 0 ? (
                    category.data.map((item) => (
                      <Box>
                        <FormControlLabel
                          key={item.id}
                          control={
                            <Checkbox
                              checked={selectedCategories.includes(item.id)} // Проверяем, выбран ли ID
                              onChange={() => handleCheckboxChange(item.id)}
                            />
                          }
                          label={item.name}
                        />
                      </Box>
                    ))
                  ) : (
                    <p>Данных нет</p>
                  )}
                </Box>
              </Box>
              <Typography variant="h6">Характеристики</Typography>
              {Array.isArray(characteristics) &&
                characteristics.map((char) => {
                  if (char.data_type === "bool") {
                    return (
                      <Box key={char.id}>
                        <Typography>{char.name}</Typography>
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={characteristicValues[char.id] === true}
                              onChange={() => handleValueChange(char.id, true)}
                            />
                          }
                          label="Да"
                        />
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={characteristicValues[char.id] === false}
                              onChange={() => handleValueChange(char.id, false)}
                            />
                          }
                          label="Нет"
                        />
                      </Box>
                    );
                  } else {
                    return (
                      <Box>
                        <Typography>{char.name}:</Typography>
                        <TextField
                          key={char.id}
                          label={`Значение для ${char.name}`}
                          value={characteristicValues[char.id] || ""}
                          onChange={(e) =>
                            handleValueChange(char.id, e.target.value)
                          }
                          fullWidth
                          margin="normal"
                        />
                      </Box>
                    );
                  }
                })}
              <Button type="submit" variant="contained" color="primary">
                Создать продукт
              </Button>
            </Box>
          </Container>
        </Paper>
      </Container>
    </Box>
  );
}
