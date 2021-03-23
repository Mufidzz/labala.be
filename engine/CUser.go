package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ofcode.dev/labala-backend/structs"
)

func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		data structs.User
	)

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	data, err = CreateUser(idb, data)
	if err != nil {
		c.JSON(http.StatusFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) GetUserProfile(c *gin.Context) {

}
