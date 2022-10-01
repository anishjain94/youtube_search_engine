package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"youtube_search_engine/infra/environment"
	"youtube_search_engine/modules/youtube"

	"github.com/gorilla/mux"
)

func InitializeApiRestServer() {

	router := mux.NewRouter().StrictSlash(true)

	initializeApiRoutes(router)

	print("HTTP Api Server starting : " + fmt.Sprint(environment.PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", environment.PORT), router))
}

func initializeApiRoutes(router *mux.Router) {

	initHealthRoutes(router.PathPrefix("/health").Subrouter())

	youtube.InitUserRoutes(router.PathPrefix("/user").Subrouter())

}

func initHealthRoutes(router *mux.Router) {
	router.StrictSlash(true)
	router.HandleFunc("", healthCheck).Methods(http.MethodGet)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	msg := "Rest Server Up and runnning"
	json.NewEncoder(w).Encode(msg)
}
