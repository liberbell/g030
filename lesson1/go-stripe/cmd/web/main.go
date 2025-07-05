package main

import (
	"flag"
	"html/template"
	"log"
	"os"
)

const version = "1.0.0"
const cssVersion = "1"

type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "Port", 4000, "Server Port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application envirnment {decelopment|environment}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
}
