package core

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/IMQS/nf"
	"gotest.tools/assert"
)

const httpPort = 2000

func startServer() *Service {
	s := NewService()
	s.config.HttpPort = httpPort
	s.config.DB = nf.TestDBConfig("assetcore")
	s.Initialize()
	go s.ListenAndServe()
	time.Sleep(5 * time.Millisecond) // should actually try ping until it works
	return s
}

func baseURL() string {
	return fmt.Sprintf("http://localhost:%v", httpPort)
}

func TestPing(t *testing.T) {
	startServer()
	resp, err := http.DefaultClient.Get(baseURL() + "/ping")
	if err != nil {
		t.Fatalf("Failed to ping: %v", err)
	}
	defer resp.Body.Close()
	assert.Assert(t, resp.StatusCode == 200)
}

func TestAddType(t *testing.T) {
	startServer()
}
