package servicios

import (
	"net/http"

	"proyecto-MarketplaceScraper/modelos"
)

func ConsultarDummyJSON() ([]modelos.Producto, error) {
	url := "https://dummyjson.com/products"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var productos []modelos.Producto
	return productos, nil

}
