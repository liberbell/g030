package main

import (
	"flag"
	"log"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	smtp struct {
		host     string
		port     int
		username string
		password string
	}
	frontend string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 5000, "Server port to listen on")
	flag.StringVar(&cfg.smtp.host, "smtphost", "sandbox.smtp.mailtrap.io", "smtp host")
	cfg.smtp.username = os.Getenv("MAIL_User")
	cfg.smtp.password = os.Getenv("MAIL_Pass")
	flag.IntVar(&cfg.smtp.port, "smtpport", 587, "Smtp Port")
	flag.StringVar(&cfg.frontend, "frontend", "http://localhost:4000", "url to front end")

	flag.Parse()
}
