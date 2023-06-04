# Splitwise

## Requirements with Entities defined
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

1. User Table:
   - No direct relationships with other tables.

2. Expense Table:
   - One-to-Many relationship with User Table (payer_id references User.id).
   - One-to-Many relationship with Participant Table (Expense.id references Participant.expense_id).

3. Participant Table:
   - Many-to-One relationship with Expense Table (expense_id references Expense.id).
   - Many-to-One relationship with User Table (user_id references User.id).

4. Settlement Table:
   - Many-to-One relationship with User Table (payer_id references User.id).
   - Many-to-One relationship with User Table (receiver_id references User.id).

5. Group Table:
   - No direct relationships with other tables.

6. GroupMember Table:
   - Many-to-One relationship with Group Table (group_id references Group.id).
   - Many-to-One relationship with User Table (user_id references User.id).

7. GroupExpense Table:
   - Many-to-One relationship with Group Table (group_id references Group.id).
   - Many-to-One relationship with Expense Table (expense_id references Expense.id).

8. Currency Table:
   - No direct relationships with other tables.



## Classes involved

1. User:
   - Responsible for user management and authentication.
   - Attributes: id, username, password, email, profileInfo, createdAt, updatedAt.
   - Methods: register(), login(), updateProfile(), ...

2. Expense:
   - Represents an individual expense.
   - Attributes: id, amount, description, payer (User), participants (User[]), createdAt, updatedAt.
   - Methods: createExpense(), addParticipant(), calculateBalances(), ...

3. Group:
   - Manages groups and shared expenses.
   - Attributes: id, name, members (User[]), expenses (Expense[]), createdAt, updatedAt.
   - Methods: createGroup(), addMember(), shareExpense(), ...

4. Payment:
   - Handles expense settlement between users.
   - Attributes: payer (User), receiver (User), amount, createdAt, updatedAt.
   - Methods: settleDebts(), calculateBalances(), ...

5. CurrencyConverter:
   - Performs currency conversions.
   - Methods: convert(), getExchangeRate(), ...

6. ExpenseTracker:
   - Tracks personal expenses and provides spending analysis.
   - Methods: viewPersonalExpenses(), trackSpending(), filterExpenses(), ...

7. NotificationService:
   - Handles real-time or delayed notifications.
   - Methods: sendNotification(), subscribeToActivity(), ...

8. Database:
   - Manages the persistence and retrieval of data.
   - Methods: saveUser(), getUser(), saveExpense(), getExpense(), ...



## Class diagram 

```
+-----------------------------------+
|              User                 |
+-----------------------------------+
| - id: String                      |
| - username: String                |
| - password: String                |
| - email: String                   |
| - profileInfo: String             |
| - createdAt: Date                 |
| - updatedAt: Date                 |
|-----------------------------------|
| +register(): void                 |
| +login(): void                    |
| +createExpense(): void            |
| +addParticipant(): void           |
| +settleDebts(): void              |
| +viewExpenses(): void             |
| +trackSpending(): void            |
| +filterExpenses(): void           |
| +notify(): void                   |
| +performCurrencyConversion(): void|
+-----------------------------------+

+-----------------------------------+
|            Expense                |
+-----------------------------------+
| - id: String                      |
| - amount: double                  |
| - description: String             |
| - payer: User                     |
| - participants: User[]            |
| - createdAt: Date                 |
| - updatedAt: Date                 |
|-----------------------------------|
| +calculateBalances(): void        |
| +displayBalances(): void          |
+-----------------------------------+

+-----------------------------------+
|              Group                |
+-----------------------------------+
| - id: String                      |
| - name: String                    |
| - members: User[]                 |
| - expenses: Expense[]             |
| - createdAt: Date                 |
| - updatedAt: Date                 |
|-----------------------------------|
| +createGroup(): void              |
| +manageGroup(): void              |
| +shareExpenses(): void            |
| +viewGroupExpenses(): void        |
+-----------------------------------+

+-----------------------------------+
|           Payment                 |
+-----------------------------------+
| - payer: User                     |
| - receiver: User                  |
| - amount: double                  |
| - createdAt: Date                 |
| - updatedAt: Date                 |
+-----------------------------------+

```



## Use case diagram
```
+-----------------------------------+
|         Splitwise Use Cases        |
+-----------------------------------+
| - User Management                 |
| - Expense Creation                |
| - Expense Settlement              |
| - Group Management                |
| - Expense Tracking                |
| - Notifications                   |
| - Currency Conversion             |
| - Privacy and Security            |
+-----------------------------------+
```



## Detailed Sequence diagram 


1. User Management (Registration, Login, Authentication):

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Authentication|              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User requests |              |                 |              |                   |
|    registration  |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Authentication|              |                 |              |                   |
|    service       |              |                 |              |                   |
|    verifies user |              |                 |              |                   |
|    credentials   |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. User receives |              |                 |              |                   |
|    registration  |              |                 |              |                   |
|    confirmation  |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. User requests |              |                 |              |                   |
|    login         |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 5. Authentication|              |                 |              |                   |
|    service       |              |                 |              |                   |
|    validates     |              |                 |              |                   |
|    user login    |              |                 |              |                   |
|    credentials   |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 6. User logged   |              |                 |              |                   |
|    in            |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

