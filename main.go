package main

import "go-music-api/internal/app"

func main() {
	app.App().Run("localhost:8000")
}
