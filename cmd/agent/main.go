package main

import (
	"flag"
	"log"
	"os"

	"github.com/robfig/cron/v3"
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

	var metrics []agent.Metric

	c := cron.New(cron.WithSeconds())
	_, err = c.AddFunc(cfg.Agent.PollInterval, func() { metrics = agent.CollectMetrics() })
	if err != nil {
		log.Fatal(err)
	}
	_, err = c.AddFunc(cfg.Agent.ReportInterval, func() { err = agent.SendMetrics(cfg, metrics) })
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
	select {}
}
