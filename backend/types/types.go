package types

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	// note that this isn't the best way to handle quantity
	// because it's not atomic (in ACID), but it's good enough for this example
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}

type CartCheckoutItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userID"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"orderID"`
	ProductID int       `json:"productID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProducts() ([]*Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items" validate:"required"`
}

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description" validate:"required"`
	Amount      float64   `json:"amount" validate:"required"`
	PaidBy      int       `json:"paidBy" validate:"required"`
	GroupID     int       `json:"groupId" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ExpenseSplit struct {
	ID        int       `json:"id"`
	ExpenseID int       `json:"expenseId" validate:"required"`
	UserID    int       `json:"userId" validate:"required"`
	Amount    float64   `json:"amount" validate:"required"`
	Status    string    `json:"status" validate:"required,oneof=pending paid"`
	CreatedAt time.Time `json:"createdAt"`
}

type Group struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	CreatedBy   int       `json:"createdBy" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GroupMember struct {
	ID        int       `json:"id"`
	GroupID   int       `json:"groupId" validate:"required"`
	UserID    int       `json:"userId" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateExpensePayload struct {
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	GroupID     int     `json:"groupId" validate:"required"`
	Splits      []struct {
		UserID int     `json:"userId" validate:"required"`
		Amount float64 `json:"amount" validate:"required"`
	} `json:"splits" validate:"required"`
}

type CreateGroupPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Members     []int  `json:"members" validate:"required"`
}

type ExpenseStore interface {
	CreateExpense(Expense) (int, error)
	CreateExpenseSplit(ExpenseSplit) error
	GetExpensesByGroup(groupID int) ([]Expense, error)
	GetExpenseSplits(expenseID int) ([]ExpenseSplit, error)
	UpdateExpenseSplitStatus(splitID int, status string) error
}

type GroupStore interface {
	CreateGroup(Group) (int, error)
	AddGroupMember(GroupMember) error
	GetGroupByID(id int) (*Group, error)
	GetGroupMembers(groupID int) ([]int, error)
}
