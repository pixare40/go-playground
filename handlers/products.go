package handlers

import (
	"log"
	"net/http"
	"pixare40/hello/data"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
	}

	if r.Method == http.MethodPut {
		p.logger.Println("Handle PUT Products", r.URL.Path)
		regex := regexp.MustCompile(`/([0-9]+)`)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
			
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		
		p.logger.Println("Got id", id)
	}	
}

func (p *Products) getProducts(w http.ResponseWriter) {
	products := data.GetProducts()

	err := products.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

	products := data.GetProducts()
	err = products.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}