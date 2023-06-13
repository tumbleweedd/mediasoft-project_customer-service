FROM golang:1.19.5

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

