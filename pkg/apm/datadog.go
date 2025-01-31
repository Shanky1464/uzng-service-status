package apm

import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

func Start(opts ...tracer.StartOption) {
	tracer.Start(opts...)
}

func Stop() {
	tracer.Stop()
}
