func tagHandler(e *Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == prefix {
			var id models.Tag
			json.NewDecoder(r.Body).Decode(&id)

			if r.Method == "GET" {
				w.WriteHeader(http.StatusOK)
			}
			if r.Method == "POST" {
				w.WriteHeader(http.StatusCreated)
				for _, t := range id.Tags {
					e.db.AddTag(id.ImageID, t)
				}
			}

			resp.Count, resp.Items = e.db.GetTags(id.ImageID)
			err := json.NewEncoder(w).Encode(resp)
			if err != nil {
				log.Error(err.Error())
			}
			return
		}
	}
}