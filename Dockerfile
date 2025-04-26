# FROM golang:alpine AS build
# WORKDIR /app
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY . .
# EXPOSE 8000

# RUN go build main.go
# FROM alpine:latest 
# WORKDIR /app

# COPY --from=build ./app/main ./app/main

# ENTRYPOINT ["./app/main"]


# مرحله Build
FROM golang:alpine AS build
WORKDIR /app

# کپی فایل‌های وابستگی
COPY go.mod .
COPY go.sum .
RUN go mod download

# کپی کد پروژه
COPY . .

# باز کردن پورت (اختیاری در build مرحله)
EXPOSE 8000

# ساخت فایل اجرایی
RUN go build -o main main.go

# مرحله Production
FROM alpine:latest 
WORKDIR /app

# اضافه کردن یک کاربر غیر root
RUN adduser -D -g '' appuser

# کپی فایل اجرایی از مرحله build
COPY --from=build /app/main /app/main

# تغییر مالکیت به کاربر جدید
RUN chown -R appuser /app

# اجرای برنامه با کاربر غیر root
USER appuser

# اجرای برنامه
ENTRYPOINT ["./app/main"]
