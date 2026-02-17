package servicios

import (
	"encoding/json"
	"io"
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// aca parseo la respuesta de Platzi (array directo)
	var productosPlatzi []modelos.ProductoPlatzi
	err = json.Unmarshal(body, &productosPlatzi)
	if err != nil {
		return nil, err
	}

	var productos []modelos.Producto

	for _, prod := range productosPlatzi {
		// lo quer hago para convertir el array es tomar la primera imagen o string vacío
		var imagen string
		if len(prod.Imagenes) > 0 {
			imagen = prod.Imagenes[0]
		} else {
			imagen = ""
		}

		productos = append(productos, modelos.Producto{
			ID:          prod.ID,
			Titulo:      prod.Titulo,
			Precio:      prod.Precio,
			Descripcion: prod.Descripcion,
			Imagen:      imagen,
			Categoria:   "", // Platzi tiene objeto, lo dejamos vacío viene asi
			Raiting: modelos.Puntuaciones{
				Puntuacion: 0, // Platzi no tiene rating
				Cantidad:   0,
			},
		})
	}

	return productos, nil
}
