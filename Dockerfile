FROM golang:alpine
WORKDIR /app
COPY . .
EXPOSE 8000
RUN CGO_ENABLED=0 GOOS=linux go build -o main
COPY .env /app
CMD [ "./main" ]
