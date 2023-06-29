package storage

import "metric-collection-service/internal/model"

type BaseStorage interface {
	Write(data *[]model.Metric)
	Read() *[]model.Metric
}
