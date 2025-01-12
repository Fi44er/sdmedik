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

const SidebarFilter = () => {
  const theme = useTheme();
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [minPrice, setMinPrice] = useState(0);
  const [maxPrice, setMaxPrice] = useState(0);
  const [filtersApplied, setFiltersApplied] = useState(false);
  const { fetchFilter, filters, selectedFilters, setSelectedFilters } =
    useFilterStore();
  const { fetchProducts, products } = useProductStore();
  const { id } = useParams();
  const category_id = id;

  useEffect(() => {
    fetchFilter(category_id);
    console.log(filters);
  }, [category_id]);

  const toggleDrawer = () => {
    setDrawerOpen(!drawerOpen);
  };

  const handleCheckboxChange = (characteristicId, value) => {
    const updatedCharacteristics = filters.data.characteristics.map(
      (filter) => {
        if (filter.id === characteristicId) {
          return {
            ...filter,
            values: filter.values.map((val) =>
              val.value === value ? { ...val, checked: !val.checked } : val
            ),
          };
        }
        return filter;
      }
    );

    setSelectedFilters({
      ...filters,
      data: {
        ...filters.data,
        characteristics: updatedCharacteristics,
      },
    });

    // Принудительное обновление состояния компонента
    setFiltersApplied(true);
  };

  const handleApplyFilters = () => {
    if (!Array.isArray(filters?.data?.characteristics)) {
      console.error("Характеристики не загружены или не являются массивом");
      return;
    }

    const serializableFilters = {
      price: {
        min: minPrice,
        max: maxPrice,
      },
      characteristics: filters.data.characteristics
        .filter((filter) => filter.values.some((value) => value.checked))
        .map((filter) => ({
          characteristic_id: filter.id,
          values: filter.values
            .filter((value) => value.checked)
            .map((value) => value.value),
        })),
    };

    console.log("Данные, отправляемые на сервер:", serializableFilters);
    if (selectedFilters) {
      setSelectedFilters(serializableFilters);
    }
    fetchProducts(category_id, serializableFilters);
    toggleDrawer();
    setFiltersApplied(true);
  };

  const handleResetFilters = () => {
    setMinPrice(0);
    setMaxPrice(0);
    setFiltersApplied(false);
    // Сбросить состояние характеристик
    const resetCharacteristics = filters.data.characteristics.map((filter) => ({
      ...filter,
      values: filter.values.map((value) => ({ ...value, checked: false })),
    }));
    setSelectedFilters({
      ...filters,
      data: {
        ...filters.data,
        characteristics: resetCharacteristics,
      },
    });
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
          </Box>

          {Array.isArray(filters?.data?.characteristics) &&
          filters.data.characteristics.length > 0 ? (
            filters.data.characteristics.map((filter, index) => (
              <FormControl fullWidth sx={{ mt: 2 }} key={index}>
                <Accordion>
                  <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                    {filter.name}
                  </AccordionSummary>
                  <AccordionDetails
                    sx={{
                      maxHeight: 200,
                      overflow: "auto",
                    }}
                  >
                    {filter.values.map((value) => (
                      <FormControlLabel
                        key={value.value}
                        control={
                          <Checkbox
                            sx={{
                              color: "#00B3A4",
                              "&.Mui-checked": { color: "#00B3A4" },
                            }}
                            checked={value.checked || false}
                            onChange={() =>
                              handleCheckboxChange(filter.id, value.value)
                            }
                          />
                        }
                        label={
                          typeof value === "boolean"
                            ? value
                              ? "Да"
                              : "Нет"
                            : typeof value === "number"
                            ? `${value} см`
                            : value
                        }
                      />
                    ))}
                  </AccordionDetails>
                </Accordion>
              </FormControl>
            ))
          ) : (
            <Typography>Данных нет</Typography>
          )}
          <Box sx={{ display: "flex", justifyContent: "space-between", mt: 3 }}>
            <Button
              variant="outlined"
              sx={{
                border: "2px solid #2CC0B3",
                color: "#2CC0B3",
                height: "50px",
                width: "48%",
              }}
              onClick={handleApplyFilters}
            >
              Показать
            </Button>
            {filtersApplied && (
              <Button
                variant="contained"
                sx={{
                  background: `#c0c0c0`,
                  height: "50px",
                  color: "black",
                  width: "48%",
                }}
                onClick={handleResetFilters}
              >
                Сбросить все
              </Button>
            )}
          </Box>
        </Box>
      </Drawer>
    </Box>
  );
};

export default SidebarFilter;
