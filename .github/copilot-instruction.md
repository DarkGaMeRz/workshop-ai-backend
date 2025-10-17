# GitHub Copilot Instructions

## Project Overview
This is a **LBK Points Management System** backend API built with Go, Fiber framework, GORM ORM, and SQLite database. The system manages member points, transfers, and transaction ledgers for a loyalty program.

## Tech Stack
- **Language**: Go 1.21+
- **Web Framework**: Fiber v2
- **ORM**: GORM
- **Database**: SQLite
- **API Style**: RESTful

---

## Coding Standards

### Naming Conventions
- **IMPORTANT**: Always use `Member` instead of `User` or `Customer` for all entities, tables, and API endpoints
  - ✅ Correct: `Member`, `/members`, `member_id`
  - ❌ Incorrect: `User`, `/users`, `user_id`
  - ❌ Incorrect: `Customer`, `/customers`, `customer_id`

- **Variables & Functions**: Use camelCase for local variables, PascalCase for exported functions
  ```go
  // ✅ Good
  memberID := 123
  func GetMemberByID(c *fiber.Ctx) error

  // ❌ Bad
  member_id := 123
  func get_member_by_id(c *fiber.Ctx) error
  ```

- **Struct Fields**: Use PascalCase with proper JSON tags
  ```go
  type Transfer struct {
      ID             uint   `json:"id" gorm:"primaryKey"`
      FromMemberID   uint   `json:"from_member_id" gorm:"not null"`
      ToMemberID     uint   `json:"to_member_id" gorm:"not null"`
  }
  ```

### File Organization
```
project/
├── main.go              # Application entry point, server setup
├── database.go          # Database initialization and connection
├── models.go            # All data models (Member, Transfer, PointLedger)
├── routes.go            # HTTP route handlers
├── swagger.yml          # API documentation
├── database.md          # Database schema documentation (Mermaid ER diagram)
└── .github/
    └── copilot-instructions.md
```

### Error Handling
- Always return proper HTTP status codes
- Use consistent error response format:
  ```go
  return c.Status(404).JSON(fiber.Map{
      "error": "Member not found",
  })
  ```

### Database Patterns
- Use GORM for all database operations
- Always check for `gorm.ErrRecordNotFound` explicitly
- Use transactions for operations affecting multiple records (especially transfers)

---

## High-Level Architecture

### Core Entities
1. **Member** - Customer information and points balance
2. **Transfer** - Point transfer transactions between members
3. **PointLedger** - Append-only audit log of all point changes

### Key Features
1. **Member CRUD** - Create, Read, Update, Delete member records
2. **Point Transfers** - Transfer points between members with idempotency
3. **Transaction History** - Query point transaction history with pagination
4. **Audit Trail** - Immutable ledger of all point movements

### API Design Principles
- RESTful endpoints following standard conventions
- Idempotency for transfer operations using `idempotency_key`
- Pagination support for list endpoints
- Atomic transactions for point transfers

---

## What You SHOULD Do

### ✅ Database Operations
- Use GORM methods: `db.Find()`, `db.First()`, `db.Create()`, `db.Save()`, `db.Delete()`
- Check for errors after every database operation
- Use `gorm.ErrRecordNotFound` for 404 responses
- Apply proper indexes on foreign keys and frequently queried fields

### ✅ API Responses
- Return appropriate HTTP status codes (200, 201, 400, 404, 409, 500)
- Use consistent JSON response structure
- Include proper error messages
- Validate input data before processing

### ✅ Code Structure
- Keep route handlers in `routes.go`
- Keep models in `models.go`
- Keep database setup in `database.go`
- Add comments for complex business logic
- Use middleware for cross-cutting concerns (CORS, logging)

### ✅ Security & Validation
- Validate required fields before creating/updating records
- Check for duplicate email/phone before creating members
- Ensure point transfers don't create negative balances
- Use prepared statements (GORM handles this automatically)

