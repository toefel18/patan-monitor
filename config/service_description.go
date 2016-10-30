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
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type DisplayStats struct {
	Counters  []string `json:"counters,omitempty"`
	Samples   []string `json:"samples,omitempty"`
	Durations []string `json:"durations,omitempty"`
}

type Service struct {
	Id            string `json: "id"`
	DisplayName   string `json: "displayName"`
	StatisticsURL string `json: "statisticsUrl"`
	ModelVersion  string `json: "modelVersion"`
	DisplayStats
}

type Connection struct {
	From string
	To   string
	DisplayStats
}

type Configuration struct {
	Version     string       `json:"version"`
	Services    []Service    `json:"services"`
	Connections []Connection `json:"connections"`
}

func LoadFile(filename string) *Configuration {
	byteContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Load(byteContent)
}

func Load(byteContent []byte) *Configuration {
	config := &Configuration{}
	err := json.Unmarshal(byteContent, config)
	if err != nil {
		panic(err)
	}
	return config
}

func Print(configuration *Configuration) {
	fmt.Println(configuration.Version)
	for _, service := range configuration.Services {
		fmt.Println("Service", service.Id, service.DisplayName, service.StatisticsURL, service.ModelVersion)
		printDisplayStats(&service.DisplayStats)
	}
	for _, connection := range configuration.Connections {
		fmt.Println("Connection", connection.From, "->", connection.To)
		printDisplayStats(&connection.DisplayStats)
	}
}

func printDisplayStats(displayStats *DisplayStats) {
	printKeysOfInterest("counters", displayStats.Counters)
	printKeysOfInterest("samples", displayStats.Samples)
	printKeysOfInterest("durations", displayStats.Durations)
}

func printKeysOfInterest(groupName string, keys []string) {
	if len(keys) > 0 {
		fmt.Println("  -", groupName, keys)
	}
}