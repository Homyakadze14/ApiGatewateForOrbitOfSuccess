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
        "/auth/activate_account": {
            "post": {
                "description": "Activate account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Activate account",
                "operationId": "Activate account",
                "parameters": [
                    {
                        "description": "activate",
                        "name": "activate",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.ActivateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.ActivateAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/change_password": {
            "post": {
                "description": "Change password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Change password",
                "operationId": "Change password",
                "parameters": [
                    {
                        "description": "password",
                        "name": "password",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.ChangePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.ChangePasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "login",
                        "name": "login",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "operationId": "Logout",
                "parameters": [
                    {
                        "description": "logout",
                        "name": "logout",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.LogoutResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh token",
                "operationId": "Refresh token",
                "parameters": [
                    {
                        "description": "refresh",
                        "name": "refresh",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.RefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.RefreshResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "operationId": "Register",
                "parameters": [
                    {
                        "description": "register",
                        "name": "register",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/send_password_link": {
            "post": {
                "description": "Send password link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Send password link",
                "operationId": "Send password link",
                "parameters": [
                    {
                        "description": "email",
                        "name": "email",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.SendPasswordLinkRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authv1.SendPasswordLinkResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "Update user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user info",
                "operationId": "Update user info",
                "parameters": [
                    {
                        "description": "info",
                        "name": "info",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/entities.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userv1.UpdateInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info",
                "operationId": "Get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userv1.GetInfoResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        }
    },
    "definitions": {
        "authv1.ActivateAccountResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "authv1.ChangePasswordResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "authv1.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "authv1.LogoutResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "authv1.RefreshResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "authv1.RegisterResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "authv1.SendPasswordLinkResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "entities.ActivateAccountRequest": {
            "type": "object",
            "required": [
                "link"
            ],
            "properties": {
                "link": {
                    "type": "string"
                }
            }
        },
        "entities.ChangePasswordRequest": {
            "type": "object",
            "required": [
                "link",
                "password"
            ],
            "properties": {
                "link": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                }
            }
        },
        "entities.LoginRequest": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entities.LogoutRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "entities.RefreshRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "entities.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "entities.SendPasswordLinkRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "entities.UserUpdateRequest": {
            "type": "object",
            "required": [
                "firstname",
                "gender",
                "icon_url",
                "lastname",
                "middlename",
                "phone",
                "user_id"
            ],
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "icon_url": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "middlename": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "userv1.GetInfoResponse": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "icon_url": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "middlename": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "userv1.UpdateInfoResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Gatewate",
	Description:      "API Gatewate for orbit of success services",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
