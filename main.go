package main

import "go-music-api/internal/app"

func main() {
	deps := app.DefaultDeps()
	app.App(&deps).Run("localhost:8000")
}
