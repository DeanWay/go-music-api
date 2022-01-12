package main

import (
	"go-music-api/config"
	"go-music-api/internal/app"
)

func main() {
	deps := config.DefaultDeps()
	app.RestApiApp(&deps).Run("localhost:8000")
}
