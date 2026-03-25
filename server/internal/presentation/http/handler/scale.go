package handler

import (
	"github.com/getLynx-tech/lynx/internal/application"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/request"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScaleHandler struct {
	scaleApplication *application.ScaleApplication
}

func NewScaleHandler(scaleApplication *application.ScaleApplication) *ScaleHandler {
	return &ScaleHandler{
		scaleApplication: scaleApplication,
	}
}

// CreateScale FindAll godoc
// @Summary CreateScale
// @Tags scale
// @ID createScale
// @Param data body request.ScaleRequest true "Scale Request"
// @Success 200
// @Router /scales [post]
func (sh *ScaleHandler) CreateScale(c *gin.Context) {
	var body request.ScaleRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := sh.scaleApplication.CreateScale(c.Request.Context(), &value.Scale{
		Meters: body.Meters,
		Pixels: body.Pixels,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Status(http.StatusCreated)
}

// GetScale FindAll godoc
// @Summary GetScale
// @Tags scale
// @ID getScale
// @Success 200 {object} response.Scale
// @Router /scales [get]
func (sh *ScaleHandler) GetScale(c *gin.Context) {
	scale, err := sh.scaleApplication.GetScale(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, response.NewScale(scale))
}
