package youtube

import (
	"net/http"

	"github.com/gorilla/mux"

	"youtube_search_engine/common"
)

func InitUserRoutes(router *mux.Router) {
	router.StrictSlash(true)

	router.HandleFunc("", common.HandleHTTPGet(getVideosFromDb)).Methods(http.MethodGet)

}
