package servicios

import (
	"net/http"

	"proyecto-MarketplaceScraper/modelos"
)

func ConsultarPlatzi() ([]modelos.Producto, error) {
	url := "https://api.escuelajs.co/api/v1/products"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var productos []modelos.Producto
	return productos, nil
}
