MarketPlace-Scraper

Es un comparador de precios que busca productos en diferentes tiendas. Backend en Go usando goroutines y channels para hacer las consultas en paralelo.

Características

Búsqueda Concurrente

Patrón Fan-out/Fan-in: Implementación con goroutines y channels para consultar diferentes APIs al mismo tiempo.
3 APIs integradas:
  - Fake Store API
  - DummyJSON API
  - Platzi Fake Store API
Filtrado inteligente: Búsqueda por coincidencia en títulos de productos. (Aprendizaje nuevo)

Arquitectura del Proyecto
```
backend/
├── main.go              # Punto de entrada y configuración del servidor
├── modelos/
│   └── modelos.go       # Estructuras de datos unificadas
├── servicios/
│   ├── busqueda_concurrente.go  # Orquestador con goroutines
│   ├── fakestore.go             # Cliente Fake Store API
│   ├── dummyjson.go             # Cliente DummyJSON API
│   └── platzi.go                # Cliente Platzi API
└── handlers/
    ├── productos.go     # Endpoints de productos
    └── tiendas.go       # Endpoint de tiendas disponibles
```

API REST Endpoints

GET /api/tiendas
Obtiene la lista de tiendas disponibles.

POST /api/productos/buscar
Busca productos en todas las tiendas concurrentemente.

GET /api/productos/:tiendaId
Obtiene productos de una tienda específica.


Normalización de Datos
- Todos los productos de diferentes APIs se normalizan a una estructura común
- Parseo: Manejo de diferentes formatos de respuesta (DummyJSON usa thumbnail, Platzi usa array de imágenes)
- Rating: Conversión de diferentes sistemas de puntuación

Configuración CORS
- Habilitado para desarrollo con frontend. 
- Métodos permitidos: GET, POST
- Headers: Content-Type

Instalación y Uso

Lo que se necesita
- Go 1.21 o superior
- Git

Instalación

1. Clonar el repo:
```bash
git clone https://github.com/tu-usuario/MarketPlace-Scraper.git
cd MarketPlace-Scraper
```

2. Instalar dependencias:
```bash
cd backend
go mod download
```

3. Ejecutarlo:
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

Flujo de Trabajo Concurrente

```
Usuario hace búsqueda
        │
        ▼
BuscarProductosConcurrente()
        │
        ├──► goroutine 1: ConsultarFakeStore()
        ├──► goroutine 2: ConsultarDummyJSON()
        └──► goroutine 3: ConsultarPlatzi()
              │
              ▼
        Channel (ResultadoTienda)
              │
              ▼
        Agregación y Filtrado
              │
              ▼
        Respuesta al Cliente
```
