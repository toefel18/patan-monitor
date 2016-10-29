package main

import (
	"github.com/toefel18/patan-monitor/config"
)

type kixcode string

func main() {
	inputConfiguration := config.Load("/home/hestersco/projects/golang/src/github.com/toefel18/labyrinth-visualizer/input-example.json")
	config.Print(inputConfiguration);
}
