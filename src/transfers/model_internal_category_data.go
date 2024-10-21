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

// checks if the InternalCategoryData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InternalCategoryData{}

// InternalCategoryData struct for InternalCategoryData
type InternalCategoryData struct {
	// The capture's merchant reference included in the transfer.
	ModificationMerchantReference *string `json:"modificationMerchantReference,omitempty"`
	// The capture reference included in the transfer.
	ModificationPspReference *string `json:"modificationPspReference,omitempty"`
	// **internal**
	Type *string `json:"type,omitempty"`
}

// NewInternalCategoryData instantiates a new InternalCategoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalCategoryData() *InternalCategoryData {
	this := InternalCategoryData{}
	var type_ string = "internal"
	this.Type = &type_
	return &this
}

// NewInternalCategoryDataWithDefaults instantiates a new InternalCategoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalCategoryDataWithDefaults() *InternalCategoryData {
	this := InternalCategoryData{}
	var type_ string = "internal"
	this.Type = &type_
	return &this
}

// GetModificationMerchantReference returns the ModificationMerchantReference field value if set, zero value otherwise.
func (o *InternalCategoryData) GetModificationMerchantReference() string {
	if o == nil || common.IsNil(o.ModificationMerchantReference) {
		var ret string
		return ret
	}
	return *o.ModificationMerchantReference
}

// GetModificationMerchantReferenceOk returns a tuple with the ModificationMerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalCategoryData) GetModificationMerchantReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ModificationMerchantReference) {
		return nil, false
	}
	return o.ModificationMerchantReference, true
}

// HasModificationMerchantReference returns a boolean if a field has been set.
func (o *InternalCategoryData) HasModificationMerchantReference() bool {
	if o != nil && !common.IsNil(o.ModificationMerchantReference) {
		return true
	}

	return false
}

// SetModificationMerchantReference gets a reference to the given string and assigns it to the ModificationMerchantReference field.
func (o *InternalCategoryData) SetModificationMerchantReference(v string) {
	o.ModificationMerchantReference = &v
}

// GetModificationPspReference returns the ModificationPspReference field value if set, zero value otherwise.
func (o *InternalCategoryData) GetModificationPspReference() string {
	if o == nil || common.IsNil(o.ModificationPspReference) {
		var ret string
		return ret
	}
	return *o.ModificationPspReference
}

// GetModificationPspReferenceOk returns a tuple with the ModificationPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalCategoryData) GetModificationPspReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ModificationPspReference) {
		return nil, false
	}
	return o.ModificationPspReference, true
}

// HasModificationPspReference returns a boolean if a field has been set.
func (o *InternalCategoryData) HasModificationPspReference() bool {
	if o != nil && !common.IsNil(o.ModificationPspReference) {
		return true
	}

	return false
}

// SetModificationPspReference gets a reference to the given string and assigns it to the ModificationPspReference field.
func (o *InternalCategoryData) SetModificationPspReference(v string) {
	o.ModificationPspReference = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InternalCategoryData) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalCategoryData) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InternalCategoryData) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InternalCategoryData) SetType(v string) {
	o.Type = &v
}

func (o InternalCategoryData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalCategoryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ModificationMerchantReference) {
		toSerialize["modificationMerchantReference"] = o.ModificationMerchantReference
	}
	if !common.IsNil(o.ModificationPspReference) {
		toSerialize["modificationPspReference"] = o.ModificationPspReference
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInternalCategoryData struct {
	value *InternalCategoryData
	isSet bool
}

func (v NullableInternalCategoryData) Get() *InternalCategoryData {
	return v.value
}

func (v *NullableInternalCategoryData) Set(val *InternalCategoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalCategoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalCategoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalCategoryData(val *InternalCategoryData) *NullableInternalCategoryData {
	return &NullableInternalCategoryData{value: val, isSet: true}
}

func (v NullableInternalCategoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalCategoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *InternalCategoryData) isValidType() bool {
	var allowedEnumValues = []string{"internal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
