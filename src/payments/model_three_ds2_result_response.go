/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the ThreeDS2ResultResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDS2ResultResponse{}

// ThreeDS2ResultResponse struct for ThreeDS2ResultResponse
type ThreeDS2ResultResponse struct {
	ThreeDS2Result *ThreeDS2Result `json:"threeDS2Result,omitempty"`
}

// NewThreeDS2ResultResponse instantiates a new ThreeDS2ResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDS2ResultResponse() *ThreeDS2ResultResponse {
	this := ThreeDS2ResultResponse{}
	return &this
}

// NewThreeDS2ResultResponseWithDefaults instantiates a new ThreeDS2ResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDS2ResultResponseWithDefaults() *ThreeDS2ResultResponse {
	this := ThreeDS2ResultResponse{}
	return &this
}

// GetThreeDS2Result returns the ThreeDS2Result field value if set, zero value otherwise.
func (o *ThreeDS2ResultResponse) GetThreeDS2Result() ThreeDS2Result {
	if o == nil || common.IsNil(o.ThreeDS2Result) {
		var ret ThreeDS2Result
		return ret
	}
	return *o.ThreeDS2Result
}

// GetThreeDS2ResultOk returns a tuple with the ThreeDS2Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS2ResultResponse) GetThreeDS2ResultOk() (*ThreeDS2Result, bool) {
	if o == nil || common.IsNil(o.ThreeDS2Result) {
		return nil, false
	}
	return o.ThreeDS2Result, true
}

// HasThreeDS2Result returns a boolean if a field has been set.
func (o *ThreeDS2ResultResponse) HasThreeDS2Result() bool {
	if o != nil && !common.IsNil(o.ThreeDS2Result) {
		return true
	}

	return false
}

// SetThreeDS2Result gets a reference to the given ThreeDS2Result and assigns it to the ThreeDS2Result field.
func (o *ThreeDS2ResultResponse) SetThreeDS2Result(v ThreeDS2Result) {
	o.ThreeDS2Result = &v
}

func (o ThreeDS2ResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDS2ResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ThreeDS2Result) {
		toSerialize["threeDS2Result"] = o.ThreeDS2Result
	}
	return toSerialize, nil
}

type NullableThreeDS2ResultResponse struct {
	value *ThreeDS2ResultResponse
	isSet bool
}

func (v NullableThreeDS2ResultResponse) Get() *ThreeDS2ResultResponse {
	return v.value
}

func (v *NullableThreeDS2ResultResponse) Set(val *ThreeDS2ResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDS2ResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDS2ResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDS2ResultResponse(val *ThreeDS2ResultResponse) *NullableThreeDS2ResultResponse {
	return &NullableThreeDS2ResultResponse{value: val, isSet: true}
}

func (v NullableThreeDS2ResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDS2ResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
