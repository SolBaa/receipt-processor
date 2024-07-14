FROM golang:1.17-alpine AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos go.mod y go.sum y descargarlos
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto de los archivos de la aplicación
COPY . .

# Construir la aplicación
RUN go build -o main ./cmd/server

# Etapa de producción
FROM alpine:latest

# Crear un directorio en el contenedor
WORKDIR /root/

# Copiar el binario de la etapa de construcción
COPY --from=builder /app/main .

# Exponer el puerto en el que corre la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]