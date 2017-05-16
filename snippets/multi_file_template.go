func RenderAdmin(w http.ResponseWriter, r *http.Request) {
	...
	
	t, err := template.ParseFiles("web/parent.html", "web/child.html")
	// check err
	t.Execute(w, a)
}