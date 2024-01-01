package plugin

import (
	"fmt"
	"time"
)

type ExecDummy1 struct{}

func (e ExecDummy1) GetActionMetadata() map[string]map[string]string {
	return map[string]map[string]string{
		"Summarize": {
			"left":  "insert first argument",
			"right": "insert second argument",
		},
		"Multiply": {
			"Factor 1": "insert first Factor",
			"Factor 2": "insert second Factor",
		},
	}
}

func (e ExecDummy1) GetMetadata() Metadata {
	return Metadata{
		Title:    "Calculator",
		Subtitle: "Calculate 2 arguments",
	}
}

func (e ExecDummy1) Execute(act string, params map[string]string) <-chan string {
	out := make(chan string, 2)

	go func(p map[string]string) {
		time.Sleep(2 * time.Second)
		out <- fmt.Sprintf("[%s] %v", act, p)
		close(out)

	}(params)

	return out
}

type ExecDummy2 struct{}

func (e ExecDummy2) GetActionMetadata() map[string]map[string]string {
	return map[string]map[string]string{
		"Mapping": {
			"latitude":  "insert latitude",
			"longitude": "insert longitude",
		},
	}
}

func (e ExecDummy2) GetMetadata() Metadata {
	return Metadata{
		Title:    "Mapping",
		Subtitle: "Mapping Simulator",
	}
}

func (e ExecDummy2) Execute(act string, params map[string]string) <-chan string {
	out := make(chan string, 2)

	go func(p map[string]string) {
		time.Sleep(2 * time.Second)
		out <- fmt.Sprintf("[%s] %v", act, p)
		close(out)
	}(params)

	return out
}
