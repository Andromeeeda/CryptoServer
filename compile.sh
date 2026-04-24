#!/bin/bash

# Проверяем Go проект в ТВОЕЙ структуре (internal/cmd/server)
if [ -f "go.mod" ] && [ -d "internal/cmd/server" ]; then
    echo "Компиляция Go crypto сервера..."
    cd internal/cmd/server
    go build -o ../../../cryptoserver .
    cd ../../..
    echo "Компиляция cryptoserver завершена"
    exit 0
fi

# Проверяем стандартную структуру cmd/server
if [ -f "go.mod" ] && [ -d "cmd/server" ]; then
    echo "Компиляция Go crypto сервера..."
    cd cmd/server
    go build -o ../../cryptoserver .
    cd ../..
    echo "Компиляция cryptoserver завершена"
    exit 0
fi

# Старые варианты для совместимости
if [ -f "cryptoserver.go" ]; then
    echo "Компиляция Go crypto сервера..."
    go build -o cryptoserver cryptoserver.go
    exit 0
fi

if [ -f "cryptoserver.py" ]; then
    echo "Python не требует компиляции"
    exit 0
fi

if [ -f "cryptoserver.js" ]; then
    echo "JavaScript не требует компиляции"
    exit 0
fi

if [ -f "cryptoserver.cpp" ]; then
    echo "Компиляция C++ crypto сервера..."
    g++ -o cryptoserver cryptoserver.cpp -lcurl -lpthread -std=c++11
    exit 0
fi

if [ -f "cryptoserver.java" ]; then
    echo "Компиляция Java crypto сервера..."
    javac cryptoserver.java
    exit 0
fi

echo "Не найден файл cryptoserver для компиляции"
echo "Поддерживаемые файлы: cryptoserver.{cpp,go,py,js,java}"
echo "Или стандартная структура: cmd/server/main.go"
exit 1