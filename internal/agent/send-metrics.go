package agent

import (
	"fmt"
	"net/http"
)

func (a *Agent) sendMetrics() error {
	for _, m := range a.metrics {
		r, err := http.Post(
			fmt.Sprintf(
				"http://%s/update/%s/%s/%s",
				a.reportAddress,
				m.Type,
				m.Name,
				m.Value,
			),
			"text/plain",
			nil,
		)
		r.Body.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
