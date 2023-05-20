# Splitwise

## Requirements 
```
┌──────────────────────────────────────────┐
│           Splitwise Requirements          │
├──────────────────────────────────────────┤
│ 1. User Management                       │
│    - Registration, Login, Authentication│
│    - User profiles with basic info       │
├──────────────────────────────────────────┤
│ 2. Expense Creation                      │
│    - Create expenses with details        │
│    - Specify amount, description, payer │
│    - Add participants                    │
├──────────────────────────────────────────┤
│ 3. Expense Settlement                    │
│    - Settle debts with other participants│
│    - Calculate and display balances      │
│    - Support various payment methods     │
├──────────────────────────────────────────┤
│ 4. Group Management                      │
│    - Create and manage groups            │
│    - Share expenses within groups        │
│    - Visible group expenses for members  │
├──────────────────────────────────────────┤
│ 5. Expense Tracking                      │
│    - View personal expenses              │
│    - Track spending                       │
│    - Filter and sort expenses            │
├──────────────────────────────────────────┤
│ 6. Notifications                         │
│    - Real-time or minimal-delay alerts   │
│    - New expenses, settlements, activity │
├──────────────────────────────────────────┤
│ 7. Currency Conversion                   │
│    - Support multiple currencies         │
│    - Perform conversions                │
├──────────────────────────────────────────┤
│ 8. Privacy and Security                  │
│    - Secure storage of information       │
│    - Privacy settings for expenses       │
├──────────────────────────────────────────┤
```


## Tables Involved

1. User Table:
   - id (Primary Key)
   - username
   - password
   - email
   - profile_info
   - created_at
   - updated_at

2. Expense Table:
   - id (Primary Key)
   - amount
   - description
   - payer_id (Foreign Key referencing User.id)
   - created_at
   - updated_at

3. Participant Table:
   - id (Primary Key)
   - expense_id (Foreign Key referencing Expense.id)
   - user_id (Foreign Key referencing User.id)
   - share_amount
   - created_at
   - updated_at

4. Settlement Table:
   - id (Primary Key)
   - payer_id (Foreign Key referencing User.id)
   - receiver_id (Foreign Key referencing User.id)
   - amount
   - created_at
   - updated_at

5. Group Table:
   - id (Primary Key)
   - name
   - created_at
   - updated_at

6. GroupMember Table:
   - id (Primary Key)
   - group_id (Foreign Key referencing Group.id)
   - user_id (Foreign Key referencing User.id)
   - created_at
   - updated_at

7. GroupExpense Table:
   - id (Primary Key)
   - group_id (Foreign Key referencing Group.id)
   - expense_id (Foreign Key referencing Expense.id)
   - created_at
   - updated_at

8. Currency Table:
   - id (Primary Key)
   - code
   - name


## Relationship among tables
