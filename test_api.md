# Users CRUD API Testing Guide

## API Endpoints

### 1. GET /users - ดึงรายการผู้ใช้ทั้งหมด
```bash
curl -X GET http://localhost:3000/users
```

### 2. GET /users/{id} - ดึงข้อมูลผู้ใช้รายบุคคล
```bash
curl -X GET http://localhost:3000/users/1
```

### 3. POST /users - สร้างผู้ใช้ใหม่
```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "สมชาย",
    "last_name": "ใจดี", 
    "phone_number": "081-234-5678",
    "email": "somchai@example.com",
    "membership_level": "Gold",
    "points_balance": 15420
  }'
```

### 4. PUT /users/{id} - แก้ไขข้อมูลผู้ใช้
```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "สมชาย",
    "last_name": "ใจดี",
    "phone_number": "081-234-5678", 
    "email": "somchai@example.com",
    "membership_level": "Gold",
    "points_balance": 20000
  }'
```

### 5. DELETE /users/{id} - ลบผู้ใช้
```bash
curl -X DELETE http://localhost:3000/users/1
```

## User Model Fields

Based on the UI shown in the image:

- `id`: รหัสผู้ใช้ (Auto-generated)
- `first_name`: ชื่อ (Required)
- `last_name`: นามสกุล (Required) 
- `phone_number`: เบอร์โทรศัพท์ (Required, Unique)
- `email`: อีเมล (Required, Unique)
- `registration_date`: วันที่สมัครสมาชิก (Auto-generated)
- `membership_level`: ระดับสมาชิก (Bronze/Silver/Gold, Default: Bronze)
- `points_balance`: แต้มคงเหลือ (Default: 0)
- `created_at`: วันที่สร้าง (Auto-generated)
- `updated_at`: วันที่อัปเดต (Auto-generated)

## Example Response

```json
{
  "id": 1,
  "first_name": "สมชาย",
  "last_name": "ใจดี",
  "phone_number": "081-234-5678",
  "email": "somchai@example.com", 
  "registration_date": "2025-10-17T13:48:25.123456789+07:00",
  "membership_level": "Gold",
  "points_balance": 15420,
  "created_at": "2025-10-17T13:48:25.123456789+07:00",
  "updated_at": "2025-10-17T13:48:25.123456789+07:00"
}
```

## Database

- SQLite database file: `users.db` (created automatically)
- GORM auto-migration enabled
- Database connection initialized on server startup