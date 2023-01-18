package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/scheduling"
	"sl-monitor/internal/business/stations"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/config"
	"sl-monitor/internal/database"
	"sl-monitor/internal/server/auth"
	"sl-monitor/internal/server/response"
	"sl-monitor/internal/smtp"
)

func main() {
	cfg := loadCfg()
	mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From)

	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	notificationsService := prepareNotificationService(db)
	notificationsHandler := notifications.NewHandler(notificationsService)
	stationsHandler := stations.NewHandler(cfg, trafikverket.NewAPIService())
	authHandler := auth.NewHandler(cfg)

	log.Printf("starting server on %s \n", cfg.Server.Addr)

	// setup routes
	http.HandleFunc("/", response.NotFound)
	notifications.Routes(notificationsHandler)
	stations.Routes(stationsHandler)
	auth.Routes(authHandler)

	scheduler := scheduling.Scheduler{Service: notificationsService, Mailer: mailer}
	scheduler.ScheduleNotifications()

	err := runServer(cfg.Server.Addr)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}

func prepareNotificationService(db *sql.DB) *notifications.NotificationService {
	store := notifications.NewStore(db)
	return notifications.NewService(store)
}

func prepareDatabase(dbName string) *sql.DB {
	sqlite := database.NewSqlite(dbName)
	log.Println(dbName)
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
