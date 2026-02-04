package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server is running on port: 3000")
	addr := ":" + strconv.Itoa(cnf.HttpPort)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("Error starting server", wrappedMux)
		os.Exit(1)
	}
}
