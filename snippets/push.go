func indexHandler(w http.ResponseWriter, r *http.Request) {
	if pusher, ok := w.(http.Pusher); ok {
		if err := pusher.Push("/main.css", nil); err != nil {
			fmt.Println("Error:", err.Error())
		}
	}
}