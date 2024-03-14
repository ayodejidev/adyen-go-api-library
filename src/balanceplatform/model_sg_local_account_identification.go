/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the SGLocalAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SGLocalAccountIdentification{}

// SGLocalAccountIdentification struct for SGLocalAccountIdentification
type SGLocalAccountIdentification struct {
	// The 4- to 19-digit bank account number, without separators or whitespace.
	AccountNumber string `json:"accountNumber"`
	// The bank's 8- or 11-character BIC or SWIFT code.
	Bic string `json:"bic"`
	// The form factor of the account.  Possible values: **physical**, **virtual**. Default value: **physical**.
	FormFactor common.NullableString `json:"formFactor,omitempty"`
	// **sgLocal**
	Type *string `json:"type,omitempty"`
}

// NewSGLocalAccountIdentification instantiates a new SGLocalAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSGLocalAccountIdentification(accountNumber string, bic string) *SGLocalAccountIdentification {
	this := SGLocalAccountIdentification{}
	this.AccountNumber = accountNumber
	this.Bic = bic
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	var type_ string = "sgLocal"
	this.Type = &type_
	return &this
}

// NewSGLocalAccountIdentificationWithDefaults instantiates a new SGLocalAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSGLocalAccountIdentificationWithDefaults() *SGLocalAccountIdentification {
	this := SGLocalAccountIdentification{}
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	var type_ string = "sgLocal"
	this.Type = &type_
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *SGLocalAccountIdentification) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *SGLocalAccountIdentification) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *SGLocalAccountIdentification) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetBic returns the Bic field value
func (o *SGLocalAccountIdentification) GetBic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bic
}

// GetBicOk returns a tuple with the Bic field value
// and a boolean to check if the value has been set.
func (o *SGLocalAccountIdentification) GetBicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bic, true
}

// SetBic sets field value
func (o *SGLocalAccountIdentification) SetBic(v string) {
	o.Bic = v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SGLocalAccountIdentification) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor.Get()) {
		var ret string
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SGLocalAccountIdentification) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *SGLocalAccountIdentification) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullableString and assigns it to the FormFactor field.
func (o *SGLocalAccountIdentification) SetFormFactor(v string) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *SGLocalAccountIdentification) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *SGLocalAccountIdentification) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SGLocalAccountIdentification) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SGLocalAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SGLocalAccountIdentification) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SGLocalAccountIdentification) SetType(v string) {
	o.Type = &v
}

func (o SGLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SGLocalAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountNumber"] = o.AccountNumber
	toSerialize["bic"] = o.Bic
	if o.FormFactor.IsSet() {
		toSerialize["formFactor"] = o.FormFactor.Get()
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSGLocalAccountIdentification struct {
	value *SGLocalAccountIdentification
	isSet bool
}

func (v NullableSGLocalAccountIdentification) Get() *SGLocalAccountIdentification {
	return v.value
}

func (v *NullableSGLocalAccountIdentification) Set(val *SGLocalAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableSGLocalAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableSGLocalAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSGLocalAccountIdentification(val *SGLocalAccountIdentification) *NullableSGLocalAccountIdentification {
	return &NullableSGLocalAccountIdentification{value: val, isSet: true}
}

func (v NullableSGLocalAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSGLocalAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SGLocalAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"sgLocal"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
