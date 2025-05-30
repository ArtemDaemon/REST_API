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
        "/api/add_item": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Принимает JSON-структуру и добавляет её в базу",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Добавить элемент",
                "parameters": [
                    {
                        "description": "Элемент данных",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Item"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверный JSON",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при добавлении записи",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/count": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает общее число записей в базе данных",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Получить количество элементов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "integer"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка при запросе количества",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/get_item": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает первую запись с указанной датой",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Получить элемент по дате",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Год (например, 2022)",
                        "name": "year",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Item"
                        }
                    },
                    "400": {
                        "description": "Параметр 'year' обязателен",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Элемент не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при запросе",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/last_created_at": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Возвращает максимальное значение поля created_at",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Получить дату последнего элемента",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка при получении даты",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.Item": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "string"
                },
                "country_iso3_code": {
                    "type": "string"
                },
                "country_value": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "decimal": {
                    "type": "integer"
                },
                "indicator_id": {
                    "type": "string"
                },
                "indicator_value": {
                    "type": "string"
                },
                "obs_status": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                },
                "value": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Aithorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Пример REST API",
	Description:      "Это REST API с авторизацией через Bearer-токен",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
