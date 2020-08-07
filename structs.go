package main

//StoreEntityBody is the body of store entity request
type StoreEntityBody struct {
	Latitude  float64
	Longitude float64
}

//InvalidEntityStoreInput are will contain errors when checking API inputs
type InvalidEntityStoreInput struct {
	Errors string
}

//InvalidInputError are will contain errors when checking API inputs
type InvalidInputError struct {
	Errors string
}
