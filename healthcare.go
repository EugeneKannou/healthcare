package main

import (
	"fmt"
)

type Healthcare struct {
	collectors map[string]func() (ComponentHealth, error)
	stateful   map[string]*ComponentHealth
}

func New() *Healthcare {
	return &Healthcare{
		collectors: make(map[string]func() (ComponentHealth, error)),
		stateful:   make(map[string]*ComponentHealth),
	}
}

func (h *Healthcare) CollectSpecific(name string) (Health, error) {
	if collector, ok := h.collectors[name]; ok {
		health, err := collector()
		if err != nil {
			return nil, err
		}

		return Health{
			name: health,
		}, nil

	}
	if stateful, ok := h.stateful[name]; ok {
		health := stateful

		return Health{
			name: *health,
		}, nil
	}

	return Health{}, fmt.Errorf("component %v doesnt have any collectors or recievers", name)
}

func (h *Healthcare) Collect() Health {
	stateful := h.collectStateful()
	stateless := h.collectStateless()
	health := make(Health)

	for name, component := range stateful {
		health[name] = component
	}

	for name, component := range stateless {
		health[name] = component
	}

	return health
}
