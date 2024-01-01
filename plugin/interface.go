package plugin

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

	Execute(act string, params map[string]string) <-chan string
}

type Metadata struct {
	Title    string
	Subtitle string
}
