package expense

import (
	"errors"
	"slices"
	"time"
)

func All() []Expense {
	return getAll()
}

func Create(title string, amount float64, paymentMethod string, categoryName string) (Expense, error) {
	expense := newExpense(title, amount, paymentMethod, categoryName)
	save(expense)
	return expense, nil
}

func UpdateAmount(id int64, newAmount float64) error {
	expense, err := findByID(id)
	if err != nil {
		return err
	}
	expense.Amount = newAmount
	expense.UpdatedAt = time.Now().UTC()
	save(expense)
	return nil
}

func UpdateTitle(id int64, newTitle string) error {
	expense, err := findByID(id)
	if err != nil {
		return err
	}
	expense.Title = newTitle
	expense.UpdatedAt = time.Now().UTC()
	save(expense)
	return nil
}

func Find(id int64) (Expense, error) {
	return findByID(id)
}

func LastFive() ([]Expense, error) {
	expensesList := getAll()

	slices.SortFunc(expensesList, func(e1 Expense, e2 Expense) int {
		if e1.CreatedAt.Before(e2.CreatedAt) {
			return -1
		} else if e2.CreatedAt.Before(e1.CreatedAt) {
			return 1
		} else {
			return 0
		}
	})

	if len(expensesList) < 5 {
		return nil, errors.New("please log atleast 5 transactions")
	}
	return expensesList[len(expensesList)-5:], nil
}

func AverageByCategory() map[string]float64 {
	expenses := getAll()
	categoryExpensesMap := make(map[string][]Expense)

	for _, e := range expenses {
		category := e.CategoryName
		categoryExpensesMap[category] = append(categoryExpensesMap[category], e)
	}

	categoryAverageMap := make(map[string]float64)

	for cat, exps := range categoryExpensesMap {
		var categorySum float64
		for _, v := range exps {
			categorySum += v.Amount
		}
		categoryAverageMap[cat] = categorySum / float64(len(exps))
	}

	return categoryAverageMap
}

func Remove(id int64) error {
	return remove(id)
}
