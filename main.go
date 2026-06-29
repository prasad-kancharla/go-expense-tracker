package main

import (
	"expense-tracker/expense"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Expense Tracker")
	fmt.Println("You have 5 options, Please enter one of the numbers from  1 to 6")

	i := 1
	for i > 0 {

		var input int
		fmt.Println("Enter your choice")
		showChoices()
		fmt.Scanln(&input)
		switch input {
		case 1:
			addExpenseFlow()
		case 2:
			fmt.Println(expense.All())
		case 3:
			fmt.Println(expense.LastFive())
		case 4:
			deleteExpenseFlow()
		case 5:
			updateExpenseFlow()
		case 6:
			i = 0
		default:
			fmt.Println("Please enter valid input")
		}
	}

}

func showChoices() {
	fmt.Println("1. Add an expense")
	fmt.Println("2. View All Expenses")
	fmt.Println("3. Get Last 5 Expenses")
	fmt.Println("4. Delete an expense")
	fmt.Println("5. Update an expense amount or title")
	fmt.Println("6. Exit")
}

func addExpenseFlow() {
	fmt.Println("Please enter title, amount, payment method, and category")
	var title string
	var amount float64
	var paymentMethod string
	var category string

	fmt.Scan(&title, &amount, &paymentMethod, &category)
	e, err := expense.Create(title, amount, paymentMethod, category)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Expense created successfully")
	fmt.Println(e)
}

func deleteExpenseFlow() {
	fmt.Println("Enter id of expense to delete")
	var id int64
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Only numbers are allowed")
		return
	}
	err = expense.Remove(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Expense deleted successfully")
}

func updateExpenseFlow() {
	fmt.Println("Enter id of expense you want to update")
	var id int64
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Only numbers are allowed")
		return
	}
	fmt.Println(`If you want to update amount type "amount" or else type "title"`)
	var updateType string
	fmt.Scanln(&updateType)
	updateType = strings.ToLower(updateType)
	switch updateType {
	case "amount":
		fmt.Println("Enter amount")
		var amount float64
		_, err = fmt.Scanln(&amount)
		if err != nil {
			fmt.Println("Only numbers are allowed")
			return
		}
		err = expense.UpdateAmount(id, amount)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Expense amount updated successfully")
	case "title":
		fmt.Println("Enter the new title")
		var newTitle string
		fmt.Scanln(&newTitle)
		err = expense.UpdateTitle(id, newTitle)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Expense title updated successfully")
	default:
		fmt.Println("Valid values are amount or title")
	}

}
