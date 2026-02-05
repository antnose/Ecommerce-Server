package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
)

func Serve() {
	cnf := config.GetConfig()

	rest.Start(cnf)
	// Check Github
}

// [Golang] 054 - Into The Actual Project Structure 1.04.07
