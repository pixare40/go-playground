package handlers

import (
	"log"
	"net/http"
	"pixare40/hello/data"
)


type Products struct {
	logger *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w)
	}
}

func (p *Products) getProducts(w http.ResponseWriter) {
	products := data.GetProducts()

	err := products.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}