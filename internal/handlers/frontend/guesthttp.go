package frontend

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/cmd/bootstrap"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain"
)

func (h *HTTPHandler) CrudGuestsHTMX(c *gin.Context) {
	guests, _ := h.guestSrv.FindGuests()
	data := map[string][]domain.Guest{
		"guests": guests,
	}
	tmpl := template.Must(template.ParseFiles("htmx/crudguests.html"))
	tmpl.Execute(c.Writer, data)
}

// AddGuest
func (h *HTTPHandler) AddGuest(c *gin.Context) {
	guest := domain.Guest{}
	if err := c.ShouldBind(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, _ := c.Get("user_id")
	username, _ := c.Get("username")
	guest.CreatedBy = user_id.(string)
	guest.User.User = username.(string)
	guest.User.ID = user_id.(string)

	guest, err := h.guestSrv.CreateGuest(guest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("HX-Redirect", "/crudguests")
}

// DeleteGuest
func (h *HTTPHandler) DeleteGuest(c *gin.Context) {
	id := c.Param("id")
	err := h.guestSrv.DeleteGuest(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("HX-Redirect", "/crudguests")
}

// FetchGuest
func (h *HTTPHandler) FetchGuestHTMX(c *gin.Context) {
	id := c.Param("id")
	guest, err := h.guestSrv.FetchGuest(id)
	if err != nil {
		// Select a random joke
		randomIndex := rand.Intn(len(bootstrap.Jokes))
		randomJoke := bootstrap.Jokes[randomIndex]

		c.String(http.StatusNotFound, randomJoke)
		return
	}
	data := map[string]domain.Guest{
		"guest": guest,
	}
	tmpl := template.Must(template.ParseFiles("htmx/guest.html"))
	tmpl.Execute(c.Writer, data)
}

type GuestConfirmationRequest struct {
	Attend   string `json:"attend" binding:"required" form:"attend"`
	Response string `json:"response" binding:"required" form:"response"`
}

// AttendConfirmation
func (h *HTTPHandler) AttendConfirmation(c *gin.Context) {
	id := c.Param("id")
	attend := GuestConfirmationRequest{}
	if err := c.ShouldBind(&attend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.guestSrv.AttendConfirmation(id, attend.Attend == "true", attend.Response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("HX-Redirect", "/guest/"+id)
}

// FetchGuest
func (h *HTTPHandler) GiftPreferencesHTMX(c *gin.Context) {

	tmpl := template.Must(template.ParseFiles("htmx/gift-preferences.html"))
	tmpl.Execute(c.Writer, nil)
}
