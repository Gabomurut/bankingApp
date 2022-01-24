package main

import (
	"github.com/bankingApp/app"
	"github.com/bankingApp/logger"
)

func main() {

	logger.Info("Starting the app...")
	app.Start()

}
