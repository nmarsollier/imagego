<!-- cSpell:language es -->

### Si queres sabes mas sobre mi:

[Nestor Marsollier](https://github.com/nmarsollier/profile)

# Image Service en GO

Image Service en GO reemplaza la version realizada en Node del proyecto [ecommerce](https://github.com/nmarsollier/ecommerce).

Si bien esta desarrollado con fines académicos, si se refinan los detalles puede utilizarse en producción.

Este microservicio recibe y almacena imágenes en formato base64 en una base de datos redis.

El cliente puede solicitar las imágenes en distintos tamaños, cada tamaño se ajusta y se guarda en la base de datos para una mejor velocidad de acceso en futuras llamadas.

Las imágenes pueden recuperarse en formato base64 o bien en formato jpeg.

[Documentación de API](./README-API.md)

La documentación de las api también se pueden consultar desde el home del microservicio
que una vez levantado el servidor se puede navegar en [localhost:3001](http://localhost:3001/docs/index.html)

El servidor GraphQL puede navegar en [localhost:4001](http://localhost:4001/)

## Directorios

- **iamge:** Logica de negocio del agregado image
- **security:** Validaciones de usuario contra el MS de Auth
- **graph:** Servidor y Controllers GraphQL federation server
- **rabbit:** Servidor y Controllers RabbitMQ
- **rest:** Servidor y Controllers Rest
- **tools:** Herramientas varias

## Dependencias

### Auth

Las imágenes solo pueden subirse y descargarse por usuarios autenticados, ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

## Requisitos

Go [golang.org](https://golang.org/doc/install)

## Configuración inicial

establecer variables de entorno (consultar documentación de la version instalada)

Para descargar el proyecto correctamente hay que ejecutar :

```bash
git clone https://github.com/nmarsollier/imagego $GOPATH/src/github.com/nmarsollier/imagego
```

Una vez descargado, tendremos el código fuente del proyecto en la carpeta

```bash
cd $GOPATH/src/github.com/nmarsollier/imagego
```

## Instalar Librerías requeridas

```bash
git config core.hooksPath .githooks
go install github.com/swaggo/gin-swagger/swaggerFiles
go install github.com/swaggo/gin-swagger
go install github.com/swaggo/swag/cmd/swag
go install github.com/golang/mock/mockgen@v1.6.0
go install github.com/99designs/gqlgen@v0.17.56
```

## Build y ejecución

```bash
go install
imagego
```

## Redis

Las imágenes se almacenan en una instancia de Redis. Seguir los pasos de instalación desde la pagina oficial [redis.io](https://redis.io/download)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## RabbitMQ

Solo usuarios autorizados pueden subir y descargar imágenes. El microservicio [Auth](https://github.com/nmarsollier/ecommerce) es el que identifica usuarios. Auth notifica con un broadcast los logouts en la aplicación para que se vacíen los caches locales de usuario.

Seguir los pasos de instalación en la pagina oficial de [RabbitMQ](https://www.rabbitmq.com/)

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## Apidoc

Usamos [swaggo](https://github.com/swaggo/swag)

Requisitos

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

La documentacion la generamos con el comando

```bash
swag init
```

Para generar el archivo README-API.md

Requisito

```bash
sudo npm install -g swagger-markdown
```

y ejecutamos

```bash
npx swagger-markdown -i ./docs/swagger.yaml -o README-API.md
```

## Configuración del servidor

Este servidor usa las siguientes variables de entorno para configuración :

RABBIT_URL : Url de rabbit (default amqp://localhost)
REDIS_URL : Url de redis (default localhost:6379)
PORT : Puerto (default 3000)
AUTH_SERVICE_URL : Secret para password (default http://localhost:3001)
GQL_PORT : Puerto GraphQL (default 4001)

## Docker

Estos comandos son para dockerizar el repositorio una vez descargado localmente.

### Build

```bash
docker build -t dev-image-go .
```

### El contenedor

Mac | Windows

```bash
docker run -it --name dev-image-go -p 3001:3001 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego dev-image-go
```

Linux

```bash
docker run -it --add-host host.docker.internal:172.17.0.1 --name dev-image-go -p 3001:3001 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego dev-image-go
```
