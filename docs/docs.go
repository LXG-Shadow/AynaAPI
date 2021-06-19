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
        "/api/v1/novel/content": {
            "get": {
                "description": "获取小说章节内容",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novel"
                ],
                "summary": "get novel content",
                "parameters": [
                    {
                        "type": "string",
                        "description": "content url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "use cache",
                        "name": "cache",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "https://www.linovelib.com/novel/2342/133318.html",
                        "schema": {
                            "$ref": "#/definitions/app.AppJsonResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/novel/info": {
            "get": {
                "description": "获取小说简介",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novel"
                ],
                "summary": "get novel info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "info page url",
                        "name": "url",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "use cache",
                        "name": "cache",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "https://www.linovelib.com/novel/8.html",
                        "schema": {
                            "$ref": "#/definitions/app.AppJsonResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/novel/providerlist": {
            "get": {
                "description": "获取来源列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novel"
                ],
                "summary": "get provider list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.AppJsonResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/novel/rulelist": {
            "get": {
                "description": "获取来源规则",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novel"
                ],
                "summary": "get provider rules",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.AppJsonResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/novel/search": {
            "get": {
                "description": "搜索小说",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Novel"
                ],
                "summary": "search novel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "use cache",
                        "name": "cache",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "诡秘之主",
                        "schema": {
                            "$ref": "#/definitions/app.AppJsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.AppJsonResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
