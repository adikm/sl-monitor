package main

import (
	"flag"
	"log"
	"os"
	"sl-monitor/internal/config"
	"sl-monitor/internal/server"
	"sl-monitor/internal/smtp"
)

type application struct {
	logger *log.Logger
	config *config.Config
	mailer *smtp.Mailer
}

func main() {
	cfg := loadCfg()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
	mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)
	app := &application{logger: logger, config: cfg, mailer: mailer}

	logger.Printf("starting server on %s", cfg.Server.Addr)

	app.routes()
	err := server.Run(cfg.Server.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}

func loadCfg() *config.Config {
	var configFile string
	flag.StringVar(&configFile, "file", "config.yml", "config file (yaml)")
	flag.Parse()
	cfg := config.Load(&configFile)
	return cfg
}
