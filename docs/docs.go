// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Daniel Petersen"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/recipes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get list of recipes",
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Recipe"
                            }
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new recipe",
                "parameters": [
                    {
                        "description": "Recipe to add",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    }
                }
            }
        },
        "/recipes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get one recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the recipe",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    },
                    "404": {
                        "description": "Invalid recipe ID",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update an existing recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the recipe",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated recipe",
                        "name": "recipe",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Recipe"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "404": {
                        "description": "Invalid recipe ID",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "delete an existing recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the recipe",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "404": {
                        "description": "Invalid recipe ID",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/main.Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Message": {
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
        "main.Recipe": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "instructions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "publishedAt": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Recipes API",
	Description:      "This is a sample recipes api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}