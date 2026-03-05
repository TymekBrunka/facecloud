package netguard

import (
	cfg "fcserver/config"
	"net/http"
)

type GuardFunc func(http.ResponseWriter, *http.Request) error
type Guarded_Func func(*cfg.Config, http.ResponseWriter, *http.Request)

type GuardGroup struct {
	guards  []GuardFunc
	handler Guarded_Func
	state   *cfg.Config
}

func (g GuardGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	for _, guard := range g.guards {
		err = guard(w, r)
		if err != nil {
			return
		}
	}

	if err == nil {
		g.handler(g.state, w, r)
	}
}

func Path(state *cfg.Config, guards []GuardFunc, path string, handler Guarded_Func) {
	handler_ := GuardGroup{
		guards: guards,
		handler: handler,
		state: state,
	}
	http.Handle(path, handler_)
}
