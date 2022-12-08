package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/stations"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/config"
	"sl-monitor/internal/database"
	"sl-monitor/internal/server/auth"
	"sl-monitor/internal/server/response"
)

func main() {
	cfg := loadCfg()
	//mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	notificationsHandler := prepareNotificationHandler(db)
	stationsHandler := stations.NewHandler(cfg, trafikverket.NewAPIService())
	authHandler := auth.NewHandler(cfg)

	log.Printf("starting server on %s \n", cfg.Server.Addr)

	// setup routes
	http.HandleFunc("/", response.NotFound)
	notifications.Routes(notificationsHandler)
	stations.Routes(stationsHandler)
	auth.Routes(authHandler)

	err := runServer(cfg.Server.Addr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}

func prepareNotificationHandler(db *sql.DB) *notifications.Handler {
	notificationsStore := notifications.NewStore(db)
	notificationsHandler := notifications.NewHandler(notificationsStore)
	return notificationsHandler
}

func prepareDatabase(dbName string) *sql.DB {
	sqlite := database.NewSqlite(dbName)
	db, err := sqlite.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = sqlite.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func loadCfg() *config.Config {
	var configFile string
	flag.StringVar(&configFile, "file", "config.yml", "config file (yaml)")
	flag.Parse()
	cfg := config.Load(&configFile)
	return cfg
}
