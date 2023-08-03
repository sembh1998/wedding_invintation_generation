package frontend

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
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
	guest.CreatedBy = user_id.(string)

	guest, err := h.guestSrv.CreateGuest(guest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl := template.Must(template.ParseFiles("htmx/crudguests.html"))
	tmpl.ExecuteTemplate(c.Writer, "guest-list-element", guest)

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
