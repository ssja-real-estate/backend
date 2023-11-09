FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8000

RUN go build main.go
FROM alpine:latest 
WORKDIR /

COPY --from=build ./app/main ./main

ENTRYPOINT ["./main"]
