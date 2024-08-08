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
| 400 | Bad Request | [custerror.ErrValidation](#custerrorerrvalidation) |
| 401 | Unauthorized | [custerror.ErrCustom](#custerrorerrcustom) |
| 404 | Not Found | [custerror.ErrCustom](#custerrorerrcustom) |
| 500 | Internal Server Error | [custerror.ErrCustom](#custerrorerrcustom) |

---
### /v1/image/:imageID

#### GET
##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Informacion de la Imagen | [image.Image](#imageimage) |
| 400 | Bad Request | [custerror.ErrValidation](#custerrorerrvalidation) |
| 404 | Not Found | [custerror.ErrCustom](#custerrorerrcustom) |
| 500 | Internal Server Error | [custerror.ErrCustom](#custerrorerrcustom) |

### /v1/image/:imageID/jpeg

#### GET
##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Imagen | file |
| 400 | Bad Request | [custerror.ErrValidation](#custerrorerrvalidation) |
| 404 | Not Found | [custerror.ErrCustom](#custerrorerrcustom) |
| 500 | Internal Server Error | [custerror.ErrCustom](#custerrorerrcustom) |

---
### Models

#### custerror.ErrCustom

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### custerror.ErrField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### custerror.ErrValidation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [custerror.ErrField](#custerrorerrfield) ] |  | No |

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