### ✅ Documentation
- Update `swagger.yml` when adding/modifying endpoints
- Update `database.md` when changing database schema
- Use Mermaid ER diagrams for database documentation
- Keep README up to date with setup instructions

---

## What You MUST NOT Do

### ❌ Naming Violations
- **NEVER** use `User`, `user`, `users` - always use `Member`, `member`, `members`
- **NEVER** use `Customer`, `customer`, `customers` - use `Member` instead
- Do not mix naming conventions (snake_case and camelCase in same context)

### ❌ Database Anti-Patterns
- Do not use raw SQL queries unless absolutely necessary
- Do not forget to check for errors after database operations
- Do not perform cascading deletes without explicit business logic
- Do not skip database migrations

### ❌ Security Issues
- Do not store sensitive data in plain text
- Do not expose internal error details to API clients
- Do not skip input validation
- Do not allow SQL injection (use GORM parameterized queries)

### ❌ API Design Mistakes
- Do not return inconsistent response formats
- Do not skip pagination for list endpoints
- Do not forget idempotency for critical operations
- Do not expose database IDs without considering security

### ❌ Code Organization
- Do not put business logic in `main.go`
- Do not mix database code with HTTP handlers
- Do not create circular dependencies
- Do not duplicate code - extract common functions

### ❌ Performance Issues
- Do not perform N+1 queries - use proper joins/preloads
- Do not load entire tables without pagination
- Do not skip database indexes on foreign keys
- Do not forget to close database connections

---

## Testing Guidelines

### Manual Testing
- Test all CRUD operations for each entity
- Test edge cases (empty lists, not found, duplicates)
- Test point transfer validations (negative balance, same sender/receiver)
- Test pagination parameters

### API Testing
- Use provided `test_crud.sh` script for basic testing
- Verify response status codes
- Verify response JSON structure
- Test error scenarios

---

## Git Workflow

### Commit Messages
- Use clear, descriptive commit messages
- Format: `[Type] Brief description`
- Types: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`

Examples:
```
feat: Add point transfer API with idempotency
fix: Prevent negative balance in transfers
docs: Update database.md with PointLedger table
refactor: Extract validation functions
```

### Before Committing
- Run `go mod tidy` to clean up dependencies
- Ensure code compiles without errors
- Update documentation if schema/API changed
- Test critical paths manually

---

## Performance Considerations

- Use database indexes on frequently queried columns
- Implement pagination for list endpoints (default 20 items/page)
- Consider caching for frequently accessed data
- Use connection pooling for database connections
- Profile and optimize slow queries

---

## Future Enhancements (Out of Scope for Now)

- Authentication & Authorization
- Rate limiting
- Webhook notifications
- Background job processing
- Multi-database support (PostgreSQL, MySQL)
- Advanced analytics and reporting
- Point expiration system
- Referral program

---

## Quick Reference

### Common Commands
```bash
# Run the server
go run main.go database.go models.go routes.go

# Clean up dependencies
go mod tidy

# Test API
./test_crud.sh

# Git operations
git add .
git commit -m "Your message"
git push
```

### Example API Calls
```bash
# Create member
curl -X POST http://localhost:3000/members \
  -H "Content-Type: application/json" \
  -d '{"first_name":"John","last_name":"Doe","email":"john@example.com","phone_number":"0812345678"}'

# Get all members
curl http://localhost:3000/members

# Transfer points
curl -X POST http://localhost:3000/transfers \
  -H "Content-Type: application/json" \
  -d '{"from_member_id":1,"to_member_id":2,"amount":100,"note":"Thanks!"}'
```

---

## Support & Resources

- **Fiber Documentation**: https://docs.gofiber.io/
- **GORM Documentation**: https://gorm.io/docs/
- **Mermaid Diagrams**: https://mermaid.js.org/
- **OpenAPI Spec**: See `swagger.yml`

---

**Remember**: When in doubt, prioritize:
1. Code clarity over cleverness
2. Consistency over personal preference
3. Member safety and data integrity
4. Proper error handling and logging