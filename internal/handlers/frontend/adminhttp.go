package frontend

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"

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

type LoginTracker struct {
	attempts   map[string]int
	lastFailed map[string]time.Time
	mutex      sync.Mutex
	timeout    time.Duration
	maxFailed  int
}

var loginTracker = LoginTracker{
	attempts:   make(map[string]int),
	lastFailed: make(map[string]time.Time),
	timeout:    5 * time.Minute,
	maxFailed:  10,
}

func (h *HTTPHandler) Login(c *gin.Context) {
	req := LoginRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ip := c.ClientIP()

	loginTracker.mutex.Lock()
	defer loginTracker.mutex.Unlock()

	// Check if the IP has a recent failed attempt
	lastFailedTime, exists := loginTracker.lastFailed[ip]
	if exists && time.Since(lastFailedTime) < loginTracker.timeout {
		log.Printf("Login failed for user %s from IP %s. Too many attempts. Remaining time: %v", req.Username, ip, loginTracker.timeout-time.Since(lastFailedTime))
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many login attempts. Please try again later."})
		return
	}

	if exists && time.Since(lastFailedTime) >= loginTracker.timeout {
		log.Printf("Time since last failed login for user %s from IP %s is %v. Resetting.", req.Username, ip, time.Since(lastFailedTime).String())
		delete(loginTracker.lastFailed, ip)
		delete(loginTracker.attempts, ip)
	}

	attempts, exists := loginTracker.attempts[ip]
	if exists && attempts >= loginTracker.maxFailed {
		loginTracker.lastFailed[ip] = time.Now()
		log.Printf("Max failed login attempts reached for user %s from IP %s. Blocking for %v", req.Username, ip, loginTracker.timeout.String())
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many login attempts. Please try again later."})
		return
	}

	token, err := h.userSrv.Login(req.Username, req.Password)
	if err != nil {
		attempts++
		loginTracker.attempts[ip] = attempts
		log.Printf("Login failed for user %s from IP %s. %d attempts so far.", req.Username, ip, attempts)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	delete(loginTracker.attempts, ip)
	delete(loginTracker.lastFailed, ip)

	domain := c.Request.Host
	log.Println(domain)
	c.SetCookie("wedding_cookie", token, 3600*24, "/", domain, false, true)
	c.Header("HX-Redirect", "/crudguests")
}
