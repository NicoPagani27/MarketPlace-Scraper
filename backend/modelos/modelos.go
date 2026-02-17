package modelos

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
