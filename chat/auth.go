package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")
	if err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		// some other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// success - call the next handler
	h.next.ServeHTTP(w, r)
}

 func MustAuth(handler http.Handler) http.Handler {
	 return &authHandler{next: handler}
 }

 //handles third party login process
 //format: /auth/{action}/{third-party provider}
 func loginHandler(w http.ResponseWriter, r *http.Request) {
	 segs := strings.Split(r.URL.Path, "/")
	 action := segs[2]
	 provider := segs[3]
	 switch action {
	 case "login":
		 log.Println("TODO handle login for", provider)
	 default:
		 w.WriteHeader(http.StatusNotFound)
		 fmt.Fprintf(w, "Auth action %s not supported", action)
	 }
 }
