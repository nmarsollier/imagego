define({ "api": [
  {
    "type": "post",
    "url": "/v1/image",
    "title": "Crear Imagen",
    "name": "Crear_Imagen",
    "group": "Imagen",
    "description": "<p>Agrega una nueva imagen al servidor.</p>",
    "examples": [
      {
        "title": "Body",
        "content": "{\n  \"image\" : \"{Imagen en formato Base 64}\"\n}",
        "type": "json"
      },
      {
        "title": "Header Autorización",
        "content": "Authorization=bearer {token}",
        "type": "String"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Respuesta",
          "content": "HTTP/1.1 200 OK\n{\n  \"id\": \"{Id de imagen}\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routes/controllers.go",
    "groupTitle": "Imagen",
    "error": {
      "examples": [
        {
          "title": "401 Unauthorized",
          "content": "HTTP/1.1 401 Unauthorized",
          "type": "json"
        },
        {
          "title": "400 Bad Request",
          "content": "HTTP/1.1 400 Bad Request\n{\n   \"messages\" : [\n     {\n       \"path\" : \"{Nombre de la propiedad}\",\n       \"message\" : \"{Motivo del error}\"\n     },\n     ...\n  ]\n}",
          "type": "json"
        },
        {
          "title": "500 Server Error",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n   \"error\" : \"Not Found\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/image/:id",
    "title": "Obtener Imagen",
    "name": "Obtener_Imagen",
    "group": "Imagen",
    "description": "<p>Obtiene una imagen del servidor en formato base64</p>",
    "success": {
      "examples": [
        {
          "title": "Respuesta",
          "content": "{\n  \"id\": \"{Id de imagen}\",\n  \"image\" : \"{Imagen en formato Base 64}\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routes/controllers.go",
    "groupTitle": "Imagen",
    "examples": [
      {
        "title": "Size : Parametro url o header",
        "content": "Size=[160|320|640|800|1024|1200]",
        "type": "String"
      },
      {
        "title": "Header Autorización",
        "content": "Authorization=bearer {token}",
        "type": "String"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "401 Unauthorized",
          "content": "HTTP/1.1 401 Unauthorized",
          "type": "json"
        },
        {
          "title": "400 Bad Request",
          "content": "HTTP/1.1 400 Bad Request\n{\n   \"messages\" : [\n     {\n       \"path\" : \"{Nombre de la propiedad}\",\n       \"message\" : \"{Motivo del error}\"\n     },\n     ...\n  ]\n}",
          "type": "json"
        },
        {
          "title": "500 Server Error",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n   \"error\" : \"Not Found\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "get",
    "url": "/v1/image/:id/jpeg",
    "title": "Obtener Imagen Jpeg",
    "name": "Obtener_Imagen_Jpeg",
    "group": "Imagen",
    "description": "<p>Obtiene una imagen del servidor en formato jpeg.</p>",
    "success": {
      "examples": [
        {
          "title": "Respuesta",
          "content": "Imagen en formato jpeg",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./routes/controllers.go",
    "groupTitle": "Imagen",
    "examples": [
      {
        "title": "Size : Parametro url o header",
        "content": "Size=[160|320|640|800|1024|1200]",
        "type": "String"
      },
      {
        "title": "Header Autorización",
        "content": "Authorization=bearer {token}",
        "type": "String"
      }
    ],
    "error": {
      "examples": [
        {
          "title": "401 Unauthorized",
          "content": "HTTP/1.1 401 Unauthorized",
          "type": "json"
        },
        {
          "title": "400 Bad Request",
          "content": "HTTP/1.1 400 Bad Request\n{\n   \"messages\" : [\n     {\n       \"path\" : \"{Nombre de la propiedad}\",\n       \"message\" : \"{Motivo del error}\"\n     },\n     ...\n  ]\n}",
          "type": "json"
        },
        {
          "title": "500 Server Error",
          "content": "HTTP/1.1 500 Internal Server Error\n{\n   \"error\" : \"Not Found\"\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "fanout",
    "url": "auth/logout",
    "title": "Logout de Usuarios",
    "group": "RabbitMQ_GET",
    "description": "<p>Escucha de mensajes logout desde auth.</p>",
    "success": {
      "examples": [
        {
          "title": "Mensaje",
          "content": "{\n   \"type\": \"logout\",\n   \"message\": \"{tokenId}\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbit.go",
    "groupTitle": "RabbitMQ_GET",
    "name": "FanoutAuthLogout"
  }
] });
