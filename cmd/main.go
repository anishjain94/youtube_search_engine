package main

import (
	"youtube_search_engine/infra/database"
	"youtube_search_engine/infra/environment"
	"youtube_search_engine/infra/rest"
)

func main() {
	environment.InitializeEnvs()
	database.InitializeGorm()

	rest.InitializeApiRestServer()
}
