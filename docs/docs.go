// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/sign-in": {
            "post": {
                "description": "signIn checks the users login and password and returns token if the user exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "signIn",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.signRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.signInResponse"
                        }
                    },
                    "400": {
                        "description": "Incorrect login or password",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Oops something went wrong(( Sorry, please try again later",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/safe/auth/changePassword": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "ChangePassword changed user password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "safe"
                ],
                "summary": "ChangePassword",
                "parameters": [
                    {
                        "description": "change pass struct",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.changePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "incorrectPassword or the id doesn't exist",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Oops something went wrong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/safe/business/test": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "test is used to test your jwt token. It return \"hello\" if your token is suitable",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "safe"
                ],
                "summary": "test",
                "responses": {
                    "200": {
                        "description": "hello",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized!!",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "something went wrong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.changePasswordRequest": {
            "type": "object",
            "properties": {
                "newpassword": {
                    "type": "string"
                },
                "oldpassword": {
                    "type": "string"
                }
            }
        },
        "http.signInResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Id - user id",
                    "type": "integer"
                },
                "token": {
                    "description": "token - generated new token for the user",
                    "type": "string"
                }
            }
        },
        "http.signRequest": {
            "type": "object",
            "properties": {
                "login": {
                    "description": "login - user login (required, not null)",
                    "type": "string"
                },
                "password": {
                    "description": "password - user password (required, not null)",
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "coursestj.herokuapp.com",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Courses - Course management system",
	Description: "This is a simple course management system",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
