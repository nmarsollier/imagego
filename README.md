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

```bash
export GO111MODULE=on
export GOFLAGS=-mod=vendor
```

Para descargar el proyecto correctamente hay que ejecutar :

```bash
go get github.com/nmarsollier/imagego
```

Una vez descargado, tendremos el código fuente del proyecto en la carpeta

```bash
cd $GOPATH/src/github.com/nmarsollier/imagego
```

## Instalar Librerías requeridas

```bash
go mod download
go mod vendor
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

## Archivo config.json

Este archivo permite configurar imagego, ver ejemplos en config-example.json.
imagego busca el archivo "./config.json". Podemos definir el archivo su ruta completa ejecutando

```bash
imagego [path_to_config.json]
```

Para mas detalles ver el archivo tools/env/env.go