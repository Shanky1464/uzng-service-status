package metrics

import (
	"github.com/DataDog/datadog-go/statsd"
)

type Metrics = statsd.Client

func New(conf Config) (*Metrics, error) {
	client, err := statsd.New(
		conf.Host,
		statsd.WithNamespace("uzng_service_status"),
		statsd.WithTags([]string{"app:uzng-service-status"}),
	)
	return client, err
}
