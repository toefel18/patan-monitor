package main

import (
	"github.com/toefel18/patan-monitor/config"
)

type kixcode string

func main() {
	inputConfiguration := config.LoadFile("input-example.json")
	config.Print(inputConfiguration);
}
