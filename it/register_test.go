package it

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createJson(body any) *bytes.Buffer {
	data, _ := json.Marshal(body)
	return bytes.NewBuffer(data)
}

func TestRunner(t *testing.T) {
	t.Run("RegisterUser", RegisterUser)
}

func RegisterUser(t *testing.T) {
	request := createJson(map[string]string{
		"username": "testtt",
		"password": "siemankoludzie",
	})

	resp, err := http.Post("http://go:8080/v1/auth/register", "application/json", request)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, resp.StatusCode, 201)
}
