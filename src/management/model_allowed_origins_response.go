/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the AllowedOriginsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllowedOriginsResponse{}

// AllowedOriginsResponse struct for AllowedOriginsResponse
type AllowedOriginsResponse struct {
	// List of allowed origins.
	Data []AllowedOrigin `json:"data,omitempty"`
}

// NewAllowedOriginsResponse instantiates a new AllowedOriginsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedOriginsResponse() *AllowedOriginsResponse {
	this := AllowedOriginsResponse{}
	return &this
}

// NewAllowedOriginsResponseWithDefaults instantiates a new AllowedOriginsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedOriginsResponseWithDefaults() *AllowedOriginsResponse {
	this := AllowedOriginsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AllowedOriginsResponse) GetData() []AllowedOrigin {
	if o == nil || common.IsNil(o.Data) {
		var ret []AllowedOrigin
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedOriginsResponse) GetDataOk() ([]AllowedOrigin, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AllowedOriginsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AllowedOrigin and assigns it to the Data field.
func (o *AllowedOriginsResponse) SetData(v []AllowedOrigin) {
	o.Data = v
}

func (o AllowedOriginsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowedOriginsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAllowedOriginsResponse struct {
	value *AllowedOriginsResponse
	isSet bool
}

func (v NullableAllowedOriginsResponse) Get() *AllowedOriginsResponse {
	return v.value
}

func (v *NullableAllowedOriginsResponse) Set(val *AllowedOriginsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedOriginsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedOriginsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedOriginsResponse(val *AllowedOriginsResponse) *NullableAllowedOriginsResponse {
	return &NullableAllowedOriginsResponse{value: val, isSet: true}
}

func (v NullableAllowedOriginsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedOriginsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
