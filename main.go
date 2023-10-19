package main

import "net/http"

func main() {
	http.HandleFunc('/api/v1/user', func(w http.ResponseWriter, r *http.Request)){
		w.Write([]byte("ОК"))
	}
	w.Write([]byte("ОК"))

	http.ListenAndServe("localhost:8080", nil)
}