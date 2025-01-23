package handlers

import (
	"io"
	"log"
	"net/http"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}


func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := io.ReadAll(r.Body)
	
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Hello, World! " + string(d)))
}