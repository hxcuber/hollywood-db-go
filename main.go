package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hxcuber/hollywood-db-go/controllers"
	"github.com/hxcuber/hollywood-db-go/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)
	controllers.Init()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	fmt.Printf("Starting server on port %d\n", 3333)
	err := http.ListenAndServe(":3333", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
