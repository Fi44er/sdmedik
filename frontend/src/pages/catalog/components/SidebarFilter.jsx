import React, { useState } from "react";
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
} from "@mui/material";
import FilterListIcon from "@mui/icons-material/FilterList";
import CloseIcon from "@mui/icons-material/Close";

const SidebarFilter = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down("sm")); // Определяем, мобильное ли устройство
  const [drawerOpen, setDrawerOpen] = useState(false);
  const [priceRange, setPriceRange] = useState([20000, 30000]);
  const [selectedOffer, setSelectedOffer] = useState("");
  const [selectedColor, setSelectedColor] = useState("");
  const [selectedBatteryType, setSelectedBatteryType] = useState("");
  const [selectedMotorPower, setSelectedMotorPower] = useState("");

  const toggleDrawer = () => {
    setDrawerOpen(!drawerOpen);
  };

  const handlePriceChange = (event, newValue) => {
    setPriceRange(newValue);
  };

  const handleApplyFilters = () => {
    // Здесь вы можете добавить логику применения фильтров
    toggleDrawer(); // Закрытие меню после применения фильтров
  };

  return (
    <Box sx={{ display: "flex" }}>
      {isMobile && (
        <IconButton onClick={toggleDrawer}>
          <FilterListIcon />
        </IconButton>
      )}

      {/* Drawer для мобильной версии */}
      <Drawer
        anchor="left"
        open={drawerOpen}
        onClose={toggleDrawer}
        sx={{ "& .MuiDrawer-paper": { width: "100vw", height: "100vh" } }} // Занимает весь экран
      >
        <Box sx={{ padding: 2, height: "100vh", position: "relative" }}>
          <IconButton
            onClick={toggleDrawer}
            sx={{ position: "absolute", right: 16, top: 16 }}
          >
            <CloseIcon />
          </IconButton>
          <Typography variant="h6">Фильтрация</Typography>
          <Typography gutterBottom>Цена</Typography>
          <Slider
            value={priceRange}
            onChange={handlePriceChange}
            valueLabelDisplay="auto"
            min={0}
            max={50000}
          />
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Наши предложения</InputLabel>
            <Select
              value={selectedOffer}
              onChange={(e) => setSelectedOffer(e.target.value)}
            >
              <MenuItem value={1}>Предложение 1</MenuItem>
              <MenuItem value={2}>Предложение 2</MenuItem>
              <MenuItem value={3}>Предложение 3</MenuItem>
            </Select>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Цвет рамы</InputLabel>
            <Select
              value={selectedColor}
              onChange={(e) => setSelectedColor(e.target.value)}
            >
              <MenuItem value="red">Красный</MenuItem>
              <MenuItem value="blue">Синий</MenuItem>
              <MenuItem value="green">Зеленый</MenuItem>
            </Select>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Тип аккумулятора</InputLabel>
            <Select
              value={selectedBatteryType}
              onChange={(e) => setSelectedBatteryType(e.target.value)}
            >
              <MenuItem value="li-ion">Li-ion</MenuItem>
              <MenuItem value="lead-acid">Свинцово-кислотный</MenuItem>
            </Select>
          </FormControl>
          <FormControl fullWidth sx={{ mt: 2 }}>
            <InputLabel>Мощность электромотора</InputLabel>
            <Select
              value={selectedMotorPower}
              onChange={(e) => setSelectedMotorPower(e.target.value)}
            >
              <MenuItem value="250w">250W</MenuItem>
              <MenuItem value="500w">500W</MenuItem>
              <MenuItem value="1000w">1000W</MenuItem>
            </Select>
          </FormControl>
          <Button variant="contained" sx={{ mt: 3 }} onClick={toggleDrawer}>
            Применить фильтры
          </Button>
        </Box>
      </Drawer>

      {/* Меню фильтров для десктопа всегда отображается */}
      <Box
        sx={{
          width: "250px",
          display: isMobile ? "none" : "block", // Скрыть на мобильной версии
          padding: 2,
        }}
      >
        <Typography variant="h6">Фильтрация</Typography>
        <Typography gutterBottom>Цена</Typography>
        <Slider
          value={priceRange}
          onChange={handlePriceChange}
          valueLabelDisplay="auto"
          min={0}
          max={50000}
        />
        <FormControl fullWidth sx={{ mt: 2 }}>
          <InputLabel>Наши предложения</InputLabel>
          <Select
            value={selectedOffer}
            onChange={(e) => setSelectedOffer(e.target.value)}
          >
            <MenuItem value={1}>Предложение 1</MenuItem>
            <MenuItem value={2}>Предложение 2</MenuItem>
            <MenuItem value={3}>Предложение 3</MenuItem>
          </Select>
        </FormControl>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <InputLabel>Цвет рамы</InputLabel>
          <Select
            value={selectedColor}
            onChange={(e) => setSelectedColor(e.target.value)}
          >
            <MenuItem value="red">Красный</MenuItem>
            <MenuItem value="blue">Синий</MenuItem>
            <MenuItem value="green">Зеленый</MenuItem>
          </Select>
        </FormControl>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <InputLabel>Тип аккумулятора</InputLabel>
          <Select
            value={selectedBatteryType}
            onChange={(e) => setSelectedBatteryType(e.target.value)}
          >
            <MenuItem value="li-ion">Li-ion</MenuItem>
            <MenuItem value="lead-acid">Свинцово-кислотный</MenuItem>
          </Select>
        </FormControl>
        <FormControl fullWidth sx={{ mt: 2 }}>
          <InputLabel>Мощность электромотора</InputLabel>
          <Select
            value={selectedMotorPower}
            onChange={(e) => setSelectedMotorPower(e.target.value)}
          >
            <MenuItem value="250w">250W</MenuItem>
            <MenuItem value="500w">500W</MenuItem>
            <MenuItem value="1000w">1000W</MenuItem>
          </Select>
        </FormControl>
        <Button variant="contained" sx={{ mt: 3 }} onClick={toggleDrawer}>
          Применить фильтры
        </Button>
      </Box>
    </Box>
  );
};

export default SidebarFilter;
