# Image Service en GO

Este Microservicio de imágenes reemplaza al del proyecto

[Microservicios Image](https://github.com/nmarsollier/ecommerce)

También puede usarse fuera del proyecto, sin embargo hay que limpiar el proyecto para que no use rabbit.

## Requisitos

Go 1.10  <https://golang.org/doc/install>

Dep <https://github.com/golang/dep>

## Configuración inicial

establecer variables de entorno (consultar documentación de la version instalada)

```bash
export GOPATH="$HOME/go"
export GOROOT=/usr/local/go
export PATH="$PATH:$GOPATH/bin:$GOROOT/bin"
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
dep ensure
```

## Build y ejecución

```bash
go install github.com/nmarsollier/imagego
imagego
```

## Redis

Seguir los pasos de instalación desde la pagina oficial

<https://redis.io/download>

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## RabbitMQ

Seguir los pasos de instalación en la pagina oficial

<https://www.rabbitmq.com/>

No se requiere ninguna configuración adicional, solo levantarlo luego de instalarlo.

## Apidoc

Apidoc es una herramienta para proyectos node, para que funcione correctamente hay que instalarla globalmente con

```bash
npm install apidoc -g
```

La documentación necesita ser generada manualmente ejecutando la siguiente linea en la carpeta imagego :

```bash
apidoc -s src -o www
```

Esto nos genera una carpeta public con la documentación, esta carpeta debe estar presente desde donde se ejecute imagego, imagego busca ./www para localizarlo, aunque se puede configurar desde el archivo de properties.

## Archivo config.json

Este archivo permite configurar imagego, ver ejemplos en config-example.json.
imagego busca el archivo "./config.json". Podemos definir el archivo su ruta completa ejecutando

```bash
imagego [path_to_config.json]
```

Para mas detalles ver el archivo tools/env/env.go