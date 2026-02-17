package servicios

import (
	"encoding/json"
	"io"
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respuesta modelos.RespuestaDummyJSON
	// aprendi que el unmarcshal es para parsear el json y lo convierte a struct de go
	err = json.Unmarshal(body, &respuesta)
	if err != nil {
		return nil, err
	}
	var productos []modelos.Producto

	for _, prod := range respuesta.Productos {
		productos = append(productos, modelos.Producto{
			ID:          prod.ID,
			Titulo:      prod.Titulo,
			Precio:      prod.Precio,
			Descripcion: prod.Descripcion,
			Imagen:      prod.ImagenChica,
			Categoria:   prod.Categoria,
			Raiting: modelos.Puntuaciones{
				Puntuacion: prod.Rating,
				Cantidad:   0, // DummyJSON no tiene cantidad
			},
		})
	}

	return productos, nil
}
