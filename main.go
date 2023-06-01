package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	controller "github.com/spider/internalportalbackend/controller"
)

func main() {

	// Session := controller.SetupDB()
	// if(Session == nil){
	// 	fmt.Println("DB connection failed")
	// 	return
	// }

	router := chi.NewRouter()
	router.Mount("/users", http.HandlerFunc(controller.Getuser))
	router.Mount("/projects", http.HandlerFunc(controller.Getprojects))
	fmt.Println("Server started on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", router))
}
