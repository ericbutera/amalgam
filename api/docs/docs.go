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
                        "type": "string",
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
                "parameters": [
                    {
                        "description": "feed data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.CreateFeedRequest"
                        }
                    }
                ],
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
                        "type": "string",
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
            "put": {
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
                        "type": "string",
                        "description": "Feed ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "feed data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.UpdateFeedRequest"
                        }
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
                        "type": "string",
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
            "required": [
                "feedId",
                "id",
                "url"
            ],
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
                "feedId": {
                    "type": "string",
                    "example": "1"
                },
                "guid": {
                    "type": "string",
                    "example": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
                },
                "id": {
                    "description": "ID        uint           ` + "`" + `gorm:\"primarykey\" json:\"id\" binding:\"required\" example:\"1\"` + "`" + `",
                    "type": "string",
                    "example": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
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
            "required": [
                "id",
                "url"
            ],
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "description": "ID        uint           ` + "`" + `gorm:\"primarykey\" json:\"id\" binding:\"required\" example:\"1\"` + "`" + `",
                    "type": "string",
                    "example": "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
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
        "server.CreateFeed": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "My Feed"
                },
                "url": {
                    "type": "string",
                    "example": "https://example.com/feed.xml"
                }
            }
        },
        "server.CreateFeedRequest": {
            "type": "object",
            "properties": {
                "feed": {
                    "$ref": "#/definitions/server.CreateFeed"
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
                "id": {
                    "type": "string"
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
        },
        "server.UpdateFeed": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "My Feed"
                },
                "url": {
                    "type": "string",
                    "example": "https://example.com/feed.xml"
                }
            }
        },
        "server.UpdateFeedRequest": {
            "type": "object",
            "properties": {
                "feed": {
                    "$ref": "#/definitions/server.UpdateFeed"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
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
