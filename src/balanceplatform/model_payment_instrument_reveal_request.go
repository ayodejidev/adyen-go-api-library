/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the PaymentInstrumentRevealRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PaymentInstrumentRevealRequest{}

// PaymentInstrumentRevealRequest struct for PaymentInstrumentRevealRequest
type PaymentInstrumentRevealRequest struct {
	// The symmetric session key that you encrypted with the [public key](https://docs.adyen.com/api-explorer/balanceplatform/2/get/publicKey) that you received from Adyen.
	EncryptedKey string `json:"encryptedKey"`
	// The unique identifier of the payment instrument, which is the card for which you are managing the PIN.
	PaymentInstrumentId string `json:"paymentInstrumentId"`
}

// NewPaymentInstrumentRevealRequest instantiates a new PaymentInstrumentRevealRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInstrumentRevealRequest(encryptedKey string, paymentInstrumentId string) *PaymentInstrumentRevealRequest {
	this := PaymentInstrumentRevealRequest{}
	this.EncryptedKey = encryptedKey
	this.PaymentInstrumentId = paymentInstrumentId
	return &this
}

// NewPaymentInstrumentRevealRequestWithDefaults instantiates a new PaymentInstrumentRevealRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInstrumentRevealRequestWithDefaults() *PaymentInstrumentRevealRequest {
	this := PaymentInstrumentRevealRequest{}
	return &this
}

// GetEncryptedKey returns the EncryptedKey field value
func (o *PaymentInstrumentRevealRequest) GetEncryptedKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptedKey
}

// GetEncryptedKeyOk returns a tuple with the EncryptedKey field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRevealRequest) GetEncryptedKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptedKey, true
}

// SetEncryptedKey sets field value
func (o *PaymentInstrumentRevealRequest) SetEncryptedKey(v string) {
	o.EncryptedKey = v
}

// GetPaymentInstrumentId returns the PaymentInstrumentId field value
func (o *PaymentInstrumentRevealRequest) GetPaymentInstrumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentInstrumentId
}

// GetPaymentInstrumentIdOk returns a tuple with the PaymentInstrumentId field value
// and a boolean to check if the value has been set.
func (o *PaymentInstrumentRevealRequest) GetPaymentInstrumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentInstrumentId, true
}

// SetPaymentInstrumentId sets field value
func (o *PaymentInstrumentRevealRequest) SetPaymentInstrumentId(v string) {
	o.PaymentInstrumentId = v
}

func (o PaymentInstrumentRevealRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentInstrumentRevealRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encryptedKey"] = o.EncryptedKey
	toSerialize["paymentInstrumentId"] = o.PaymentInstrumentId
	return toSerialize, nil
}

type NullablePaymentInstrumentRevealRequest struct {
	value *PaymentInstrumentRevealRequest
	isSet bool
}

func (v NullablePaymentInstrumentRevealRequest) Get() *PaymentInstrumentRevealRequest {
	return v.value
}

func (v *NullablePaymentInstrumentRevealRequest) Set(val *PaymentInstrumentRevealRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInstrumentRevealRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInstrumentRevealRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInstrumentRevealRequest(val *PaymentInstrumentRevealRequest) *NullablePaymentInstrumentRevealRequest {
	return &NullablePaymentInstrumentRevealRequest{value: val, isSet: true}
}

func (v NullablePaymentInstrumentRevealRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInstrumentRevealRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
