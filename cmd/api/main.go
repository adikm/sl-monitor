package main

import (
	"database/sql"
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
	//mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)
	jsonCommon := internal.NewJsonCommon(logger)

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	//app := &application{logger: logger, config: cfg, mailer: mailer, jsonCommon: jsonCommon}

	notificationsHandler := prepareNotificationHandler(db, jsonCommon)
	stationsHandler := stations.NewHandler(cfg, jsonCommon)

	logger.Printf("starting server on %s", cfg.Server.Addr)

	notifications.Routes(notificationsHandler)
	stations.Routes(stationsHandler)

	err := request.Run(cfg.Server.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}

func prepareNotificationHandler(db *sql.DB, jsonCommon *internal.JsonCommon) *notifications.NotificationHandler {
	notificationsStore := notifications.NewStore(db)
	notificationsHandler := notifications.NewHandler(notificationsStore, jsonCommon)
	return notificationsHandler
}

func prepareDatabase(dbName string) *sql.DB {
	sqlite := database.NewSqlite(dbName)
	db := sqlite.Connect()
	sqlite.Migrate(db)
	return db
}

func loadCfg() *config.Config {
	var configFile string
	flag.StringVar(&configFile, "file", "config.yml", "config file (yaml)")
	flag.Parse()
	cfg := config.Load(&configFile)
	return cfg
}
