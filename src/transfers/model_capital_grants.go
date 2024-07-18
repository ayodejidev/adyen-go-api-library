/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the CapitalGrants type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapitalGrants{}

// CapitalGrants struct for CapitalGrants
type CapitalGrants struct {
	// The unique identifier of the grant.
	Grants []CapitalGrant `json:"grants"`
}

// NewCapitalGrants instantiates a new CapitalGrants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalGrants(grants []CapitalGrant) *CapitalGrants {
	this := CapitalGrants{}
	this.Grants = grants
	return &this
}

// NewCapitalGrantsWithDefaults instantiates a new CapitalGrants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalGrantsWithDefaults() *CapitalGrants {
	this := CapitalGrants{}
	return &this
}

// GetGrants returns the Grants field value
func (o *CapitalGrants) GetGrants() []CapitalGrant {
	if o == nil {
		var ret []CapitalGrant
		return ret
	}

	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value
// and a boolean to check if the value has been set.
func (o *CapitalGrants) GetGrantsOk() ([]CapitalGrant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Grants, true
}

// SetGrants sets field value
func (o *CapitalGrants) SetGrants(v []CapitalGrant) {
	o.Grants = v
}

func (o CapitalGrants) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapitalGrants) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grants"] = o.Grants
	return toSerialize, nil
}

type NullableCapitalGrants struct {
	value *CapitalGrants
	isSet bool
}

func (v NullableCapitalGrants) Get() *CapitalGrants {
	return v.value
}

func (v *NullableCapitalGrants) Set(val *CapitalGrants) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalGrants) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalGrants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalGrants(val *CapitalGrants) *NullableCapitalGrants {
	return &NullableCapitalGrants{value: val, isSet: true}
}

func (v NullableCapitalGrants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalGrants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
