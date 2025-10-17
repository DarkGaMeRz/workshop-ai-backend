#!/bin/bash

# Users CRUD API Test Script
echo "=== Users CRUD API Test Script ==="
echo ""

BASE_URL="http://localhost:3000"

echo "1. Testing GET /users (ดึงรายการผู้ใช้ทั้งหมด)"
echo "curl -X GET $BASE_URL/users"
curl -X GET $BASE_URL/users | jq '.' 2>/dev/null || curl -X GET $BASE_URL/users
echo -e "\n"

echo "2. Testing POST /users (สร้างผู้ใช้ใหม่)"
echo "Creating user: มาลี ดีใจ"
USER_DATA='{
  "first_name": "มาลี",
  "last_name": "ดีใจ", 
  "phone_number": "089-876-5432",
  "email": "malee@example.com",
  "membership_level": "Silver",
  "points_balance": 8500
}'

echo "curl -X POST $BASE_URL/users -H 'Content-Type: application/json' -d '$USER_DATA'"
RESPONSE=$(curl -X POST $BASE_URL/users -H "Content-Type: application/json" -d "$USER_DATA")
echo $RESPONSE | jq '.' 2>/dev/null || echo $RESPONSE
USER_ID=$(echo $RESPONSE | grep -o '"id":[0-9]*' | cut -d':' -f2)
echo -e "\n"

echo "3. Testing GET /users/{id} (ดึงข้อมูลผู้ใช้รายบุคคล)"
echo "curl -X GET $BASE_URL/users/$USER_ID"
curl -X GET $BASE_URL/users/$USER_ID | jq '.' 2>/dev/null || curl -X GET $BASE_URL/users/$USER_ID
echo -e "\n"

echo "4. Testing PUT /users/{id} (แก้ไขข้อมูลผู้ใช้)"
echo "Updating points balance to 12000"
UPDATE_DATA='{
  "first_name": "มาลี",
  "last_name": "ดีใจ",
  "phone_number": "089-876-5432", 
  "email": "malee@example.com",
  "membership_level": "Gold",
  "points_balance": 12000
}'

echo "curl -X PUT $BASE_URL/users/$USER_ID -H 'Content-Type: application/json' -d '$UPDATE_DATA'"
curl -X PUT $BASE_URL/users/$USER_ID -H "Content-Type: application/json" -d "$UPDATE_DATA" | jq '.' 2>/dev/null || curl -X PUT $BASE_URL/users/$USER_ID -H "Content-Type: application/json" -d "$UPDATE_DATA"
echo -e "\n"

echo "5. Testing GET /users (ดูรายการผู้ใช้ทั้งหมดหลังอัปเดต)"
echo "curl -X GET $BASE_URL/users"
curl -X GET $BASE_URL/users | jq '.' 2>/dev/null || curl -X GET $BASE_URL/users
echo -e "\n"

echo "6. Testing DELETE /users/{id} (ลบผู้ใช้)"
echo "curl -X DELETE $BASE_URL/users/$USER_ID"
curl -X DELETE $BASE_URL/users/$USER_ID | jq '.' 2>/dev/null || curl -X DELETE $BASE_URL/users/$USER_ID
echo -e "\n"

echo "7. Testing GET /users (ดูรายการผู้ใช้ทั้งหมดหลังลบ)"
echo "curl -X GET $BASE_URL/users"
curl -X GET $BASE_URL/users | jq '.' 2>/dev/null || curl -X GET $BASE_URL/users
echo -e "\n"

echo "=== Test Complete ==="