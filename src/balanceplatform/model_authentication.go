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

// checks if the Authentication type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Authentication{}

// Authentication struct for Authentication
type Authentication struct {
	// The email address where the one-time password (OTP) is sent.
	Email *string `json:"email,omitempty"`
	// The password used for 3D Secure password-based authentication. The value must be between 1 to 30 characters and must only contain the following supported characters.  * Characters between **a-z**, **A-Z**, and **0-9**  * Special characters: **äöüßÄÖÜ+-*_/ç%()=?!~#'\",;:$&àùòâôûáúó**
	Password *string `json:"password,omitempty"`
	Phone    *Phone  `json:"phone,omitempty"`
}

// NewAuthentication instantiates a new Authentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthentication() *Authentication {
	this := Authentication{}
	return &this
}

// NewAuthenticationWithDefaults instantiates a new Authentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationWithDefaults() *Authentication {
	this := Authentication{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Authentication) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Authentication) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Authentication) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Authentication) GetPassword() string {
	if o == nil || common.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetPasswordOk() (*string, bool) {
	if o == nil || common.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Authentication) HasPassword() bool {
	if o != nil && !common.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Authentication) SetPassword(v string) {
	o.Password = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Authentication) GetPhone() Phone {
	if o == nil || common.IsNil(o.Phone) {
		var ret Phone
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authentication) GetPhoneOk() (*Phone, bool) {
	if o == nil || common.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Authentication) HasPhone() bool {
	if o != nil && !common.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given Phone and assigns it to the Phone field.
func (o *Authentication) SetPhone(v Phone) {
	o.Phone = &v
}

func (o Authentication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !common.IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	return toSerialize, nil
}

type NullableAuthentication struct {
	value *Authentication
	isSet bool
}

func (v NullableAuthentication) Get() *Authentication {
	return v.value
}

func (v *NullableAuthentication) Set(val *Authentication) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthentication(val *Authentication) *NullableAuthentication {
	return &NullableAuthentication{value: val, isSet: true}
}

func (v NullableAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
