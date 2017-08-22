package main

import (
	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
	"github.com/twosevenska/forge/server"
	"github.com/twosevenska/forge/splashes"
)

var conf server.Config

func main() {
	debug := flag.Bool("debug", false, "Display verbose debug output")

	flag.Parse()

	err := envconfig.Process("aesir", &conf)
	if err != nil {
		log.Fatalf("Failed to load Server Config: %s", err.Error())
	}

	if *debug || conf.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug Enabled")
		conf.Debug = true
	} else {
		log.SetLevel(log.InfoLevel)
	}

	splashes.Splash()

	log.Info("The flames are now roaring in the forge")

	server.Run(conf)

	log.Info("The flames have been extinguished.")
	log.Exit(0)
}
