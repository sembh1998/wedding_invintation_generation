package userhdl

import (
	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
)

type HTTPHandler struct {
	userSrv ports.UserSrv
}

func New(userSrv ports.UserSrv) *HTTPHandler {
	return &HTTPHandler{
		userSrv: userSrv,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *HTTPHandler) Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userSrv.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
