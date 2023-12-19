# Microblog API:
- Microblog API es una aplicación simple desarrollada en Golang que permite a los usuarios enviar mensajes, 
seguir a otros usuarios y ver publicaciones de los usuarios a los que siguen.

## Ejecucion de la API:

  ### Requisitos:
  - Asegúrate de tener Go instalado en tu máquina. Además, necesitarás un servidor MySQL en ejecución.

  ### Repositorio:
  - Cloná el repositorio en tu máquina.

  ### Configuracion de la base de Datos : 
  - Ejecuta los siguientes comandos SQL para crear las tablas necesarias:
  ```sql
  CREATE TABLE followers (
      id INT AUTO_INCREMENT PRIMARY KEY,
      user_id INT NOT NULL,
      follower_id INT NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
      CONSTRAINT fk_follower FOREIGN KEY (follower_id) REFERENCES users (id),
      UNIQUE KEY unique_followers (user_id, follower_id)
  );

  CREATE TABLE users (
      id INT AUTO_INCREMENT PRIMARY KEY,
      username VARCHAR(50) NOT NULL,
      timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

  CREATE TABLE messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    text VARCHAR(250) NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );
  ```
  - Asegúrate de crear algunos usuarios manualmente en la base de datos antes de probar la API, ya que no hay un servicio de creación de usuarios implementado.
  
### Compilacion y Ejecucion:
- Exporta las siguientes variables de entorno en tu consola:
```sh
export GORM_DRIVER="mysql"
export SQL_CONNECTION="root:@tcp(127.0.0.1:3306)/blogging_uala"
export BD_NAME="blogging_uala"
export TABLE_MESSAGES="messages"
export TABLE_FOLLOWERS="followers"
export TABLE_USERS="users"
```
- Ejecuta la API con el siguiente comando:
```sh
go run main.go
```

### Endpoints :
- Enviar Mensajes:
  - Metodo POST
  - URL: http://localhost:8080/microblog/send
  - Cuerpo (JSON):
  ```json
  {
    "username": "nombre_usuario",
    "text": "contenido_del_mensaje"
  }
  ```

- Seguir a otro usuario:
  - Metodo POST
  - URL: http://localhost:8080/microblog/follow
  - Cuerpo (JSON):
  ```json
  {
  "username": "nombre_usuario",
  "followUsername": "usuario_a_seguir"
  }
  ```

- Obtener mensajes de los usuarios que sigo:
  - Metodo GET
  - URL: http://localhost:8080/microblog/messages
  - Cuerpo (JSON):
  ```json
  {
  "username": "nombre_usuario"
  }
  ```

## Base de Datos:
- MySQL como sistema de gestión de bases de datos relacional.
- Estructura básica de la base de datos:
  - **Tabla: users**
    - id (int, clave primaria, autoincremental)
    - username (varchar(50), único)
    - timestamp (TIMESTAMP)

  - **Tabla: messages**
    - id (int, clave primaria, autoincremental)
    - username (varchar(50))
    - text (varchar(255))
    - timestamp (timestamp)

  - **Tabla: followers**
    - id (int, clave primaria, autoincremental)
    - user_id (int, clave externa)
    - follower_id (int, clave externa)

## Flujos de Trabajo:
- La API cuenta con 3 endpoints: 
  - `/send` para publicar mensajes.
  - `/follow` para seguir a otros usuarios.
  - `/messages` para obtener publicaciones de los usuarios seguidos.

## Seguridad:
- No se han implementado medidas de seguridad en este ejercicio.

## Despliegue:
- No hay información de despliegue, ya que es un ejercicio simple.

## Servicios Adicionales:
- No se han integrado servicios adicionales.

## Detalles:
- Este es un proyecto de ejemplo para practicar conceptos básicos de desarrollo en Golang y MySQL.

