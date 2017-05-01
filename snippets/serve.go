log.Fatal(http.ListenAndServe(":8080", mux))

log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", mux))