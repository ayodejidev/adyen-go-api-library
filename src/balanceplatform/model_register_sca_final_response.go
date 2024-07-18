/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the RegisterSCAFinalResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RegisterSCAFinalResponse{}

// RegisterSCAFinalResponse struct for RegisterSCAFinalResponse
type RegisterSCAFinalResponse struct {
	// Specifies if the registration was initiated successfully.
	Success *bool `json:"success,omitempty"`
}

// NewRegisterSCAFinalResponse instantiates a new RegisterSCAFinalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSCAFinalResponse() *RegisterSCAFinalResponse {
	this := RegisterSCAFinalResponse{}
	return &this
}

// NewRegisterSCAFinalResponseWithDefaults instantiates a new RegisterSCAFinalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSCAFinalResponseWithDefaults() *RegisterSCAFinalResponse {
	this := RegisterSCAFinalResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RegisterSCAFinalResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterSCAFinalResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RegisterSCAFinalResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RegisterSCAFinalResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o RegisterSCAFinalResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterSCAFinalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableRegisterSCAFinalResponse struct {
	value *RegisterSCAFinalResponse
	isSet bool
}

func (v NullableRegisterSCAFinalResponse) Get() *RegisterSCAFinalResponse {
	return v.value
}

func (v *NullableRegisterSCAFinalResponse) Set(val *RegisterSCAFinalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSCAFinalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSCAFinalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSCAFinalResponse(val *RegisterSCAFinalResponse) *NullableRegisterSCAFinalResponse {
	return &NullableRegisterSCAFinalResponse{value: val, isSet: true}
}

func (v NullableRegisterSCAFinalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSCAFinalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
