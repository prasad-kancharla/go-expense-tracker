package expense

import (
	"errors"
)

var expenses = make(map[int64]Expense)

func save(expense Expense) {
	expenses[expense.ID] = expense
}

func getAll() []Expense {
	expenseList := make([]Expense, 0, len(expenses))
	for _, v := range expenses {
		expenseList = append(expenseList, v)
	}
	return expenseList
}

func findByID(id int64) (Expense, error) {
	expense, idPresent := expenses[id]
	if !idPresent {
		return Expense{}, errors.New("invalid expense id")
	}
	return expense, nil
}

func remove(id int64) error {
	_, idPresent := expenses[id]
	if !idPresent {
		return errors.New("invalid expense id")
	}
	delete(expenses, id)
	return nil
}
