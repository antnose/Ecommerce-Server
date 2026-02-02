package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

// Get All Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflight(w, r)

	sendData(w, productList, 200)
}

// Create Product
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflight(w, r)

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

// All Cors Method
func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

// Handle Preflight Request
func handlePreflight(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /hello", http.HandlerFunc(handler))

	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))

	mux.Handle("GET /products", http.HandlerFunc(getProducts))

	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port: 3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red",
		Price:       123,
		ImgUrl:      "https://imgs.search.brave.com/tirWQ6LwYqy7qryevI2CPhs3cPJopGXtyC2AB8_uxD8/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9wbHVz/LnVuc3BsYXNoLmNv/bS9wcmVtaXVtX3Bo/b3RvLTE2NzA1MTIx/ODEwNjEtZTI0Mjgy/ZjdlZTc4P2ZtPWpw/ZyZxPTYwJnc9MzAw/MCZhdXRvPWZvcm1h/dCZmaXQ9Y3JvcCZp/eGxpYj1yYi00LjEu/MCZpeGlkPU0zd3hN/akEzZkRCOE1IeHpa/V0Z5WTJoOE1YeDhi/M0poYm1kbGZHVnVm/REI4ZkRCOGZId3c",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Orange is green",
		Price:       123,
		ImgUrl:      "https://imgs.search.brave.com/l_kAL4yktCPYyQgrzzc0AYapzr13rq1OvgSOgXjpjlo/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly90aHVt/YnMuZHJlYW1zdGlt/ZS5jb20vYi9hcHBs/ZS1pcGhvbmUtaW9z/LXVwZ3JhZGUtc3Rh/cnQtc2NyZWVuLWhv/dXN0b24tdGV4YXMt/dXNhLXNlcHQtY2xv/c2UtdXAtaXBob25l/LXN0YXJ0LXNjcmVl/bi1hcHBsZS1pY29u/LWlvcy11cGdyYWRl/LTEwMDUxODgyNC5q/cGc",
	}

	prd3 := Product{
		ID:          2,
		Title:       "Lychee",
		Description: "Orange is White",
		Price:       123,
		ImgUrl:      "https://imgs.search.brave.com/cQjLAoBi6YOD42C2muCTwCSW037CcrlLavWpPaW0xeM/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvMjE2/MjUyNTA0NC9waG90/by9zbWFsbC1yaXBl/LXRyb3BpY2FsLWx5/Y2hlZS1iZXJyaWVz/LWluLWEtbWluaW1h/bGlzdC1zdHlsZS1s/b3cta2V5LmpwZz9z/PTYxMng2MTImdz0w/Jms9MjAmYz1hLTZm/YmtxMFRISUVod2xy/VDVEMnI4eHdEb0hQ/MkZkMlVvYjVPbzNC/VGdZPQ",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
}
