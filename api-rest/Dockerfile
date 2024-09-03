FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy
RUN go get github.com/confluentinc/confluent-kafka-go/kafka

COPY . .

RUN go build -o projeto-go

EXPOSE 8000 50051

CMD ["./projeto-go"]
