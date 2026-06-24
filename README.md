# Sistema de Gestión de E-commerce

## Descripción

Sistema de Gestión de E-commerce desarrollado en Go para la asignatura de Programación Orientada a Objetos.

El proyecto implementa una API REST para la gestión de productos, carrito de compras y pedidos, aplicando los principios fundamentales de Programación Orientada a Objetos, estructuras de datos, encapsulación, interfaces, concurrencia y pruebas unitarias.


## Autor

**Luis Miguel Vasconez**


## Tecnologías Utilizadas

* Go
* Git
* GitHub
* API REST
* JSON
* HTTP
* Concurrencia con Mutex
* Testing Unitario


## Objetivos del Proyecto

Desarrollar un sistema de gestión de comercio electrónico aplicando:

* Programación Orientada a Objetos.
* Encapsulación.
* Interfaces.
* Modularización.
* Manejo de errores.
* Estructuras de datos dinámicas.
* Servicios Web REST.
* Concurrencia.
* Pruebas unitarias.


## Arquitectura del Proyecto

sistema_gestion_ecommerce_POO/
│
├── main.go
├── go.mod
│
├── handlers/
│   ├── product_handler.go
│   ├── cart_handler.go
│   └── order_handler.go
│
├── interfaces/
│   └── searchable.go
│
├── models/
│   ├── product.go
│   ├── cart_item.go
│   ├── order.go
│   └── user.go
│
├── services/
│   ├── product_service.go
│   ├── cart_service.go
│   ├── order_service.go
│   └── user_service.go
│
├── utils/
│   ├── helpers.go
│   └── validators.go
│
├── tests/
│   ├── product_test.go
│   ├── cart_test.go
│   └── order_test.go
│
└── README.md


## Funcionalidades Implementadas

### Gestión de Productos

* Consultar productos.
* Crear productos.
* Buscar productos por nombre.
* Gestionar stock.

### Gestión de Carrito

* Añadir productos al carrito.
* Consultar carrito.
* Calcular total de compra.
* Vaciar carrito.

### Gestión de Pedidos

* Generar pedidos.
* Consultar historial de pedidos.

### Gestión de Usuarios

* Modelado de usuarios mediante Programación Orientada a Objetos.


## Endpoints Disponibles

### Productos

```http
GET /products
```

Obtiene todos los productos registrados.

```http
POST /products
```

Crea un nuevo producto.

```http
GET /products/search?name=Producto
```

Busca un producto por nombre.

---

### Carrito

```http
GET /cart
```

Obtiene el contenido actual del carrito.

```http
POST /cart
```

Añade un producto al carrito.

```http
DELETE /cart
```

Vacía completamente el carrito.

---

### Pedidos

```http
POST /checkout
```

Genera un pedido utilizando los productos del carrito.

```http
GET /orders
```

Obtiene el historial de pedidos realizados.

---

## Programación Orientada a Objetos Aplicada

El proyecto implementa los siguientes conceptos:

### Encapsulación

Los atributos de las entidades se encuentran protegidos mediante campos privados y métodos getter.

Ejemplo:

* Product
* User
* Order
* CartItem

### Interfaces

Se implementa una interfaz para definir comportamientos reutilizables dentro del sistema.

### Modularización

El sistema se encuentra dividido en:

* Models
* Services
* Handlers
* Interfaces
* Utils
* Tests

---

## Concurrencia

Para garantizar la integridad de los datos en accesos simultáneos, se implementó:

```go
sync.Mutex
```

en los servicios:

* ProductService
* CartService
* OrderService

---

## Pruebas Unitarias

El proyecto incluye pruebas unitarias para validar el funcionamiento de los servicios principales.

Pruebas implementadas:

* TestAddProduct()
* TestFindProductByName()
* TestAddToCart()
* TestCalculateTotal()
* TestCreateOrder()

Ejecutar pruebas:

```bash
go test ./...
```

---

## Ejecución del Proyecto

### Compilar

```bash
go build .
```

### Ejecutar

```bash
go run .
```

Servidor:

```plaintext
http://localhost:8080
```

---

## Ejemplos de Uso

Consultar productos:

```bash
curl http://localhost:8080/products
```

Buscar producto:

```bash
curl "http://localhost:8080/products/search?name=Mouse"
```

Agregar producto al carrito:

```bash
curl -X POST \
-H "Content-Type: application/json" \
-d '{"product_name":"Mouse"}' \
http://localhost:8080/cart
```

Consultar carrito:

```bash
curl http://localhost:8080/cart
```

Generar pedido:

```bash
curl -X POST http://localhost:8080/checkout
```

---

## Limitaciones Actuales

* No utiliza base de datos.
* Los datos se almacenan temporalmente en memoria.
* No incluye autenticación de usuarios.
* No incluye pasarela de pagos.
* No incluye persistencia permanente.

---

## Trabajo Académico

Proyecto desarrollado como parte de la asignatura de Programación Orientada a Objetos, aplicando los conceptos estudiados durante las Unidades 1, 2, 3 y 4 del curso.
