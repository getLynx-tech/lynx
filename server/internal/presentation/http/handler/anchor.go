package handler

import (
	"github.com/getLynx-tech/lynx/internal/application"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/request"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnchorHandler struct {
	anchorApplication *application.AnchorApplication
}

func NewAnchorHandler(
	anchorApplication *application.AnchorApplication,
) *AnchorHandler {
	return &AnchorHandler{
		anchorApplication: anchorApplication,
	}
}

// UpsertAnchors FindAll godoc
// @Summary UpsertAnchors
// @Tags anchor
// @ID upsertAnchors
// @Param data body request.AnchorsRequest true "Anchors Request"
// @Success 201
// @Router /anchors [post]
func (ah *AnchorHandler) UpsertAnchors(c *gin.Context) {
	var body request.AnchorsRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	anchors := make([]*value.Anchor, len(body.Anchors))
	for i, anchor := range body.Anchors {
		anchors[i] = &value.Anchor{
			Id: anchor.Id,
			X:  anchor.X,
			Y:  anchor.Y,
		}
	}

	err := ah.anchorApplication.UpsertAnchors(c.Request.Context(), anchors)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Status(http.StatusCreated)
}

// GetAllAnchors FindAll godoc
// @Summary GetAllAnchors
// @Tags anchor
// @ID getAllAnchors
// @Success 200 {array} response.Anchor
// @Router /anchors [get]
func (ah *AnchorHandler) GetAllAnchors(c *gin.Context) {
	anchors, err := ah.anchorApplication.GetAllAnchors(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, response.NewAnchors(anchors))
}
