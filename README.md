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
  - `/timeline/user` para obtener publicaciones de los usuarios seguidos.

## Seguridad:
- No se han implementado medidas de seguridad en este ejercicio.

## Despliegue:
- No hay información de despliegue, ya que es un ejercicio simple.

## Servicios Adicionales:
- No se han integrado servicios adicionales.

## Detalles:
- Este es un proyecto de ejemplo para practicar conceptos básicos de desarrollo en Golang y MySQL.

## Env
```sh

export GORM_DRIVER="mysql"
export SQL_CONNECTION="root:@tcp(127.0.0.1:3306)/blogging_uala"
export TABLE_TWEETS="tweets"
export BD_NAME="bloggin_uala"
export TABLE_MESSAGES="messages"
export TABLE_FOLLOWERS="followers"
export TABLE_USERS="users"
```

<br/>

