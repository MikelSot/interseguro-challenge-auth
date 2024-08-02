FROM golang:1.22
LABEL authors="miguel soto"

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

# Comando para iniciar la aplicación
CMD ["./main"]
