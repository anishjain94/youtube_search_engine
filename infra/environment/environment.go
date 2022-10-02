package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT        int
	ENV         string
)

func InitializeEnvs(variant ...string) {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/youtube_search_engine/config/dev.env"))
	if err != nil {
		log.Fatal("unable to load " + ENV + ".env file")
	}

	ENV = os.Getenv("ENV")
	PORT, _ = strconv.Atoi(os.Getenv("APP_PORT"))
}
