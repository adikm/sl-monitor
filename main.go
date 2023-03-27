package main

import (
	"database/sql"
	"flag"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"sl-monitor/internal/business/notifications"
	"sl-monitor/internal/business/scheduling"
	"sl-monitor/internal/business/stations"
	"sl-monitor/internal/business/stations/trafikverket"
	"sl-monitor/internal/business/users"
	"sl-monitor/internal/config"
	"sl-monitor/internal/database"
	"sl-monitor/internal/server/auth"
	"sl-monitor/internal/server/response"
	"sl-monitor/internal/smtp"
)

func main() {
	cfg := loadCfg()

	/*
		DATABASE
	*/
	db := prepareDatabase(cfg.Database.Name)
	defer db.Close()

	/*
		BUSINESS
	*/
	usersService := users.NewService(users.NewStore(db))
	mailer := smtp.NewMailer(cfg.Mail.SmtpHost, cfg.Mail.SmtpPort, cfg.Mail.From, cfg.Mail.Password, cfg.Mail.From, usersService)
	notificationsService := prepareNotificationService(db)
	notificationsHandler := notifications.NewHandler(notificationsService)
	tvService := trafikverket.NewAPIService(cfg.TrafficAPI.AuthKey)
	stationsHandler := stations.NewHandler(tvService)
	usersHandler := users.NewHandler(usersService)
	authHandler := auth.NewHandler(cfg)

	log.Printf("starting server on %s \n", cfg.Server.Addr)

	/*
		ROUTES
	*/
	http.HandleFunc("/", response.NotFound)
	users.Routes(usersHandler)
	notifications.Routes(notificationsHandler)
	stations.Routes(stationsHandler)
	auth.Routes(authHandler)

	/*
		SCHEDULING
	*/
	sender := scheduling.NewSender(usersService, tvService)
	scheduler := scheduling.NewScheduler(notificationsService, sender, mailer)

	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", //every day at midnight
		func() {
			scheduler.ScheduleNotifications()
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Start()

	err = runServer(cfg.Server.Addr)

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
