# Database Setup Guide

## PostgreSQL Configuration

### 1. Cài đặt PostgreSQL Driver

```bash
go get gorm.io/driver/postgres
```

### 2. Cấu hình Environment Variables

Tạo file `.env` trong thư mục gốc với nội dung:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=ecommerce_mini
DB_SSLMODE=disable
```

### 3. Tạo Database trong pgAdmin 4

1. Mở pgAdmin 4
2. Kết nối đến PostgreSQL server
3. Tạo database mới tên `ecommerce_mini`
4. Hoặc có thể để code tự động tạo database

### 4. Sử dụng trong code

```go
package main

import (
    "gin/config/db"
)

func main() {
    // Kết nối database
    db.ConnectDatabase()

    // Sử dụng database
    database := db.GetDB()
    // ... your code here
}
```

### 5. Các Model được Auto Migrate

- User (với gorm.Model)
- Product
- Order
- OrderDetail

### 6. Quan hệ giữa các bảng

- User -> Order (1:N)
- Order -> OrderDetail (1:N)
- Product -> OrderDetail (1:N)
