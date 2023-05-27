package main

import "time"

// Expense represents an expense with its details
type Expense struct {
	ID           int
	Amount       float64
	Description  string
	Payer        *User
	Participants []*User
	CreatedAt    time.Time
}
