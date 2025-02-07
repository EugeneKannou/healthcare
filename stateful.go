package main

import "fmt"

func (h *Healthcare) AddStateful(name string) (*ComponentHealth, error) {
	if _, ok := h.stateful[name]; ok {
		return nil, fmt.Errorf("reciever for component %v already exists", name)
	}
	if _, ok := h.collectors[name]; ok {
		return nil, fmt.Errorf("component %v already has stateless collector", name)
	}

	h.stateful[name] = &ComponentHealth{}

	return h.stateful[name], nil
}

func (h *Healthcare) RemoveStateful(name string) {
	delete(h.stateful, name)
}

func (h *Healthcare) collectStateful() Health {
	health := make(Health)
	for name, component := range h.stateful {
		health[name] = *component
	}

	return health
}
