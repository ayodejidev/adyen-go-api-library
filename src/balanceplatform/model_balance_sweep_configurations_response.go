/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the BalanceSweepConfigurationsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceSweepConfigurationsResponse{}

// BalanceSweepConfigurationsResponse struct for BalanceSweepConfigurationsResponse
type BalanceSweepConfigurationsResponse struct {
	// Indicates whether there are more items on the next page.
	HasNext bool `json:"hasNext"`
	// Indicates whether there are more items on the previous page.
	HasPrevious bool `json:"hasPrevious"`
	// List of sweeps associated with the balance account.
	Sweeps []SweepConfigurationV2 `json:"sweeps"`
}

// NewBalanceSweepConfigurationsResponse instantiates a new BalanceSweepConfigurationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceSweepConfigurationsResponse(hasNext bool, hasPrevious bool, sweeps []SweepConfigurationV2) *BalanceSweepConfigurationsResponse {
	this := BalanceSweepConfigurationsResponse{}
	this.HasNext = hasNext
	this.HasPrevious = hasPrevious
	this.Sweeps = sweeps
	return &this
}

// NewBalanceSweepConfigurationsResponseWithDefaults instantiates a new BalanceSweepConfigurationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceSweepConfigurationsResponseWithDefaults() *BalanceSweepConfigurationsResponse {
	this := BalanceSweepConfigurationsResponse{}
	return &this
}

// GetHasNext returns the HasNext field value
func (o *BalanceSweepConfigurationsResponse) GetHasNext() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value
// and a boolean to check if the value has been set.
func (o *BalanceSweepConfigurationsResponse) GetHasNextOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasNext, true
}

// SetHasNext sets field value
func (o *BalanceSweepConfigurationsResponse) SetHasNext(v bool) {
	o.HasNext = v
}

// GetHasPrevious returns the HasPrevious field value
func (o *BalanceSweepConfigurationsResponse) GetHasPrevious() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPrevious
}

// GetHasPreviousOk returns a tuple with the HasPrevious field value
// and a boolean to check if the value has been set.
func (o *BalanceSweepConfigurationsResponse) GetHasPreviousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPrevious, true
}

// SetHasPrevious sets field value
func (o *BalanceSweepConfigurationsResponse) SetHasPrevious(v bool) {
	o.HasPrevious = v
}

// GetSweeps returns the Sweeps field value
func (o *BalanceSweepConfigurationsResponse) GetSweeps() []SweepConfigurationV2 {
	if o == nil {
		var ret []SweepConfigurationV2
		return ret
	}

	return o.Sweeps
}

// GetSweepsOk returns a tuple with the Sweeps field value
// and a boolean to check if the value has been set.
func (o *BalanceSweepConfigurationsResponse) GetSweepsOk() ([]SweepConfigurationV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sweeps, true
}

// SetSweeps sets field value
func (o *BalanceSweepConfigurationsResponse) SetSweeps(v []SweepConfigurationV2) {
	o.Sweeps = v
}

func (o BalanceSweepConfigurationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceSweepConfigurationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hasNext"] = o.HasNext
	toSerialize["hasPrevious"] = o.HasPrevious
	toSerialize["sweeps"] = o.Sweeps
	return toSerialize, nil
}

type NullableBalanceSweepConfigurationsResponse struct {
	value *BalanceSweepConfigurationsResponse
	isSet bool
}

func (v NullableBalanceSweepConfigurationsResponse) Get() *BalanceSweepConfigurationsResponse {
	return v.value
}

func (v *NullableBalanceSweepConfigurationsResponse) Set(val *BalanceSweepConfigurationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceSweepConfigurationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceSweepConfigurationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceSweepConfigurationsResponse(val *BalanceSweepConfigurationsResponse) *NullableBalanceSweepConfigurationsResponse {
	return &NullableBalanceSweepConfigurationsResponse{value: val, isSet: true}
}

func (v NullableBalanceSweepConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceSweepConfigurationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
