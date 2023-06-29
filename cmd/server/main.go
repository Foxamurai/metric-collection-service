package main

import (
	"flag"
	"log"
	"os"

	"metric-collection-service/internal/config"
	"metric-collection-service/internal/server"
	"metric-collection-service/internal/server/storage"
)

func main() {
	pwd, _ := os.Getwd()
	cfgPath := flag.String("c", pwd+"/etc/config.yml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.Init(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	var stg storage.BaseStorage = storage.Init()
	srv := server.Init(cfg, &stg)
	log.Fatal(srv.Run())
}
