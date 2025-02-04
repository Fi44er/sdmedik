import React, { useEffect, useRef, useState } from "react";
import {
  InputBase,
  Button,
  Box,
  MenuItem,
  Typography,
  CircularProgress,
} from "@mui/material";
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

  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const [isSuggestionsVisible, setIsSuggestionsVisible] = useState(false); // Состояние видимости подсказок
  const inputRef = useRef(null);
  const searchBoxRef = useRef(null); // Ref для контейнера поиска

  // Обрабатываем ввод в поле поиска
  const handleSearchInput = (query) => {
    setSearchQuery(query ?? "");
    if (query.trim().length) {
      setIsLoading(true);
      setError(null);
      debouncedSearchProducts(query);
      setIsSuggestionsVisible(true); // Показываем подсказки при вводе
    } else {
      setSearchSuggestions([]);
      setIsSuggestionsVisible(false); // Скрываем подсказки, если запрос пустой
    }
  };

  // Дебаунсированная функция для выполнения поиска товаров
  const debouncedSearchProducts = useRef(
    debounce(async (query) => {
      try {
        const suggestions = await searchProducts(query); // Выполняем поиск товаров
        setSearchSuggestions(suggestions ?? []);
      } catch (error) {
        console.error("Ошибка при получении подсказок:", error);
        setError("Произошла ошибка при поиске. Пожалуйста, попробуйте снова.");
        setSearchSuggestions([]); // Установите пустой массив, если произошла ошибка
      } finally {
        setIsLoading(false);
      }
    }, DEBOUNCE_DELAY)
  ).current;

  // Обработчик клика по подсказке
  const handleSuggestionClick = (suggestion) => {
    window.location.href = `/product/${suggestion.id}`;
    setIsSuggestionsVisible(false); // Скрываем подсказки после выбора
  };

  // Обработчик клика вне компонента
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (
        searchBoxRef.current &&
        !searchBoxRef.current.contains(event.target)
      ) {
        setIsSuggestionsVisible(false); // Скрываем подсказки, если клик вне компонента
      }
    };

    // Добавляем обработчик события клика
    document.addEventListener("click", handleClickOutside);

    // Убираем обработчик при размонтировании компонента
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  }, []);

  // Отменяем предыдущие запросы при размонтаже компонента
  useEffect(() => {
    return () => {
      debouncedSearchProducts.cancel();
    };
  }, [debouncedSearchProducts]);

  return (
    <Box
      ref={searchBoxRef} // Привязываем ref к контейнеру поиска
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
        onFocus={() => setIsSuggestionsVisible(true)} // Показываем подсказки при фокусе
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
      {isSuggestionsVisible && searchQuery && (
        <Box
          sx={{
            position: "absolute",
            top: "60px",
            left: 0,
            width: "100%",
            backgroundColor: "white",
            border: "1px solid #e0e0e0",
            borderRadius: "8px",
            boxShadow: "0px 4px 15px rgba(0, 0, 0, 0.1)",
            zIndex: 1000,
            maxHeight: "300px",
            overflowY: "auto",
            transition: "opacity 0.3s ease, transform 0.3s ease",
            opacity: isSuggestionsVisible ? 1 : 0,
            transform: isSuggestionsVisible
              ? "translateY(0)"
              : "translateY(-10px)",
          }}
        >
          {isLoading ? (
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                padding: "15px",
              }}
            >
              <CircularProgress color="success" size={24} />
            </Box>
          ) : error ? (
            <Typography
              sx={{
                padding: "15px 20px",
                color: "error.main",
                fontFamily: "Arial, sans-serif",
              }}
            >
              {error}
            </Typography>
          ) : (searchSuggestions ?? []).length > 0 ? (
            searchSuggestions.map((suggestion, index) => (
              <MenuItem
                key={index}
                onClick={() => handleSuggestionClick(suggestion)}
                sx={{
                  padding: "12px 20px",
                  display: "flex",
                  alignItems: "center",
                  gap: "10px",
                  "&:hover": { backgroundColor: "#f5f5f5" },
                  transition: "background-color 0.2s ease",
                }}
              >
                <Box>
                  <Typography
                    variant="body1"
                    sx={{
                      fontWeight: 500,
                      color: "black",
                      fontFamily: "Arial, sans-serif",
                    }}
                  >
                    {suggestion.name}
                  </Typography>
                  {suggestion.description && (
                    <Typography
                      variant="body2"
                      sx={{
                        color: "text.secondary",
                        fontFamily: "Arial, sans-serif",
                      }}
                    >
                      {suggestion.description}
                    </Typography>
                  )}
                </Box>
              </MenuItem>
            ))
          ) : (
            <Typography
              sx={{
                padding: "15px 20px",
                color: "text.secondary",
                fontFamily: "Arial, sans-serif",
              }}
            >
              Ничего не найдено
            </Typography>
          )}
        </Box>
      )}
    </Box>
  );
}
