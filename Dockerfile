FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8000
# RUN CGO_ENABLED=0 GOOS=linux go build -o main
# RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

RUN go build main.go
# CMD [ "./main" ]
FROM alpine:latest 
WORKDIR /

COPY --from=build ./app/main ./main

ENTRYPOINT ["./main"]
