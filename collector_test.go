package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthcareStateless(t *testing.T) {
	collectorHealthy := func() (ComponentHealth, error) {
		return ComponentHealth{
			Status: StatusHealthy,
		}, nil
	}
	collectorUnhealthy := func() (ComponentHealth, error) {
		return ComponentHealth{
			Status: StatusUnhealthy,
		}, nil
	}

	expected := Health{
		"sampleHealthy":   ComponentHealth{Status: StatusHealthy},
		"sampleUnhealthy": ComponentHealth{Status: StatusUnhealthy},
	}

	healthcare := New()

	err := healthcare.AddStateless("sampleHealthy", collectorHealthy)
	require.NoError(t, err)
	err = healthcare.AddStateless("sampleUnhealthy", collectorUnhealthy)
	require.NoError(t, err)

	health := healthcare.Collect()

	require.Equal(t, expected, health)
}

func TestHealthcareStateful(t *testing.T) {
	expected := Health{
		"sampleHealthy": ComponentHealth{Status: StatusHealthy},
	}

	healthcare := New()

	receiver, err := healthcare.AddStateful("sampleHealthy")
	require.NoError(t, err)
	receiver.SetHealthy("")

	health := healthcare.Collect()
	require.Equal(t, expected, health)
}
