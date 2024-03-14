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

// checks if the NumberAndBicAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NumberAndBicAccountIdentification{}

// NumberAndBicAccountIdentification struct for NumberAndBicAccountIdentification
type NumberAndBicAccountIdentification struct {
	// The bank account number, without separators or whitespace. The length and format depends on the bank or country.
	AccountNumber                string                        `json:"accountNumber"`
	AdditionalBankIdentification *AdditionalBankIdentification `json:"additionalBankIdentification,omitempty"`
	// The bank's 8- or 11-character BIC or SWIFT code.
	Bic string `json:"bic"`
	// Business accounts with a `formFactor` value of **physical** are business accounts issued under the central bank of that country. The default value is **physical** for NL, US, and UK business accounts.   Adyen creates a local IBAN for business accounts when the `formFactor` value is set to **virtual**. The local IBANs that are supported are for DE and FR, which reference a physical NL account, with funds being routed through the central bank of NL.
	FormFactor common.NullableString `json:"formFactor,omitempty"`
	// **numberAndBic**
	Type string `json:"type"`
}

// NewNumberAndBicAccountIdentification instantiates a new NumberAndBicAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNumberAndBicAccountIdentification(accountNumber string, bic string, type_ string) *NumberAndBicAccountIdentification {
	this := NumberAndBicAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Bic = bic
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	this.Type = type_
	return &this
}

// NewNumberAndBicAccountIdentificationWithDefaults instantiates a new NumberAndBicAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNumberAndBicAccountIdentificationWithDefaults() *NumberAndBicAccountIdentification {
	this := NumberAndBicAccountIdentification{}
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	var type_ string = "numberAndBic"
	this.Type = type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *NumberAndBicAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *NumberAndBicAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAdditionalBankIdentification returns the AdditionalBankIdentification field value if set, zero value otherwise.
func (o *NumberAndBicAccountIdentification) GetAdditionalBankIdentification() AdditionalBankIdentification {
	if o == nil || common.IsNil(o.AdditionalBankIdentification) {
		var ret AdditionalBankIdentification
		return ret
	}
	return *o.AdditionalBankIdentification
}

// GetAdditionalBankIdentificationOk returns a tuple with the AdditionalBankIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetAdditionalBankIdentificationOk() (*AdditionalBankIdentification, bool) {
	if o == nil || common.IsNil(o.AdditionalBankIdentification) {
		return nil, false
	}
	return o.AdditionalBankIdentification, true
}

// HasAdditionalBankIdentification returns a boolean if a field has been set.
func (o *NumberAndBicAccountIdentification) HasAdditionalBankIdentification() bool {
	if o != nil && !common.IsNil(o.AdditionalBankIdentification) {
		return true
	}

	return false
}

// SetAdditionalBankIdentification gets a reference to the given AdditionalBankIdentification and assigns it to the AdditionalBankIdentification field.
func (o *NumberAndBicAccountIdentification) SetAdditionalBankIdentification(v AdditionalBankIdentification) {
	o.AdditionalBankIdentification = &v
}

// GetBic returns the Bic field value
func (o *NumberAndBicAccountIdentification) GetBic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bic
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetBicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bic, true
}

// SetBic sets field value
func (o *NumberAndBicAccountIdentification) SetBic(v string) {
	o.Bic = v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NumberAndBicAccountIdentification) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor.Get()) {
		var ret string
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NumberAndBicAccountIdentification) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *NumberAndBicAccountIdentification) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullableString and assigns it to the FormFactor field.
func (o *NumberAndBicAccountIdentification) SetFormFactor(v string) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *NumberAndBicAccountIdentification) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *NumberAndBicAccountIdentification) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetType returns the Type field value
func (o *NumberAndBicAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NumberAndBicAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NumberAndBicAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o NumberAndBicAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NumberAndBicAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	if !common.IsNil(o.AdditionalBankIdentification) {
		toSerialize["additionalBankIdentification"] = o.AdditionalBankIdentification
	}
	toSerialize["bic"] = o.Bic
	if o.FormFactor.IsSet() {
		toSerialize["formFactor"] = o.FormFactor.Get()
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNumberAndBicAccountIdentification struct {
	value *NumberAndBicAccountIdentification
	isSet bool
}

func (v NullableNumberAndBicAccountIdentification) Get() *NumberAndBicAccountIdentification {
	return v.value
}

func (v *NullableNumberAndBicAccountIdentification) Set(val *NumberAndBicAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableNumberAndBicAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableNumberAndBicAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumberAndBicAccountIdentification(val *NumberAndBicAccountIdentification) *NullableNumberAndBicAccountIdentification {
	return &NullableNumberAndBicAccountIdentification{value: val, isSet: true}
}

func (v NullableNumberAndBicAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumberAndBicAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *NumberAndBicAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"numberAndBic"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
