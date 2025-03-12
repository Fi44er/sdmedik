import {
  Box,
  Typography,
  TextField,
  Checkbox,
  FormControlLabel,
  Button,
  Container,
  Paper,
  Grid,
  InputAdornment,
  IconButton,
  Avatar,
  CircularProgress,
} from "@mui/material";
import React, { useEffect, useState } from "react";
import useCategoryStore from "../../../store/categoryStore";
import useProductStore from "../../../store/productStore";
import { useParams } from "react-router-dom";
import { Delete as DeleteIcon, DirtyLens } from "@mui/icons-material";
import { urlPictures } from "../../../constants/constants";

export default function UpdateProduct() {
  const { fetchCategory, category } = useCategoryStore();
  const { updateProduct, fetchProductById, products } = useProductStore();
  const { id } = useParams();

  const [product, setProduct] = useState({
    article: "",
    category_ids: [],
    characteristic_values: [],
    description: "",
    name: "",
    images: [],
    price: 0,
    del_images: [],
  });

  const [characteristics, setCharacteristics] = useState([]);
  const [selectedCategories, setSelectedCategories] = useState([]);
  const [characteristicValues, setCharacteristicValues] = useState({});
  const [catalogs, setCatalogs] = useState([]);
  const [delImages, setDelImages] = useState({});
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchCategory();
  }, []);

  useEffect(() => {
    fetchProductById(id);
  }, [id]);

  const handleCatalogChange = (event) => {
    const { value, checked } = event.target;

    let updatedCatlogs = [...catalogs]; // Копируем текущее состояние

    if (checked) {
      updatedCatlogs.push(Number(value)); // Добавляем ID каталога
    } else {
      updatedCatlogs = updatedCatlogs.filter((log) => log !== Number(value)); // Удаляем ID каталога
    }

    setCatalogs(updatedCatlogs); // Обновляем состояние
  };

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
      images: [...prevProduct.images, ...files],
    }));
  };

  const handleRemoveImage = (image) => {
    setProduct((prevProduct) => ({
      ...prevProduct,
      images: prevProduct.images.filter((img) => img !== image),
      del_images: [
        ...prevProduct.del_images,
        {
          id: image.id,
          name: image.name,
        },
      ],
    }));
    setDelImages((prevDelImages) => ({
      ...prevDelImages,
      [image.id]: true, // Помечаем изображение как удалённое
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

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
      price: product.price,
      catalogs: catalogs,
      del_images: product.del_images,
    };
    console.log(productData);

    const formData = new FormData();
    formData.append("json", JSON.stringify(productData));
    product.images.forEach((file) => {
      formData.append("files", file);
    });

    await updateProduct(id, formData);
    setLoading(false);
  };

  return (
    <Box sx={{ mt: 5, mb: 5 }}>
      <Container maxWidth="md">
        <Paper elevation={3} sx={{ p: 3 }}>
          <Typography variant="h4" align="center" gutterBottom>
            Редактирование продукта
          </Typography>
          <Box component="form" onSubmit={handleSubmit}>
            <Grid container spacing={3}>
              {/* Основная информация */}
              <Grid item xs={12}>
                <TextField
                  // label="Название"
                  value={product.name}
                  onChange={(e) =>
                    setProduct({ ...product, name: e.target.value })
                  }
                  fullWidth
                  margin="normal"
                  placeholder={products.data?.name}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  // label="Артикул"
                  value={product.article}
                  onChange={(e) =>
                    setProduct({ ...product, article: e.target.value })
                  }
                  fullWidth
                  margin="normal"
                  placeholder={products.data?.article}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  label="Цена"
                  value={product.price}
                  onChange={(e) => {
                    const priceValue = parseFloat(e.target.value);
                    setProduct({
                      ...product,
                      price: isNaN(priceValue) ? 0 : priceValue,
                    });
                  }}
                  placeholder={products.data?.price}
                  fullWidth
                  margin="normal"
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">₽</InputAdornment>
                    ),
                  }}
                />
              </Grid>
              <Grid item xs={12}>
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
                  placeholder={products.data?.description}
                />
              </Grid>

              {/* Управление изображениями */}
              <Grid item xs={12}>
                <input
                  type="file"
                  multiple
                  onChange={handleFileChange}
                  accept="image/*"
                />
                <Box sx={{ display: "flex", flexWrap: "wrap", mt: 2 }}>
                  {products &&
                    products.data &&
                    products.data.images.map((e) => (
                      <Box
                        key={e.id}
                        sx={{ position: "relative", mr: 1, mb: 1 }}
                      >
                        <Avatar
                          src={`${urlPictures}/${e.name}`}
                          alt="preview"
                          sx={{
                            width: 100,
                            height: 100,
                            position: "relative",
                            opacity: delImages[e.id] ? 0.5 : 1, // Уменьшаем прозрачность, если изображение удалено
                            "&::before": delImages[e.id]
                              ? {
                                  content: '""',
                                  position: "absolute",
                                  top: "50%",
                                  left: 0,
                                  right: 0,
                                  height: "2px",
                                  backgroundColor: "red",
                                  transform: "rotate(-45deg)",
                                }
                              : null, // Перечеркивание, если изображение удалено
                          }}
                        />
                        <IconButton
                          onClick={() => handleRemoveImage(e)}
                          sx={{ position: "absolute", top: 0, right: 0 }}
                        >
                          <DeleteIcon />
                        </IconButton>
                      </Box>
                    ))}
                </Box>
              </Grid>

              {/* Категории */}
              <Grid item xs={12}>
                <Typography variant="h6">Категории</Typography>
                <Box sx={{ display: "flex", flexWrap: "wrap" }}>
                  {Array.isArray(category.data) && category.data.length > 0 ? (
                    category.data.map((item) => (
                      <FormControlLabel
                        key={item.id}
                        control={
                          <Checkbox
                            checked={selectedCategories.includes(item.id)}
                            onChange={() => handleCheckboxChange(item.id)}
                          />
                        }
                        label={item.name}
                      />
                    ))
                  ) : (
                    <p>Данных нет</p>
                  )}
                </Box>
              </Grid>

              <Grid item xs={12}>
                <label>
                  Каталог
                  <Checkbox
                    value={1}
                    onChange={(event) => handleCatalogChange(event)}
                  />
                </label>
                <label>
                  Каталог по электроному сертификату
                  <Checkbox
                    value={2}
                    onChange={(event) => handleCatalogChange(event)}
                  />
                </label>
              </Grid>

              {/* Характеристики */}
              <Grid item xs={12}>
                <Typography variant="h6">Характеристики</Typography>
                {Array.isArray(characteristics) &&
                  characteristics.map((char) => (
                    <Box key={char.id} sx={{ mb: 2 }}>
                      <Typography>{char.name}:</Typography>
                      {char.data_type === "bool" ? (
                        <Box>
                          <FormControlLabel
                            control={
                              <Checkbox
                                checked={characteristicValues[char.id] === true}
                                onChange={() =>
                                  handleValueChange(char.id, true)
                                }
                              />
                            }
                            label="Да"
                          />
                          <FormControlLabel
                            control={
                              <Checkbox
                                checked={
                                  characteristicValues[char.id] === false
                                }
                                onChange={() =>
                                  handleValueChange(char.id, false)
                                }
                              />
                            }
                            label="Нет"
                          />
                        </Box>
                      ) : (
                        <TextField
                          label={`Значение для ${char.name}`}
                          value={characteristicValues[char.id] || ""}
                          onChange={(e) =>
                            handleValueChange(char.id, e.target.value)
                          }
                          fullWidth
                          margin="normal"
                        />
                      )}
                    </Box>
                  ))}
              </Grid>
            </Grid>

            {/* Кнопки управления */}
            <Box
              sx={{ display: "flex", justifyContent: "space-between", mt: 3 }}
            >
              <Button
                variant="outlined"
                color="secondary"
                onClick={() =>
                  setProduct({
                    article: "",
                    category_ids: [],
                    characteristic_values: [],
                    description: "",
                    name: "",
                    images: [],
                    price: 0,
                  })
                }
              >
                Сбросить
              </Button>
              <Button
                type="submit"
                variant="contained"
                color="primary"
                disabled={loading}
              >
                {loading ? (
                  <CircularProgress size={24} />
                ) : (
                  "Сохранить изменения"
                )}
              </Button>
            </Box>
          </Box>
        </Paper>
      </Container>
    </Box>
  );
}
