/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the SplitConfigurationList type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SplitConfigurationList{}

// SplitConfigurationList struct for SplitConfigurationList
type SplitConfigurationList struct {
	// List of split configurations applied to the stores under the merchant account.
	Data []SplitConfiguration `json:"data,omitempty"`
}

// NewSplitConfigurationList instantiates a new SplitConfigurationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSplitConfigurationList() *SplitConfigurationList {
	this := SplitConfigurationList{}
	return &this
}

// NewSplitConfigurationListWithDefaults instantiates a new SplitConfigurationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSplitConfigurationListWithDefaults() *SplitConfigurationList {
	this := SplitConfigurationList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SplitConfigurationList) GetData() []SplitConfiguration {
	if o == nil || common.IsNil(o.Data) {
		var ret []SplitConfiguration
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SplitConfigurationList) GetDataOk() ([]SplitConfiguration, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SplitConfigurationList) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SplitConfiguration and assigns it to the Data field.
func (o *SplitConfigurationList) SetData(v []SplitConfiguration) {
	o.Data = v
}

func (o SplitConfigurationList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SplitConfigurationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSplitConfigurationList struct {
	value *SplitConfigurationList
	isSet bool
}

func (v NullableSplitConfigurationList) Get() *SplitConfigurationList {
	return v.value
}

func (v *NullableSplitConfigurationList) Set(val *SplitConfigurationList) {
	v.value = val
	v.isSet = true
}

func (v NullableSplitConfigurationList) IsSet() bool {
	return v.isSet
}

func (v *NullableSplitConfigurationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSplitConfigurationList(val *SplitConfigurationList) *NullableSplitConfigurationList {
	return &NullableSplitConfigurationList{value: val, isSet: true}
}

func (v NullableSplitConfigurationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSplitConfigurationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
