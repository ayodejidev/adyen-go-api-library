/*
Adyen Checkout API

API version: 70
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// checks if the SubMerchant2 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubMerchant2{}

// SubMerchant2 struct for SubMerchant2
type SubMerchant2 struct {
	Address *Address `json:"address,omitempty"`
	Id      *string  `json:"id,omitempty"`
	Mcc     *string  `json:"mcc,omitempty"`
	Name    *string  `json:"name,omitempty"`
	TaxId   *string  `json:"taxId,omitempty"`
}

// NewSubMerchant2 instantiates a new SubMerchant2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubMerchant2() *SubMerchant2 {
	this := SubMerchant2{}
	return &this
}

// NewSubMerchant2WithDefaults instantiates a new SubMerchant2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubMerchant2WithDefaults() *SubMerchant2 {
	this := SubMerchant2{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *SubMerchant2) GetAddress() Address {
	if o == nil || common.IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubMerchant2) GetAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *SubMerchant2) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *SubMerchant2) SetAddress(v Address) {
	o.Address = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubMerchant2) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubMerchant2) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubMerchant2) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubMerchant2) SetId(v string) {
	o.Id = &v
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *SubMerchant2) GetMcc() string {
	if o == nil || common.IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubMerchant2) GetMccOk() (*string, bool) {
	if o == nil || common.IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *SubMerchant2) HasMcc() bool {
	if o != nil && !common.IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *SubMerchant2) SetMcc(v string) {
	o.Mcc = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubMerchant2) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubMerchant2) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubMerchant2) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubMerchant2) SetName(v string) {
	o.Name = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *SubMerchant2) GetTaxId() string {
	if o == nil || common.IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubMerchant2) GetTaxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *SubMerchant2) HasTaxId() bool {
	if o != nil && !common.IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *SubMerchant2) SetTaxId(v string) {
	o.TaxId = &v
}

func (o SubMerchant2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubMerchant2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.TaxId) {
		toSerialize["taxId"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableSubMerchant2 struct {
	value *SubMerchant2
	isSet bool
}

func (v NullableSubMerchant2) Get() *SubMerchant2 {
	return v.value
}

func (v *NullableSubMerchant2) Set(val *SubMerchant2) {
	v.value = val
	v.isSet = true
}

func (v NullableSubMerchant2) IsSet() bool {
	return v.isSet
}

func (v *NullableSubMerchant2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubMerchant2(val *SubMerchant2) *NullableSubMerchant2 {
	return &NullableSubMerchant2{value: val, isSet: true}
}

func (v NullableSubMerchant2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubMerchant2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
