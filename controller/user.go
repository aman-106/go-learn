package controller

import "net/http"

type userCtler struct{}

func (uc userCtler) method(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("skddkdk"))
}

