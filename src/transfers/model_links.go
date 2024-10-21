/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the Links type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Links{}

// Links struct for Links
type Links struct {
	Next *Link `json:"next,omitempty"`
	Prev *Link `json:"prev,omitempty"`
}

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks() *Links {
	this := Links{}
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Links) GetNext() Link {
	if o == nil || common.IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetNextOk() (*Link, bool) {
	if o == nil || common.IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Links) HasNext() bool {
	if o != nil && !common.IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *Links) SetNext(v Link) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *Links) GetPrev() Link {
	if o == nil || common.IsNil(o.Prev) {
		var ret Link
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Links) GetPrevOk() (*Link, bool) {
	if o == nil || common.IsNil(o.Prev) {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *Links) HasPrev() bool {
	if o != nil && !common.IsNil(o.Prev) {
		return true
	}

	return false
}

// SetPrev gets a reference to the given Link and assigns it to the Prev field.
func (o *Links) SetPrev(v Link) {
	o.Prev = &v
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Links) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !common.IsNil(o.Prev) {
		toSerialize["prev"] = o.Prev
	}
	return toSerialize, nil
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
