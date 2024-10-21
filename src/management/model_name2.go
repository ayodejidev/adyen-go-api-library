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

// checks if the Name2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Name2{}

// Name2 struct for Name2
type Name2 struct {
	// The first name.
	FirstName *string `json:"firstName,omitempty"`
	// The last name.
	LastName *string `json:"lastName,omitempty"`
}

// NewName2 instantiates a new Name2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewName2() *Name2 {
	this := Name2{}
	return &this
}

// NewName2WithDefaults instantiates a new Name2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewName2WithDefaults() *Name2 {
	this := Name2{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Name2) GetFirstName() string {
	if o == nil || common.IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name2) GetFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Name2) HasFirstName() bool {
	if o != nil && !common.IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Name2) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Name2) GetLastName() string {
	if o == nil || common.IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name2) GetLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Name2) HasLastName() bool {
	if o != nil && !common.IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Name2) SetLastName(v string) {
	o.LastName = &v
}

func (o Name2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Name2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !common.IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	return toSerialize, nil
}

type NullableName2 struct {
	value *Name2
	isSet bool
}

func (v NullableName2) Get() *Name2 {
	return v.value
}

func (v *NullableName2) Set(val *Name2) {
	v.value = val
	v.isSet = true
}

func (v NullableName2) IsSet() bool {
	return v.isSet
}

func (v *NullableName2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableName2(val *Name2) *NullableName2 {
	return &NullableName2{value: val, isSet: true}
}

func (v NullableName2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableName2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
