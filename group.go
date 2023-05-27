package main

import "time"

// Group represents a group in the system
type Group struct {
	ID        int
	Name      string
	Members   []*User
	Expenses  []*Expense
	CreatedAt time.Time
}
