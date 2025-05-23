// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * auth-service API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package generated




type ResponseMessage struct {

	// the message
	Message string `json:"message,omitempty"`
}

// AssertResponseMessageRequired checks if the required fields are not zero-ed
func AssertResponseMessageRequired(obj ResponseMessage) error {
	return nil
}

// AssertResponseMessageConstraints checks if the values respects the defined constraints
func AssertResponseMessageConstraints(obj ResponseMessage) error {
	return nil
}
