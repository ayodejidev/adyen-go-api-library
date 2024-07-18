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

// checks if the MerchantAcquirerPair type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantAcquirerPair{}

// MerchantAcquirerPair struct for MerchantAcquirerPair
type MerchantAcquirerPair struct {
	// The acquirer ID.
	AcquirerId *string `json:"acquirerId,omitempty"`
	// The merchant identification number (MID).
	MerchantId *string `json:"merchantId,omitempty"`
}

// NewMerchantAcquirerPair instantiates a new MerchantAcquirerPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAcquirerPair() *MerchantAcquirerPair {
	this := MerchantAcquirerPair{}
	return &this
}

// NewMerchantAcquirerPairWithDefaults instantiates a new MerchantAcquirerPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAcquirerPairWithDefaults() *MerchantAcquirerPair {
	this := MerchantAcquirerPair{}
	return &this
}

// GetAcquirerId returns the AcquirerId field value if set, zero value otherwise.
func (o *MerchantAcquirerPair) GetAcquirerId() string {
	if o == nil || common.IsNil(o.AcquirerId) {
		var ret string
		return ret
	}
	return *o.AcquirerId
}

// GetAcquirerIdOk returns a tuple with the AcquirerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAcquirerPair) GetAcquirerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AcquirerId) {
		return nil, false
	}
	return o.AcquirerId, true
}

// HasAcquirerId returns a boolean if a field has been set.
func (o *MerchantAcquirerPair) HasAcquirerId() bool {
	if o != nil && !common.IsNil(o.AcquirerId) {
		return true
	}

	return false
}

// SetAcquirerId gets a reference to the given string and assigns it to the AcquirerId field.
func (o *MerchantAcquirerPair) SetAcquirerId(v string) {
	o.AcquirerId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *MerchantAcquirerPair) GetMerchantId() string {
	if o == nil || common.IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAcquirerPair) GetMerchantIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *MerchantAcquirerPair) HasMerchantId() bool {
	if o != nil && !common.IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *MerchantAcquirerPair) SetMerchantId(v string) {
	o.MerchantId = &v
}

func (o MerchantAcquirerPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantAcquirerPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AcquirerId) {
		toSerialize["acquirerId"] = o.AcquirerId
	}
	if !common.IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	return toSerialize, nil
}

type NullableMerchantAcquirerPair struct {
	value *MerchantAcquirerPair
	isSet bool
}

func (v NullableMerchantAcquirerPair) Get() *MerchantAcquirerPair {
	return v.value
}

func (v *NullableMerchantAcquirerPair) Set(val *MerchantAcquirerPair) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAcquirerPair) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAcquirerPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAcquirerPair(val *MerchantAcquirerPair) *NullableMerchantAcquirerPair {
	return &NullableMerchantAcquirerPair{value: val, isSet: true}
}

func (v NullableMerchantAcquirerPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAcquirerPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
