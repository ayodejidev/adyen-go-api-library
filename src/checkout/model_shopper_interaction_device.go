/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the ShopperInteractionDevice type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ShopperInteractionDevice{}

// ShopperInteractionDevice struct for ShopperInteractionDevice
type ShopperInteractionDevice struct {
	// Locale on the shopper interaction device.
	Locale *string `json:"locale,omitempty"`
	// Operating system running on the shopper interaction device.
	Os *string `json:"os,omitempty"`
	// Version of the operating system on the shopper interaction device.
	OsVersion *string `json:"osVersion,omitempty"`
}

// NewShopperInteractionDevice instantiates a new ShopperInteractionDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShopperInteractionDevice() *ShopperInteractionDevice {
	this := ShopperInteractionDevice{}
	return &this
}

// NewShopperInteractionDeviceWithDefaults instantiates a new ShopperInteractionDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShopperInteractionDeviceWithDefaults() *ShopperInteractionDevice {
	this := ShopperInteractionDevice{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ShopperInteractionDevice) GetLocale() string {
	if o == nil || common.IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperInteractionDevice) GetLocaleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ShopperInteractionDevice) HasLocale() bool {
	if o != nil && !common.IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ShopperInteractionDevice) SetLocale(v string) {
	o.Locale = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ShopperInteractionDevice) GetOs() string {
	if o == nil || common.IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperInteractionDevice) GetOsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ShopperInteractionDevice) HasOs() bool {
	if o != nil && !common.IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ShopperInteractionDevice) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *ShopperInteractionDevice) GetOsVersion() string {
	if o == nil || common.IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShopperInteractionDevice) GetOsVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *ShopperInteractionDevice) HasOsVersion() bool {
	if o != nil && !common.IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *ShopperInteractionDevice) SetOsVersion(v string) {
	o.OsVersion = &v
}

func (o ShopperInteractionDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShopperInteractionDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !common.IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !common.IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	return toSerialize, nil
}

type NullableShopperInteractionDevice struct {
	value *ShopperInteractionDevice
	isSet bool
}

func (v NullableShopperInteractionDevice) Get() *ShopperInteractionDevice {
	return v.value
}

func (v *NullableShopperInteractionDevice) Set(val *ShopperInteractionDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableShopperInteractionDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableShopperInteractionDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShopperInteractionDevice(val *ShopperInteractionDevice) *NullableShopperInteractionDevice {
	return &NullableShopperInteractionDevice{value: val, isSet: true}
}

func (v NullableShopperInteractionDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShopperInteractionDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
