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

Rabbit Message

##### Description

Listens for logout messages from auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | General message structure | Yes | [rabbit.message](#rabbitmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### /v1/image

#### POST
##### Summary

Save image

##### Description

Adds a new image to the server.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| image | body | Image in base64 | Yes | [rest.NewRequest](#restnewrequest) |
| Authorization | header | bearer {token} | Yes | string |
| correlation_id | header | Logging Correlation Id | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Image | [rest.NewImageResponse](#restnewimageresponse) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/image/:imageID

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
| 200 | Image Information | [image.Image](#imageimage) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/image/:imageID/jpeg

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
