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
package main

import (
	"github.com/toefel18/patan-monitor/config"
	"github.com/toefel18/patan-monitor/remote"
	"fmt"
)

func main() {
	inputConfiguration := config.LoadFile("input-example.json")
	config.Print(inputConfiguration);

	// read from test stub with data:
	// {"timestampTaken":1477825694462,"durations":{"inbound.restcall.put.htmessages.failed":{"sampleCount":4,"minimum":0.121489,"maximum":3.227217,"mean":0.9263502499999999,"stdDeviation":1.5344550190629846}},"occurrences":{},"samples":{}}
	snapshot, err := remote.GetSnapshot("http://localhost:8080/stats")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(snapshot.Timestamp())
	duration, present := snapshot.FindDuration("inbound.restcall.put.htmessages.failed")
	if present {
		fmt.Println("found in snapshot:", duration.SampleCount())
	} else {
		fmt.Println("not in snapshot", duration.SampleCount())
	}
}