2. Expense Creation:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Expense       |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User creates  |              |                 |              |                   |
|    an expense    |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Expense       |              |                 |              |                   |
|    object        |              |                 |              |                   |
|    created       |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. User specifies|              |                 |              |                   |
|    expense       |              |                 |              |                   |
|    details       |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. Expense       |              |                 |              |                   |
|    details       |              |                 |              |                   |
|    recorded      |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 5. User adds     |              |                 |              |                   |
|    participants  |              |                 |              |                   |
|    to the        |              |                 |              |                   |
|    expense       |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 6. Participants  |              |                 |              |                   |
|    added         |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 7. Expense       |              |                 |              |                   |
|    participants  |              |                 |              |                   |
|    recorded      |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 8. Expense       |              |                 |              |                   |
|    creation      |              |                 |              |                   |
|    completed     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

3. Expense Settlement:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Expense       |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User selects  |              |                 |              |                   |
|    an expense    |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Expense       |              |                 |              |                   |
|    details       |              |                 |              |                   |
|    retrieved     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. User selects  |              |                 |              |                   |
|    settlement    |              |                 |              |                   |
|                  |              |                 |
|    option        |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. Settlement    |              |                 |              |                   |
|    process       |              |                 |              |                   |
|    initiated     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 5. Debts         |              |                 |              |                   |
|    calculated    |              |                 |              |                   |
|    and displayed |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 6. User selects  |              |                 |              |                   |
|    payment       |              |                 |              |                   |
|    method        |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 7. Payment       |              |                 |              |                   |
|    process       |              |                 |              |                   |
|    completed     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 8. Settlement    |              |                 |              |                   |
|    process       |              |                 |              |                   |
|    completed     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

4. Group Management:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Group         |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User creates  |              |                 |              |                   |
|    a group       |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Group details |              |                 |              |                   |
|    entered       |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. Group         |              |                 |              |                   |
|    created       |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. User adds     |              |                 |              |                   |
|    members       |              |                 |              |                   |
|    to the group 
|                  |              |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 5. Members       |              |                 |              |                   |
|    added         |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 6. Group         |              |                 |              |                   |
|    members       |              |                 |              |                   |
|    recorded      |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 7. User shares   |              |                 |              |                   |
|    an expense    |              |                 |              |                   |
|    within the    |              |                 |              |                   |
|    group         |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 8. Expense       |              |                 |              |                   |
|    shared        |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

5. Expense Tracking:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Expense       |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User views    |              |                 |              |                   |
|    personal      |              |                 |              |                   |
|    expenses      |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Personal      |              |                 |              |                   |
|    expenses      |              |                 |              |                   |
|    retrieved     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. User tracks   |              |                 |              |                   |
|    spending      |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. Spending      |              |                 |              |                   |
|    tracked       |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 5. User filters  |              |                 |              |                   |
|    and sorts     |              |                 |              |                   |
|    expenses      |              |                 |              |                   |
|----------------
->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 6. Filtered and  |              |                 |              |                   |
|    sorted        |              |                 |              |                   |
|    expenses      |              |                 |              |                   |
|    displayed     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

6. Notifications:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Splitwise     |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User performs |              |                 |              |                   |
|    an action     |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Notification  |              |                 |              |                   |
|    triggered     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. Notification  |              |                 |              |                   |
|    sent to user  |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. User receives |              |                 |              |                   |
|    notification  |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

7. Currency Conversion:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Splitwise     |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User selects  |              |                 |              |                   |
|    currency      |              |                 |              |                   |
|    conversion    |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Currency      |              |                 |              |                   |
|    conversion    |              |                 |              |                   |
|    performed     |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. Converted     |              |                 |              |                   |
|    amount        |              |                 |              |                   |
|    displayed     |              |                 |              |                   |
|<

-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```

8. Privacy and Security:

```
+------------------+              +-----------------+              +-------------------+
|      User        |              |   Splitwise     |              |      Database     |
+------------------+              +-----------------+              +-------------------+
|                  |              |                 |              |                   |
| 1. User provides |              |                 |              |                   |
|    sensitive     |              |                 |              |                   |
|    information   |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 2. Information   |              |                 |              |                   |
|    securely       |              |                 |              |                   |
|    stored        |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
| 3. User sets     |              |                 |              |                   |
|    privacy       |              |                 |              |                   |
|    settings      |              |                 |              |                   |
|----------------->|              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
|                  |              |                 |              |                   |
| 4. Privacy       |              |                 |              |                   |
|    settings      |              |                 |              |                   |
|    updated       |              |                 |              |                   |
|<-----------------|              |                 |              |                   |
|                  |              |                 |              |                   |
+------------------+              +-----------------+              +-------------------+
```
