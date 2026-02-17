package handlers

import (
	"proyecto-MarketplaceScraper/modelos"
	"proyecto-MarketplaceScraper/servicios"

	"github.com/gin-gonic/gin"
)

func BuscarProductos(c *gin.Context) {
	var peticion modelos.BusquedaRequest

	if err := c.BindJSON(&peticion); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	productos, err := servicios.BuscarProductosConcurrente(peticion.Query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"query":     peticion.Query,
		"total":     len(productos),
		"productos": productos,
	})
}
