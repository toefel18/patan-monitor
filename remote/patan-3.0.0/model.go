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
package patan_3_0_0

import (
    "github.com/toefel18/patan-monitor/remote/common"
)

type Snapshot struct {
    TimestampCreated int64                   `json:"timestampTaken"`
    Durations        map[string]Distribution `json:"durations"`
    Counters         map[string]int64        `json:"counters"`
    Samples          map[string]Distribution `json:"samples"`
}

func (s *Snapshot) Timestamp() int64 {
    return s.TimestampCreated
}

func (s *Snapshot) FindDuration(key string) (common.Distribution, bool) {
    item, found := s.Durations[key]
    return &item, found
}

func (s *Snapshot) FindCounter(key string) (int64, bool) {
    ctr, found := s.Counters[key]
    return ctr, found
}

func (s *Snapshot) FindSample(key string) (common.Distribution, bool) {
    item, found := s.Samples[key]
    return &item, found
}

type Distribution struct {
    Samples      int64   `json:"sampleCount"`
    Minimum      float64 `json:"minimum"`
    Maximum      float64 `json:"maximum"`
    Avg          float64 `json:"mean"`
    StdDeviation float64 `json:"stdDeviation"`
}

func (d *Distribution) SampleCount() int64 {
    return d.Samples
}

func (d *Distribution) Min() float64 {
    return d.Minimum
}

func (d *Distribution) Max() float64 {
    return d.Maximum
}

func (d *Distribution) Mean() float64 {
    return d.Avg
}

func (d *Distribution) StdDev() float64 {
    return d.StdDeviation
}
