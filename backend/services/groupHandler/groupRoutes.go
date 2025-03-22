package grouphandler

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
	store types.GroupStore
}

func NewHandler(store types.GroupStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	groups := router.Group("/groups")
	groups.Use(auth.WithJWTAuth(nil)) // Apply JWT auth middleware
	{
		groups.POST("", h.createGroup)
		groups.GET("/:groupId", h.getGroup)
		groups.GET("/:groupId/members", h.getGroupMembers)
	}
}

func (h *Handler) createGroup(c *gin.Context) {
	var payload types.CreateGroupPayload
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

	// Create the group
	group := types.Group{
		Name:        payload.Name,
		Description: payload.Description,
		CreatedBy:   userID,
	}

	groupID, err := h.store.CreateGroup(group)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// Add the creator as a member
	member := types.GroupMember{
		GroupID: groupID,
		UserID:  userID,
	}

	if err := h.store.AddGroupMember(member); err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	// Add other members
	for _, memberID := range payload.Members {
		member := types.GroupMember{
			GroupID: groupID,
			UserID:  memberID,
		}

		if err := h.store.AddGroupMember(member); err != nil {
			utils.WriteError(c, http.StatusInternalServerError, err)
			return
		}
	}

	utils.WriteJSON(c, http.StatusCreated, map[string]int{"groupId": groupID})
}

func (h *Handler) getGroup(c *gin.Context) {
	groupID, err := utils.ParseIntParam(c, "groupId")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	group, err := h.store.GetGroupByID(groupID)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, group)
}

func (h *Handler) getGroupMembers(c *gin.Context) {
	groupID, err := utils.ParseIntParam(c, "groupId")
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, err)
		return
	}

	members, err := h.store.GetGroupMembers(groupID)
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(c, http.StatusOK, members)
}
