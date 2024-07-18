/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the PayoutSettingsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PayoutSettingsResponse{}

// PayoutSettingsResponse struct for PayoutSettingsResponse
type PayoutSettingsResponse struct {
	// The list of payout accounts.
	Data []PayoutSettings `json:"data,omitempty"`
}

// NewPayoutSettingsResponse instantiates a new PayoutSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayoutSettingsResponse() *PayoutSettingsResponse {
	this := PayoutSettingsResponse{}
	return &this
}

// NewPayoutSettingsResponseWithDefaults instantiates a new PayoutSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayoutSettingsResponseWithDefaults() *PayoutSettingsResponse {
	this := PayoutSettingsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PayoutSettingsResponse) GetData() []PayoutSettings {
	if o == nil || common.IsNil(o.Data) {
		var ret []PayoutSettings
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayoutSettingsResponse) GetDataOk() ([]PayoutSettings, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PayoutSettingsResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PayoutSettings and assigns it to the Data field.
func (o *PayoutSettingsResponse) SetData(v []PayoutSettings) {
	o.Data = v
}

func (o PayoutSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PayoutSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePayoutSettingsResponse struct {
	value *PayoutSettingsResponse
	isSet bool
}

func (v NullablePayoutSettingsResponse) Get() *PayoutSettingsResponse {
	return v.value
}

func (v *NullablePayoutSettingsResponse) Set(val *PayoutSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePayoutSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePayoutSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayoutSettingsResponse(val *PayoutSettingsResponse) *NullablePayoutSettingsResponse {
	return &NullablePayoutSettingsResponse{value: val, isSet: true}
}

func (v NullablePayoutSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayoutSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
