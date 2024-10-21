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

// checks if the Contact type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Contact{}

// Contact struct for Contact
type Contact struct {
	// The individual's email address.
	Email *string `json:"email,omitempty"`
	// The individual's first name.
	FirstName *string `json:"firstName,omitempty"`
	// The infix in the individual's name, if any.
	Infix *string `json:"infix,omitempty"`
	// The individual's last name.
	LastName *string `json:"lastName,omitempty"`
	// The individual's phone number, specified as 10-14 digits with an optional `+` prefix.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// NewContact instantiates a new Contact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContact() *Contact {
	this := Contact{}
	return &this
}

// NewContactWithDefaults instantiates a new Contact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactWithDefaults() *Contact {
	this := Contact{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Contact) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Contact) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Contact) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Contact) GetFirstName() string {
	if o == nil || common.IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Contact) HasFirstName() bool {
	if o != nil && !common.IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Contact) SetFirstName(v string) {
	o.FirstName = &v
}

// GetInfix returns the Infix field value if set, zero value otherwise.
func (o *Contact) GetInfix() string {
	if o == nil || common.IsNil(o.Infix) {
		var ret string
		return ret
	}
	return *o.Infix
}

// GetInfixOk returns a tuple with the Infix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetInfixOk() (*string, bool) {
	if o == nil || common.IsNil(o.Infix) {
		return nil, false
	}
	return o.Infix, true
}

// HasInfix returns a boolean if a field has been set.
func (o *Contact) HasInfix() bool {
	if o != nil && !common.IsNil(o.Infix) {
		return true
	}

	return false
}

// SetInfix gets a reference to the given string and assigns it to the Infix field.
func (o *Contact) SetInfix(v string) {
	o.Infix = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Contact) GetLastName() string {
	if o == nil || common.IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Contact) HasLastName() bool {
	if o != nil && !common.IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Contact) SetLastName(v string) {
	o.LastName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Contact) GetPhoneNumber() string {
	if o == nil || common.IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contact) GetPhoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Contact) HasPhoneNumber() bool {
	if o != nil && !common.IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Contact) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o Contact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Contact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !common.IsNil(o.Infix) {
		toSerialize["infix"] = o.Infix
	}
	if !common.IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !common.IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableContact struct {
	value *Contact
	isSet bool
}

func (v NullableContact) Get() *Contact {
	return v.value
}

func (v *NullableContact) Set(val *Contact) {
	v.value = val
	v.isSet = true
}

func (v NullableContact) IsSet() bool {
	return v.isSet
}

func (v *NullableContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContact(val *Contact) *NullableContact {
	return &NullableContact{value: val, isSet: true}
}

func (v NullableContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
