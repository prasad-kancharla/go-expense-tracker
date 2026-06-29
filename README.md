# go-expense-tracker

A command-line expense tracker built in Go to learn idiomatic Go project structure, package design, and naming conventions.

## Project Structure

```
expense-tracker/
├── go.mod
├── main.go              # Entry point and CLI interaction
└── expense/
    ├── model.go         # Expense type and constructor
    ├── repository.go    # In-memory storage (unexported)
    └── service.go       # Public API — business logic
```

## Features

- Add an expense (title, amount, payment method, category)
- View all expenses
- Get last 5 expenses (sorted by date)
- Update an expense amount or title
- Delete an expense

## Requirements

- [Go 1.21+](https://go.dev/dl/)

## Running

```bash
go run main.go
```

## Usage

When prompted, enter a number to choose an action:

```
1. Add an expense
2. View All Expenses
3. Get Last 5 Expenses
4. Delete an expense
5. Update an expense amount or title
6. Exit
```

When adding an expense, enter the values space-separated:

```
title amount paymentMethod category
Coffee 4.50 card food
```

## Notes

- Data is stored in memory and is lost when the program exits (no database).
- This project is intentionally simple — focused on learning Go conventions rather than production features.
