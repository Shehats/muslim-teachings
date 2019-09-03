package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muslim-teachings/api-orchestrator/src/main/controllers"
	// "github.com/muslim-teachings/api-orchestrator/src/main/controllers"
	// "github.com/muslim-teachings/api-orchestrator/src/main/middleware"
)

const (
	// POST methods
	POST = "POST"
	// PUT handlers
	PUT = "PUT"
	// GET hanlders
	GET = "GET"
	// DELETE handlers
	DELETE = "DELETE"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := mux.NewRouter()

	// router.NotFoundHandler = Handle404()

	{
		router.HandleFunc("/teachings", controllers.GetTeachings).Methods(GET)
	}

	{
		router.HandleFunc("/quran", controllers.GetQuran).Methods(GET)
	}

	http.ListenAndServe(":9000", router)
}
