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

// checks if the BankAccountModel type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BankAccountModel{}

// BankAccountModel struct for BankAccountModel
type BankAccountModel struct {
	// Business accounts with a `formFactor` value of **physical** are business accounts issued under the central bank of that country. The default value is **physical** for NL, US, and UK business accounts.   Adyen creates a local IBAN for business accounts when the `formFactor` value is set to **virtual**. The local IBANs that are supported are for DE and FR, which reference a physical NL account, with funds being routed through the central bank of NL.
	FormFactor common.NullableString `json:"formFactor,omitempty"`
}

// NewBankAccountModel instantiates a new BankAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountModel() *BankAccountModel {
	this := BankAccountModel{}
	var formFactor = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	return &this
}

// NewBankAccountModelWithDefaults instantiates a new BankAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountModelWithDefaults() *BankAccountModel {
	this := BankAccountModel{}
	var formFactor = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	return &this
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankAccountModel) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor.Get()) {
		var ret string
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankAccountModel) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *BankAccountModel) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullableString and assigns it to the FormFactor field.
func (o *BankAccountModel) SetFormFactor(v string) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *BankAccountModel) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *BankAccountModel) UnsetFormFactor() {
	o.FormFactor.Unset()
}

func (o BankAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FormFactor.IsSet() {
		toSerialize["formFactor"] = o.FormFactor.Get()
	}
	return toSerialize, nil
}

type NullableBankAccountModel struct {
	value *BankAccountModel
	isSet bool
}

func (v NullableBankAccountModel) Get() *BankAccountModel {
	return v.value
}

func (v *NullableBankAccountModel) Set(val *BankAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountModel(val *BankAccountModel) *NullableBankAccountModel {
	return &NullableBankAccountModel{value: val, isSet: true}
}

func (v NullableBankAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BankAccountModel) isValidFormFactor() bool {
	var allowedEnumValues = []string{"physical", "unknown", "virtual"}
	for _, allowed := range allowedEnumValues {
		if o.GetFormFactor() == allowed {
			return true
		}
	}
	return false
}
