package main

import (
	"banner-display-service/src/internal/app"
)

const configPath = "src/config/config.yml"

func main() {
	app.Run(configPath)
}