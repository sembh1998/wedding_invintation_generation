package frontend

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
)

type HTTPHandler struct {
	guestSrv ports.GuestSrv
	userSrv  ports.UserSrv
}

func New(guestSrv ports.GuestSrv, userSrv ports.UserSrv) *HTTPHandler {
	return &HTTPHandler{
		guestSrv: guestSrv,
		userSrv:  userSrv,
	}
}
func (h *HTTPHandler) LoginHTMX(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("htmx/login.html"))
	tmpl.Execute(c.Writer, nil)
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
	// relocate to /crudguests
	// Set HX-Redirect header for successful login
	c.Header("HX-Redirect", "/crudguests")
}
