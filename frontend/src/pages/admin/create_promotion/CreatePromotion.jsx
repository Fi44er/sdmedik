import React, { useState } from "react";
import {
  Button,
  TextField,
  Typography,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Container,
} from "@mui/material";
import usePromotionStore from "../../../store/promotionStore";

export default function CreatePromotion() {
  const { createPromotion } = usePromotionStore();

  // Состояния для каждого поля формы
  const [type, setType] = useState("product_discount");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [targetId, setTargetId] = useState("");
  const [conditionType, setConditionType] = useState("min_quantity");
  const [conditionValue, setConditionValue] = useState("");
  const [rewardType, setRewardType] = useState("percentage");
  const [rewardValue, setRewardValue] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    // Формирование объекта для отправки
    const payload = {
      condition: {
        type: conditionType,
        value: conditionValue,
      },
      description: description,
      end_date: endDate,
      name: name,
      reward: {
        type: rewardType,
        value: parseFloat(rewardValue), // Приведение к числу
      },
      start_date: startDate,
      target_id: targetId,
      type: type,
    };

    // Вызов функции createPromotion с объектом payload
    await createPromotion(payload);
  };

  return (
    <Container maxWidth="sm" sx={{ mt: 5 }}>
      <Typography variant="h4" sx={{ mb: 3, textAlign: "center" }}>
        Создать Акцию
      </Typography>
      <form onSubmit={handleSubmit}>
        <FormControl fullWidth sx={{ mb: 2 }}>
          <InputLabel>Тип акции</InputLabel>
          <Select
            value={type}
            onChange={(e) => setType(e.target.value)}
            required
          >
            <MenuItem value="product_discount">Скидка на товар</MenuItem>
            <MenuItem value="category_discount">Скидка на категорию</MenuItem>
            <MenuItem value="buy_n_get_m">Купи N, получи M</MenuItem>
          </Select>
        </FormControl>
        <Typography variant="h4" sx={{ mb: 3, textAlign: "center" }}>
          Условия и вознаграждения
        </Typography>
        <TextField
          fullWidth
          label="Название"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
          sx={{ mb: 2 }}
        />
        <TextField
          fullWidth
          label="Описание"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
          sx={{ mb: 2 }}
        />
        <TextField
          fullWidth
          label="ID товара или категории"
          value={targetId}
          onChange={(e) => setTargetId(e.target.value)}
          required
          sx={{ mb: 2 }}
        />
        <FormControl fullWidth sx={{ mb: 2 }}>
          <InputLabel>Тип условия</InputLabel>
          <Select
            value={conditionType}
            onChange={(e) => setConditionType(e.target.value)}
            required
          >
            <MenuItem value="min_quantity">Минимальное количество</MenuItem>
            <MenuItem value="buy_n">Купить энное количество товара</MenuItem>
          </Select>
        </FormControl>
        <TextField
          fullWidth
          label="Минимальное количество"
          type="number"
          value={conditionValue}
          onChange={(e) => setConditionValue(e.target.value)}
          required
          sx={{ mb: 2 }}
        />
        <FormControl fullWidth sx={{ mb: 2 }}>
          <InputLabel>Тип вознаграждения</InputLabel>
          <Select
            value={rewardType}
            onChange={(e) => setRewardType(e.target.value)}
            required
          >
            <MenuItem value="percentage">Процент</MenuItem>
 <MenuItem value="fixed">Фиксированная сумма</MenuItem>
          </Select>
        </FormControl>
        <TextField
          fullWidth
          label="Значение вознаграждения"
          type="number"
          value={rewardValue}
          onChange={(e) => setRewardValue(e.target.value)}
          required
          sx={{ mb: 2 }}
        />
        <TextField
          fullWidth
          label="Дата начала"
          type="date"
          value={startDate}
          onChange={(e) => setStartDate(e.target.value)}
          required
          sx={{ mb: 2 }}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <TextField
          fullWidth
          label="Дата окончания"
          type="date"
          value={endDate}
          onChange={(e) => setEndDate(e.target.value)}
          required
          sx={{ mb: 2 }}
          InputLabelProps={{
            shrink: true,
          }}
        />
        <Button type="submit" variant="contained" color="primary">
          Создать Акцию
        </Button>
      </form>
    </Container>
  );
}