package main

import "fmt"

func (h *Healthcare) AddStateless(name string, collector func() (ComponentHealth, error)) error {
	if _, ok := h.collectors[name]; ok {
		return fmt.Errorf("collector for component %v already exists", name)
	}
	if _, ok := h.stateful[name]; ok {
		return fmt.Errorf("component %v already has stateful reciever", name)
	}

	h.collectors[name] = collector

	return nil
}

func (h *Healthcare) RemoveStateless(name string) {
	delete(h.collectors, name)
}

func (h *Healthcare) collectStateless() Health {
	health := make(Health)
	for name, collector := range h.collectors {
		componentHealth, err := collector()
		if err != nil {
			health[name] = ComponentHealth{
				Status:  StatusUnhealthy,
				Message: err.Error(),
			}
			continue
		}
		health[name] = componentHealth
	}

	return health
}
