import React, { useEffect, useRef, useState } from "react";
import { InputBase, Button, Box, MenuItem, Typography } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import { debounce } from "lodash";
import useSearchStore from "../../store/serchStore";

const DEBOUNCE_DELAY = 250; // Задержка перед выполнением запроса

export default function Search() {
  const {
    searchQuery,
    searchSuggestions,
    setSearchQuery,
    setSearchSuggestions,
    searchProducts,
  } = useSearchStore();

  const inputRef = useRef(null);

  // Обрабатываем ввод в поле поиска
  const handleSearchInput = (query) => {
    setSearchQuery(query === null || query === undefined ? "" : query); // Устанавливаем пустую строку, если query равен null или undefined
    if ((query || "").trim().length) {
      // Используем дебаунс для оптимизации запросов
      debouncedSearchProducts(query);
    }
  };

  // Дебаунсированная функция для выполнения поиска товаров
  const debouncedSearchProducts = useRef(
    debounce(async (query) => {
      try {
        const suggestions = await searchProducts(query); // Выполняем поиск товаров
        setSearchSuggestions(suggestions); // Обновляем подсказки
      } catch (error) {
        console.error("Ошибка при получении подсказок:", error);
      }
    }, DEBOUNCE_DELAY)
  ).current;

  // Обработчик клика по подсказке
  const handleSuggestionClick = (suggestion) => {
    window.location.href = `/product/${suggestion.id}`;
  };

  // Отменяем предыдущие запросы при размонтаже компонента
  useEffect(
    () => () => {
      debouncedSearchProducts.cancel();
    },
    []
  );

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        width: "100%",
        maxWidth: "500px",
        position: "relative",
      }}
    >
      {/* Поле поиска */}
      <InputBase
        ref={inputRef}
        type="text"
        placeholder="Поиск по товарам"
        value={searchQuery}
        sx={{
          height: "53px",
          width: "100%",
          border: "2px solid #87EBEB",
          borderRight: "none",
          paddingLeft: "20px",
          fontSize: "16px",
          outline: "none",
          backgroundColor: "#FAFAFA",
        }}
        onChange={(e) => handleSearchInput(e.target.value)}
      />

      {/* Кнопка поиска */}
      <Button
        variant="contained"
        sx={{
          height: "53px",
          borderTopLeftRadius: "0",
          borderBottomLeftRadius: "0",
          borderTopRightRadius: "10px",
          borderBottomRightRadius: "10px",
          backgroundColor: "#00B3A4",
          "&:hover": {
            backgroundColor: "#009688",
          },
        }}
      >
        <SearchIcon fontSize="large" />
      </Button>

      {/* Выпадающий список с подсказками */}
      <Box
        sx={{
          position: "absolute",
          top: "60px",
          left: 0,
          width: "100%",
          backgroundColor: "white",
          border: "1px solid #ddd",
          borderRadius: "5px",
          boxShadow: "0px 4px 10px rgba(0, 0, 0, 0.1)",
          zIndex: 1000,
          maxHeight: "300px",
          overflowY: "auto",
          transition: "opacity 0.3s ease, transform 0.3s ease",
          // opacity: searchSuggestions.length,
          // transform: searchSuggestions.length
          //   ? "translateY(0)"
          //   : "translateY(-10px)",
          // pointerEvents: searchSuggestions.length ? "auto" : "none",
        }}
      >
        {searchSuggestions &&
          searchSuggestions.map((suggestion, index) => (
            <MenuItem
              key={index}
              onClick={() => handleSuggestionClick(suggestion)}
              sx={{
                padding: "10px 20px",
                display: "flex",
                alignItems: "center",
                gap: "10px",
                "&:hover": { backgroundColor: "#f5f5f5" },
              }}
            >
              <Box>
                <Typography
                  variant="body1"
                  sx={{ fontWeight: 500, color: "black" }}
                >
                  {suggestion.name}
                </Typography>
              </Box>
            </MenuItem>
          ))}
      </Box>
    </Box>
  );
}
