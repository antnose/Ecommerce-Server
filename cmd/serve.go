package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products",
		middleware.Logger(
			http.HandlerFunc(handlers.GetProducts),
		),
	)

	mux.Handle("POST /products",
		middleware.Logger(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle("GET /products/{productId}",
		middleware.Logger(
			http.HandlerFunc(handlers.GetProductById),
		),
	)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server is running on port: 3000")
	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
