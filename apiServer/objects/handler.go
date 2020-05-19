package objects

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Access-Control-Allow-Headers,Content-Length,Accept,Authorization,X-Requested-With,Digest")
	w.Header().Add("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")

	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	if r.Method == "OPTIONS" {
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
