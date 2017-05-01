// EnvHandler - A type for our environment accepting handlers
type EnvHandler func(env *Env, writer http.ResponseWriter, req *http.Request) http.HandlerFunc

// ContextRouter - A struct for wrapping context around a handler
type ContextRouter struct {
	*Env
	Adapters []Adapter

	H EnvHandler
}

func (cr ContextRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := Adapt(cr.H(cr.Env, w, r), cr.Adapters...)

	f(w, r)
}

func mapAPIRoutes(mux *http.serveMux) {
	mapRoute(env, "/tag/", tagHandler, defaultJSONAdapters, mux)
}

func mapRoute(e *Env, r string, h EnvHandler, a []Adapter, m *http.ServeMux) {
	m.Handle(e.config.APIBase+r, ContextRouter{e, a, h})
}