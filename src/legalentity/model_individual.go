/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the Individual type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Individual{}

// Individual struct for Individual
type Individual struct {
	BirthData *BirthData `json:"birthData,omitempty"`
	// The email address of the legal entity.
	Email              *string             `json:"email,omitempty"`
	IdentificationData *IdentificationData `json:"identificationData,omitempty"`
	Name               Name                `json:"name"`
	// The individual's nationality.
	Nationality        *string      `json:"nationality,omitempty"`
	Phone              *PhoneNumber `json:"phone,omitempty"`
	ResidentialAddress Address      `json:"residentialAddress"`
	// The tax information of the individual.
	TaxInformation []TaxInformation `json:"taxInformation,omitempty"`
	WebData        *WebData         `json:"webData,omitempty"`
}

// NewIndividual instantiates a new Individual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividual(name Name, residentialAddress Address) *Individual {
	this := Individual{}
	this.Name = name
	this.ResidentialAddress = residentialAddress
	return &this
}

// NewIndividualWithDefaults instantiates a new Individual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualWithDefaults() *Individual {
	this := Individual{}
	return &this
}

// GetBirthData returns the BirthData field value if set, zero value otherwise.
func (o *Individual) GetBirthData() BirthData {
	if o == nil || common.IsNil(o.BirthData) {
		var ret BirthData
		return ret
	}
	return *o.BirthData
}

// GetBirthDataOk returns a tuple with the BirthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetBirthDataOk() (*BirthData, bool) {
	if o == nil || common.IsNil(o.BirthData) {
		return nil, false
	}
	return o.BirthData, true
}

// HasBirthData returns a boolean if a field has been set.
func (o *Individual) HasBirthData() bool {
	if o != nil && !common.IsNil(o.BirthData) {
		return true
	}

	return false
}

// SetBirthData gets a reference to the given BirthData and assigns it to the BirthData field.
func (o *Individual) SetBirthData(v BirthData) {
	o.BirthData = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Individual) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Individual) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Individual) SetEmail(v string) {
	o.Email = &v
}

// GetIdentificationData returns the IdentificationData field value if set, zero value otherwise.
func (o *Individual) GetIdentificationData() IdentificationData {
	if o == nil || common.IsNil(o.IdentificationData) {
		var ret IdentificationData
		return ret
	}
	return *o.IdentificationData
}

// GetIdentificationDataOk returns a tuple with the IdentificationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetIdentificationDataOk() (*IdentificationData, bool) {
	if o == nil || common.IsNil(o.IdentificationData) {
		return nil, false
	}
	return o.IdentificationData, true
}

// HasIdentificationData returns a boolean if a field has been set.
func (o *Individual) HasIdentificationData() bool {
	if o != nil && !common.IsNil(o.IdentificationData) {
		return true
	}

	return false
}

// SetIdentificationData gets a reference to the given IdentificationData and assigns it to the IdentificationData field.
func (o *Individual) SetIdentificationData(v IdentificationData) {
	o.IdentificationData = &v
}

// GetName returns the Name field value
func (o *Individual) GetName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Individual) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Individual) SetName(v Name) {
	o.Name = v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *Individual) GetNationality() string {
	if o == nil || common.IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetNationalityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *Individual) HasNationality() bool {
	if o != nil && !common.IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *Individual) SetNationality(v string) {
	o.Nationality = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Individual) GetPhone() PhoneNumber {
	if o == nil || common.IsNil(o.Phone) {
		var ret PhoneNumber
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetPhoneOk() (*PhoneNumber, bool) {
	if o == nil || common.IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Individual) HasPhone() bool {
	if o != nil && !common.IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given PhoneNumber and assigns it to the Phone field.
func (o *Individual) SetPhone(v PhoneNumber) {
	o.Phone = &v
}

// GetResidentialAddress returns the ResidentialAddress field value
func (o *Individual) GetResidentialAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ResidentialAddress
}

// GetResidentialAddressOk returns a tuple with the ResidentialAddress field value
// and a boolean to check if the value has been set.
func (o *Individual) GetResidentialAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResidentialAddress, true
}

// SetResidentialAddress sets field value
func (o *Individual) SetResidentialAddress(v Address) {
	o.ResidentialAddress = v
}

// GetTaxInformation returns the TaxInformation field value if set, zero value otherwise.
func (o *Individual) GetTaxInformation() []TaxInformation {
	if o == nil || common.IsNil(o.TaxInformation) {
		var ret []TaxInformation
		return ret
	}
	return o.TaxInformation
}

// GetTaxInformationOk returns a tuple with the TaxInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetTaxInformationOk() ([]TaxInformation, bool) {
	if o == nil || common.IsNil(o.TaxInformation) {
		return nil, false
	}
	return o.TaxInformation, true
}

// HasTaxInformation returns a boolean if a field has been set.
func (o *Individual) HasTaxInformation() bool {
	if o != nil && !common.IsNil(o.TaxInformation) {
		return true
	}

	return false
}

// SetTaxInformation gets a reference to the given []TaxInformation and assigns it to the TaxInformation field.
func (o *Individual) SetTaxInformation(v []TaxInformation) {
	o.TaxInformation = v
}

// GetWebData returns the WebData field value if set, zero value otherwise.
func (o *Individual) GetWebData() WebData {
	if o == nil || common.IsNil(o.WebData) {
		var ret WebData
		return ret
	}
	return *o.WebData
}

// GetWebDataOk returns a tuple with the WebData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Individual) GetWebDataOk() (*WebData, bool) {
	if o == nil || common.IsNil(o.WebData) {
		return nil, false
	}
	return o.WebData, true
}

// HasWebData returns a boolean if a field has been set.
func (o *Individual) HasWebData() bool {
	if o != nil && !common.IsNil(o.WebData) {
		return true
	}

	return false
}

// SetWebData gets a reference to the given WebData and assigns it to the WebData field.
func (o *Individual) SetWebData(v WebData) {
	o.WebData = &v
}

func (o Individual) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Individual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BirthData) {
		toSerialize["birthData"] = o.BirthData
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.IdentificationData) {
		toSerialize["identificationData"] = o.IdentificationData
	}
	toSerialize["name"] = o.Name
	if !common.IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !common.IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	toSerialize["residentialAddress"] = o.ResidentialAddress
	if !common.IsNil(o.TaxInformation) {
		toSerialize["taxInformation"] = o.TaxInformation
	}
	if !common.IsNil(o.WebData) {
		toSerialize["webData"] = o.WebData
	}
	return toSerialize, nil
}

type NullableIndividual struct {
	value *Individual
	isSet bool
}

func (v NullableIndividual) Get() *Individual {
	return v.value
}

func (v *NullableIndividual) Set(val *Individual) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividual) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividual(val *Individual) *NullableIndividual {
	return &NullableIndividual{value: val, isSet: true}
}

func (v NullableIndividual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
