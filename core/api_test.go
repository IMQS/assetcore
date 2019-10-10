package core

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

const httpPort = 2000

func startServer() *Service {
	s := NewService()
	s.config.HttpPort = httpPort
	s.config.DB.Driver = "postgres"
	s.config.DB.Database = "unittest_assetcore"
	s.config.DB.Username = "unittest_user"
	s.config.DB.Password = "unittest_password"
	go s.ListenAndServe()
	return s
}

func baseURL() string {
	return fmt.Sprintf("http://localhost:%v", httpPort)
}

func TestPing(t *testing.T) {
	startServer()
	resp, err := http.DefaultClient.Get(baseURL() + "/ping")
	assert.Assert(t, err == nil)
	defer resp.Body.Close()
	assert.Assert(t, resp.StatusCode == 200)
}
