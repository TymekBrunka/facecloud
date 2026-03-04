package netguard

import (
	"net/http"
)

type GuardFunc func(http.ResponseWriter, *http.Request) (error)

type GuardGroup struct {
	Guards  []GuardFunc
	Handler http.Handler
}

func (g GuardGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	for _, guard := range g.Guards {
		err = guard(w, r)
		if err != nil {
			return
		}
	}

	if err == nil {
		g.Handler.ServeHTTP(w, r)
	}
}
