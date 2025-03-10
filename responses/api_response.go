// Package responses defines the response payload structures for the API Contact Form application.
//
// It includes the APIResponse struct for standard API responses and the ContactResponse struct
// for representing contact data in responses. Additionally, it provides helper functions
// to convert models to response formats.
//
// Author: Adit
package responses

// APIResponse represents the standard structure for API responses.
type APIResponse struct {
	// Code is a string representing the status code of the response.
	Code string `json:"code"`
	// Message provides a human-readable message about the response.
	Message string `json:"message"`
	// Data holds the payload of the response, which can be any type.
	Data interface{} `json:"data"`
}
