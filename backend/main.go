package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Puntuaciones struct {
	Puntuacion float64 `json:"rate"`
	Cantidad   int     `json:"count"`
}

type Producto struct {
	ID          int          `json:"id"`
	Titulo      string       `json:"title"`
	Precio      float64      `json:"price"`
	Descripcion string       `json:"description"`
	Imagen      string       `json:"image"`
	Categoria   string       `json:"category"`
	Raiting     Puntuaciones `json:"rating"`
}

type Tienda struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	URL    string `json:"url"`
}

type BusquedaRequest struct {
	Query string `json:"query"`
}

func buscarProductos(query string) ([]Producto, error) {
	productos, err := consultarFakeStore()
	if err != nil {
		return nil, err
	}
	var resultados []Producto
	for _, producto := range productos {
		if strings.Contains(strings.ToLower(producto.Titulo), strings.ToLower(query)) {
			resultados = append(resultados, producto)
		}
	}
	return resultados, nil
}

func consultarFakeStore() ([]Producto, error) {
	url := "https://fakestoreapi.com/products"
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var productos []Producto
	err = json.Unmarshal(body, &productos)

	if err != nil {
		return nil, err
	}
	return productos, nil

}

var tiendas = []Tienda{

	{ID: 1, Nombre: "Fake Store", URL: "https://fakestoreapi.com/"},
	{ID: 2, Nombre: "Dummy JSON", URL: "https://dummyjson.com/"},
	{ID: 3, Nombre: "Platzi Store", URL: "https://fakeapi.platzi.com/"},
}

func main() {
	fmt.Println("Pruebo que funcione")

	router := gin.Default()

	router.POST("/api/productos/buscar", func(c *gin.Context) {
		// Tuve que hacer un if y crear la variable peticion para que con el BindJSON
		// (aprendi que sirve para parsear datos, parsear es convertir datos de un formato a otro)
		// parsee todos los datos de la peticion y no fiera error, tambien aprendi que gin.H es un mapa de string y se usa para enviar respuestas JSON personalizadas
		var peticion BusquedaRequest

		if err := c.BindJSON(&peticion); err != nil {
			c.JSON(400, gin.H{"error": "JSON inválido"})
			return
		}

		productos, err := buscarProductos(peticion.Query)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al buscar productos"})
			return
		}

		c.JSON(200, gin.H{
			"query":     peticion.Query,
			"total":     len(productos),
			"productos": productos,
		})

	})
	router.GET("/api/tiendas", func(c *gin.Context) { c.JSON(200, tiendas) })
	fmt.Println("corriendo en http://localhost:8080")
	router.Run(":8080")
}
