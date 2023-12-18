
# API Almacenes Autorizador
<br/>

## Descripción :memo:

Api RESTful en Golang utilizado para otorgar autorizacion a los clientes warehouse que se integran por API's a Andreani.

<br/>

## Instalación :wrench:

* `go mod tidy`
* `go build`
* `go run main.go`

<br/>

## Testing :triangular_flag_on_post:


* En el project ejectuar `go test ./...`
* Para ver el coverage de los tests: `go test ./... -cover`

<br/>

## Endpoints 🔗

Default

[GET] http://localhost:8080/api/doc/index.html

<br/>

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

## Maintainers :man_firefighter:

Proyecto_APIsClientes@andreani.com
