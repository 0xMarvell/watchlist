package config

import (
	"github.com/0xMarvell/watchlist/pkg/utils"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables
func LoadEnv() {
	dotEnvErr := godotenv.Load()
	utils.CheckErr("error loading env variables: ", dotEnvErr)
}
