/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the DetailsRequestAuthenticationData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DetailsRequestAuthenticationData{}

// DetailsRequestAuthenticationData struct for DetailsRequestAuthenticationData
type DetailsRequestAuthenticationData struct {
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation. Default: *false**.
	AuthenticationOnly *bool `json:"authenticationOnly,omitempty"`
}

// NewDetailsRequestAuthenticationData instantiates a new DetailsRequestAuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailsRequestAuthenticationData() *DetailsRequestAuthenticationData {
	this := DetailsRequestAuthenticationData{}
	var authenticationOnly bool = false
	this.AuthenticationOnly = &authenticationOnly
	return &this
}

// NewDetailsRequestAuthenticationDataWithDefaults instantiates a new DetailsRequestAuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailsRequestAuthenticationDataWithDefaults() *DetailsRequestAuthenticationData {
	this := DetailsRequestAuthenticationData{}
	var authenticationOnly bool = false
	this.AuthenticationOnly = &authenticationOnly
	return &this
}

// GetAuthenticationOnly returns the AuthenticationOnly field value if set, zero value otherwise.
func (o *DetailsRequestAuthenticationData) GetAuthenticationOnly() bool {
	if o == nil || common.IsNil(o.AuthenticationOnly) {
		var ret bool
		return ret
	}
	return *o.AuthenticationOnly
}

// GetAuthenticationOnlyOk returns a tuple with the AuthenticationOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailsRequestAuthenticationData) GetAuthenticationOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AuthenticationOnly) {
		return nil, false
	}
	return o.AuthenticationOnly, true
}

// HasAuthenticationOnly returns a boolean if a field has been set.
func (o *DetailsRequestAuthenticationData) HasAuthenticationOnly() bool {
	if o != nil && !common.IsNil(o.AuthenticationOnly) {
		return true
	}

	return false
}

// SetAuthenticationOnly gets a reference to the given bool and assigns it to the AuthenticationOnly field.
func (o *DetailsRequestAuthenticationData) SetAuthenticationOnly(v bool) {
	o.AuthenticationOnly = &v
}

func (o DetailsRequestAuthenticationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailsRequestAuthenticationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AuthenticationOnly) {
		toSerialize["authenticationOnly"] = o.AuthenticationOnly
	}
	return toSerialize, nil
}

type NullableDetailsRequestAuthenticationData struct {
	value *DetailsRequestAuthenticationData
	isSet bool
}

func (v NullableDetailsRequestAuthenticationData) Get() *DetailsRequestAuthenticationData {
	return v.value
}

func (v *NullableDetailsRequestAuthenticationData) Set(val *DetailsRequestAuthenticationData) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailsRequestAuthenticationData) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailsRequestAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailsRequestAuthenticationData(val *DetailsRequestAuthenticationData) *NullableDetailsRequestAuthenticationData {
	return &NullableDetailsRequestAuthenticationData{value: val, isSet: true}
}

func (v NullableDetailsRequestAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailsRequestAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
