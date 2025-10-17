# Database Schema Documentation

## Entity Relationship Diagram

```mermaid
erDiagram
    User {
        uint id PK
        string first_name
        string last_name
        string phone_number
        string email
        datetime registration_date
        string membership_level
        int points_balance
        datetime created_at
        datetime updated_at
    }

    Transfer {
        uint id PK
        uint from_user_id FK
        uint to_user_id FK
        int amount
        string status
        string note
        string idempotency_key UNIQUE
        datetime created_at
        datetime updated_at
        datetime completed_at
        string fail_reason
    }

    PointLedger {
        uint id PK
        uint user_id FK
        int change
        int balance_after
        string event_type
        uint transfer_id FK
        string reference
        string metadata
        datetime created_at
    }

    User ||--o{ Transfer : "sends/receives"
    User ||--o{ PointLedger : "has"
    Transfer ||--o{ PointLedger : "records"
```

## Table Descriptions

### User
Stores customer information and loyalty points balance.

**Fields:**
- `id`: Primary key
- `first_name`: User's first name
- `last_name`: User's last name
- `phone_number`: Contact phone number (unique)
- `email`: Email address (unique)
- `registration_date`: Date of registration
- `membership_level`: Membership tier (Bronze, Silver, Gold, Platinum)
- `points_balance`: Current points balance
- `created_at`: Record creation timestamp
- `updated_at`: Record update timestamp

### Transfer
Records point transfer transactions between users.

**Fields:**
- `id`: Primary key
- `from_user_id`: Foreign key to User (sender)
- `to_user_id`: Foreign key to User (receiver)
- `amount`: Number of points transferred
- `status`: Transfer status (pending, processing, completed, failed, cancelled, reversed)
- `note`: Optional note/message
- `idempotency_key`: Unique key for idempotent operations
- `created_at`: Transfer creation timestamp
- `updated_at`: Transfer update timestamp
- `completed_at`: Transfer completion timestamp
- `fail_reason`: Reason if transfer failed

**Status Values:**
- `pending`: Transfer initiated, awaiting processing
- `processing`: Transfer in progress
- `completed`: Transfer successful
- `failed`: Transfer failed
- `cancelled`: Transfer cancelled
- `reversed`: Transfer reversed/rolled back

### PointLedger
Append-only ledger tracking all point balance changes.

**Fields:**
- `id`: Primary key
- `user_id`: Foreign key to User
- `change`: Points added (+) or deducted (-)
- `balance_after`: Balance after this transaction
- `event_type`: Type of event (transfer_out, transfer_in, adjust, earn, redeem)
- `transfer_id`: Foreign key to Transfer (if applicable)
- `reference`: Optional reference/tracking ID
- `metadata`: Additional JSON metadata
- `created_at`: Ledger entry timestamp

**Event Types:**
- `transfer_out`: Points sent to another user
- `transfer_in`: Points received from another user
- `adjust`: Manual adjustment by admin
- `earn`: Points earned from activities
- `redeem`: Points redeemed for rewards

## Relationships

1. **User → Transfer**: One user can send or receive multiple transfers
   - `from_user_id` links to the sender
   - `to_user_id` links to the receiver

2. **User → PointLedger**: One user has multiple ledger entries tracking their point history

3. **Transfer → PointLedger**: One transfer can generate multiple ledger entries (one for sender, one for receiver)

## Indexes

### Transfer Table
- `idx_transfers_from`: Index on `from_user_id`
- `idx_transfers_to`: Index on `to_user_id`
- `idx_transfers_created`: Index on `created_at`
- Unique index on `idempotency_key`

### PointLedger Table
- `idx_ledger_user`: Index on `user_id`
- `idx_ledger_transfer`: Index on `transfer_id`
- `idx_ledger_created`: Index on `created_at`
