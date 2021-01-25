package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// HostName ...
var HostName string

// StandardFields ...
var StandardFields = log.Fields{
	"hostname": "MyPC",
	"appname":  "WebVideo",
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)

	// Get Hostname for updating Log StandardFields
	HostName, err := os.Hostname()
	if err != nil {
		log.WithFields(StandardFields).Infof("Error in getting the Hostname: %v", err)
	}
	// Check if loading environment variables from .env file is required
	if os.Getenv("LOAD_CONFIG_FILE") == "" {
		log.WithFields(StandardFields).Infof("util.StandardFields")
		// Load environment variables from .env file
		err = godotenv.Load()
		if err != nil {
			log.WithFields(StandardFields).Infof("Error in reading the coonfig file: %v", err)
		}
	}

	StandardFields = log.Fields{
		"hostname": HostName,
		"appname":  os.Getenv("APPNAME"),
	}
}

func main() {
	log.WithFields(StandardFields).Infof("Starting WebVideo Server at Port: %s", os.Getenv("PORT"))
	http.Handle("/", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
