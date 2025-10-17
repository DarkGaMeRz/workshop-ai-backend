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

    User ||--o{ Transfer : "from_user_id"
    User ||--o{ Transfer : "to_user_id"
    User ||--o{ PointLedger : "user_id"
    Transfer ||--o{ PointLedger : "transfer_id"
```
