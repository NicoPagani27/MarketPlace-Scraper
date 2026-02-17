package servicios

import (
	"proyecto-MarketplaceScraper/modelos"
	"strings"
	"sync"
)

type ResultadoTienda struct {
	Tienda    string
	Productos []modelos.Producto
	Error     error
}

func BuscarProductosConcurrente(busqueda string) ([]modelos.Producto, error) {
	resultados := make(chan ResultadoTienda, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		productos, err := ConsultarFakeStore()
		resultados <- ResultadoTienda{Tienda: "Fake Store", Productos: productos, Error: err}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		productos, err := ConsultarDummyJSON()
		resultados <- ResultadoTienda{Tienda: "Dummy JSON", Productos: productos, Error: err}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		productos, err := ConsultarPlatzi()
		resultados <- ResultadoTienda{Tienda: "Platzi Store", Productos: productos, Error: err}
	}()

	go func() {
		wg.Wait()
		close(resultados)
	}()

	var todosProductos []modelos.Producto
	// aprendi los ... sacá los elementos de esta lista y agregá cada uno individualmente
	for resultado := range resultados {
		todosProductos = append(todosProductos, resultado.Productos...)
	}
	var productosFiltrados []modelos.Producto
	for _, producto := range todosProductos {
		if strings.Contains(strings.ToLower(producto.Titulo), strings.ToLower(busqueda)) {
			productosFiltrados = append(productosFiltrados, producto)
		}
	}

	return productosFiltrados, nil
}
