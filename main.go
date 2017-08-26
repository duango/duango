package main

import (
	log "github.com/sirupsen/logrus"
	"runtime"
	"github.com/spf13/pflag"
	"fmt"
	"os"
	"strings"
)

func main() {


	configPath := *pflag.StringP("config","c","config.toml","Path of Duango's config file.")
	getVersion := *pflag.Bool("version", false, "Print Duango's server version.")
	logLevel := *pflag.StringP("level", "l", "debug", "Log level.")

	if getVersion {
		fmt.Printf("Duango, a Minecraft server. ver: %s\n", VERSION)
		os.Exit(0)
	}

	// set log level for logrus
	switch (strings.ToLower(logLevel)) {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		fmt.Println("log level only can be set as debug/error/info/panic/warn/fatal")
		os.Exit(1)
	}

	log.WithFields(log.Fields{
		"CPUNum": runtime.GOMAXPROCS(-1),
		"configPath": configPath,
		"LogLevel": logLevel,
	}).Info("Server Info")

	runtime.GOMAXPROCS(runtime.NumCPU())

	server := NewServer()

	server.ListenAndServer()
}

