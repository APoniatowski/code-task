# Builder
FROM golang:alpine as builder

ARG VERSION

WORKDIR /go/src/app

COPY . .

RUN mkdir -p /go/bin/

RUN go mod tidy
RUN go get -d -v ./...
RUN go build -o /go/bin/task -v ./cmd/code-task/code-task.go


# Actual image
FROM alpine:latest

COPY --from=builder /go/bin/task /task

ENV HOSTNAME=$HOSTNAME
ENV POSTGRES_USER=$POSTGRES_USER
ENV POSTGRES_PASSWORD=$POSTGRES_PASSWORD
ENV POSTGRES_DB=$POSTGRES_DB

ARG VERSION
LABEL Name=code-task Version=VERSION

ENTRYPOINT ["/task"]
EXPOSE 8080
