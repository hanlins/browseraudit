package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func InterSetHandler(w http.ResponseWriter, r *http.Request) {
	DontCache(&w)

	session := store.Get(w, r)

	id := mux.Vars(r)["id"]
	val := mux.Vars(r)["val"]

	// If ?serveOnly=true, just serve the iframe: don't change the test
	// state in the session cookie
	r.ParseForm()
	if r.Form["serveOnly"] == nil || r.Form["serveOnly"][0] != "true" {
		session.Set("inter"+id, val)
	}
}

func InterGetHandler(w http.ResponseWriter, r *http.Request) {
	DontCache(&w)

	session := store.Get(w, r)

	id := mux.Vars(r)["id"]

	result, err := session.Get("inter" + id)
	if err != nil {
		log.Printf("nil result sop%s", id)
		result = "nil"
	}

	fmt.Fprintf(w, "%s", result)
}

func RegisterPassHandler(w http.ResponseWriter, r *http.Request) {
	DontCache(&w)

	id := mux.Vars(r)["id"]
	session := store.Get(w, r)
	session.Set("register"+id, "pass")

	http.ServeFile(w, r, "./static/pixel.png")
}

func RegisterFailHandler(w http.ResponseWriter, r *http.Request) {
	DontCache(&w)

	id := mux.Vars(r)["id"]
	session := store.Get(w, r)
	session.Set("register"+id, "fail")

	http.ServeFile(w, r, "./static/pixel.png")
}

func RegisterResultHandler(w http.ResponseWriter, r *http.Request) {
	DontCache(&w)

	session := store.Get(w, r)

	result, err := session.Get("register" + mux.Vars(r)["id"])
	if err != nil {
		log.Printf("nil result frameoptions%s", mux.Vars(r)["id"])
		result = "nil"
	}

	fmt.Fprintf(w, "%s", result)
}
