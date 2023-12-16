package util

import (
	"Go-ToDo/postgres"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
)

func SetupTestApp(method string, target string, handler fiber.Handler) *fiber.App {
	app := fiber.New()
	postgres.Connect()

	switch method {
	case fiber.MethodGet:
		app.Get(target, handler)
	case fiber.MethodPost:
		app.Post(target, handler)
	case fiber.MethodPut:
		app.Put(target, handler)
	case fiber.MethodDelete:
		app.Delete(target, handler)
	}

	return app
}

func CreateHttpTestRequestWithBody(app *fiber.App, method string, target string, body interface{}) (*http.Response, error) {
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(method, target, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	return app.Test(req)
}

func CreateHttpTestRequestWithoutBody(app *fiber.App, method string, target string) (*http.Response, error) {
	req := httptest.NewRequest(method, target, nil)
	req.Header.Set("Content-Type", "application/json")
	return app.Test(req)
}
