package expensehandler

import (
	"database/sql"
	"fmt"
	"splisense/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateExpense(expense types.Expense) (int, error) {
	result, err := s.db.Exec(
		"INSERT INTO expenses (description, amount, paid_by, group_id) VALUES (?, ?, ?, ?)",
		expense.Description,
		expense.Amount,
		expense.PaidBy,
		expense.GroupID,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create expense: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %v", err)
	}

	return int(id), nil
}

func (s *Store) CreateExpenseSplit(split types.ExpenseSplit) error {
	_, err := s.db.Exec(
		"INSERT INTO expense_splits (expense_id, user_id, amount, status) VALUES (?, ?, ?, ?)",
		split.ExpenseID,
		split.UserID,
		split.Amount,
		split.Status,
	)
	if err != nil {
		return fmt.Errorf("failed to create expense split: %v", err)
	}

	return nil
}

func (s *Store) GetExpensesByGroup(groupID int) ([]types.Expense, error) {
	rows, err := s.db.Query(
		"SELECT id, description, amount, paid_by, group_id, created_at FROM expenses WHERE group_id = ?",
		groupID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query expenses: %v", err)
	}
	defer rows.Close()

	var expenses []types.Expense
	for rows.Next() {
		var expense types.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Description,
			&expense.Amount,
			&expense.PaidBy,
			&expense.GroupID,
			&expense.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense: %v", err)
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (s *Store) GetExpenseSplits(expenseID int) ([]types.ExpenseSplit, error) {
	rows, err := s.db.Query(
		"SELECT id, expense_id, user_id, amount, status, created_at FROM expense_splits WHERE expense_id = ?",
		expenseID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query expense splits: %v", err)
	}
	defer rows.Close()

	var splits []types.ExpenseSplit
	for rows.Next() {
		var split types.ExpenseSplit
		err := rows.Scan(
			&split.ID,
			&split.ExpenseID,
			&split.UserID,
			&split.Amount,
			&split.Status,
			&split.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan expense split: %v", err)
		}
		splits = append(splits, split)
	}

	return splits, nil
}

func (s *Store) UpdateExpenseSplitStatus(splitID int, status string) error {
	_, err := s.db.Exec(
		"UPDATE expense_splits SET status = ? WHERE id = ?",
		status,
		splitID,
	)
	if err != nil {
		return fmt.Errorf("failed to update expense split status: %v", err)
	}

	return nil
}
