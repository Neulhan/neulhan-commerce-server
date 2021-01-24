FROM golang

WORKDIR /app

COPY go.mod go.sum main.go prd.env ./

COPY src ./src

RUN go mod download

RUN go build -o main

ENV GIN_MODE=release

EXPOSE 8081

CMD ["./main"]
