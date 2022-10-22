package main

import (
	"flag"
	"log"
	"os"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server"
)

type application struct {
	logger *log.Logger
	config *config.Config
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "file", "config.yml", "config file (yaml)")
	flag.Parse()
	cfg := config.Load(&configFile)

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	app := &application{logger: logger, config: cfg}

	logger.Printf("starting server on %s", cfg.Server.Addr)

	app.routes()
	err := server.Run(cfg.Server.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}
