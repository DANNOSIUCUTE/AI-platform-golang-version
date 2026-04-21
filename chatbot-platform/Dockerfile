FROM golang:1.21-alpine

# Cài đặt thư viện cần thiết trên alpine
RUN apk add --no-cache build-base

WORKDIR /app

# Sao chép các tệp quản lý dependency
COPY go.mod go.sum ./

# Tải xuống thư viện
RUN go mod download

# Sao chép toàn bộ mã nguồn theo cấu trúc mới
COPY . .

# Build mã nguồn thành một file thực thi (Binary)
RUN go build -o main .

# Mở port 3000 ra môi trường ngoài
EXPOSE 3000

# Chạy ứng dụng
CMD ["./main"]
