package userhandler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRouts(router *gin.RouterGroup) {
	router.POST("/login", h.loginUser)
	router.POST("/register", h.registerUser)
}

func (h *Handler) loginUser(c *gin.Context) {

}

func (h *Handler) registerUser(c *gin.Context) {

}
