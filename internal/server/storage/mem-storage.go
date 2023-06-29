package storage

import "metric-collection-service/internal/model"

type MemStorage struct {
	data []model.Metric
}

func Init() *MemStorage {
	return &MemStorage{
		data: make([]model.Metric, 0, 30),
	}
}

func (s *MemStorage) Write(data *[]model.Metric) {
	s.data = *data
}

func (s *MemStorage) Read() *[]model.Metric {
	return &s.data
}
