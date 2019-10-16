package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/IMQS/log"
	"github.com/IMQS/nf/nftest"
	"github.com/IMQS/serviceauth"
	"github.com/IMQS/serviceauth/permissions"
)

const httpPort = 2000

func startServer(t *testing.T) *Service {
	s := NewService()
	s.config.HttpPort = httpPort
	s.config.DB = nftest.MakeDBConfig("assetcore")
	s.log = log.NewTesting(t)
	serviceauth.ActivateMockToken(1, "user@example.com", []int{permissions.PermEnabled, permissions.PermAdmin})
	s.Initialize()
	go s.ListenAndServe()
	nftest.PingUntil200(t, time.Second, baseURL()+"/ping")
	return s
}

func baseURL() string {
	return fmt.Sprintf("http://localhost:%v", httpPort)
}

func TestAssetTypeCRUD(t *testing.T) {
	startServer(t)
}
