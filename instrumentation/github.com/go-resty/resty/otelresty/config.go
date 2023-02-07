// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otelresty // import "go.opentelemetry.io/contrib/instrumentation/github.com/go-resty/resty/otelresty"

import (
	"github.com/go-resty/resty/v2"

	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// Skipper is a function used to determine if no span should be created for a request.
type Skipper func(*resty.Request) bool

// config is used to configure the go-resty middleware.
type config struct {
	TracerProvider oteltrace.TracerProvider
	Propagators    propagation.TextMapPropagator
	Skipper        Skipper
}

// Option applies a configuration value.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithSkipper specifies a skipper function to determine if the middleware
// should not create a span for a determined request. If not specified,
// a span will always be created.
func WithSkipper(skipper Skipper) Option {
	return optionFunc(func(c *config) {
		if skipper != nil {
			c.Skipper = skipper
		}
	})
}

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	return optionFunc(func(cfg *config) {
		if propagators != nil {
			cfg.Propagators = propagators
		}
	})
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider oteltrace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}
