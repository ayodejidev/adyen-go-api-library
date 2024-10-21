/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the NotificationUrl type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NotificationUrl{}

// NotificationUrl struct for NotificationUrl
type NotificationUrl struct {
	// One or more local URLs to send notifications to when using Terminal API.
	LocalUrls []Url `json:"localUrls,omitempty"`
	// One or more public URLs to send notifications to when using Terminal API.
	PublicUrls []Url `json:"publicUrls,omitempty"`
}

// NewNotificationUrl instantiates a new NotificationUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationUrl() *NotificationUrl {
	this := NotificationUrl{}
	return &this
}

// NewNotificationUrlWithDefaults instantiates a new NotificationUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationUrlWithDefaults() *NotificationUrl {
	this := NotificationUrl{}
	return &this
}

// GetLocalUrls returns the LocalUrls field value if set, zero value otherwise.
func (o *NotificationUrl) GetLocalUrls() []Url {
	if o == nil || common.IsNil(o.LocalUrls) {
		var ret []Url
		return ret
	}
	return o.LocalUrls
}

// GetLocalUrlsOk returns a tuple with the LocalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationUrl) GetLocalUrlsOk() ([]Url, bool) {
	if o == nil || common.IsNil(o.LocalUrls) {
		return nil, false
	}
	return o.LocalUrls, true
}

// HasLocalUrls returns a boolean if a field has been set.
func (o *NotificationUrl) HasLocalUrls() bool {
	if o != nil && !common.IsNil(o.LocalUrls) {
		return true
	}

	return false
}

// SetLocalUrls gets a reference to the given []Url and assigns it to the LocalUrls field.
func (o *NotificationUrl) SetLocalUrls(v []Url) {
	o.LocalUrls = v
}

// GetPublicUrls returns the PublicUrls field value if set, zero value otherwise.
func (o *NotificationUrl) GetPublicUrls() []Url {
	if o == nil || common.IsNil(o.PublicUrls) {
		var ret []Url
		return ret
	}
	return o.PublicUrls
}

// GetPublicUrlsOk returns a tuple with the PublicUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationUrl) GetPublicUrlsOk() ([]Url, bool) {
	if o == nil || common.IsNil(o.PublicUrls) {
		return nil, false
	}
	return o.PublicUrls, true
}

// HasPublicUrls returns a boolean if a field has been set.
func (o *NotificationUrl) HasPublicUrls() bool {
	if o != nil && !common.IsNil(o.PublicUrls) {
		return true
	}

	return false
}

// SetPublicUrls gets a reference to the given []Url and assigns it to the PublicUrls field.
func (o *NotificationUrl) SetPublicUrls(v []Url) {
	o.PublicUrls = v
}

func (o NotificationUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.LocalUrls) {
		toSerialize["localUrls"] = o.LocalUrls
	}
	if !common.IsNil(o.PublicUrls) {
		toSerialize["publicUrls"] = o.PublicUrls
	}
	return toSerialize, nil
}

type NullableNotificationUrl struct {
	value *NotificationUrl
	isSet bool
}

func (v NullableNotificationUrl) Get() *NotificationUrl {
	return v.value
}

func (v *NullableNotificationUrl) Set(val *NotificationUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationUrl(val *NotificationUrl) *NullableNotificationUrl {
	return &NullableNotificationUrl{value: val, isSet: true}
}

func (v NullableNotificationUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
