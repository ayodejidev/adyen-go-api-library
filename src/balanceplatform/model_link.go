/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the Link type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Link{}

// Link struct for Link
type Link struct {
	First    *Href `json:"first,omitempty"`
	Last     *Href `json:"last,omitempty"`
	Next     *Href `json:"next,omitempty"`
	Previous *Href `json:"previous,omitempty"`
	Self     *Href `json:"self,omitempty"`
}

// NewLink instantiates a new Link object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLink() *Link {
	this := Link{}
	return &this
}

// NewLinkWithDefaults instantiates a new Link object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkWithDefaults() *Link {
	this := Link{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *Link) GetFirst() Href {
	if o == nil || common.IsNil(o.First) {
		var ret Href
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetFirstOk() (*Href, bool) {
	if o == nil || common.IsNil(o.First) {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *Link) HasFirst() bool {
	if o != nil && !common.IsNil(o.First) {
		return true
	}

	return false
}

// SetFirst gets a reference to the given Href and assigns it to the First field.
func (o *Link) SetFirst(v Href) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *Link) GetLast() Href {
	if o == nil || common.IsNil(o.Last) {
		var ret Href
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetLastOk() (*Href, bool) {
	if o == nil || common.IsNil(o.Last) {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *Link) HasLast() bool {
	if o != nil && !common.IsNil(o.Last) {
		return true
	}

	return false
}

// SetLast gets a reference to the given Href and assigns it to the Last field.
func (o *Link) SetLast(v Href) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *Link) GetNext() Href {
	if o == nil || common.IsNil(o.Next) {
		var ret Href
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetNextOk() (*Href, bool) {
	if o == nil || common.IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *Link) HasNext() bool {
	if o != nil && !common.IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Href and assigns it to the Next field.
func (o *Link) SetNext(v Href) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *Link) GetPrevious() Href {
	if o == nil || common.IsNil(o.Previous) {
		var ret Href
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetPreviousOk() (*Href, bool) {
	if o == nil || common.IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *Link) HasPrevious() bool {
	if o != nil && !common.IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given Href and assigns it to the Previous field.
func (o *Link) SetPrevious(v Href) {
	o.Previous = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Link) GetSelf() Href {
	if o == nil || common.IsNil(o.Self) {
		var ret Href
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Link) GetSelfOk() (*Href, bool) {
	if o == nil || common.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Link) HasSelf() bool {
	if o != nil && !common.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Href and assigns it to the Self field.
func (o *Link) SetSelf(v Href) {
	o.Self = &v
}

func (o Link) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Link) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.First) {
		toSerialize["first"] = o.First
	}
	if !common.IsNil(o.Last) {
		toSerialize["last"] = o.Last
	}
	if !common.IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !common.IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	if !common.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableLink struct {
	value *Link
	isSet bool
}

func (v NullableLink) Get() *Link {
	return v.value
}

func (v *NullableLink) Set(val *Link) {
	v.value = val
	v.isSet = true
}

func (v NullableLink) IsSet() bool {
	return v.isSet
}

func (v *NullableLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLink(val *Link) *NullableLink {
	return &NullableLink{value: val, isSet: true}
}

func (v NullableLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
