
# Ecommerce System – Go
Sistema de Gestión de E-commerce desarrollado en Go para la materia de Programación Orientada a Objetos.

# Autor
Luis Miguel Vasconez


## Tecnologías Utilizadas
Go
Git
GitHub

## Descripción
Este proyecto consiste en el desarrollo de un sistema de gestión de e-commerce mediante consola, utilizando el lenguaje Go y aplicando conceptos fundamentales de programación vistos en la Unidad 1.

El sistema permite gestionar:

Productos
Carrito de compras
Pedidos
Usuarios

El proyecto fue desarrollado utilizando programación funcional y una arquitectura modular basada en funciones y paquetes.

## Funcionalidades

### Gestión de Productos
Agregar productos
Mostrar productos
Buscar productos
Filtrar productos
Gestionar stock

### Gestión de Carrito
Añadir productos al carrito
Eliminar productos
Mostrar carrito
Calcular total

### Gestión de Pedidos
Confirmar compras
Generar pedidos
Mostrar historial

### Gestión de Usuarios
Registro básico de usuarios
Gestión de perfiles

## Estructura del Proyecto

Ecommerce-System/
│
├── main.go
├── go.mod
│
├── models/
│   ├── product.go
│   ├── cart.go
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
│   ├── validators.go
│   └── helpers.go
│
├── data/
│   └── mock_data.go
│
└── README.md

## Paquetes Utilizados
### Paquetes Estándar de Go
Paquete	        Uso
fmt:	        Entrada y salida de datos
strings:	    Manipulación de texto
strconv:	    Conversión de tipos
time:	        Manejo de fechas
bufio:	        Lectura desde consola
os:	            Interacción con el sistema
errors:	        Manejo de errores

## Programación Funcional Aplicada

El proyecto implementa programación funcional mediante:

Funciones reutilizables
Funciones puras
Funciones anónimas
Manipulación de slices
Separación de lógica

## Flujo del Sistema
1. Ver productos
2. Agregar producto
3. Buscar producto
4. Añadir al carrito
5. Ver carrito
6. Comprar
7. Ver pedidos
8. Salir

## Alcance del Proyecto
### Incluye
Gestión básica de productos
Gestión de carrito
Gestión de pedidos
Control de stock
Arquitectura modular
Interacción mediante consola

### No incluye
Base de datos
Pasarela de pagos
Interfaz web
APIs
Persistencia de datos
