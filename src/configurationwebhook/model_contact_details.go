/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the ContactDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ContactDetails{}

// ContactDetails struct for ContactDetails
type ContactDetails struct {
	Address Address `json:"address"`
	// The email address of the account holder.
	Email string `json:"email"`
	Phone Phone  `json:"phone"`
	// The URL of the account holder's website.
	WebAddress *string `json:"webAddress,omitempty"`
}

// NewContactDetails instantiates a new ContactDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactDetails(address Address, email string, phone Phone) *ContactDetails {
	this := ContactDetails{}
	this.Address = address
	this.Email = email
	this.Phone = phone
	return &this
}

// NewContactDetailsWithDefaults instantiates a new ContactDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactDetailsWithDefaults() *ContactDetails {
	this := ContactDetails{}
	return &this
}

// GetAddress returns the Address field value
func (o *ContactDetails) GetAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ContactDetails) SetAddress(v Address) {
	o.Address = v
}

// GetEmail returns the Email field value
func (o *ContactDetails) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ContactDetails) SetEmail(v string) {
	o.Email = v
}

// GetPhone returns the Phone field value
func (o *ContactDetails) GetPhone() Phone {
	if o == nil {
		var ret Phone
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetPhoneOk() (*Phone, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *ContactDetails) SetPhone(v Phone) {
	o.Phone = v
}

// GetWebAddress returns the WebAddress field value if set, zero value otherwise.
func (o *ContactDetails) GetWebAddress() string {
	if o == nil || common.IsNil(o.WebAddress) {
		var ret string
		return ret
	}
	return *o.WebAddress
}

// GetWebAddressOk returns a tuple with the WebAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetWebAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WebAddress) {
		return nil, false
	}
	return o.WebAddress, true
}

// HasWebAddress returns a boolean if a field has been set.
func (o *ContactDetails) HasWebAddress() bool {
	if o != nil && !common.IsNil(o.WebAddress) {
		return true
	}

	return false
}

// SetWebAddress gets a reference to the given string and assigns it to the WebAddress field.
func (o *ContactDetails) SetWebAddress(v string) {
	o.WebAddress = &v
}

func (o ContactDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["email"] = o.Email
	toSerialize["phone"] = o.Phone
	if !common.IsNil(o.WebAddress) {
		toSerialize["webAddress"] = o.WebAddress
	}
	return toSerialize, nil
}

type NullableContactDetails struct {
	value *ContactDetails
	isSet bool
}

func (v NullableContactDetails) Get() *ContactDetails {
	return v.value
}

func (v *NullableContactDetails) Set(val *ContactDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableContactDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableContactDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactDetails(val *ContactDetails) *NullableContactDetails {
	return &NullableContactDetails{value: val, isSet: true}
}

func (v NullableContactDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
