package main

import (
	"github.com/vipinkavlar/gofiber-boilerplate/core/engine"
)

func main() {

	//setup various configuration for app
	//config.LoadAllConfigs(".env")

	engine.Start()
}