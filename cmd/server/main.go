package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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

	http.HandleFunc("/", MetricsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port), nil))
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}
