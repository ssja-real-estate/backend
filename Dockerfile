FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8000

RUN go build main.go
FROM alpine:latest 
WORKDIR /app

COPY --from=build ./app/main ./app/main

ENTRYPOINT ["./app/main"]
