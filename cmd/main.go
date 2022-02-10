package main

import (
	routes "github.com/galigaribaldi/PersonalBooks/pkg/routers"
)

func main() {
	// Our server will live in the routes package
	routes.Run()
}
