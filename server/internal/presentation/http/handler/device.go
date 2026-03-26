package handler

import (
	"encoding/json"
	"github.com/getLynx-tech/lynx/internal/application"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/request"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type DeviceHandler struct {
	deviceApplication *application.DeviceApplication
}

func NewDeviceHandler(deviceApplication *application.DeviceApplication) *DeviceHandler {
	return &DeviceHandler{
		deviceApplication: deviceApplication,
	}
}

// GetAllDevices FindAll godoc
// @Summary GetAllDevices
// @Tags device
// @ID getAllDevices
// @Success 200 {array} response.Device
// @Router /devices [get]
func (dh *DeviceHandler) GetAllDevices(c *gin.Context) {
	devices, err := dh.deviceApplication.GetAllDevices(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response.NewDevices(devices))
}

// UpsertPosition FindAll godoc
// @Summary UpsertDevicePosition
// @Tags device
// @ID upsertDevicePosition
// @Param data body request.DeviceRequest true "Device Request"
// @Success 200
// @Router /devices/position [post]
func (dh *DeviceHandler) UpsertPosition(c *gin.Context) {
	var body request.DeviceRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	b, _ := json.Marshal(body)
	log.Printf("%s", string(b))

	readings := make([]*value.Reading, len(body.Readings))
	for i, reading := range body.Readings {
		readings[i] = &value.Reading{
			Id:   reading.Id,
			Rssi: reading.Rssi,
			Dist: reading.Dist,
		}
	}

	device := &value.Device{
		DeviceId: body.DeviceId,
		IsActive: body.IsActive,
		Readings: readings,
	}
	err := dh.deviceApplication.UpsertDevicePosition(c.Request.Context(), device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.Status(http.StatusOK)
}
