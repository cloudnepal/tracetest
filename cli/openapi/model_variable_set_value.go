/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VariableSetValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableSetValue{}

// VariableSetValue struct for VariableSetValue
type VariableSetValue struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// NewVariableSetValue instantiates a new VariableSetValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableSetValue() *VariableSetValue {
	this := VariableSetValue{}
	var type_ string = "raw"
	this.Type = &type_
	return &this
}

// NewVariableSetValueWithDefaults instantiates a new VariableSetValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableSetValueWithDefaults() *VariableSetValue {
	this := VariableSetValue{}
	var type_ string = "raw"
	this.Type = &type_
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VariableSetValue) GetKey() string {
	if o == nil || isNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSetValue) GetKeyOk() (*string, bool) {
	if o == nil || isNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VariableSetValue) HasKey() bool {
	if o != nil && !isNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VariableSetValue) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VariableSetValue) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSetValue) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VariableSetValue) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *VariableSetValue) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VariableSetValue) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableSetValue) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VariableSetValue) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VariableSetValue) SetType(v string) {
	o.Type = &v
}

func (o VariableSetValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableSetValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVariableSetValue struct {
	value *VariableSetValue
	isSet bool
}

func (v NullableVariableSetValue) Get() *VariableSetValue {
	return v.value
}

func (v *NullableVariableSetValue) Set(val *VariableSetValue) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableSetValue) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableSetValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableSetValue(val *VariableSetValue) *NullableVariableSetValue {
	return &NullableVariableSetValue{value: val, isSet: true}
}

func (v NullableVariableSetValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableSetValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
