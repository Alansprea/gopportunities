package main

import (
	"github.com/Alansprea/gopportunities/config"
	"github.com/Alansprea/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {

	logger = *config.GetLogger(("main"))
	// Inicialize Conifgs

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error : v%", err)
		return
	}
	// Inicialize Router
	router.Inicialize()
}
