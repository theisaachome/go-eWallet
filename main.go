package main

import (
	"github.com/theisaachome/go-eWallet/app"
	"github.com/theisaachome/go-eWallet/logger"
)

func main() {

	logger.Info("Starting eWallet application.....")

	app.StartApplication()
}
