# ImageGo
Microservicio de Imagenes.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

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
| image | body | Imagen en base64 | Yes | [routes.NewRequest](#routesnewrequest) |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Imagen | [routes.NewImageResponse](#routesnewimageresponse) |
| 400 | Bad Request | [apperr.ErrValidation](#apperrerrvalidation) |
| 401 | Unauthorized | [apperr.ErrCustom](#apperrerrcustom) |
| 404 | Not Found | [apperr.ErrCustom](#apperrerrcustom) |
| 500 | Internal Server Error | [apperr.ErrCustom](#apperrerrcustom) |

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
| 400 | Bad Request | [apperr.ErrValidation](#apperrerrvalidation) |
| 404 | Not Found | [apperr.ErrCustom](#apperrerrcustom) |
| 500 | Internal Server Error | [apperr.ErrCustom](#apperrerrcustom) |

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
| 400 | Bad Request | [apperr.ErrValidation](#apperrerrvalidation) |
| 404 | Not Found | [apperr.ErrCustom](#apperrerrcustom) |
| 500 | Internal Server Error | [apperr.ErrCustom](#apperrerrcustom) |

---
### Models

#### apperr.ErrCustom

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### apperr.ErrField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### apperr.ErrValidation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [apperr.ErrField](#apperrerrfield) ] |  | No |

#### image.Image

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | Yes |
| image | string |  | Yes |

#### routes.NewImageResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | Yes |

#### routes.NewRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| image | string |  | Yes |
