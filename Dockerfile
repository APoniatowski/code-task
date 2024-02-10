# Builder
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o main .


# Actual image
FROM alpine:latest

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
