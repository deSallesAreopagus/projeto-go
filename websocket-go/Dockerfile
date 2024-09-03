FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o projeto-go

EXPOSE 8080 50052

CMD ["./projeto-go"]
