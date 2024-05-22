# Используйте официальный образ golang в качестве базового
FROM golang:1.22.3

# Установите рабочую директорию внутри контейнера
WORKDIR /app

# Скопируйте go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загрузите зависимости
RUN go mod download

# Скопируйте остальные файлы в рабочую директорию
COPY . .

# Соберите приложение
RUN go build -o warehouse-api ./cmd/warehouse-api

# Запустите приложение
CMD ["./warehouse-api"]