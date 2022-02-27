package main

import (
	"github.com/vipinkavlar/gofiber-boilerplate/core/bootstrap"
)

func main() {

	//setup various configuration for app
	//config.LoadAllConfigs(".env")

	bootstrap.Start()
}