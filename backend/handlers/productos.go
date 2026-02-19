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

func ObtenerProductosPorTienda(peticion *gin.Context) {
	tiendaId := peticion.Param("tiendaId")
	switch tiendaId {
	case "fakestore":
		productos, err := servicios.ConsultarFakeStore()
		if err != nil {
			peticion.JSON(500, gin.H{"error": err.Error()})
			return
		}
		peticion.JSON(200, gin.H{"tienda": tiendaId, "total": len(productos), "productos": productos})
	case "dummyjson":
		productos, err := servicios.ConsultarDummyJSON()
		if err != nil {
			peticion.JSON(500, gin.H{"error": err.Error()})
			return
		}
		peticion.JSON(200, gin.H{"tienda": tiendaId, "total": len(productos), "productos": productos})
	case "platzi":
		productos, err := servicios.ConsultarPlatzi()
		if err != nil {
			peticion.JSON(500, gin.H{"error": err.Error()})
			return
		}
		peticion.JSON(200, gin.H{"tienda": tiendaId, "total": len(productos), "productos": productos})
	default:
		peticion.JSON(404, gin.H{"error": "Tienda no encontrada"})
	}
}
