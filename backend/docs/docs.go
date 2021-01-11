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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/cleanernames": {
            "get": {
                "description": "list cleanername entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List cleanername entities",
                "operationId": "list-cleanername",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.CleanerName"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/cleanernames/{id}": {
            "get": {
                "description": "get cleanername by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a cleanername entity by ID",
                "operationId": "get-cleanername",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CleanerName ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.CleanerName"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/cleaningrooms": {
            "get": {
                "description": "list cleaningroom entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List cleaningroom entities",
                "operationId": "list-cleaningroom",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.CleaningRoom"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create cleaningroom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create cleaningroom",
                "operationId": "create-cleaningroom",
                "parameters": [
                    {
                        "description": "CleaningRoom entity",
                        "name": "cleaningroom",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.CleaningRoom"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.CleaningRoom"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/cleaningrooms/{id}": {
            "get": {
                "description": "get cleaningroom by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a cleaningroom entity by ID",
                "operationId": "get-cleaningroom",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CleaningRoom ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.CleaningRoom"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/deposits": {
            "get": {
                "description": "list deposit entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List deposit entities",
                "operationId": "list-deposit",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Deposit"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create deposit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create deposit",
                "operationId": "create-deposit",
                "parameters": [
                    {
                        "description": "Deposit entity",
                        "name": "deposit",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.Deposit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Deposit"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employee/{id}": {
            "delete": {
                "description": "get employee by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a employee entity by ID",
                "operationId": "delete-employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employees": {
            "get": {
                "description": "list employee entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List employee entities",
                "operationId": "list-employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Employee"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create employee",
                "operationId": "create-employee",
                "parameters": [
                    {
                        "description": "Employee entity",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "get": {
                "description": "get employee by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a employee entity by ID",
                "operationId": "get-employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/lengthtimes": {
            "get": {
                "description": "list lengthtime entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List lengthtime entities",
                "operationId": "list-lengthtime",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.LengthTime"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/lengthtimes/{id}": {
            "get": {
                "description": "get lengthtime by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a lengthtime entity by ID",
                "operationId": "get-lengthtime",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "LengthTime ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.LengthTime"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/statusd/{id}": {
            "delete": {
                "description": "get statusd by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a statusd entity by ID",
                "operationId": "delete-statusd",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Statusd ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/statusds": {
            "get": {
                "description": "list statusd entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List statusd entities",
                "operationId": "list-statusd",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Statusd"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create statusd",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create statusd",
                "operationId": "create-statusd",
                "parameters": [
                    {
                        "description": "Statusd entity",
                        "name": "statusd",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Statusd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Statusd"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/statusds/{id}": {
            "get": {
                "description": "get statusd by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a statusd entity by ID",
                "operationId": "get-statusd",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Statusd ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Statusd"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Deposit": {
            "type": "object",
            "properties": {
                "added": {
                    "type": "string"
                },
                "employee": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "statusd": {
                    "type": "integer"
                }
            }
        },
        "ent.CleanerName": {
            "type": "object",
            "properties": {
                "cleanername": {
                    "description": "Cleanername holds the value of the \"cleanername\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the CleanerNameQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.CleanerNameEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.CleanerNameEdges": {
            "type": "object",
            "properties": {
                "cleaningrooms": {
                    "description": "Cleaningrooms holds the value of the cleaningrooms edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.CleaningRoom"
                    }
                }
            }
        },
        "ent.CleaningRoom": {
            "type": "object",
            "properties": {
                "dateandstarttime": {
                    "description": "Dateandstarttime holds the value of the \"dateandstarttime\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the CleaningRoomQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.CleaningRoomEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "note": {
                    "description": "Note holds the value of the \"note\" field.",
                    "type": "string"
                }
            }
        },
        "ent.CleaningRoomEdges": {
            "type": "object",
            "properties": {
                "cleanerName": {
                    "description": "CleanerName holds the value of the CleanerName edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.CleanerName"
                },
                "lengthTime": {
                    "description": "LengthTime holds the value of the LengthTime edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.LengthTime"
                }
            }
        },
        "ent.Deposit": {
            "type": "object",
            "properties": {
                "addedtime": {
                    "description": "Addedtime holds the value of the \"addedtime\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the DepositQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.DepositEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "info": {
                    "description": "Info holds the value of the \"info\" field.",
                    "type": "string"
                }
            }
        },
        "ent.DepositEdges": {
            "type": "object",
            "properties": {
                "employee": {
                    "description": "Employee holds the value of the Employee edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Employee"
                },
                "statusd": {
                    "description": "Statusd holds the value of the Statusd edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Statusd"
                }
            }
        },
        "ent.Employee": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the EmployeeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.EmployeeEdges"
                },
                "employeeemail": {
                    "description": "Employeeemail holds the value of the \"employeeemail\" field.",
                    "type": "string"
                },
                "employeename": {
                    "description": "Employeename holds the value of the \"employeename\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "password": {
                    "description": "Password holds the value of the \"password\" field.",
                    "type": "string"
                }
            }
        },
        "ent.EmployeeEdges": {
            "type": "object",
            "properties": {
                "employees": {
                    "description": "Employees holds the value of the employees edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Deposit"
                    }
                }
            }
        },
        "ent.LengthTime": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the LengthTimeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.LengthTimeEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "lengthtime": {
                    "description": "Lengthtime holds the value of the \"lengthtime\" field.",
                    "type": "string"
                }
            }
        },
        "ent.LengthTimeEdges": {
            "type": "object",
            "properties": {
                "cleaningrooms": {
                    "description": "Cleaningrooms holds the value of the cleaningrooms edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.CleaningRoom"
                    }
                }
            }
        },
        "ent.Statusd": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the StatusdQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.StatusdEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "statusdname": {
                    "description": "Statusdname holds the value of the \"statusdname\" field.",
                    "type": "string"
                }
            }
        },
        "ent.StatusdEdges": {
            "type": "object",
            "properties": {
                "statusds": {
                    "description": "Statusds holds the value of the statusds edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Deposit"
                    }
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "authorizationUrl": "",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "authorizationUrl": "",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "SUT SA Example API",
	Description: "This is a sample server for SUT SE 2563",
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
