package guesthdl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/domain"
	"github.com/sembh1998/wedding_invitation_generation/internal/core/ports"
)

type HTTPHandler struct {
	guestSrv ports.GuestSrv
}

func New(guestSrv ports.GuestSrv) *HTTPHandler {
	return &HTTPHandler{
		guestSrv: guestSrv,
	}
}

func (h *HTTPHandler) CreateGuest(c *gin.Context) {
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

	c.JSON(http.StatusCreated, gin.H{"guest": guest})
}

func (h *HTTPHandler) FindGuests(c *gin.Context) {
	guests, err := h.guestSrv.FindGuests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"guests": guests})
}

func (h *HTTPHandler) FetchGuest(c *gin.Context) {
	id := c.Param("id")
	guest, err := h.guestSrv.FetchGuest(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"guest": guest})
}

func (h *HTTPHandler) UpdateGuest(c *gin.Context) {
	id := c.Param("id")
	guest := domain.Guest{}
	if err := c.ShouldBind(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guest.ID = id

	guest, err := h.guestSrv.UpdateGuest(guest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"guest": guest})
}

func (h *HTTPHandler) DeleteGuest(c *gin.Context) {
	id := c.Param("id")
	err := h.guestSrv.DeleteGuest(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guest deleted successfully"})
}

type GuestConfirmationRequest struct {
	Attend   string `json:"attend" binding:"required" form:"attend"`
	Response string `json:"response" binding:"required" form:"response"`
}

func (h *HTTPHandler) AttendConfirmation(c *gin.Context) {
	id := c.Param("id")
	attend := GuestConfirmationRequest{}
	if err := c.ShouldBind(&attend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guest, err := h.guestSrv.AttendConfirmation(id, attend.Attend == "true", attend.Response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"guest": guest})
}
