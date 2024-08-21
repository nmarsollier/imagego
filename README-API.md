# ImageGo
Microservicio de Imagenes.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [rabbit.message](#rabbitmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### /v1/image

#### POST
##### Summary

Guardar imagen

##### Description

Agrega una nueva imagen al servidor.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| image | body | Imagen en base64 | Yes | [rest.NewRequest](#restnewrequest) |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Imagen | [rest.NewImageResponse](#restnewimageresponse) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/image/:imageID

#### GET
##### Summary

Obtener imagen

##### Description

Obtiene una imagen del servidor en formato base64

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Size | path | [160\|320\|640\|800\|1024\|1200] | Yes | string |
| imageID | path | ID de la imagen | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Informacion de la Imagen | [image.Image](#imageimage) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/image/:imageID/jpeg

#### GET
##### Summary

Obtener jpeg

##### Description

Obtiene una imagen del servidor en formato jpeg.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Size | path | [160\|320\|640\|800\|1024\|1200] | Yes | string |
| imageID | path | ID de la imagen | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Imagen | file |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

---
### Models

#### errs.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errs.errField](#errserrfield) ] |  | No |

#### errs.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### image.Image

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | Yes |
| image | string |  | Yes |

#### rabbit.message

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| correlation_id | string | *Example:* `"123123"` | No |
| message | string | *Example:* `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"` | No |

#### rest.NewImageResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | Yes |

#### rest.NewRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| image | string |  | Yes |

#### server.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
