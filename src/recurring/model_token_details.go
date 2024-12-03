/*
Adyen Recurring API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package recurring

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the TokenDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenDetails{}

// TokenDetails struct for TokenDetails
type TokenDetails struct {
	TokenData     *map[string]string `json:"tokenData,omitempty"`
	TokenDataType *string            `json:"tokenDataType,omitempty"`
}

// NewTokenDetails instantiates a new TokenDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenDetails() *TokenDetails {
	this := TokenDetails{}
	return &this
}

// NewTokenDetailsWithDefaults instantiates a new TokenDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDetailsWithDefaults() *TokenDetails {
	this := TokenDetails{}
	return &this
}

// GetTokenData returns the TokenData field value if set, zero value otherwise.
func (o *TokenDetails) GetTokenData() map[string]string {
	if o == nil || common.IsNil(o.TokenData) {
		var ret map[string]string
		return ret
	}
	return *o.TokenData
}

// GetTokenDataOk returns a tuple with the TokenData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetTokenDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.TokenData) {
		return nil, false
	}
	return o.TokenData, true
}

// HasTokenData returns a boolean if a field has been set.
func (o *TokenDetails) HasTokenData() bool {
	if o != nil && !common.IsNil(o.TokenData) {
		return true
	}

	return false
}

// SetTokenData gets a reference to the given map[string]string and assigns it to the TokenData field.
func (o *TokenDetails) SetTokenData(v map[string]string) {
	o.TokenData = &v
}

// GetTokenDataType returns the TokenDataType field value if set, zero value otherwise.
func (o *TokenDetails) GetTokenDataType() string {
	if o == nil || common.IsNil(o.TokenDataType) {
		var ret string
		return ret
	}
	return *o.TokenDataType
}

// GetTokenDataTypeOk returns a tuple with the TokenDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetTokenDataTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenDataType) {
		return nil, false
	}
	return o.TokenDataType, true
}

// HasTokenDataType returns a boolean if a field has been set.
func (o *TokenDetails) HasTokenDataType() bool {
	if o != nil && !common.IsNil(o.TokenDataType) {
		return true
	}

	return false
}

// SetTokenDataType gets a reference to the given string and assigns it to the TokenDataType field.
func (o *TokenDetails) SetTokenDataType(v string) {
	o.TokenDataType = &v
}

func (o TokenDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TokenData) {
		toSerialize["tokenData"] = o.TokenData
	}
	if !common.IsNil(o.TokenDataType) {
		toSerialize["tokenDataType"] = o.TokenDataType
	}
	return toSerialize, nil
}

type NullableTokenDetails struct {
	value *TokenDetails
	isSet bool
}

func (v NullableTokenDetails) Get() *TokenDetails {
	return v.value
}

func (v *NullableTokenDetails) Set(val *TokenDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenDetails(val *TokenDetails) *NullableTokenDetails {
	return &NullableTokenDetails{value: val, isSet: true}
}

func (v NullableTokenDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
