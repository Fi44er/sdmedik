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
import { Delete as DeleteIcon } from "@mui/icons-material";
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
  const [isFetching, setIsFetching] = useState(true);

  // Загрузка категорий и данных продукта
  useEffect(() => {
    fetchCategory();
    fetchProductById(id).then(() => setIsFetching(false));
  }, [id, fetchCategory, fetchProductById]);

  // Инициализация формы данными из products.data
  useEffect(() => {
    if (products.data && !isFetching) {
      setProduct({
        article: products.data.article || "",
        category_ids: products.data.categories?.map((cat) => cat.id) || [],
        characteristic_values: products.data.characteristic || [],
        description: products.data.description || "",
        name: products.data.name || "",
        images: products.data.images || [],
        price: products.data.price || 0,
        del_images: [],
      });

      setSelectedCategories(
        products.data.categories?.map((cat) => cat.id) || []
      );
      setCatalogs([products.data.catalogs] || []);

      // Инициализация характеристик
      const initialCharValues = {};
      products.data.characteristic?.forEach((char) => {
        initialCharValues[char.id] = char.value;
      });
      setCharacteristicValues(initialCharValues);

      // Инициализация характеристик категорий
      const selectedCats = products.data.categories?.map((cat) => cat.id) || [];
      const allCharacteristics = category.data
        ?.filter((cat) => selectedCats.includes(cat.id))
        .flatMap((cat) => cat.characteristic || []);
      setCharacteristics([...new Set(allCharacteristics)] || []);
    }
  }, [products.data, isFetching, category.data]);

  const handleCatalogChange = (event) => {
    const { value, checked } = event.target;
    let updatedCatalogs = [...catalogs];
    if (checked) {
      updatedCatalogs.push(Number(value));
    } else {
      updatedCatalogs = updatedCatalogs.filter((log) => log !== Number(value));
    }
    setCatalogs(updatedCatalogs);
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
        { id: image.id, name: image.name },
      ],
    }));
    setDelImages((prevDelImages) => ({
      ...prevDelImages,
      [image.id]: true,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    const productData = {
      // article: product.article,
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

    const formData = new FormData();
    formData.append("json", JSON.stringify(productData));
    product.images.forEach((file) => {
      if (file instanceof File) {
        formData.append("files", file);
      }
    });

    await updateProduct(id, formData);
    setLoading(false);
  };

  if (isFetching) {
    return (
      <Box sx={{ display: "flex", justifyContent: "center", mt: 5 }}>
        <CircularProgress />
      </Box>
    );
  }

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
                  label="Название"
                  value={product.name}
                  onChange={(e) =>
                    setProduct({ ...product, name: e.target.value })
                  }
                  fullWidth
                  margin="normal"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  label="Артикул"
                  value={product.article}
                  onChange={(e) =>
                    setProduct({ ...product, article: e.target.value })
                  }
                  fullWidth
                  margin="normal"
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
                  {product.images.map((image, index) => (
                    <Box
                      key={image.id || index}
                      sx={{ position: "relative", mr: 1, mb: 1 }}
                    >
                      <Avatar
                        src={
                          image instanceof File
                            ? URL.createObjectURL(image)
                            : `${urlPictures}/${image.name}`
                        }
                        alt="preview"
                        sx={{
                          width: 100,
                          height: 100,
                          position: "relative",
                          opacity: delImages[image.id] ? 0.5 : 1,
                          "&::before": delImages[image.id]
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
                            : null,
                        }}
                      />
                      <IconButton
                        onClick={() => handleRemoveImage(image)}
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
                    checked={catalogs.includes(1)}
                    onChange={handleCatalogChange}
                  />
                </label>
                <label>
                  Каталог по электронному сертификату
                  <Checkbox
                    value={2}
                    checked={catalogs.includes(2)}
                    onChange={handleCatalogChange}
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
                                checked={
                                  characteristicValues[char.id] === "true" ||
                                  characteristicValues[char.id] === true
                                }
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
                                  characteristicValues[char.id] === "false" ||
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

            <Box
              sx={{ display: "flex", justifyContent: "space-between", mt: 3 }}
            >
              <Button
                variant="outlined"
                color="secondary"
                onClick={() => {
                  setProduct({
                    article: products.data?.article || "",
                    category_ids:
                      products.data?.categories?.map((cat) => cat.id) || [],
                    characteristic_values: products.data?.characteristic || [],
                    description: products.data?.description || "",
                    name: products.data?.name || "",
                    images: products.data?.images || [],
                    price: products.data?.price || 0,
                    del_images: [],
                  });
                  setSelectedCategories(
                    products.data?.categories?.map((cat) => cat.id) || []
                  );
                  setCatalogs([products.data?.catalogs] || []);
                  const initialCharValues = {};
                  products.data?.characteristic?.forEach((char) => {
                    initialCharValues[char.id] = char.value;
                  });
                  setCharacteristicValues(initialCharValues);
                }}
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
