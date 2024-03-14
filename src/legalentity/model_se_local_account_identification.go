/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the SELocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SELocalAccountIdentification{}

// SELocalAccountIdentification struct for SELocalAccountIdentification
type SELocalAccountIdentification struct {
	// The 7- to 10-digit bank account number ([Bankkontonummer](https://sv.wikipedia.org/wiki/Bankkonto)), without the clearing number, separators, or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The 4- to 5-digit clearing number ([Clearingnummer](https://sv.wikipedia.org/wiki/Clearingnummer)), without separators or whitespace.
	ClearingNumber string `json:"clearingNumber"`
	// Business accounts with a `formFactor` value of **physical** are business accounts issued under the central bank of that country. The default value is **physical** for NL, US, and UK business accounts.   Adyen creates a local IBAN for business accounts when the `formFactor` value is set to **virtual**. The local IBANs that are supported are for DE and FR, which reference a physical NL account, with funds being routed through the central bank of NL.
	FormFactor common.NullableString `json:"formFactor,omitempty"`
	// **seLocal**
	Type string `json:"type"`
}

// NewSELocalAccountIdentification instantiates a new SELocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSELocalAccountIdentification(accountNumber string, clearingNumber string, type_ string) *SELocalAccountIdentification {
	this := SELocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.ClearingNumber = clearingNumber
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	this.Type = type_
	return &this
}

// NewSELocalAccountIdentificationWithDefaults instantiates a new SELocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSELocalAccountIdentificationWithDefaults() *SELocalAccountIdentification {
	this := SELocalAccountIdentification{}
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	var type_ string = "seLocal"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *SELocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *SELocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *SELocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetClearingNumber returns the ClearingNumber field value
func (o *SELocalAccountIdentification) GetClearingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClearingNumber
}

// GetClearingNumberOk returns a tuple with the ClearingNumber field value
// and a boolean to check if the value has been set.
func (o *SELocalAccountIdentification) GetClearingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClearingNumber, true
}

// SetClearingNumber sets field value
func (o *SELocalAccountIdentification) SetClearingNumber(v string) {
	o.ClearingNumber = v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SELocalAccountIdentification) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor.Get()) {
		var ret string
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SELocalAccountIdentification) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *SELocalAccountIdentification) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullableString and assigns it to the FormFactor field.
func (o *SELocalAccountIdentification) SetFormFactor(v string) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *SELocalAccountIdentification) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *SELocalAccountIdentification) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetType returns the Type field value
func (o *SELocalAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SELocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SELocalAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o SELocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SELocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["clearingNumber"] = o.ClearingNumber
	if o.FormFactor.IsSet() {
		toSerialize["formFactor"] = o.FormFactor.Get()
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSELocalAccountIdentification struct {
	value *SELocalAccountIdentification
	isSet bool
}

func (v NullableSELocalAccountIdentification) Get() *SELocalAccountIdentification {
	return v.value
}

func (v *NullableSELocalAccountIdentification) Set(val *SELocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableSELocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableSELocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSELocalAccountIdentification(val *SELocalAccountIdentification) *NullableSELocalAccountIdentification {
	return &NullableSELocalAccountIdentification{value: val, isSet: true}
}

func (v NullableSELocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSELocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SELocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"seLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
