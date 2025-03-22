package expensehandler

import (
	"fmt"
	"net/http"
	"splisense/services/auth"
	"splisense/types"
	"splisense/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.ExpenseStore
}

func NewHandler(store types.ExpenseStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	expenses := router.Group("/expenses")
	expenses.Use(auth.WithJWTAuth(nil)) // Apply JWT auth middleware
	{
		expenses.POST("", h.createExpense)
		expenses.GET("/group/:groupId", h.getExpensesByGroup)
		expenses.GET("/:expenseId/splits", h.getExpenseSplits)
		expenses.PUT("/splits/:splitId/status", h.updateExpenseSplitStatus)
	}
}

func (h *Handler) createExpense(c *gin.Context) {
	var payload types.CreateExpensePayload
	if err := utils.ParseJSON(c, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Get the user ID from the context (set by JWT middleware)
	userID := auth.GetUserIDFromContext(c.Request.Context())
	if userID == -1 {
		utils.WriteError(c, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	// Create the expense
	expense := types.Expense{
		Description: payload.Description,
		Amount:      payload.Amount,
		PaidBy:      userID,
		GroupID:     payload.GroupID,
	}

	expenseID, err := h.store.CreateExpense(expense)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// Create expense splits
	for _, split := range payload.Splits {
		expenseSplit := types.ExpenseSplit{
			ExpenseID: expenseID,
			UserID:    split.UserID,
			Amount:    split.Amount,
			Status:    "pending",
		}

		if err := h.store.CreateExpenseSplit(expenseSplit); err != nil {
			utils.WriteError(c, http.StatusInternalServerError, err)
			return
		}
	}

	utils.WriteJSON(c, http.StatusCreated, map[string]int{"expenseId": expenseID})
}

func (h *Handler) getExpensesByGroup(c *gin.Context) {
	groupID, err := utils.ParseIntParam(c, "groupId")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	expenses, err := h.store.GetExpensesByGroup(groupID)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, expenses)
}

func (h *Handler) getExpenseSplits(c *gin.Context) {
	expenseID, err := utils.ParseIntParam(c, "expenseId")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	splits, err := h.store.GetExpenseSplits(expenseID)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, splits)
}

func (h *Handler) updateExpenseSplitStatus(c *gin.Context) {
	splitID, err := utils.ParseIntParam(c, "splitId")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	var payload struct {
		Status string `json:"status" validate:"required,oneof=pending paid"`
	}
	if err := utils.ParseJSON(c, &payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	if err := h.store.UpdateExpenseSplitStatus(splitID, payload.Status); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, nil)
}
