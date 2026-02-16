package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	router.GET("/api/tiendas", func(c *gin.Context) { c.JSON(200, tiendas) })
	fmt.Println("corriendo en http://localhost:8080")
	router.Run(":8080")
}
