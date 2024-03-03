FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o news-api .

EXPOSE 8080

CMD ["./news-api"]