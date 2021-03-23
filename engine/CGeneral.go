package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) AvailableSoon(c *gin.Context) {
	c.JSON(http.StatusNoContent, "Available Soon")
}
