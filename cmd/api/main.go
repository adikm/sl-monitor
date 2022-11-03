package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"sl-monitor/internal/config"
	"sl-monitor/internal/database"
	"sl-monitor/internal/server"
	"sl-monitor/internal/smtp"
)

type application struct {
	logger   *log.Logger
	config   *config.Config
	mailer   *smtp.Mailer
	database *sql.DB
}

func main() {
	cfg := loadCfg()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
	mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	app := &application{logger: logger, config: cfg, mailer: mailer, database: db}

	logger.Printf("starting server on %s", cfg.Server.Addr)

	app.routes()
	err := server.Run(cfg.Server.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}

func prepareDatabase(dbName string) *sql.DB {
	db := database.Connect(dbName)
	database.Migrate(db, dbName)
	return db
}

func loadCfg() *config.Config {
	var configFile string
	flag.StringVar(&configFile, "file", "config.yml", "config file (yaml)")
	flag.Parse()
	cfg := config.Load(&configFile)
	return cfg
}
