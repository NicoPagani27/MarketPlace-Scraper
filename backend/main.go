package main

import (
	"fmt"

	"proyecto-MarketplaceScraper/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("iniciando servidor")

	router := gin.Default()

	router.GET("/api/tiendas", handlers.ObtenerTiendas)
	router.POST("/api/productos/buscar", handlers.BuscarProductos)

	fmt.Println("corriendo en http://localhost:8080")
	router.Run(":8080")
}
