package plugin

import (
	"fmt"
	"time"
)

type Executor interface {
	// Get Key-Value object, with key as action to execute .
	// and Value as key-Value from value-question for the paramters.
	// 	example:
	// 	CallAmbulance(string hospitalName, int quantity).
	// 	GetActionMetadata() --> map[string]map[string]string{
	// 		"CallAmbulance": map[string]sttring{
	// 			"hospitalName" : "what is hospital name ?",
	// 			"quantity" : "how much ambulance do you need ?",
	// 		}
	// 	}
	GetActionMetadata() map[string]map[string]string

	GetMetadata() Metadata

	Execute(act string, params map[string]string) <-chan string
}

type Metadata struct {
	Title    string
	Subtitle string
}

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
