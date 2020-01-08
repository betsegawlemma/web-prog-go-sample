package handler

import "net/http"

// About handles requests on route /about
func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

// Contact handles requests on route /contact
func Contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact"))
}
