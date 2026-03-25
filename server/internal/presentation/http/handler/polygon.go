package handler

import (
	"github.com/getLynx-tech/lynx/internal/application"
	"github.com/getLynx-tech/lynx/internal/domain/value"
	"github.com/getLynx-tech/lynx/internal/presentation/http/dto/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PolygonHandler struct {
	polygonApplication *application.PolygonApplication
}

func NewPolygonHandler(polygonApplication *application.PolygonApplication) *PolygonHandler {
	return &PolygonHandler{
		polygonApplication: polygonApplication,
	}
}

// CreatePolygon FindAll godoc
// @Summary Create Polygon
// @Tags polygon
// @ID createPolygon
// @Param X-User-ID header string true "User ID"
// @Param data body request.CreatePolygon true "Create Polygon"
// @Success 201
// @Router /polygon [post]
func (ph *PolygonHandler) CreatePolygon(c *gin.Context) {
	var polygon request.CreatePolygon
	if err := c.BindJSON(&polygon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	vertices := make([]*value.Vertex, len(polygon.Vertices))
	for i, vertex := range polygon.Vertices {
		vertices[i] = &value.Vertex{
			Longitude: vertex.Longitude,
			Latitude:  vertex.Latitude,
		}
	}

	polygonValue := &value.Polygon{
		Id:       polygon.Id,
		Vertices: vertices,
	}

	err := ph.polygonApplication.CreatePolygon(c.Request.Context(), polygonValue)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusCreated)
}
