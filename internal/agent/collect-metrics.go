package agent

import (
	"fmt"
	"math/rand"
	"runtime"

	"metric-collection-service/internal/model"
)

func (a *Agent) collectMetrics() {
	a.pollCount++

	memStat := runtime.MemStats{}
	runtime.ReadMemStats(&memStat)

	a.metrics = []model.Metric{
		{
			Name:  "Alloc",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.Alloc)),
		},
		{
			Name:  "BuckHashSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.BuckHashSys)),
		},
		{
			Name:  "Frees",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.Frees)),
		},
		{
			Name:  "GCCPUFraction",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", memStat.GCCPUFraction),
		},
		{
			Name:  "GCSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.GCSys)),
		},
		{
			Name:  "HeapAlloc",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.HeapAlloc)),
		},
		{
			Name:  "HeapIdle",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.HeapIdle)),
		},
		{
			Name:  "HeapInuse",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.HeapInuse)),
		},
		{
			Name:  "HeapObjects",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.HeapObjects)),
		},
		{
			Name:  "HeapReleased",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.HeapReleased)),
		},
		{
			Name:  "HeapSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.HeapSys)),
		},
		{
			Name:  "LastGC",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.LastGC)),
		},
		{
			Name:  "Lookups",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.Lookups)),
		},
		{
			Name:  "MCacheInuse",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.MCacheInuse)),
		},
		{
			Name:  "MCacheSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.MCacheSys)),
		},
		{
			Name:  "MSpanInuse",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.MSpanInuse)),
		},
		{
			Name:  "MSpanSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.MSpanSys)),
		},
		{
			Name:  "Mallocs",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.Mallocs)),
		},
		{
			Name:  "NextGC",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.NextGC)),
		},
		{
			Name:  "NumForcedGC",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.NumForcedGC)),
		},
		{
			Name:  "NumGC",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.NumGC)),
		},
		{
			Name:  "OtherSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.OtherSys)),
		},
		{
			Name:  "PauseTotalNs",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.PauseTotalNs)),
		},
		{
			Name:  "StackInuse",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.StackInuse)),
		},
		{
			Name:  "StackSys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.StackSys)),
		},
		{
			Name:  "Sys",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.Sys)),
		},
		{
			Name:  "TotalAlloc",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", float64(memStat.TotalAlloc)),
		},
		{
			Name:  "PollCount",
			Type:  "counter",
			Value: fmt.Sprintf("%d", a.pollCount),
		},
		{
			Name:  "RandomValue",
			Type:  "gauge",
			Value: fmt.Sprintf("%f", rand.Float64()),
		},
	}
}
