package main

import "time"

// User represents a user in the system
type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}
