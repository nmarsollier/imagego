### Si queres sabes mas sobre mi:
[Nestor Marsollier](https://github.com/nmarsollier/profile)

# Image Service en GO

Image Service en GO reemplaza la version realizada en Node del proyecto  [ecommerce](https://github.com/nmarsollier/ecommerce).

Si bien esta desarrollado con fines académicos, si se refinan los detalles puede utilizarse en producción.

Este microservicio recibe y almacena imágenes en formato base64 en una base de datos redis.

El cliente puede solicitar las imágenes en distintos tamaños, cada tamaño se ajusta y se guarda en la base de datos para una mejor velocidad de acceso en futuras llamadas.

Las imágenes pueden recuperarse en formato base64 o bien en formato jpeg.

[Documentación de API](./README-API.md)

## Dependencias

### Auth

Las imágenes solo pueden subirse y descargarse por usuarios autenticados, ver la arquitectura de microservicios de [ecommerce](https://github.com/nmarsollier/ecommerce).

## Requisitos

Go 1.14  [golang.org](https://golang.org/doc/install)

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
go mod download
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

Apidoc es una herramienta que genera documentación de apis para proyectos node (ver [Apidoc](http://apidocjs.com/)).

El microservicio muestra la documentación como archivos estáticos si se abre en un browser la raíz del servidor [localhost:3001](http://localhost:3001/)

Ademas se genera la documentación en formato markdown.

Para que funcione correctamente hay que instalarla globalmente con

```bash
npm install apidoc -g
npm install -g apidoc-markdown2
```

La documentación necesita ser generada manualmente ejecutando la siguiente linea en la carpeta imagego :

```bash
apidoc -o www
apidoc-markdown2 -p www -o README-API.md
```

Esto nos genera una carpeta con la documentación, esta carpeta debe estar presente desde donde se ejecute imagego, imagego busca ./www para localizarlo, aunque se puede configurar desde el archivo de properties.

## Configuración del servidor

Este servidor usa las siguientes variables de entorno para configuración :

RABBIT_URL : Url de rabbit (default amqp://localhost)
REDIS_URL : Url de redis (default localhost:6379)
PORT : Puerto (default 3000)
WWW_PATH : Path donde se ubica la documentación apidoc (default www)
AUTH_SERVICE_URL : Secret para password (default http://localhost:3000)

## Docker

### Build

```bash
docker build -t dev-image-go .
```

### El contenedor

```bash
# Mac | Windows
docker run -it --name dev-image-go -p 3001:3001 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego dev-image-go

# Linux
docker run -it --add-host host.docker.internal:172.17.0.1 --name dev-image-go -p 3001:3001 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego dev-image-go
```

### Debug en Docker

Existe un archivo Docker.debug, hay que armar la imagen usando ese archivo.

```bash
docker build -t debug-image-go -f Dockerfile.debug .
```

```bash
# Mac | Windows
docker run -it --name debug-image-go -p 3000:3000 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego debug-image-go

# Linux
docker run -it --add-host host.docker.internal:172.17.0.1 --name debug-image-go -p 3000:3000 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego debug-image-go
```

El archivo launch.json debe contener lo siguiente

```bash
{
    "version": "0.2.0",
    "configurations": [
          {
                "name": "Debug en Docker",
                "type": "go",
                "request": "launch",
                "mode": "remote",
                "remotePath": "/go/src/github.com/nmarsollier/imagego",
                "port": 40001,
                "host": "127.0.0.1",
                "program": "${workspaceRoot}",
                "showLog": true
          }
    ]
}
```

En el menú run start debugging se conecta a docker.
