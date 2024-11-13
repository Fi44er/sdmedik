#!/bin/bash

# Папки
MODEL_DIR="../internal/model"
REPO_DIR="../internal/repository"

# Проверка существования папки model
if [ ! -d "$MODEL_DIR" ]; then
  echo "Папка $MODEL_DIR не найдена!"
  exit 1
fi

# Создание структуры каталогов в repository
for model_file in "$MODEL_DIR"/*.go; do
  # Извлечение имени файла без расширения
  model_name=$(basename "$model_file" .go)

if [ ! -d "$REPO_DIR/$model_name" ]; then
  mkdir "$REPO_DIR/$model_name"
fi

if [ ! -d "$REPO_DIR/$model_name/converter" ]; then
  mkdir "$REPO_DIR/$model_name/converter"
fi

if [ ! -d "$REPO_DIR/$model_name/model" ]; then
  mkdir "$REPO_DIR/$model_name/model"
fi
  
if [ ! -f "$REPO_DIR/$model_name/repository.go" ]; then 
    cat <<EOL > "$REPO_DIR/$model_name/repository.go"
package $model_name

// Здесь будет написана ллогика работы с бд
EOL
fi

if [ ! -f "$REPO_DIR/$model_name/converter/$model_name.go" ]; then
    cat <<EOL > "$REPO_DIR/$model_name/converter/$model_name.go"
package converter

// Здесь будет логика конвертации данных
EOL
fi

if [ ! -f "$REPO_DIR/$model_name/model/$model_name.go" ]; then
  # Создание файла модели в соответствующем каталоге
  cat <<EOL > "$REPO_DIR/$model_name/model/$model_name.go"
package model

// Здесь будет определение структуры для $model_name
EOL
fi

  echo "Создана структура для $model_name в $REPO_DIR/$model_name/"
done

echo "Структура репозитория успешно создана!"
