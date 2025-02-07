package main

type ComponentHealth struct {
	Status  Status `json:"status"`
	Message string `json:"message,omitempty"`
}

type Health map[string]ComponentHealth

func (c *ComponentHealth) SetUnhealthy(message string) {
	c.Status = StatusUnhealthy
	c.Message = message
}

func (c *ComponentHealth) SetHealthy(message string) {
	c.Status = StatusHealthy
	c.Message = message
}

func (c *ComponentHealth) SetUnknown(message string) {
	c.Status = StatusUnknown
	c.Message = message
}
