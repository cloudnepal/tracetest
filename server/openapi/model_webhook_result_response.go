/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type WebhookResultResponse struct {
	StatusCode int32 `json:"statusCode,omitempty"`

	Status string `json:"status,omitempty"`

	Body string `json:"body,omitempty"`

	Headers []HttpHeader `json:"headers,omitempty"`

	Error string `json:"error,omitempty"`
}

// AssertWebhookResultResponseRequired checks if the required fields are not zero-ed
func AssertWebhookResultResponseRequired(obj WebhookResultResponse) error {
	for _, el := range obj.Headers {
		if err := AssertHttpHeaderRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseWebhookResultResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of WebhookResultResponse (e.g. [][]WebhookResultResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseWebhookResultResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aWebhookResultResponse, ok := obj.(WebhookResultResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertWebhookResultResponseRequired(aWebhookResultResponse)
	})
}
