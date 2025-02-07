package main

type Status string

const StatusHealthy Status = "HEALTHY"
const StatusUnhealthy Status = "UNHEALTHY"
const StatusUnknown Status = "UNKNOWN"

func (s Status) String() string {
	return string(s)
}
