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

// checks if the CardBrandDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardBrandDetails{}

// CardBrandDetails struct for CardBrandDetails
type CardBrandDetails struct {
	// Indicates if you support the card brand.
	Supported *bool `json:"supported,omitempty"`
	// The name of the card brand.
	Type *string `json:"type,omitempty"`
}

// NewCardBrandDetails instantiates a new CardBrandDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardBrandDetails() *CardBrandDetails {
	this := CardBrandDetails{}
	return &this
}

// NewCardBrandDetailsWithDefaults instantiates a new CardBrandDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardBrandDetailsWithDefaults() *CardBrandDetails {
	this := CardBrandDetails{}
	return &this
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *CardBrandDetails) GetSupported() bool {
	if o == nil || common.IsNil(o.Supported) {
		var ret bool
		return ret
	}
	return *o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBrandDetails) GetSupportedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Supported) {
		return nil, false
	}
	return o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *CardBrandDetails) HasSupported() bool {
	if o != nil && !common.IsNil(o.Supported) {
		return true
	}

	return false
}

// SetSupported gets a reference to the given bool and assigns it to the Supported field.
func (o *CardBrandDetails) SetSupported(v bool) {
	o.Supported = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CardBrandDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardBrandDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CardBrandDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CardBrandDetails) SetType(v string) {
	o.Type = &v
}

func (o CardBrandDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardBrandDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Supported) {
		toSerialize["supported"] = o.Supported
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCardBrandDetails struct {
	value *CardBrandDetails
	isSet bool
}

func (v NullableCardBrandDetails) Get() *CardBrandDetails {
	return v.value
}

func (v *NullableCardBrandDetails) Set(val *CardBrandDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCardBrandDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCardBrandDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardBrandDetails(val *CardBrandDetails) *NullableCardBrandDetails {
	return &NullableCardBrandDetails{value: val, isSet: true}
}

func (v NullableCardBrandDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardBrandDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
