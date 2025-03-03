import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  Drawer,
  IconButton,
  Typography,
  FormControlLabel,
  Checkbox,
  Accordion,
  AccordionSummary,
  AccordionDetails,
  TextField,
  Slider,
  styled,
  useMediaQuery,
  useTheme,
} from "@mui/material";
import FilterListIcon from "@mui/icons-material/FilterList";
import CloseIcon from "@mui/icons-material/Close";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import useFilterStore from "../../../store/filterStore";
import { useParams } from "react-router-dom";
import useProductStore from "../../../store/productStore";

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
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [minPrice, setMinPrice] = useState(0);
  const [maxPrice, setMaxPrice] = useState(100000); // Максимальная цена по умолчанию
  const { fetchFilter, filters } = useFilterStore();
  const { fetchProducts } = useProductStore();
  const [selectedValues, setSelectedValues] = useState([]);
  const { id } = useParams();
  const category_id = id;
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down("sm"));

  useEffect(() => {
    fetchFilter(category_id);
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

  const handlePriceChange = (event, newValue) => {
    setMinPrice(newValue[0]);
    setMaxPrice(newValue[1]);
  };

  const handleApplyFilters = async () => {
    const filterData = {
      price: {
        min: minPrice,
        max: maxPrice,
      },
      characteristics: selectedValues
        .filter((characteristic) => characteristic.values.length > 0)
        .map((characteristic) => ({
          characteristic_id: characteristic.characteristic_id,
          values: characteristic.values.map((value) => value.toString()),
        })),
    };

    const jsonData = JSON.stringify(filterData);
    fetchProducts(category_id, jsonData);
    toggleDrawer();
  };

  const handleResetFilters = () => {
    setSelectedValues([]);
    setMinPrice(0);
    setMaxPrice(100000);
    fetchProducts(category_id, null);
    toggleDrawer();
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
            "&:hover": {
              backgroundColor: "#009B8A",
            },
          }}
          onClick={toggleDrawer}
        >
          Фильтрация
          <FilterListIcon sx={{ ml: 1 }} />
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
            overflowY: "auto",
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
            sx={{ fontWeight: "bold", color: "#00B3A4", mb: 2 }}
          >
            Фильтрация
          </Typography>

          {/* Ценовой диапазон */}
          <Box sx={{ mb: 3 }}>
            <Typography variant="body1" sx={{ fontWeight: "bold", mb: 1 }}>
              Цена
            </Typography>
            <Slider
              value={[minPrice, maxPrice]}
              onChange={handlePriceChange}
              valueLabelDisplay="auto"
              min={0}
              max={100000}
              sx={{ color: "#00B3A4" }}
            />
            <Box
              sx={{
                display: "flex",
                justifyContent: "space-between",
                gap: 2,
                mt: 2,
              }}
            >
              <CustomTextField
                variant="outlined"
                placeholder="От"
                value={minPrice}
                onChange={(e) => setMinPrice(Number(e.target.value))}
                sx={{ width: "48%" }}
              />
              <CustomTextField
                variant="outlined"
                placeholder="До"
                value={maxPrice}
                onChange={(e) => setMaxPrice(Number(e.target.value))}
                sx={{ width: "48%" }}
              />
            </Box>
          </Box>

          {/* Фильтры по характеристикам */}
          {filters &&
            filters.data &&
            filters.data.characteristics &&
            filters.data.characteristics.length > 0 ? (
            filters.data.characteristics.map((char) => (
              <Accordion key={char.id} defaultExpanded sx={{ mb: 2 }}>
                <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                  <Typography sx={{ fontWeight: "bold" }}>{char.name}</Typography>
                </AccordionSummary>
                <AccordionDetails>
                  {char.values.map((value) => (
                    <FormControlLabel
                      key={`${char.id}-${value}`}
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
                          <>{value ? "Есть" : "Нет"}</>
                        ) : (
                          value
                        )
                      }
                    />
                  ))}
                </AccordionDetails>
              </Accordion>
            ))
          ) : (
            <Typography>Нет доступных фильтров</Typography>
          )}

          {/* Кнопки "Применить" и "Сбросить" */}
          <Box
            sx={{
              display: "flex",
              justifyContent: "space-between",
              gap: 2,
              mt: 3,
              position: "sticky",
              bottom: 0,
              backgroundColor: "#fff",
              padding: 2,
              boxShadow: "0px -2px 4px rgba(0, 0, 0, 0.1)",
            }}
          >
            <Button
              sx={{
                background: "#00B3A4",
                color: "white",
                flex: 1,
                "&:hover": {
                  backgroundColor: "#009B8A",
                },
              }}
              onClick={handleApplyFilters}
            >
              Применить
            </Button>
            <Button
              sx={{
                background: "#E74C3C",
                color: "white",
                flex: 1,
                "&:hover": {
                  backgroundColor: "#C0392B",
                },
              }}
              onClick={handleResetFilters}
            >
              Сбросить
            </Button>
          </Box>
        </Box>
      </Drawer>
    </Box>
  );
};

export default SidebarFilter;