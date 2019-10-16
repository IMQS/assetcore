package core

import (
	"net/http"

	"github.com/IMQS/nf"
	"github.com/IMQS/nf/nfdb"
	"github.com/IMQS/serviceauth"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	nf.SendPong(w)
}

func (s *Service) assetList(w http.ResponseWriter, r *http.Request, p httprouter.Params, auth *serviceauth.Token) {
	query := ""
	args := []interface{}{}

	qtype := r.FormValue("type")
	if qtype != "" {
		query = "type IN " + nfdb.SQLCleanIDList(qtype)
	}

	assets := []asset{}
	s.db.Where(query, args...).Find(&assets)
	nf.SendJSON(w, assets)
}

func (s *Service) assetAdd(w http.ResponseWriter, r *http.Request, p httprouter.Params, auth *serviceauth.Token) {
	assets := []jsonInputAsset{}
	nf.ReadJSON(r, &assets)
}
