# Chatbot Platform - Golang Edition

Đây là dự án nền tảng Chatbot ứng dụng mô hình của Anthropic (Claude AI), được thiết kế với chuẩn kiến trúc **Clean Architecture** (MVC Golang API) và môi trường đóng gói **Docker**.

## Chức năng chính
- **Hệ thống tài khoản:** Đăng nhập, đăng ký bảo mật qua **JWT Token**.
- **Chatbot thông minh:** Giao tiếp trực tiếp với Claude 3 Sonnet của Anthropic.
- **Lưu trữ lịch sử:** Toàn bộ cuộc hội thoại được lưu vết tự động vào MySQL.
- **Auto-migration:** Kịch bản tự tạo sơ đồ Database và cấu trúc bảng khi khởi động app.

## Kiến trúc công nghệ
* **Ngôn ngữ Core:** Golang 1.21+
* **Web Framework:** [Gin-gonic](https://gin-gonic.com/) (Routing, Middleware siêu tốc)
* **ORM Database:** [GORM](https://gorm.io/) + MySQL 8.0 Driver
* **Bảo vệ API:** Golang-JWT `v5` (Access Token Authentication)
* **Frontend:** HTML, CSS, JavaScript thuần túy (Kết xuất tĩnh từ Server)
* **Infra:** Docker & Docker Compose

## Cấu trúc Thư mục

```text
chatbot-platform/
├── main.go                     # Điểm vào ứng dụng (Entrypoint), nạp Route và biến
├── database/
│   ├── db.go                   # Logic kết nối Connection Pool & Auto Migrate
├── models/
│   ├── user.go                 # Entity Người dùng
│   └── chat.go                 # Entity Lịch sử Chat 
├── controllers/
│   ├── auth.go                 # Điểm nhận Request Đăng ký/Đăng nhập 
│   └── chat.go                 # Điểm nhận Request Chat 
├── middlewares/
│   └── auth_middleware.go      # Chốt chặn Auth (Bảo vệ API bằng JWT)
├── services/
│   ├── auth_service.go         # Component chứa nghiệp vụ sinh Access Token 
│   └── claude_service.go       # Component chứa logic Call HTTP lên Claude Server
├── docker-compose.yml          # Container hóa Web App + Database
└── Dockerfile                  # Đóng gói App Core
```

## Hướng dẫn Cài đặt (Bằng Docker)

Dự án này ứng dụng mô hình của hệ thống thực tế (Architecture: UngDungDangVienCBSV5). **Không cần cài đặt Golang hay XAMPP/MySQL thủ công**. Chỉ cần dùng **Docker**.

### Bước 1: Khởi chạy dự án
Mở Terminal ở thư mục gốc của dự án (`/chatbot-platform`) và gõ:
```bash
docker-compose up -d --build
```
*Lưu ý: Có thể đưa biến Token Claude thật vào môi trường trực tiếp từ Docker-Compose hoặc file .env nếu sử dụng lâu dài.*

### Bước 2: Truy cập Giao diện
- Địa chỉ truy cập: `http://localhost:3000`
- Có thể **Đăng ký** một tài khoản bất kỳ trên giao diện. Sau khi bấn đăng nhập, hệ thống sẽ tự trả về một mã JWT và trình duyệt tự động gán Bearer Token vào chuỗi bộ nhớ của trang Web (Local Storage).
- Trải nghiệm tính năng Chat an toàn. Nếu thiếu Token, giao diện sẽ bị Server chặn.

## Cơ chế Bảo mật JWT Authentication
Tất cả các truy xuất dữ liệu trong API Group ngoài vùng Login đều được thông qua cổng An ninh `middlewares.JWTAuth()`.  Khi người dùng gọi API `/api/chatbots/1/chat`, họ sẽ đính kèm Header:
`Authorization: Bearer <JWT_Mã_Bảo_Mật>` 
Hệ thống Golang lập tức xác thực nó. Nếu không có hoặc Token đã hết hạn 24 tiếng, mã sẽ trả về HTTP Status `401 Unauthorized` để buộc user đăng nhập lại.
