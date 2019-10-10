package core

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

func startServer() *Service {
	s := NewService()
	go s.ListenAndServe()
	return s
}

func baseURL(s *Service) string {
	return fmt.Sprintf("http://localhost:%v", s.httpPort)
}

func TestPing(t *testing.T) {
	s := startServer()
	resp, err := http.DefaultClient.Get(baseURL(s) + "/ping")
	assert.Assert(t, err == nil)
	defer resp.Body.Close()
	assert.Assert(t, resp.StatusCode == 200)
}
