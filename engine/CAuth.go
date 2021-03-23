package engine

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"ofcode.dev/labala-backend/security/SJWT"
	"ofcode.dev/labala-backend/structs"
)

func (idb *InDB) Login(c *gin.Context) {
	var (
		data     structs.User
		dataUser structs.User
	)

	err := c.BindJSON(&dataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = GetSingleUser(idb, "email", dataUser.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(dataUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	token, err := SJWT.Generate(SJWT.Body{Uid: data.ID, Acs: data.Access})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)

}
