import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  Drawer,
  IconButton,
  Slider,
  Typography,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  useMediaQuery,
  useTheme,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  FormControlLabel,
  Checkbox,
  TextField,
  styled,
  Paper,
} from "@mui/material";
import FilterListIcon from "@mui/icons-material/FilterList";
import CloseIcon from "@mui/icons-material/Close";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import useFilterStore from "../../../store/filterStore";
import { useParams } from "react-router-dom";
import useProductStore from "../../../store/productStore";
import stringify from "json-stringify-safe";

const CustomTextField = styled(TextField)({
  "& .MuiOutlinedInput-root": {
    "& fieldset": {
      borderColor: "#26BDB8",
    },
    "&:hover fieldset": {
      borderColor: "#26BDB8",
    },
    "&.Mui-focused fieldset": {
      borderColor: "#26BDB8",
    },
  },
});

const SidebarFilter = ({ setFilters }) => {
  const theme = useTheme();
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [minPrice, setMinPrice] = useState(0);
  const [maxPrice, setMaxPrice] = useState(0);
  const [filtersApplied, setFiltersApplied] = useState(false);
  const { fetchFilter, filters } = useFilterStore();
  const { fetchProducts, products } = useProductStore();
  const [selectedValues, setSelectedValues] = useState([]);
  const { id } = useParams();
  const category_id = id;

  useEffect(() => {
    fetchFilter(category_id);
    console.log(filters);
  }, [category_id]);

  useEffect(() => {
    if (drawerOpen && filters && filters.data && filters.data.characteristics) {
      const initialCharacteristics = filters.data.characteristics.map(
        (filter) => ({
          characteristic_id: filter.id,
          values: [],
        })
      );
      setSelectedValues(initialCharacteristics);
    }
  }, [drawerOpen, filters]);

  const toggleDrawer = () => {
    setDrawerOpen(!drawerOpen);
  };

  const handleChangeCheckbox = (event, characteristicId, value) => {
    const updatedSelectedValues = [...selectedValues];
    const index = updatedSelectedValues.findIndex(
      (item) => item.characteristic_id === characteristicId
    );

    if (index !== -1) {
      const currentCharacteristic = updatedSelectedValues[index];
      if (event.target.checked) {
        currentCharacteristic.values.push(value);
      } else {
        currentCharacteristic.values = currentCharacteristic.values.filter(
          (v) => v !== value
        );
      }
      setSelectedValues(updatedSelectedValues);
    }
  };

  const handleApplyFilters = async () => {
    // Сбор данных фильтрации
    const filterData = {
      price: {
        min: minPrice,
        max: maxPrice,
      },
      characteristics: selectedValues.filter((characteristic) =>
        characteristic.values.length > 0 ? true : false
      ),
    };

    // Оборачивание каждого значения в объект
    const cleanedFilterData = {
      ...filterData,
      characteristics: filterData.characteristics.map((char) => ({
        characteristic_id: char.characteristic_id,
        values: char.values.map((val) => String(val)), // Каждое значение становится объектом
      })),
    };

    // Преобразование данных в JSON
    const jsonData = JSON.stringify(cleanedFilterData);

    // Создание FormData для отправки
    const formData = new FormData();
    formData.append("filters", jsonData);

    console.log("Результат фильтрации:", jsonData);

    fetchProducts(category_id, jsonData);
  };

  const handleResetFilters = () => {
    // Сбрасываем все состояния фильтров
    setSelectedValues([]);
    setMinPrice(0);
    setMaxPrice(0);

    // Отправляем GET-запрос без фильтров
    fetchProducts(category_id, null); // или пустой объект, если требуется
  };

  return (
    <Box sx={{ display: "flex" }}>
      <Box sx={{ mt: 5 }}>
        <Button
          sx={{
            background: "#00B3A4",
            color: "white",
            height: "50px",
            width: "150px",
          }}
          onClick={toggleDrawer}
        >
          Фильтрация
          <FilterListIcon />
        </Button>
      </Box>

      <Drawer
        anchor="left"
        open={drawerOpen}
        onClose={toggleDrawer}
        sx={{
          "& .MuiDrawer-paper": {
            width: { xs: "100%", sm: "100%", md: "350px" },
            height: "100vh",
          },
        }}
      >
        <Box sx={{ padding: 2, height: "100vh", position: "relative" }}>
          <IconButton
            onClick={toggleDrawer}
            sx={{ position: "absolute", right: 16, top: 16 }}
          >
            <CloseIcon />
          </IconButton>
          <Typography
            variant="h6"
            sx={{ fontWeight: "bold", color: "#00B3A4" }}
          >
            Фильтрация
          </Typography>
          <Box sx={{ mt: 2 }}>
            <Box sx={{ mb: 2 }}>
              <Typography variant="body">Цена</Typography>
            </Box>
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                gridGap: "20px",
              }}
            >
              <CustomTextField
                variant="outlined"
                placeholder="От"
                onChange={(e) => setMinPrice(Number(e.target.value))}
                sx={{ width: "48%", mt: 2, color: "#00B3A4" }}
              />
              <CustomTextField
                variant="outlined"
                placeholder="До"
                onChange={(e) => setMaxPrice(Number(e.target.value))}
                sx={{ width: "48%", mt: 2, color: "#00B3A4" }}
              />
            </Box>
            {filters &&
              filters.data &&
              filters.data.characteristics.map((char) => (
                <Accordion sx={{ mt: 2, mb: 2 }} key={char.id} defaultExpanded>
                  <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                    <Typography>{char.name}</Typography>
                  </AccordionSummary>
                  <AccordionDetails>
                    {char.values.map((value, index) => (
                      <FormControlLabel
                        key={`${char.id}-${value}`} // Unique key for each checkbox
                        control={
                          <Checkbox
                            sx={{
                              color: "#00B3A4",
                              "&.Mui-checked": { color: "#00B3A4" },
                            }}
                            checked={selectedValues.some(
                              (c) =>
                                c.characteristic_id === char.id &&
                                c.values.includes(value)
                            )}
                            onChange={(e) =>
                              handleChangeCheckbox(e, char.id, value)
                            }
                          />
                        }
                        label={
                          typeof value === "boolean" ? (
                            <>{value ? "Есть" : "Нету"}</>
                          ) : (
                            value
                          )
                        }
                      />
                    ))}
                  </AccordionDetails>
                </Accordion>
              ))}
          </Box>
          <Box
            sx={{
              display: "flex",
              justifyContent: "flex-end",
              alignItems: "center",
              mt: 2,
              gridGap: 30,
            }}
          >
            <Button
              sx={{
                background: "#00B3A4",
                color: "white",
                height: "50px",
                width: { xs: "30%", md: "150px" },
              }}
              onClick={() => {
                handleApplyFilters();
                toggleDrawer();
              }}
            >
              Применить фильтры
            </Button>
            <Button
              sx={{
                background: "#E74C3C",
                color: "white",
                height: "50px",
                width: { xs: "30%", md: "150px" },
              }}
              onClick={() => {
                handleResetFilters();
                toggleDrawer();
              }}
            >
              Сбросить фильтры
            </Button>
          </Box>
        </Box>
      </Drawer>
    </Box>
  );
};

export default SidebarFilter;
