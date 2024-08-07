/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Summary{}

// Summary struct for Summary
type Summary struct {
	Runs        *int32     `json:"runs,omitempty"`
	LastState   *string    `json:"lastState,omitempty"`
	LastRunTime *time.Time `json:"lastRunTime,omitempty"`
}

// NewSummary instantiates a new Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummary() *Summary {
	this := Summary{}
	return &this
}

// NewSummaryWithDefaults instantiates a new Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryWithDefaults() *Summary {
	this := Summary{}
	return &this
}

// GetRuns returns the Runs field value if set, zero value otherwise.
func (o *Summary) GetRuns() int32 {
	if o == nil || isNil(o.Runs) {
		var ret int32
		return ret
	}
	return *o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetRunsOk() (*int32, bool) {
	if o == nil || isNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *Summary) HasRuns() bool {
	if o != nil && !isNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given int32 and assigns it to the Runs field.
func (o *Summary) SetRuns(v int32) {
	o.Runs = &v
}

// GetLastState returns the LastState field value if set, zero value otherwise.
func (o *Summary) GetLastState() string {
	if o == nil || isNil(o.LastState) {
		var ret string
		return ret
	}
	return *o.LastState
}

// GetLastStateOk returns a tuple with the LastState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetLastStateOk() (*string, bool) {
	if o == nil || isNil(o.LastState) {
		return nil, false
	}
	return o.LastState, true
}

// HasLastState returns a boolean if a field has been set.
func (o *Summary) HasLastState() bool {
	if o != nil && !isNil(o.LastState) {
		return true
	}

	return false
}

// SetLastState gets a reference to the given string and assigns it to the LastState field.
func (o *Summary) SetLastState(v string) {
	o.LastState = &v
}

// GetLastRunTime returns the LastRunTime field value if set, zero value otherwise.
func (o *Summary) GetLastRunTime() time.Time {
	if o == nil || isNil(o.LastRunTime) {
		var ret time.Time
		return ret
	}
	return *o.LastRunTime
}

// GetLastRunTimeOk returns a tuple with the LastRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetLastRunTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastRunTime) {
		return nil, false
	}
	return o.LastRunTime, true
}

// HasLastRunTime returns a boolean if a field has been set.
func (o *Summary) HasLastRunTime() bool {
	if o != nil && !isNil(o.LastRunTime) {
		return true
	}

	return false
}

// SetLastRunTime gets a reference to the given time.Time and assigns it to the LastRunTime field.
func (o *Summary) SetLastRunTime(v time.Time) {
	o.LastRunTime = &v
}

func (o Summary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Summary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Runs) {
		toSerialize["runs"] = o.Runs
	}
	if !isNil(o.LastState) {
		toSerialize["lastState"] = o.LastState
	}
	if !isNil(o.LastRunTime) {
		toSerialize["lastRunTime"] = o.LastRunTime
	}
	return toSerialize, nil
}

type NullableSummary struct {
	value *Summary
	isSet bool
}

func (v NullableSummary) Get() *Summary {
	return v.value
}

func (v *NullableSummary) Set(val *Summary) {
	v.value = val
	v.isSet = true
}

func (v NullableSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummary(val *Summary) *NullableSummary {
	return &NullableSummary{value: val, isSet: true}
}

func (v NullableSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
