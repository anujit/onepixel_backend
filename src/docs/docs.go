// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/championswimmer/onepixel_backend",
        "contact": {
            "name": "Arnav Gupta",
            "email": "dev@championswimmer.in"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register new user",
                "operationId": "register-user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.UserResponse"
                        }
                    },
                    "400": {
                        "description": "The request body is not valid",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "User with this email already exists",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "email and password are required to create user",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/:id": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user info",
                "operationId": "get-user-info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "Update user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user info",
                "operationId": "update-user-info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/users/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "operationId": "login-user",
                "responses": {}
            }
        },
        "/urls": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all urls",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "urls"
                ],
                "summary": "Get all urls",
                "responses": {
                    "200": {
                        "description": "GetAllUsers",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dtos.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Something went wrong"
                },
                "status": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "dtos.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@test.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "API Key": {
            "type": "apiKey",
            "name": "X-API-Key",
            "in": "header"
        },
        "BearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "127.0.0.1:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "onepixel API",
	Description:      "1px.li URL Shortner API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
