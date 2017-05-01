func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	for _, route := range m.Routes {
		if route.URL == url {
			route.Handler(w, r)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}