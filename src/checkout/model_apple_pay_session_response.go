/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the ApplePaySessionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplePaySessionResponse{}

// ApplePaySessionResponse struct for ApplePaySessionResponse
type ApplePaySessionResponse struct {
	// Base64 encoded data you need to [complete the Apple Pay merchant validation](https://docs.adyen.com/payment-methods/apple-pay/api-only?tab=adyen-certificate-validation_1#complete-apple-pay-session-validation).
	Data string `json:"data"`
}

// NewApplePaySessionResponse instantiates a new ApplePaySessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplePaySessionResponse(data string) *ApplePaySessionResponse {
	this := ApplePaySessionResponse{}
	this.Data = data
	return &this
}

// NewApplePaySessionResponseWithDefaults instantiates a new ApplePaySessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplePaySessionResponseWithDefaults() *ApplePaySessionResponse {
	this := ApplePaySessionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ApplePaySessionResponse) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ApplePaySessionResponse) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ApplePaySessionResponse) SetData(v string) {
	o.Data = v
}

func (o ApplePaySessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplePaySessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableApplePaySessionResponse struct {
	value *ApplePaySessionResponse
	isSet bool
}

func (v NullableApplePaySessionResponse) Get() *ApplePaySessionResponse {
	return v.value
}

func (v *NullableApplePaySessionResponse) Set(val *ApplePaySessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplePaySessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplePaySessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplePaySessionResponse(val *ApplePaySessionResponse) *NullableApplePaySessionResponse {
	return &NullableApplePaySessionResponse{value: val, isSet: true}
}

func (v NullableApplePaySessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplePaySessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
