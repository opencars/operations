package main

import (
	"flag"

	_ "github.com/lib/pq"

	"github.com/opencars/operations/pkg/apiserver"
	"github.com/opencars/operations/pkg/config"
	"github.com/opencars/operations/pkg/logger"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	conf, err := config.New(configPath)
	if err != nil {
		logger.Fatal(err)
	}

	if err := apiserver.Start(":8080", conf); err != nil {
		logger.Fatal(err)
	}
}
