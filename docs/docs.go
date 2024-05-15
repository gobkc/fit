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
        "/p/health": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public apis"
                ],
                "summary": "k8s health check",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/p/list-cate": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public apis"
                ],
                "summary": "List notebook categories",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/rest.ListCateResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/p/new-cate": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public apis"
                ],
                "summary": "Add a new cate",
                "parameters": [
                    {
                        "description": "request parameters, must be fill in",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.NewCateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/rest.NewCateResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/p/new-note": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public apis"
                ],
                "summary": "Add a new note",
                "parameters": [
                    {
                        "description": "request parameters, must be fill in",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.NewNoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/rest.NewNoteResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/p/version": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public apis"
                ],
                "summary": "Get current api version",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/rest.VersionResponse"
                        }
                    }
                }
            }
        },
        "/p/{cate}/list-note": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "public apis"
                ],
                "summary": "List notebooks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "cate",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "$ref": "#/definitions/rest.ListNoteResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "driver.NoteInstance": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_time": {
                    "type": "string"
                }
            }
        },
        "rest.ListCateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "error": {
                    "type": "integer"
                },
                "more": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "rest.ListNoteResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/driver.NoteInstance"
                    }
                },
                "error": {
                    "type": "integer"
                },
                "more": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "rest.NewCateRequest": {
            "type": "object",
            "properties": {
                "cate": {
                    "type": "string"
                }
            }
        },
        "rest.NewCateResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "integer"
                },
                "more": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "parameters": {
                    "$ref": "#/definitions/rest.NewCateRequest"
                }
            }
        },
        "rest.NewNoteRequest": {
            "type": "object",
            "properties": {
                "cate": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "rest.NewNoteResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "integer"
                },
                "more": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "parameters": {
                    "$ref": "#/definitions/rest.NewNoteRequest"
                }
            }
        },
        "rest.VersionResponse": {
            "type": "object",
            "properties": {
                "apiVersion": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
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
