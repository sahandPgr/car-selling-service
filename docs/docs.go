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
        "/v1/countries/": {
            "post": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Create a country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries"
                ],
                "summary": "Create a country",
                "parameters": [
                    {
                        "description": "Create a country",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.CreateUpdateCountryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.CountryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/countries/{id}": {
            "get": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Get a country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries"
                ],
                "summary": "Get a country",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Get a country",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.CountryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Update a country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries"
                ],
                "summary": "Update a country",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Update a country",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update a country",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.CreateUpdateCountryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.CountryResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AuthBearer": []
                    }
                ],
                "description": "Delete a country",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Countries"
                ],
                "summary": "Delete a country",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete a country",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/login-by-username": {
            "post": {
                "description": "LoginByUsername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "LoginByUsername",
                "parameters": [
                    {
                        "description": "LoginByUsernameRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.LoginByUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/register-by-username": {
            "post": {
                "description": "RegisterByUsername",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterByUsername",
                "parameters": [
                    {
                        "description": "RegisterUserByUsernameRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.RegisterUserByUsernameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/register-login-by-phone": {
            "post": {
                "description": "RegisterLoginByMobile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "RegisterLoginByMobile",
                "parameters": [
                    {
                        "description": "RegisterLoginByMobileRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.RegisterLoginByMobileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/send-otp": {
            "post": {
                "description": "Send otp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Send otp",
                "parameters": [
                    {
                        "description": "GetOtpRequest",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_dto.GetOtpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    },
                    "409": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_sahandPgr_car-selling-service_api_dto.CountryResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_dto.CreateUpdateCountryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 2
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_dto.GetOtpRequest": {
            "type": "object",
            "required": [
                "mobileNumber"
            ],
            "properties": {
                "mobileNumber": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_dto.LoginByUsernameRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_dto.RegisterLoginByMobileRequest": {
            "type": "object",
            "required": [
                "otp",
                "phoneMobile"
            ],
            "properties": {
                "otp": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
                },
                "phoneMobile": {
                    "type": "string",
                    "maxLength": 11,
                    "minLength": 11
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_dto.RegisterUserByUsernameRequest": {
            "type": "object",
            "required": [
                "firstName:",
                "lastName:",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "minLength": 6
                },
                "firstName:": {
                    "type": "string",
                    "minLength": 3
                },
                "lastName:": {
                    "type": "string",
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_helper.BaseHttpResponse": {
            "type": "object",
            "properties": {
                "error": {},
                "result": {},
                "statusCode:": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                },
                "validationErrors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_sahandPgr_car-selling-service_api_validations.ValidationError"
                    }
                }
            }
        },
        "github_com_sahandPgr_car-selling-service_api_validations.ValidationError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "param": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthBearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
