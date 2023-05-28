package main

import (
	"fmt"
	"time"
)

// Splitwise represents the main application
type Splitwise struct {
	Users      []*User
	Groups     []*Group
	Expenses   []*Expense
	Balances   map[*User]float64
	Currencies map[string]float64
}

// RegisterUser registers a new user
func (s *Splitwise) RegisterUser(username, email, password string) *User {
	user := &User{
		ID:        len(s.Users) + 1,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	s.Users = append(s.Users, user)
	s.Balances[user] = 0.0

	return user
}

// LoginUser authenticates a user
func (s *Splitwise) LoginUser(email, password string) (*User, error) {
	for _, user := range s.Users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, fmt.Errorf("invalid email or password")
}

// CreateExpense creates a new expense
func (s *Splitwise) CreateExpense(amount float64, description string, payer *User, participants []*User) *Expense {
	expense := &Expense{
		ID:           len(s.Expenses) + 1,
		Amount:       amount,
		Description:  description,
		Payer:        payer,
		Participants: participants,
		CreatedAt:    time.Now(),
	}

	s.Expenses = append(s.Expenses, expense)

	// Update balances
	equalShare := amount / float64(len(participants)+1)
	s.Balances[payer] += amount - equalShare
	for _, participant := range participants {
		s.Balances[participant] -= equalShare
	}

	return expense
}

// CreateGroup creates a new group
func (s *Splitwise) CreateGroup(name string, members []*User) *Group {
	group := &Group{
		ID:        len(s.Groups) + 1,
		Name:      name,
		Members:   members,
		Expenses:  []*Expense{},
		CreatedAt: time.Now(),
	}

	s.Groups = append(s.Groups, group)

	return group
}

// SettleExpenses settles debts between users
func (s *Splitwise) SettleExpenses() {
	for user, balance := range s.Balances {
		if balance < 0 {
			for otherUser, otherBalance := range s.Balances {
				if otherBalance > 0 {
					amount := -balance
					if amount > otherBalance {
						amount = otherBalance
					}
					s.Balances[user] += amount
					s.Balances[otherUser] -= amount
					fmt.Printf("Settled debt between %s and %s: %.2f\n", user.Username, otherUser.Username, amount)
					if s.Balances[user] >= 0 {
						break
					}
				}
			}
		}
	}
}

// TrackSpending calculates and returns the total spent by a user
func (s *Splitwise) TrackSpending(user *User) float64 {
	totalSpent := 0.0
	for _, expense := range s.Expenses {
		if expense.Payer == user {
			totalSpent += expense.Amount
		}
	}
	return totalSpent
}

// FilterExpenses returns expenses filtered by the specified conditions
func (s *Splitwise) FilterExpenses(user *User, minAmount float64, maxAmount float64, startDate time.Time,
	endDate time.Time) []*Expense {
	filteredExpenses := []*Expense{}
	for _, expense := range s.Expenses {
		if expense.Payer == user &&
			expense.Amount >= minAmount &&
			expense.Amount <= maxAmount &&
			expense.CreatedAt.After(startDate) &&
			expense.CreatedAt.Before(endDate) {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}
	return filteredExpenses
}
