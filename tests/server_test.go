package tests

import (
	"net/http/httptest"

	"go_final_project/internal/config"
	"go_final_project/pkg/router"

	_ "github.com/mattn/go-sqlite3"
)

func createTestServer() *httptest.Server {
	config.MustLoad()

	router := router.SetupRouter()

	return httptest.NewServer(router)
}
