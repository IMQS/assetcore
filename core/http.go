package core

import (
	"fmt"
	"net/http"

	"github.com/IMQS/nf"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

// Service definition
type Service struct {
	httpPort int
	config   Config
	router   *httprouter.Router
	db       *gorm.DB
}

// NewService creates a new AssetCore service
func NewService() *Service {
	return &Service{
		httpPort: 2042,
		router:   httprouter.New(),
	}
}

// ListenAndServe runs the service
func (s *Service) ListenAndServe() {
	nf.Handle(s.router, "GET", "/ping", s.ping)
	http.ListenAndServe(fmt.Sprintf(":%v", s.httpPort), s.router)
}

// LoadConfig loads our config from the configuration service, and panics upon failure
func (s *Service) LoadConfig() {
}

func (s *Service) ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	nf.SendPong(w)
}
