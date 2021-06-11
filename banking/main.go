package main

import (
	"github.com/neoyewchuan/RestDevGo/banking/app"
	"github.com/neoyewchuan/RestDevGo/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
