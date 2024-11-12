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
        "/orders": {
            "post": {
                "description": "Create a new order in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "Order Request",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "endpoint.CreateOrderRequest": {
            "type": "object",
            "properties": {
                "order": {
                    "$ref": "#/definitions/models.Order"
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "discount": {
                    "type": "number"
                },
                "employee_id": {
                    "type": "integer"
                },
                "order_date": {
                    "type": "string"
                },
                "order_details": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderDetail"
                    }
                },
                "order_id": {
                    "description": "OrderID now uses uint",
                    "type": "integer"
                },
                "order_status_id": {
                    "type": "integer"
                },
                "order_time": {
                    "type": "string"
                },
                "order_type": {
                    "description": "Example values: \"Dine-In\", \"Takeout\", \"Delivery\"",
                    "type": "string"
                },
                "payment_method": {
                    "type": "string"
                },
                "table_id": {
                    "type": "integer"
                },
                "tax_id": {
                    "type": "integer"
                }
            }
        },
        "models.OrderDetail": {
            "type": "object",
            "properties": {
                "item_id": {
                    "type": "integer"
                },
                "order_detail_id": {
                    "type": "integer"
                },
                "order_id": {
                    "description": "OrderID now uses uint",
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "special_instructions": {
                    "type": "string"
                },
                "total_price": {
                    "type": "number"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Actual data payload (if any)"
                },
                "error": {
                    "description": "Error details (if any)"
                },
                "message": {
                    "description": "Descriptive message",
                    "type": "string"
                },
                "status": {
                    "description": "\"success\" or \"error\"",
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
