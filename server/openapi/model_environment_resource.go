/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type EnvironmentResource struct {
	Type string `json:"type,omitempty"`

	Spec Environment `json:"spec,omitempty"`
}

// AssertEnvironmentResourceRequired checks if the required fields are not zero-ed
func AssertEnvironmentResourceRequired(obj EnvironmentResource) error {
	if err := AssertEnvironmentRequired(obj.Spec); err != nil {
		return err
	}
	return nil
}

// AssertRecurseEnvironmentResourceRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EnvironmentResource (e.g. [][]EnvironmentResource), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEnvironmentResourceRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEnvironmentResource, ok := obj.(EnvironmentResource)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEnvironmentResourceRequired(aEnvironmentResource)
	})
}
