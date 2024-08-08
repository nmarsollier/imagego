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
        "/v1/image": {
            "post": {
                "description": "Agrega una nueva imagen al servidor.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Imagen"
                ],
                "summary": "Guardar imagen",
                "parameters": [
                    {
                        "description": "Imagen en base64",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routes.NewRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Imagen",
                        "schema": {
                            "$ref": "#/definitions/routes.NewImageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrValidation"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    }
                }
            }
        },
        "/v1/image/:imageID": {
            "get": {
                "responses": {
                    "200": {
                        "description": "Informacion de la Imagen",
                        "schema": {
                            "$ref": "#/definitions/image.Image"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrValidation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    }
                }
            }
        },
        "/v1/image/:imageID/jpeg": {
            "get": {
                "responses": {
                    "200": {
                        "description": "Imagen",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrValidation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/custerror.ErrCustom"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "custerror.ErrCustom": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "custerror.ErrField": {
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
        "custerror.ErrValidation": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/custerror.ErrField"
                    }
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
        "routes.NewImageResponse": {
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
        "routes.NewRequest": {
            "type": "object",
            "required": [
                "image"
            ],
            "properties": {
                "image": {
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
