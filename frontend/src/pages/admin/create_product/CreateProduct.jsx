import {
  Box,
  Typography,
  TextField,
  Checkbox,
  FormControlLabel,
  Button,
  Container,
} from "@mui/material";
import React, { useState } from "react";

export default function CreateProduct() {
  const [product, setProduct] = useState({
    article: "",
    category_ids: [],
    characteristic_values: [],
    description: "",
    name: "",
  });

  

  const [characteristics, setCharacteristics] = useState([
    { id: 1, name: "Characteristic 1" },
    { id: 2, name: "Characteristic 2" },
    { id: 3, name: "Characteristic 3" },
  ]);

  useEffect(() => {
    fetchCategory();
    console.log(category);
  }, []);

  const handleCheckboxChange = (id) => {
    setProduct((prev) => {
      const newCharacteristicValues = prev.characteristic_values.some(
        (c) => c.characteristic_id === id
      )
        ? prev.characteristic_values.filter((c) => c.characteristic_id !== id)
        : [...prev.characteristic_values, { characteristic_id: id, value: "" }];
      return { ...prev, characteristic_values: newCharacteristicValues };
    });
  };

  const handleValueChange = (id, value) => {
    setProduct((prev) => {
      const newCharacteristicValues = prev.characteristic_values.map((c) =>
        c.characteristic_id === id ? { ...c, value } : c
      );
      return { ...prev, characteristic_values: newCharacteristicValues };
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(product);
    // Здесь можно добавить логику для отправки данных на сервер
  };

  return (
    <Box component="form" onSubmit={handleSubmit}>
      <Container>
        <Typography variant="h4">Создать продукт</Typography>
        <TextField
          label="Название"
          value={product.name}
          onChange={(e) => setProduct({ ...product, name: e.target.value })}
          fullWidth
          margin="normal"
        />
        <TextField
          label="Артикул"
          value={product.article}
          onChange={(e) => setProduct({ ...product, article: e.target.value })}
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
        {Array.isArray(category.data) && category.data.length > 0 ? (
          category.data.map((item, index) => (
            <div key={index}>
              <h2>{item.name}</h2>
              <p>{item.description}</p>

            </div>
          ))
        ) : (
          <p>Данных нет</p>
        )}
        <Typography variant="h6">Характеристики</Typography>
        {characteristics.map((char) => (
          <FormControlLabel
            key={char.id}
            control={
              <Checkbox
                checked={product.characteristic_values.some(
                  (c) => c.characteristic_id === char.id
                )}
                onChange={() => handleCheckboxChange(char.id)}
              />
            }
            label={char.name}
          />
        ))}
        {product.characteristic_values.map((char) => (
          <TextField
            key={char.characteristic_id}
            label={`Значение для ${
              characteristics.find((c) => c.id === char.characteristic_id)?.name
            }`}
            value={char.value}
            onChange={(e) =>
              handleValueChange(char.characteristic_id, e.target.value)
            }
            fullWidth
            margin="normal"
          />
        ))}
        <Button type="submit" variant="contained" color="primary">
          Создать продукт
        </Button>
      </Container>
    </Box>
  );
}
