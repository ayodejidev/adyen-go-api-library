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

// checks if the RevealPinResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RevealPinResponse{}

// RevealPinResponse struct for RevealPinResponse
type RevealPinResponse struct {
	// The encrypted [PIN block](https://www.pcisecuritystandards.org/glossary/pin-block).
	EncryptedPinBlock string `json:"encryptedPinBlock"`
	// The 16-digit token that you need to extract the `encryptedPinBlock`.
	Token string `json:"token"`
}

// NewRevealPinResponse instantiates a new RevealPinResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevealPinResponse(encryptedPinBlock string, token string) *RevealPinResponse {
	this := RevealPinResponse{}
	this.EncryptedPinBlock = encryptedPinBlock
	this.Token = token
	return &this
}

// NewRevealPinResponseWithDefaults instantiates a new RevealPinResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevealPinResponseWithDefaults() *RevealPinResponse {
	this := RevealPinResponse{}
	return &this
}

// GetEncryptedPinBlock returns the EncryptedPinBlock field value
func (o *RevealPinResponse) GetEncryptedPinBlock() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedPinBlock
}

// GetEncryptedPinBlockOk returns a tuple with the EncryptedPinBlock field value
// and a boolean to check if the value has been set.
func (o *RevealPinResponse) GetEncryptedPinBlockOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedPinBlock, true
}

// SetEncryptedPinBlock sets field value
func (o *RevealPinResponse) SetEncryptedPinBlock(v string) {
	o.EncryptedPinBlock = v
}

// GetToken returns the Token field value
func (o *RevealPinResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *RevealPinResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *RevealPinResponse) SetToken(v string) {
	o.Token = v
}

func (o RevealPinResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevealPinResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encryptedPinBlock"] = o.EncryptedPinBlock
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableRevealPinResponse struct {
	value *RevealPinResponse
	isSet bool
}

func (v NullableRevealPinResponse) Get() *RevealPinResponse {
	return v.value
}

func (v *NullableRevealPinResponse) Set(val *RevealPinResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRevealPinResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRevealPinResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevealPinResponse(val *RevealPinResponse) *NullableRevealPinResponse {
	return &NullableRevealPinResponse{value: val, isSet: true}
}

func (v NullableRevealPinResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevealPinResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
