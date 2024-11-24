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
        "/health": {
            "get": {
                "tags": [
                    "General"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/verify/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Verify an Email",
                "parameters": [
                    {
                        "description": "Email Details",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.VerifyEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Email Verification Status",
                        "schema": {
                            "$ref": "#/definitions/types.VerifyEmailResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.VerifyEmailRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "xyz@example.com"
                },
                "validate_disposable_email": {
                    "type": "boolean",
                    "example": true
                },
                "validate_mx_record": {
                    "type": "boolean",
                    "example": true
                },
                "validate_regex": {
                    "type": "boolean",
                    "example": true
                },
                "validate_smtp_running": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "types.VerifyEmailResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Regex check failed"
                },
                "valid": {
                    "type": "boolean",
                    "example": false
                }
            }
        }
    },
    "externalDocs": {
        "description": "pkg.go.dev",
        "url": "https://pkg.go.dev/github.com/hsnice16/email-verifier@v0.2.1/core/service"
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
