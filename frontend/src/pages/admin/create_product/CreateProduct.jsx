import {
  Box,
  Typography,
  TextField,
  Checkbox,
  FormControlLabel,
  Button,
  Container,
  Paper,
  InputBase,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import useCategoryStore from "../../../store/categoryStore";
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
    images: [], // Добавлено состояние для изображений
  });

  const [characteristics, setCharacteristics] = useState([]);
  const [selectedCategories, setSelectedCategories] = useState([]);
  const [characteristicValues, setCharacteristicValues] = useState({});

  useEffect(() => {
    fetchCategory();
    console.log(category.data);
  }, []);

  const handleCheckboxChange = (id) => {
    setSelectedCategories((prevSelected) => {
      const isSelected = prevSelected.includes(id);
      const newSelected = isSelected
        ? prevSelected.filter((categoryId) => categoryId !== id)
        : [...prevSelected, id];

      setProduct((prevProduct) => ({
        ...prevProduct,
        category_ids: newSelected,
      }));

      const selected = category.data.find((item) => item.id === id);
      if (selected) {
        setCharacteristics((prevCharacteristics) => {
          const newCharacteristics = selected.characteristic || [];
          if (isSelected) {
            return prevCharacteristics.filter(
              (char) =>
                !newCharacteristics.some((newChar) => newChar.id === char.id)
            );
          } else {
            return [
              ...new Set([...prevCharacteristics, ...newCharacteristics]),
            ];
          }
        });
      }

      return newSelected;
    });
  };

  const handleValueChange = (id, value) => {
    setCharacteristicValues((prevValues) => ({
      ...prevValues,
      [id]: value,
    }));
  };

  const handleFileChange = (e) => {
    const files = Array.from(e.target.files);
    setProduct((prevProduct) => ({
      ...prevProduct,
      images: files,
    }));
  };
  const handleSubmit = (e) => {
    e.preventDefault();

    // Создаем объект с данными продукта
    const productData = {
      article: product.article,
      category_ids: product.category_ids,
      characteristic_values: Object.entries(characteristicValues).map(
        ([id, value]) => ({
          characteristic_id: Number(id),
          value: String(value),
        })
      ),
      description: product.description,
      name: product.name,
    };

    // Преобразуем объект в строку JSON
    const jsonData = JSON.stringify(productData);

    // Создаем FormData
    const formData = new FormData();

    // Добавляем JSON-строку в FormData
    formData.append("json", jsonData);

    // Добавляем массив изображений в FormData
    product.images.forEach((file) => {
      formData.append("files", file);
    });

    // Логируем JSON для проверки
    console.log(jsonData);

    // Отправляем данные
    createProduct(formData); // Убедитесь, что функция createProduct может обрабатывать FormData
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
              <input
                type="file"
                multiple
                onChange={handleFileChange}
                accept="image/*"
              />
              <Box>
                {product.images.map((file, index) => (
                  <img
                    key={index}
                    src={URL.createObjectURL(file)}
                    alt="preview"
                    width="100"
                  />
                ))}
              </Box>
              <Box>
                <Box>
                  <Typography variant="h5">Категории</Typography>
                </Box>
                <Box sx={{ display: "flex", flexWrap: " wrap" }}>
                  {Array.isArray(category.data) && category.data.length > 0 ? (
                    category.data.map((item) => (
                      <Box key={item.id}>
                        <FormControlLabel
                          control={
                            <Checkbox
                              checked={selectedCategories.includes(item.id)}
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
                      <Box key={char.id}>
                        <Typography>{char.name}:</Typography>
                        <TextField
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
