package core

import (
	"net/http"
	"strings"

	"github.com/IMQS/nf"
	"github.com/IMQS/serviceauth"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	nf.SendPong(w)
}

func (s *Service) listAssets(w http.ResponseWriter, r *http.Request, p httprouter.Params, auth *serviceauth.Token) {
	query := ""
	args := []string{}

	qtype := r.FormValue("type")
	if qtype != "" {
		query += "AND type = ?"
		args = append(args, qtype)
	}

	if strings.HasPrefix(query, "AND") {
		query = query[:3]
	}

	assets := []asset{}
	s.db.Where(query, args).Find(&assets)
	nf.SendJSON(w, assets)
}

func (s *Service) addAssets(w http.ResponseWriter, r *http.Request, p httprouter.Params, auth *serviceauth.Token) {
	assets := []jsonInputAsset{}
	nf.ReadJSON(r, &assets)
}
