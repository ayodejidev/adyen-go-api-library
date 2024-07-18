/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the CartesBancairesInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CartesBancairesInfo{}

// CartesBancairesInfo struct for CartesBancairesInfo
type CartesBancairesInfo struct {
	// Cartes Bancaires SIRET. Format: 14 digits.
	Siret                  string                      `json:"siret"`
	TransactionDescription *TransactionDescriptionInfo `json:"transactionDescription,omitempty"`
}

// NewCartesBancairesInfo instantiates a new CartesBancairesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCartesBancairesInfo(siret string) *CartesBancairesInfo {
	this := CartesBancairesInfo{}
	this.Siret = siret
	return &this
}

// NewCartesBancairesInfoWithDefaults instantiates a new CartesBancairesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCartesBancairesInfoWithDefaults() *CartesBancairesInfo {
	this := CartesBancairesInfo{}
	return &this
}

// GetSiret returns the Siret field value
func (o *CartesBancairesInfo) GetSiret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Siret
}

// GetSiretOk returns a tuple with the Siret field value
// and a boolean to check if the value has been set.
func (o *CartesBancairesInfo) GetSiretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Siret, true
}

// SetSiret sets field value
func (o *CartesBancairesInfo) SetSiret(v string) {
	o.Siret = v
}

// GetTransactionDescription returns the TransactionDescription field value if set, zero value otherwise.
func (o *CartesBancairesInfo) GetTransactionDescription() TransactionDescriptionInfo {
	if o == nil || common.IsNil(o.TransactionDescription) {
		var ret TransactionDescriptionInfo
		return ret
	}
	return *o.TransactionDescription
}

// GetTransactionDescriptionOk returns a tuple with the TransactionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CartesBancairesInfo) GetTransactionDescriptionOk() (*TransactionDescriptionInfo, bool) {
	if o == nil || common.IsNil(o.TransactionDescription) {
		return nil, false
	}
	return o.TransactionDescription, true
}

// HasTransactionDescription returns a boolean if a field has been set.
func (o *CartesBancairesInfo) HasTransactionDescription() bool {
	if o != nil && !common.IsNil(o.TransactionDescription) {
		return true
	}

	return false
}

// SetTransactionDescription gets a reference to the given TransactionDescriptionInfo and assigns it to the TransactionDescription field.
func (o *CartesBancairesInfo) SetTransactionDescription(v TransactionDescriptionInfo) {
	o.TransactionDescription = &v
}

func (o CartesBancairesInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CartesBancairesInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["siret"] = o.Siret
	if !common.IsNil(o.TransactionDescription) {
		toSerialize["transactionDescription"] = o.TransactionDescription
	}
	return toSerialize, nil
}

type NullableCartesBancairesInfo struct {
	value *CartesBancairesInfo
	isSet bool
}

func (v NullableCartesBancairesInfo) Get() *CartesBancairesInfo {
	return v.value
}

func (v *NullableCartesBancairesInfo) Set(val *CartesBancairesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCartesBancairesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCartesBancairesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCartesBancairesInfo(val *CartesBancairesInfo) *NullableCartesBancairesInfo {
	return &NullableCartesBancairesInfo{value: val, isSet: true}
}

func (v NullableCartesBancairesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCartesBancairesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
