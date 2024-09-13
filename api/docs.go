// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "LRT SUMSEL",
            "email": "indraadha24@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/realtime-track": {
            "get": {
                "description": "Get Realtime Track",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Realtime Track"
                ],
                "summary": "Get Realtime Track",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Latitude",
                        "name": "latitude",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Longitude",
                        "name": "longitude",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/http.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "value": {
                                            "$ref": "#/definitions/response.GeocodingResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.Error": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "http.Response": {
            "type": "object",
            "required": [
                "message"
            ],
            "properties": {
                "message": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "response.AddressResponse": {
            "type": "object",
            "properties": {
                "long_name": {
                    "type": "string"
                },
                "short_name": {
                    "type": "string"
                },
                "types": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.GeocodingResponse": {
            "type": "object",
            "properties": {
                "address_components": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.AddressResponse"
                    }
                },
                "formatted_address": {
                    "type": "string"
                },
                "geometry": {
                    "$ref": "#/definitions/response.GeometryResponse"
                },
                "place_id": {
                    "type": "string"
                },
                "types": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "response.GeometryResponse": {
            "type": "object",
            "properties": {
                "location": {
                    "$ref": "#/definitions/response.LocationResponse"
                },
                "location_type": {
                    "type": "string"
                },
                "viewport": {
                    "$ref": "#/definitions/response.ViewportResponse"
                }
            }
        },
        "response.LocationResponse": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "response.ViewportResponse": {
            "type": "object",
            "properties": {
                "northeast": {
                    "$ref": "#/definitions/response.LocationResponse"
                },
                "southwest": {
                    "$ref": "#/definitions/response.LocationResponse"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
