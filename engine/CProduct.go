package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ofcode.dev/labala-backend/structs"
)

func (idb *InDB) GetProductUsingID(c *gin.Context) {
	var (
		data interface{}
		err  error
	)
	id := c.Param("q1")

	if c.Query("ProductExchange") == "true" {
		data, err = GetSingleProductWProductExchange(idb, "id", id)

	} else {
		data, err = GetSingleProduct(idb, "id", id)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) CreateProduct(c *gin.Context) {
	var (
		data structs.ProductWProductExchange
	)

	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = CreateProduct(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) UpdateProduct(c *gin.Context) {
	var (
		data structs.Product
	)

	id := c.Param("q1")
	data, err := GetSingleProduct(idb, "id", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	data, err = UpdateProduct(idb, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func (idb *InDB) DeleteProduct(c *gin.Context) {
	var (
		data structs.Product
	)

	id := c.Param("q1")
	data, err := DeleteProduct(idb, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
