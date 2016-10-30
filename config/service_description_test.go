/*
 *
 *     Copyright 2016 Christophe Hesters
 *
 *     Licensed under the Apache License, Version 2.0 (the "License");
 *     you may not use this file except in compliance with the License.
 *     You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 *     Unless required by applicable law or agreed to in writing, software
 *     distributed under the License is distributed on an "AS IS" BASIS,
 *     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *     See the License for the specific language governing permissions and
 *     limitations under the License.
 *
 */
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
