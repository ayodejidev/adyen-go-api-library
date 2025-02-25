/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the WeChatPayPosInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WeChatPayPosInfo{}

// WeChatPayPosInfo struct for WeChatPayPosInfo
type WeChatPayPosInfo struct {
	// The name of the contact person from merchant support.
	ContactPersonName string `json:"contactPersonName"`
	// The email address of merchant support.
	Email string `json:"email"`
}

// NewWeChatPayPosInfo instantiates a new WeChatPayPosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeChatPayPosInfo(contactPersonName string, email string) *WeChatPayPosInfo {
	this := WeChatPayPosInfo{}
	this.ContactPersonName = contactPersonName
	this.Email = email
	return &this
}

// NewWeChatPayPosInfoWithDefaults instantiates a new WeChatPayPosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeChatPayPosInfoWithDefaults() *WeChatPayPosInfo {
	this := WeChatPayPosInfo{}
	return &this
}

// GetContactPersonName returns the ContactPersonName field value
func (o *WeChatPayPosInfo) GetContactPersonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactPersonName
}

// GetContactPersonNameOk returns a tuple with the ContactPersonName field value
// and a boolean to check if the value has been set.
func (o *WeChatPayPosInfo) GetContactPersonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactPersonName, true
}

// SetContactPersonName sets field value
func (o *WeChatPayPosInfo) SetContactPersonName(v string) {
	o.ContactPersonName = v
}

// GetEmail returns the Email field value
func (o *WeChatPayPosInfo) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *WeChatPayPosInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *WeChatPayPosInfo) SetEmail(v string) {
	o.Email = v
}

func (o WeChatPayPosInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeChatPayPosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contactPersonName"] = o.ContactPersonName
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableWeChatPayPosInfo struct {
	value *WeChatPayPosInfo
	isSet bool
}

func (v NullableWeChatPayPosInfo) Get() *WeChatPayPosInfo {
	return v.value
}

func (v *NullableWeChatPayPosInfo) Set(val *WeChatPayPosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWeChatPayPosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWeChatPayPosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeChatPayPosInfo(val *WeChatPayPosInfo) *NullableWeChatPayPosInfo {
	return &NullableWeChatPayPosInfo{value: val, isSet: true}
}

func (v NullableWeChatPayPosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeChatPayPosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
