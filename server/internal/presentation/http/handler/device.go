package handler

import (
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type DeviceHandler struct{}

func NewDeviceHandler() *DeviceHandler {
	return &DeviceHandler{}
}

// UpdatePosition FindAll godoc
// @Summary UpdateDevicePosition
// @Tags device
// @ID updateDevicePosition
// @Param data body request.DeviceRequest true "Device Request"
// @Success 200
// @Router /devices/position [post]
func (dh *DeviceHandler) UpdatePosition(c *gin.Context) {
	var body request.DeviceRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Printf("body: %v", body)

	c.Status(http.StatusOK)
}
