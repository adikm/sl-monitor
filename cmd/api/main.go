package main

import (
	"database/sql"
	"flag"
	"log"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/stations"
	"sl-monitor/internal/config"
	database "sl-monitor/internal/database"
	customlogger "sl-monitor/internal/logger"
	auth2 "sl-monitor/internal/server/auth"
	"sl-monitor/internal/server/request"
	"sl-monitor/internal/smtp"
)

type application struct {
	logger *log.Logger
	config *config.Config
	mailer *smtp.Mailer
}

func main() {
	cfg := loadCfg()
	logger := customlogger.GetInstance()
	//mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	//app := &application{logger: logger, config: cfg, mailer: mailer, jsonCommon: jsonCommon}

	notificationsHandler := prepareNotificationHandler(db)
	stationsHandler := stations.NewHandler(cfg)
	authHandler := auth2.NewHandler(cfg)

	logger.Printf("starting server on %s", cfg.Server.Addr)

	DefaultRoutes()
	notifications.Routes(notificationsHandler)
	stations.Routes(stationsHandler)
	auth2.Routes(authHandler)

	err := request.Run(cfg.Server.Addr)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Print("server stopped")
}

func prepareNotificationHandler(db *sql.DB) *notifications.NotificationHandler {
	notificationsStore := notifications.NewStore(db)
	notificationsHandler := notifications.NewHandler(notificationsStore)
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
