/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the IbanAccountIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IbanAccountIdentification{}

// IbanAccountIdentification struct for IbanAccountIdentification
type IbanAccountIdentification struct {
	// The form factor of the account.  Possible values: **physical**, **virtual**. Default value: **physical**.
	FormFactor common.NullableString `json:"formFactor,omitempty"`
	// The international bank account number as defined in the [ISO-13616](https://www.iso.org/standard/81090.html) standard.
	Iban string `json:"iban"`
	// **iban**
	Type string `json:"type"`
}

// NewIbanAccountIdentification instantiates a new IbanAccountIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIbanAccountIdentification(iban string, type_ string) *IbanAccountIdentification {
	this := IbanAccountIdentification{}
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	this.Iban = iban
	this.Type = type_
	return &this
}

// NewIbanAccountIdentificationWithDefaults instantiates a new IbanAccountIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIbanAccountIdentificationWithDefaults() *IbanAccountIdentification {
	this := IbanAccountIdentification{}
	var formFactor string = "physical"
	this.FormFactor = *common.NewNullableString(&formFactor)
	var type_ string = "iban"
	this.Type = type_
	return &this
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IbanAccountIdentification) GetFormFactor() string {
	if o == nil || common.IsNil(o.FormFactor.Get()) {
		var ret string
		return ret
	}
	return *o.FormFactor.Get()
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IbanAccountIdentification) GetFormFactorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormFactor.Get(), o.FormFactor.IsSet()
}

// HasFormFactor returns a boolean if a field has been set.
func (o *IbanAccountIdentification) HasFormFactor() bool {
	if o != nil && o.FormFactor.IsSet() {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given NullableString and assigns it to the FormFactor field.
func (o *IbanAccountIdentification) SetFormFactor(v string) {
	o.FormFactor.Set(&v)
}

// SetFormFactorNil sets the value for FormFactor to be an explicit nil
func (o *IbanAccountIdentification) SetFormFactorNil() {
	o.FormFactor.Set(nil)
}

// UnsetFormFactor ensures that no value is present for FormFactor, not even an explicit nil
func (o *IbanAccountIdentification) UnsetFormFactor() {
	o.FormFactor.Unset()
}

// GetIban returns the Iban field value
func (o *IbanAccountIdentification) GetIban() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
func (o *IbanAccountIdentification) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *IbanAccountIdentification) SetIban(v string) {
	o.Iban = v
}

// GetType returns the Type field value
func (o *IbanAccountIdentification) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IbanAccountIdentification) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IbanAccountIdentification) SetType(v string) {
	o.Type = v
}

func (o IbanAccountIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IbanAccountIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FormFactor.IsSet() {
		toSerialize["formFactor"] = o.FormFactor.Get()
	}
	toSerialize["iban"] = o.Iban
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableIbanAccountIdentification struct {
	value *IbanAccountIdentification
	isSet bool
}

func (v NullableIbanAccountIdentification) Get() *IbanAccountIdentification {
	return v.value
}

func (v *NullableIbanAccountIdentification) Set(val *IbanAccountIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullableIbanAccountIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullableIbanAccountIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIbanAccountIdentification(val *IbanAccountIdentification) *NullableIbanAccountIdentification {
	return &NullableIbanAccountIdentification{value: val, isSet: true}
}

func (v NullableIbanAccountIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIbanAccountIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *IbanAccountIdentification) isValidType() bool {
	var allowedEnumValues = []string{"iban"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
