package main

import (
	"database/sql"
	"flag"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/stations"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/config"
	"sl-monitor/internal/database"
	customlogger "sl-monitor/internal/logger"
	"sl-monitor/internal/server/auth"
	"sl-monitor/internal/server/request"
)

func main() {
	cfg := loadCfg()
	logger := customlogger.GetInstance()
	//mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	notificationsHandler := prepareNotificationHandler(db)
	stationsHandler := stations.NewHandler(cfg, trafikverket.NewAPIService())
	authHandler := auth.NewHandler(cfg)

	logger.Printf("starting server on %s", cfg.Server.Addr)

	DefaultRoutes()
	notifications.Routes(notificationsHandler)
	stations.Routes(stationsHandler)
	auth.Routes(authHandler)

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
