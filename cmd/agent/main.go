package main

import (
	"flag"
	"log"
	"os"

	"metric-collection-service/internal/agent"
	"metric-collection-service/internal/config"
)

func main() {
	pwd, _ := os.Getwd()
	cfgPath := flag.String("c", pwd+"/etc/config.yml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.Init(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	a := agent.Init(cfg)
	log.Fatal(a.Run())
}
