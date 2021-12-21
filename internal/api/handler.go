package api

import (
	"github.com/enkhalifapro/obi-partners/internal/types"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PartnerManager interface {
	Add(partner *types.Partner) error
}

func Handler(partnerMgr PartnerManager) http.Handler {
	r := httprouter.New()
	r.HandlerFunc(http.MethodPost, "/partner", func(w http.ResponseWriter, r *http.Request) {
		// todo: to be implemented
		// routing logic
		partnerMgr.Add(nil)
	})
	return r
}
