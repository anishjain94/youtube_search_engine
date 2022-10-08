package main

import (
	"youtube_search_engine/infra/database"
	"youtube_search_engine/infra/environment"
	"youtube_search_engine/infra/rest"
	youtube_integration "youtube_search_engine/integrations/youtube_integration"
)

func main() {
	environment.InitializeEnvs()
	database.InitializeGorm()
	youtube_integration.InitializeYoutube()

	rest.InitializeApiRestServer()
}
