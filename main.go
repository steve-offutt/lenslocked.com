package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `To get in touch, please send and email to <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>.`)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	faq := `<h1>Awesome site FAQ</h1><ul><li>question 1</li><li>question 2</li></ul>`
	fmt.Fprint(w, faq)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	faq := `<h1>NOT FOUND DAWG</h1>`
	fmt.Fprint(w, faq)
}

func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
