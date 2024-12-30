# ImageGo
Microservicio de Imagenes.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /images/:imageID

#### GET
##### Summary

Get image

##### Description

Gets an image from the server in base64 format

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| correlation_id | header | Logging Correlation Id | Yes | string |
| Size | path | [160\|320\|640\|800\|1024\|1200] | Yes | string |
| imageID | path | Image ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Image Information |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /images/:imageID/jpeg

#### GET
##### Summary

Get jpeg

##### Description

Gets an image from the server in jpeg format.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| correlation_id | header | Logging Correlation Id | Yes | string |
| Size | path | [160\|320\|640\|800\|1024\|1200] | Yes | string |
| imageID | path | Image ID | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Image | file |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /images/create

#### POST
##### Summary

Save image

##### Description

Adds a new image to the server.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| image | body | Image in base64 | Yes | [rest.NewRequest](#restnewrequest) |
| Authorization | header | Bearer {token} | Yes | string |
| correlation_id | header | Logging Correlation Id | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Image | [rest.NewImageResponse](#restnewimageresponse) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

---
### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit logout

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [rbt.InputMessage-string](#rbtinputmessage-string) |

##### Responses

| Code | Description |
| ---- | ----------- |

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

#### rbt.InputMessage-string

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| correlation_id | string | *Example:* `"123123"` | No |
| exchange | string | *Example:* `"Remote Exchange to Reply"` | No |
| message | string |  | No |
| routing_key | string | *Example:* `"Remote RoutingKey to Reply"` | No |

#### rest.NewImageResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | Yes |

#### rest.NewRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| image | string |  | Yes |

#### rst.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
