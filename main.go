package main

import (
	"fmt"
	"time"
)

func main() {
	splitwise := &Splitwise{
		Users:    []*User{},
		Groups:   []*Group{},
		Expenses: []*Expense{},
		Balances: make(map[*User]float64),
		Currencies: map[string]float64{
			"USD": 1.0,
			"EUR": 0.85,
			"GBP": 0.72,
		},
	}

	// Register users
	user1 := splitwise.RegisterUser("user1", "user1@example.com", "password1")
	user2 := splitwise.RegisterUser("user2", "user2@example.com", "password2")

	// Create an expense
	expense := splitwise.CreateExpense(50.0, "Dinner", user1, []*User{user1, user2})

	// Create a group
	group := splitwise.CreateGroup("Group1", []*User{user1, user2})

	// Settle debts
	splitwise.SettleExpenses()

	// Track spending
	totalSpent := splitwise.TrackSpending(user1)
	fmt.Printf("%s has spent: %.2f\n", user1.Username, totalSpent)

	// Filter expenses
	startDate := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, time.May, 31, 23, 59, 59, 0, time.UTC)
	filteredExpenses := splitwise.FilterExpenses(user1, 0.0, 100.0, startDate, endDate)
	fmt.Printf("Filtered expenses for %s:\n", user1.Username)
	for _, expense := range filteredExpenses {
		fmt.Printf("Expense ID: %d, Amount: %.2f, Description: %s\n", expense.ID, expense.Amount, expense.Description)
	}
}
