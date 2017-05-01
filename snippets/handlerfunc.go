// RenderCSS - Serves main.css
func RenderCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/main.css")
}

// http.ResponseWriter satisfies the io.Writer interface
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// http.ResponseWriter can also be used to write JSON
func JSONVersionHandler(w http.ResponseWriter, r *http.Request) {
	version := struct {
		Version   string
		BuildTime string
		Commit    string
	}{e.config.Version, e.config.BuildTime, e.config.Commit}
	json.NewEncoder(w).Encode(version)
}