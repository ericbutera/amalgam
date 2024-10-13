// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/articles/{id}": {
            "get": {
                "description": "view article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "view article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.ArticleResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/feeds": {
            "get": {
                "description": "list feeds",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list feeds",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.FeedsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create feed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create feed",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.FeedCreateResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/feeds/{id}": {
            "get": {
                "description": "view feed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "view feed",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 1,
                        "description": "Feed ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.FeedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/server.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "update feed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "update feed",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Feed ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.FeedUpdateResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/feeds/{id}/articles": {
            "get": {
                "description": "list articles for a feed",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list articles for a feed",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Feed ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.FeedArticlesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK"
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
        "models.Article": {
            "type": "object",
            "properties": {
                "authorEmail": {
                    "type": "string",
                    "example": "example@example.com"
                },
                "authorName": {
                    "type": "string",
                    "example": "Eric Butera"
                },
                "content": {
                    "type": "string",
                    "example": "Article content text. May contain HTML."
                },
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "feed": {
                    "$ref": "#/definitions/models.Feed"
                },
                "feedID": {
                    "type": "integer",
                    "example": 1
                },
                "guid": {
                    "type": "string",
                    "example": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "imageUrl": {
                    "type": "string",
                    "example": "https://example.com/image.jpg"
                },
                "preview": {
                    "type": "string",
                    "example": "Article preview text. May contain HTML."
                },
                "title": {
                    "type": "string",
                    "example": "Example Article"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "url": {
                    "type": "string",
                    "example": "https://example.com/"
                }
            }
        },
        "models.Feed": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Example"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "url": {
                    "type": "string",
                    "example": "https://example.com/"
                }
            }
        },
        "server.ArticleResponse": {
            "type": "object",
            "properties": {
                "article": {
                    "$ref": "#/definitions/models.Article"
                }
            }
        },
        "server.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "unable to fetch feeds"
                }
            }
        },
        "server.FeedArticlesResponse": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Article"
                    }
                }
            }
        },
        "server.FeedCreateResponse": {
            "type": "object",
            "properties": {
                "feed": {
                    "$ref": "#/definitions/models.Feed"
                }
            }
        },
        "server.FeedResponse": {
            "type": "object",
            "properties": {
                "feed": {
                    "$ref": "#/definitions/models.Feed"
                }
            }
        },
        "server.FeedUpdateResponse": {
            "type": "object",
            "properties": {
                "feed": {
                    "$ref": "#/definitions/models.Feed"
                }
            }
        },
        "server.FeedsResponse": {
            "type": "object",
            "properties": {
                "feeds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Feed"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Feed API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
