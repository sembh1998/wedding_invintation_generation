package userhdl

import (
	"net/http"

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
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}

func (h *HTTPHandler) Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userSrv.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("wedding_cookie", token, 3600*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
