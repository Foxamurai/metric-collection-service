package server

import (
	"fmt"
	"log"
	"net/http"

	"metric-collection-service/internal/config"
	"metric-collection-service/internal/server/handler"
	"metric-collection-service/internal/server/storage"
)

type MetricsServer struct {
	storage *storage.BaseStorage
	address string
	port    string
}

func Init(cfg *config.Config, storage *storage.BaseStorage) MetricsServer {
	return MetricsServer{
		storage: storage,
		address: cfg.Server.Address,
		port:    cfg.Server.Port,
	}
}

func (s MetricsServer) Run() error {
	http.HandleFunc("/update/", handler.UpdateMetrics)
	log.Printf("Server starting on %s:%s", s.address, s.port)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", s.address, s.port), nil)
}
