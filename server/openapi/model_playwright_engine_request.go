/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlaywrightEngineRequest struct {
	Target string `json:"target,omitempty"`

	Method string `json:"method,omitempty"`

	Script string `json:"script,omitempty"`
}

// AssertPlaywrightEngineRequestRequired checks if the required fields are not zero-ed
func AssertPlaywrightEngineRequestRequired(obj PlaywrightEngineRequest) error {
	return nil
}

// AssertRecursePlaywrightEngineRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlaywrightEngineRequest (e.g. [][]PlaywrightEngineRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlaywrightEngineRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlaywrightEngineRequest, ok := obj.(PlaywrightEngineRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlaywrightEngineRequestRequired(aPlaywrightEngineRequest)
	})
}
