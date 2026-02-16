package main

import (
	"fmt"
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

func main() {
	fmt.Println("Pruebo que funcione")
}
