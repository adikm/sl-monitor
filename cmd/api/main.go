package main

import (
	"flag"
	"log"
	"os"
	"sl-monitor/internal"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/stations"
	"sl-monitor/internal/config"
	database "sl-monitor/internal/database"
	"sl-monitor/internal/server/request"
	"sl-monitor/internal/smtp"
)

type application struct {
	logger     *log.Logger
	config     *config.Config
	mailer     *smtp.Mailer
	jsonCommon *internal.JsonCommon
}

func main() {
	cfg := loadCfg()
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
	mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)
	jsonCommon := &internal.JsonCommon{Logger: logger} // TODO new

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	app := &application{logger: logger, config: cfg, mailer: mailer, jsonCommon: jsonCommon}

	notificationsStore := notifications.NewStore(db.DB)
	notificationsHandler := notifications.NewHandler(notificationsStore, jsonCommon)
	stationsHandler := stations.NewHandler(cfg, jsonCommon)

	logger.Printf("starting server on %s", cfg.Server.Addr)

	app.routes(notificationsHandler, stationsHandler)
	err := request.Run(cfg.Server.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}

func prepareDatabase(dbName string) *database.DB {
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
