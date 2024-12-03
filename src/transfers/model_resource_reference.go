/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the ResourceReference type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ResourceReference{}

// ResourceReference struct for ResourceReference
type ResourceReference struct {
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the resource.
	Id *string `json:"id,omitempty"`
	// The reference for the resource.
	Reference *string `json:"reference,omitempty"`
}

// NewResourceReference instantiates a new ResourceReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceReference() *ResourceReference {
	this := ResourceReference{}
	return &this
}

// NewResourceReferenceWithDefaults instantiates a new ResourceReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceReferenceWithDefaults() *ResourceReference {
	this := ResourceReference{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceReference) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReference) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceReference) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceReference) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceReference) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReference) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceReference) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceReference) SetId(v string) {
	o.Id = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ResourceReference) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceReference) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ResourceReference) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ResourceReference) SetReference(v string) {
	o.Reference = &v
}

func (o ResourceReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableResourceReference struct {
	value *ResourceReference
	isSet bool
}

func (v NullableResourceReference) Get() *ResourceReference {
	return v.value
}

func (v *NullableResourceReference) Set(val *ResourceReference) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceReference) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceReference(val *ResourceReference) *NullableResourceReference {
	return &NullableResourceReference{value: val, isSet: true}
}

func (v NullableResourceReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
