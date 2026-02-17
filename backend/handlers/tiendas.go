package handlers

import (
	"proyecto-MarketplaceScraper/modelos"

	"github.com/gin-gonic/gin"
)

var tiendas = []modelos.Tienda{
	{ID: 1, Nombre: "Fake Store", URL: "https://fakestoreapi.com/"},
	{ID: 2, Nombre: "Dummy JSON", URL: "https://dummyjson.com/"},
	{ID: 3, Nombre: "Platzi Store", URL: "https://fakeapi.platzi.com/"},
}

func ObtenerTiendas(c *gin.Context) {
	c.JSON(200, tiendas)
}
