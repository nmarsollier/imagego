<a name="top"></a>
# ImageGo Service v0.1.0

Microservicio de Imágenes

- [Imagen](#imagen)
	- [Crear Imagen](#crear-imagen)
	- [Obtener Imagen](#obtener-imagen)
	- [Obtener Imagen Jpeg](#obtener-imagen-jpeg)
	
- [RabbitMQ_GET](#rabbitmq_get)
	- [Logout de Usuarios](#logout-de-usuarios)
	


# <a name='imagen'></a> Imagen

## <a name='crear-imagen'></a> Crear Imagen
[Back to top](#top)

<p>Agrega una nueva imagen al servidor.</p>

	POST /image




### Success Response

Respuesta

```
HTTP/1.1 200 OK
{
  "id": "{Id de imagen}"
}
```
401 Unauthorized

```
HTTP/1.1 401 Unauthorized
```
400 Bad Request

```
HTTP/1.1 400 Bad Request
{
   "messages" : [
     {
       "path" : "{Nombre de la propiedad}",
       "message" : "{Motivo del error}"
     },
     ...
  ]
}
```
500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```


## <a name='obtener-imagen'></a> Obtener Imagen
[Back to top](#top)

<p>Obtiene una imagen del servidor en formato base64</p>

	GET /image/:id




### Success Response

Respuesta

```
{
  "id": "{Id de imagen}",
  "image" : "{Imagen en formato Base 64}"
}
```
401 Unauthorized

```
HTTP/1.1 401 Unauthorized
```
400 Bad Request

```
HTTP/1.1 400 Bad Request
{
   "messages" : [
     {
       "path" : "{Nombre de la propiedad}",
       "message" : "{Motivo del error}"
     },
     ...
  ]
}
```
500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```


## <a name='obtener-imagen-jpeg'></a> Obtener Imagen Jpeg
[Back to top](#top)

<p>Obtiene una imagen del servidor en formato jpeg.</p>

	GET /image/:id/jpeg




### Success Response

Respuesta

```
Imagen en formato jpeg
```
401 Unauthorized

```
HTTP/1.1 401 Unauthorized
```
400 Bad Request

```
HTTP/1.1 400 Bad Request
{
   "messages" : [
     {
       "path" : "{Nombre de la propiedad}",
       "message" : "{Motivo del error}"
     },
     ...
  ]
}
```
500 Server Error

```
HTTP/1.1 500 Internal Server Error
{
   "error" : "Not Found"
}
```


# <a name='rabbitmq_get'></a> RabbitMQ_GET

## <a name='logout-de-usuarios'></a> Logout de Usuarios
[Back to top](#top)

<p>Escucha de mensajes logout desde auth.</p>

	FANOUT auth/logout




### Success Response

Mensaje

```
{
   "type": "logout",
   "message": "{tokenId}"
}
```


