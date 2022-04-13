package main

import (
	"dnsd/config"
	"dnsd/dns"
	"flag"
	"os"
	"strconv"
	"time"
)

var interval = flag.Int("i", getEnvAsInt("UPDATE_INTERVAL", 300), "update interval in seconds")
var configPath = flag.String("c", "config.yaml", "config path")

func main() {
	flag.Parse()
	conf, err := config.FromFile(*configPath)
	if err != nil {
		panic(err)
	}
	go dns.RunTimer(20*time.Second, time.Duration(*interval)*time.Second, conf)
}

func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return i
}
