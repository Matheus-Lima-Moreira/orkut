package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	API_URL   = ""
	API_PORT  = 0
	HASH_KEY  []byte
	BLOCK_KEY []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file")
	}

	API_PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		API_PORT = 3000
	}

	API_URL = os.Getenv("API_URL")
	HASH_KEY = []byte(os.Getenv("HASH_KEY"))
	BLOCK_KEY = []byte(os.Getenv("BLOCK_KEY"))
}
