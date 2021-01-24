FROM golang:1.13-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    TZ=Asia/Seoul

WORKDIR /build

COPY go.mod go.sum main.go prd.env ./

COPY src ./src

RUN go mod download

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /dist/main .

COPY --from=builder /build/prd.env .

EXPOSE 8081

ENTRYPOINT ["/main"]