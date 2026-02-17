package servicios

import (
	"encoding/json"
	"io"
	"net/http"
	"proyecto-MarketplaceScraper/modelos"
	"strings"
)

func ConsultarFakeStore() ([]modelos.Producto, error) {
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
	var productos []modelos.Producto
	err = json.Unmarshal(body, &productos)

	if err != nil {
		return nil, err
	}
	return productos, nil
}

func BuscarProductos(query string) ([]modelos.Producto, error) {
	productos, err := ConsultarFakeStore()
	if err != nil {
		return nil, err
	}
	var resultados []modelos.Producto
	for _, producto := range productos {
		if strings.Contains(strings.ToLower(producto.Titulo), strings.ToLower(query)) {
			resultados = append(resultados, producto)
		}
	}
	return resultados, nil
}
