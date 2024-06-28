## Proyecto de Gestión de Libros
**Nombre de la Página Web:** Biblioteca Cerebritos 

## Descripción
Este proyecto es sobre un sistema de gestión de libros electrónicos donde se pueden administrar autores, editoriales, géneros y libros. Los usuarios pueden registrarse, autenticarse y buscar libros por género para poder sus detalles.
## Grupo
Este proyecto fue desarrollado por 
[Bryan Espinoza 
Gustavo Montaguano 
Michael Tandazo 
Camila Ortega].
## Objetivo del Programa
El objetivo de este programa es proporcionar una plataforma para la gestión y consulta de libros, facilitando la administración de información relacionada con autores, editoriales, géneros y usuarios.

## Fecha
Fecha: [30 de junio de 2024]

## Tecnologías Utilizadas

- HTML, Bootstrap
- JavaScript (jQuery)
- Go (para el backend)
- MySQL (base de datos)

## Funcionalidades Principales de los Paquetes 

- **Gestión de Autores:** Permite listar todos los autores almacenados en la base de datos y buscar detalles de un autor por ID específico.
- **Gestión de Editoriales:** Lista editoriales disponibles.
- **Gestión de Géneros:** Proporciona una lista completa de géneros de libros almacenados en la base de datos y permite obtener detalles de un género por su ID.
- **Gestión de Libros:** Ofrece la capacidad de listar todos los libros registrados, incluyendo detalles como título, año de publicación, descripción, y detalles de autor, editorial y género asociados. Permite buscar libros por género específico.
- **Autenticación de Usuarios:** Permite a los usuarios registrarse en el sistema proporcionando un correo electrónico y una contraseña. Implementa validaciones de seguridad para garantizar que las contraseñas cumplan con los requisitos. Facilita la autenticación durante el inicio de sesión validando la combinación de correo electrónico y contraseña en la base de datos.

## Estructura de Carpetas y Archivos

- **main.go**: Contiene la lógica principal del servidor web, configuración de rutas, conexión a la base de datos y el manejo de solicitudes HTTP.
- **services/app.go**: Define métodos para manejar peticiones relacionadas con autores, editoriales, géneros y libros.
- **templates**: Directorio que almacena las plantillas HTML utilizadas para renderizar las diferentes páginas web de la aplicación.
- **libro, genero, editorial, autor, usuario**: Paquetes que encapsulan la lógica del programa, relacionada con los libros, géneros y usuarios respectivamente.
 
## Ejecutar el Programa 

1. Clonar el repositorio en una máquina.
2. Tener Go instalado.
3. Instalar las dependencias necesarias especificadas en el archivo go.mod.
4. Configurar una base de datos MySQL.
5. Ejecutar el servidor utilizando go run main.go.
6. Abrir el navegador y utilizar el enlace http://localhost:8080 para interactuar con la aplicación.





