package metrics

import (
	"github.com/DataDog/datadog-go/statsd"
)

type Metrics = statsd.Client

func New(conf Config) (*Metrics, error) {
	client, err := statsd.New(
		conf.Host,
		statsd.WithNamespace("hello_world_go"),
		statsd.WithTags([]string{"app:hello-world-go"}),
	)
	return client, err
}
