// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Nestor Marsollier",
            "email": "nmarsollier@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rabbit/logout": {
            "get": {
                "description": "Listens for logout messages from auth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rabbit"
                ],
                "summary": "Rabbit Message",
                "parameters": [
                    {
                        "description": "General message structure",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rabbit.message"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/image": {
            "post": {
                "description": "Adds a new image to the server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Save image",
                "parameters": [
                    {
                        "description": "Image in base64",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.NewRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Logging Correlation Id",
                        "name": "correlation_id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image",
                        "schema": {
                            "$ref": "#/definitions/rest.NewImageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/image/:imageID": {
            "get": {
                "description": "Gets an image from the server in base64 format",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Get image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Logging Correlation Id",
                        "name": "correlation_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "[160|320|640|800|1024|1200]",
                        "name": "Size",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image ID",
                        "name": "imageID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image Information",
                        "schema": {
                            "$ref": "#/definitions/image.Image"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/image/:imageID/jpeg": {
            "get": {
                "description": "Gets an image from the server in jpeg format.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/jpeg"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "Get jpeg",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Logging Correlation Id",
                        "name": "correlation_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "[160|320|640|800|1024|1200]",
                        "name": "Size",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image ID",
                        "name": "imageID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.ValidationErr"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errs.ValidationErr": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/errs.errField"
                    }
                }
            }
        },
        "errs.errField": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "image.Image": {
            "type": "object",
            "required": [
                "id",
                "image"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                }
            }
        },
        "rabbit.message": {
            "type": "object",
            "properties": {
                "correlation_id": {
                    "type": "string",
                    "example": "123123"
                },
                "message": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"
                }
            }
        },
        "rest.NewImageResponse": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "rest.NewRequest": {
            "type": "object",
            "required": [
                "image"
            ],
            "properties": {
                "image": {
                    "type": "string"
                }
            }
        },
        "server.ErrorData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3001",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "ImageGo",
	Description:      "Microservicio de Imagenes.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
