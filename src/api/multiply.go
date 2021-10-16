package api

import "github.com/kenesparta/multiplyLogic"

// Multiply Main struct to get the data from the user
type Multiply struct {
	FirstFactor  *float64 `json:"first_factor"`
	SecondFactor *float64 `json:"second_factor"`
}

// Result Shows the result from the imported repository ("github.com/kenesparta/multiplyLogic")
func (m *Multiply) Result() float64 {
	return multiplyLogic.Multiply(*m.FirstFactor, *m.SecondFactor)
}

// AreValidNumbers Verifies if the two values are set in the JSON request
// If one of these values is not set, the result is false
func (m *Multiply) AreValidNumbers() bool {
	return m.FirstFactor != nil && m.SecondFactor != nil
}

// MultiplyResponse is used to handle the result response
type MultiplyResponse struct {
	Result float64 `json:"result"`
}
