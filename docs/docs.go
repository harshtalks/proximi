// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "API Support",
            "email": "harshpareek91@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/businesses": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "# Businesses API\n\nUse this API to access businesses nearby you, based on latitude and longitude you provide.\n\n## Params\n\n### 1. Lat\n\n- Latitude of the location of the user (string value)\n- Required\n\n### 2. Long\n\n- Longitude of the location of the user (string value)\n- Required\n\n### 3. range\n\n- Range of the result\n- You can specify the range in kilometers\n- default is the area of radius .5kms to 2kms\n- optional\n\n### 4. page\n\n- This API is paginated.\n- default page number is 1\n- optional\n\n### 5. perPage\n\n- per page results\n- default is 10\n- max is 100\n- minimum is 10\n- optional\n",
                "tags": [
                    "Businesses"
                ],
                "summary": "get businesses near you",
                "parameters": [
                    {
                        "type": "string",
                        "default": "23983.2",
                        "description": "string value of latitude",
                        "name": "lat",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "-29829.09",
                        "description": "string value of longitude",
                        "name": "long",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "specify range of results",
                        "name": "range",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 10,
                        "type": "integer",
                        "description": "number of results per page",
                        "name": "perPage",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DocBusinessModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "412": {
                        "description": "Precondition Failed",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    }
                }
            }
        },
        "/api/businesses/:businessId": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "# Business\n\nUse this API to get indiviudal Business details\n\n## Params\n",
                "tags": [
                    "Businesses"
                ],
                "summary": "get details of a business",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the business",
                        "name": "businessId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "latitude of your current location",
                        "name": "lat",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "longitude of your current location",
                        "name": "long",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "\"driving\"",
                            "\"walking\"",
                            "\"public transit\""
                        ],
                        "type": "string",
                        "description": "your mode of travel",
                        "name": "travelMode",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "Signin with username and password, you will receive a token which u will have to provide in the header of subsequent requests at /api gateway",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Basic token based Auth",
                "parameters": [
                    {
                        "description": "username and password input",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SigninInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.SigninSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Signup with username and password, you will receive a token which u will have to provide in the header of subsequent requests at /api gateway",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "For new User",
                "parameters": [
                    {
                        "description": "username and password input",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.SigninInput"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.DocBusinessModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "attributes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "business_id": {
                    "type": "string"
                },
                "categories": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "geoHash": {
                    "type": "string"
                },
                "hours": {
                    "type": "object",
                    "additionalProperties": true
                },
                "is_open": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "string"
                },
                "review_count": {
                    "type": "integer"
                },
                "stars": {
                    "type": "number"
                },
                "state": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "handlers.ErrorModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "internal server error"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                },
                "status_code": {
                    "type": "integer",
                    "example": 500
                }
            }
        },
        "handlers.SigninInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "example": "harshtalks"
                }
            }
        },
        "handlers.SigninSuccessResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handlers.SuccessData"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "handlers.SuccessData": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
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
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Proximi",
	Description:      "# Proximi\n\nproximi is an api service that delivers nearby services to the users. Using Proximi, one can find out nearby businesses with specified range.\n\n## Tech Stack\n\n- Go\n- REST APIs\n- Postgres\n- ORM\n- GeoHashing\n- JWT based custom Auth\n\n## Features\n\n- Rate Limiting\n- Authentication\n- Pagination\n- Geocoding\n- Distance from the business.\n\n## Procedure\n\n- First of all, signin/signup from /auth endpoints to get the auth token (Authenticate Yourself)\n- once received the token, copy the token and click on Authorize button to login urself.\n- the format is `Bearer <your-token>`\n- once upon verification of the token, you will be able to access protected routes such as /api endpoints\n\n## Important\n\nwe have applied a rate limiter to make our service always available, and keep our downtime as zero.\ncheck header to see the limit.\n\n### Checkout our better API DOC\n\n[https://bump.sh/harshtalks/doc/proximi/](https://bump.sh/harshtalks/doc/proximi/)\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
