package main

import (
	"database/sql"
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/robfig/cron/v3"
	"log"
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
	db := prepareDatabase(cfg.Database)
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
		ROUTING/SERVER
	*/
	r := routerWithMiddleware()
	r.HandleFunc("/", response.NotFound)
	users.Routes(r, usersHandler)
	notifications.Routes(r, notificationsHandler)
	stations.Routes(r, stationsHandler)
	auth.Routes(r, authHandler)

	err := runServer(cfg.Server.Addr, r)

	if err != nil {
		log.Fatal(err)
	}

	/*
		SCHEDULING
	*/
	sender := scheduling.NewSender(usersService, tvService)
	scheduler := scheduling.NewScheduler(notificationsService, sender, mailer)

	c := cron.New()
	_, err = c.AddFunc("0 0 * * *", //every day at midnight
		func() {
			go scheduler.ScheduleNotifications()
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Start()

}

func routerWithMiddleware() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r
}

func prepareNotificationService(db *sql.DB) *notifications.NotificationService {
	store := notifications.NewStore(db)
	return notifications.NewService(store)
}

func prepareDatabase(dbConfig config.Database) *sql.DB {
	sqlite := database.NewPostgre(dbConfig)
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
