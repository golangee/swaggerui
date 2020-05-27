package swaggerui

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	http.HandleFunc("/swaggerui/", Handler("/swaggerui/", testSpec))
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}

const testSpec = `
openapi: 3.0.1
info:
  title: test
  version: ''
servers:
  - url: 'http://localhost:8080'
paths:
  /api/v1/setup/status:
    get:
      tags:
        - setup
      summary: Status returns the current setup status
      description: Status returns the current setup status. This is usually only relevant in the installation phase.
      responses:
        '200':
          description: Status represents the current setup status.
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Status'
components:
  schemas:
    Status:
      type: object
      properties:
        Id:
          type: integer
          format: int32
          description: |
            status id
        Message:
          type: string
          description: |
            a textual representation as a developer notice
      description: |
        Status represents the current setup status.
`
