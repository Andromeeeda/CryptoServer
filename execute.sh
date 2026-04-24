#!/bin/bash

# 1. Пробуем запустить скомпилированный бинарник в корне
if [ -f "cryptoserver" ]; then
    echo "Запуск скомпилированного crypto сервера..."
    exec ./cryptoserver
fi

# 2. Пробуем запустить через go run (твоя структура)
if [ -f "go.mod" ] && [ -d "internal/cmd/server" ]; then
    echo "Запуск Go crypto сервера через go run..."
    exec go run ./internal/cmd/server
fi

# 3. Пробуем стандартную структуру cmd/server
if [ -f "go.mod" ] && [ -d "cmd/server" ]; then
    echo "Запуск Go crypto сервера через go run..."
    exec go run ./cmd/server
fi

# 4. Пробуем одиночный файл cryptoserver.go
if [ -f "cryptoserver.go" ]; then
    echo "Запуск Go crypto сервера..."
    exec go run cryptoserver.go
fi

# 5. Для Python
if [ -f "cryptoserver.py" ]; then
    echo "Запуск Python crypto сервера..."
    exec python3 cryptoserver.py
fi

# 6. Для Node.js
if [ -f "cryptoserver.js" ]; then
    echo "Запуск Node.js crypto сервера..."
    exec node cryptoserver.js
fi

# 7. Для Java
if [ -f "cryptoserver.class" ]; then
    echo "Запуск Java crypto сервера..."
    exec java cryptoserver
fi

echo "Не найден исполняемый файл crypto сервера"
echo "Убедитесь что файл скомпилирован или существует internal/cmd/server/main.go"
exit 1