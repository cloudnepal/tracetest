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

// checks if the InviteEnvironment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteEnvironment{}

// InviteEnvironment struct for InviteEnvironment
type InviteEnvironment struct {
	Id          string       `json:"id"`
	Role        Role         `json:"role"`
	Environment *Environment `json:"environment,omitempty"`
}

// NewInviteEnvironment instantiates a new InviteEnvironment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteEnvironment(id string, role Role) *InviteEnvironment {
	this := InviteEnvironment{}
	this.Id = id
	this.Role = role
	return &this
}

// NewInviteEnvironmentWithDefaults instantiates a new InviteEnvironment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteEnvironmentWithDefaults() *InviteEnvironment {
	this := InviteEnvironment{}
	return &this
}

// GetId returns the Id field value
func (o *InviteEnvironment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InviteEnvironment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InviteEnvironment) SetId(v string) {
	o.Id = v
}

// GetRole returns the Role field value
func (o *InviteEnvironment) GetRole() Role {
	if o == nil {
		var ret Role
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InviteEnvironment) GetRoleOk() (*Role, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InviteEnvironment) SetRole(v Role) {
	o.Role = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *InviteEnvironment) GetEnvironment() Environment {
	if o == nil || isNil(o.Environment) {
		var ret Environment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteEnvironment) GetEnvironmentOk() (*Environment, bool) {
	if o == nil || isNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *InviteEnvironment) HasEnvironment() bool {
	if o != nil && !isNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given Environment and assigns it to the Environment field.
func (o *InviteEnvironment) SetEnvironment(v Environment) {
	o.Environment = &v
}

func (o InviteEnvironment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteEnvironment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["role"] = o.Role
	if !isNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableInviteEnvironment struct {
	value *InviteEnvironment
	isSet bool
}

func (v NullableInviteEnvironment) Get() *InviteEnvironment {
	return v.value
}

func (v *NullableInviteEnvironment) Set(val *InviteEnvironment) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteEnvironment(val *InviteEnvironment) *NullableInviteEnvironment {
	return &NullableInviteEnvironment{value: val, isSet: true}
}

func (v NullableInviteEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
