package handlers

import (
	"proyecto-MarketplaceScraper/modelos"
	"proyecto-MarketplaceScraper/servicios"

	"github.com/gin-gonic/gin"
)

func BuscarProductos(c *gin.Context) {
	var req modelos.BusquedaRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	productos, err := servicios.BuscarProductos(req.Query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"query":     req.Query,
		"total":     len(productos),
		"productos": productos,
	})
}
