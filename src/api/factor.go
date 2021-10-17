package api

import "github.com/kenesparta/multiplyLogic"

// Factor Main struct to get the data from the user
type Factor struct {
	First  *float64 `json:"first_factor"`
	Second *float64 `json:"second_factor"`
}

// Product Shows the result from the imported repository ("github.com/kenesparta/multiplyLogic")
func (f *Factor) Product() float64 {
	return multiplyLogic.Multiply(*f.First, *f.Second)
}

// AreValidNumbers Verifies if the two values are set in the JSON request
// If one of these values is not set, the result is false
func (f *Factor) AreValidNumbers() bool {
	return f.First != nil && f.Second != nil
}

// ProductResponse is used to handle the result response
type ProductResponse struct {
	Result float64 `json:"result"`
}
