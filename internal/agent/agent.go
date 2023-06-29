package agent

import (
	"fmt"

	"metric-collection-service/internal/config"
	"metric-collection-service/internal/model"

	"github.com/robfig/cron/v3"
)

type Agent struct {
	reportAddress  string
	scheduler      *cron.Cron
	pollInterval   string
	reportInterval string
	metrics        []model.Metric
	pollCount      int64
}

func Init(cfg *config.Config) *Agent {
	return &Agent{
		reportAddress:  fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port),
		scheduler:      cron.New(cron.WithSeconds()),
		pollInterval:   cfg.Agent.PollInterval,
		reportInterval: cfg.Agent.ReportInterval,
		metrics:        make([]model.Metric, 0, 30),
		pollCount:      0,
	}
}

func (a *Agent) Run() (err error) {
	_, err = a.scheduler.AddFunc(a.pollInterval, func() { a.collectMetrics() })
	if err != nil {
		return err
	}
	_, err = a.scheduler.AddFunc(a.reportInterval, func() { err = a.sendMetrics() })
	if err != nil {
		return err
	}
	a.scheduler.Start()
	select {}
}
