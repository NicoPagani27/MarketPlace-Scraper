package main

import (
	"fmt"

	"proyecto-MarketplaceScraper/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("iniciando servidor")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"}, //aprendi que es un header de peticion ej tipo de contenido application/json
		AllowCredentials: true,
	}))

	router.GET("/api/tiendas", handlers.ObtenerTiendas)
	router.POST("/api/productos/buscar", handlers.BuscarProductos)
	router.GET("/api/productos/:tiendaId", handlers.ObtenerProductosPorTienda)

	fmt.Println("corriendo en http://localhost:8080")
	router.Run(":8080")
}
