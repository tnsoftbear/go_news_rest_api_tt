// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Igor",
            "url": "http://github.com/tnsoftbear",
            "email": "myg0t@inbox.lv"
        },
        "license": {
            "name": "MIT",
            "url": "https://rem.mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/:NewsId": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete News record by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Delete News",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "News record ID",
                        "name": "NewsId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.DeleteNewsByIdResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a News record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Add News",
                "parameters": [
                    {
                        "description": "News record data",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.PostNewsAddRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.PostNewsAddResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/add-category/:NewsId/:CatId": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Assign category to some news record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Assign Category",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "News record ID",
                        "name": "NewsId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "CatId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.PostNewsAddCategoryResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/edit/:Id": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Modify the existing News record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Edit News",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "News record ID",
                        "name": "Id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "News record data",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.PostNewsEditByIdRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.PostNewsEditByIdResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve news list at some page",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "News List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Show page number (def: 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Records per page (def: 10)",
                        "name": "per-page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetNewsListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticate user and provide access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authentication",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.AccessToken"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Check service health by ping http request",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "infra"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.GetPingResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AccessToken": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "controller.DeleteNewsByIdResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.GetNewsListResponse": {
            "type": "object",
            "properties": {
                "News": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.News"
                    }
                },
                "Success": {
                    "type": "boolean"
                }
            }
        },
        "controller.GetPingResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.PostNewsAddCategoryResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "controller.PostNewsAddRequest": {
            "type": "object",
            "properties": {
                "Categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "Content": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "controller.PostNewsAddResponse": {
            "type": "object",
            "properties": {
                "Categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "Message": {
                    "type": "string"
                },
                "News": {
                    "$ref": "#/definitions/model.News"
                },
                "Success": {
                    "type": "boolean"
                }
            }
        },
        "controller.PostNewsEditByIdRequest": {
            "type": "object",
            "properties": {
                "Categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "Content": {
                    "type": "string"
                },
                "Id": {
                    "type": "integer"
                },
                "Title": {
                    "type": "string"
                }
            }
        },
        "controller.PostNewsEditByIdResponse": {
            "type": "object",
            "properties": {
                "Categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "Message": {
                    "type": "string"
                },
                "News": {
                    "$ref": "#/definitions/model.News"
                },
                "Success": {
                    "type": "boolean"
                }
            }
        },
        "model.News": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "content": {
                    "type": "string"
                },
                "id": {
                    "description": "primary key",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "REST API details",
        "url": "https://github.com/tnsoftbear/go_news_rest_api_tt"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:4000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "News service",
	Description:      "This is a testing task for implementing JSON REST API with fiber and reform.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}