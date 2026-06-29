package expense

import "time"

var expenseID int = 1

type Expense struct {
	ID            int64
	Title         string
	Amount        float64
	CategoryName  string
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// type Category struct {
// 	Id   int64
// 	Name string
// }

func newExpense(title string, amount float64, paymentMethod string, categoryName string) Expense {
	newExpense := Expense{
		ID:            int64(expenseID),
		Title:         title,
		Amount:        amount,
		PaymentMethod: paymentMethod,
		CategoryName:  categoryName,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
	expenseID++
	return newExpense
}
