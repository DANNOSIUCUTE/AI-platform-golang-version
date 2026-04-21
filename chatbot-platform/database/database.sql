-- Tạo cơ sở dữ liệu nếu chưa tồn tại
CREATE DATABASE IF NOT EXISTS `chatbot_platform` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Mở cơ sở dữ liệu vừa tạo để sử dụng
USE `chatbot_platform`;

-- Tạo bảng thông tin người dùng (Users)
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(50) UNIQUE NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(100),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Thêm một vài dữ liệu mẫu để bạn test đăng nhập ngay lập tức
INSERT INTO `users` (`username`, `password`, `email`) VALUES
('testuser', '123456', 'test@example.com'),
('admin', 'admin123', 'admin@example.com');
