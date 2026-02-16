package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func main() {
	fmt.Println("Pruebo que funcione")

	productos, err := consultarFakeStore()
	if err != nil {
		fmt.Println("Error al consultar la API:", err)
		return
	}
	fmt.Println("Productos obtenidos:", len(productos))

	if len(productos) > 0 {
		primer := productos[0]
		fmt.Println("Primer producto:")
		fmt.Printf("ID: %d\n", primer.ID)
		fmt.Printf("Título: %s\n", primer.Titulo)
		fmt.Printf("Precio: $%.2f\n", primer.Precio)
		fmt.Printf("Categoría: %s\n", primer.Categoria)
	}
}
