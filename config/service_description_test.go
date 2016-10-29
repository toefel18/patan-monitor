package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var INPUT_CONFIG = []byte(`
{
    "version": "0.0.1",
    "services": [
        {
            "id": "as",
            "displayName": "AddressService",
            "statisticsUrl": "http://localhost:8080/statistics",
            "modelVersion": "1",
            "counters": ["invocations"],
            "samples": ["some.sample"],
            "durations": ["duration"]
        },
        {
            "id": "flow",
            "displayName": "Flow",
            "statisticsUrl": "http://localhost:8081/statistics",
            "modelVersion": "1",
            "counters": ["invocations"],
            "samples": ["some.sample"],
            "durations": ["duration"]
        }
    ],
    "connections": [
        {
            "from": "flow",
            "to": "as",
            "durations": ["duration"]
        }
    ]
}`)

func TestLoadAndPrint (t *testing.T) {
    configuration := Load(INPUT_CONFIG)
    assert.Equal(t, "0.0.1", configuration.Version)
    assert.Equal(t, 2, len(configuration.Services))
    assert.Equal(t, 1, len(configuration.Connections))
    assert.Equal(t, "flow", configuration.Connections[0].From)
    assert.Equal(t, "as", configuration.Connections[0].To)
    Print(configuration)
}
