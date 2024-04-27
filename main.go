package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shreyash2101/online-job-portal/router"
)

func main() {
	fmt.Println("Welcome to online job portal")
	
	r := router.Router()
	fmt.Println("Server is getting started at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}