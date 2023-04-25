// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product/createproduct": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "how to Create product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "create product example",
                "parameters": [
                    {
                        "description": "Input Data product",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/reponse.WebResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/product.Product"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/product/deleteproduct/{id}": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "how to Delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "delete product example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/reponse.WebResponseSuccess"
                        }
                    }
                }
            }
        },
        "/product/getall": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "how to get all Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get All Product example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/reponse.WebResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/product.Product"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/product/getone/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "how to get product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Get product example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/reponse.WebResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/product.Product"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/product/updateproduct/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "how to update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "update product example",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path"
                    },
                    {
                        "description": "Input Data product",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.ProductCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/reponse.WebResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/product.Product"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "do login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login example",
                "parameters": [
                    {
                        "description": "Login User",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserAuthenticate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/reponse.WebResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/token.Tokens"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Input tidak valid",
                        "schema": {
                            "$ref": "#/definitions/reponse.WebResponseFailed"
                        }
                    },
                    "500": {
                        "description": "Error Server Side",
                        "schema": {
                            "$ref": "#/definitions/reponse.WebResponseFailed"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "do register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register example",
                "parameters": [
                    {
                        "description": "Input Data User",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/reponse.WebResponseSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user.UserCreateResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Input tidak valid",
                        "schema": {
                            "$ref": "#/definitions/reponse.WebResponseFailed"
                        }
                    },
                    "500": {
                        "description": "Error Server Side",
                        "schema": {
                            "$ref": "#/definitions/reponse.WebResponseFailed"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "product.ProductCreateRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "reponse.WebResponseFailed": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "reponse.WebResponseSuccess": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "token.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "user.UserAuthenticate": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UserCreateRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UserCreateResponse": {
            "type": "object",
            "properties": {
                "Username": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:9999",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}