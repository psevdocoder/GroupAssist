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
        "/api/queues": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "Create queue",
                "operationId": "create-queue",
                "parameters": [
                    {
                        "description": "Queue data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Queue"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Queue"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/queues/by_subject/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "Get all queues by subject id",
                "operationId": "get-all-queues",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subject ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Queue"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/queues/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "Get queue by ID",
                "operationId": "get-queue-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Queue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Queue"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "Delete queue",
                "operationId": "delete-queue",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Queue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "patch": {
                "description": "Allowed to use any field provided in the input body",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "queues"
                ],
                "summary": "Update queue",
                "operationId": "update-queue",
                "parameters": [
                    {
                        "description": "Queue",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateQueueInput"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Queue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/subjects": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subjects"
                ],
                "summary": "Get all subjects",
                "operationId": "get-all-subjects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Subject"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subjects"
                ],
                "summary": "Create subject",
                "operationId": "create-subject",
                "parameters": [
                    {
                        "description": "Subject data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Subject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Subject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subjects"
                ],
                "summary": "Delete subject",
                "operationId": "delete-subject",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subject ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/api/subjects/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subjects"
                ],
                "summary": "Get subject by ID",
                "operationId": "get-subject-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subject ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Subject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subjects"
                ],
                "summary": "Update subject",
                "operationId": "update-subject",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Subject ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Subject data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateSubjectInput"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Queue": {
            "type": "object",
            "required": [
                "subject_id",
                "title"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_open": {
                    "type": "boolean"
                },
                "subject_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.Subject": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.UpdateQueueInput": {
            "type": "object",
            "properties": {
                "is_open": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "domain.UpdateSubjectInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Group Assistant API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}