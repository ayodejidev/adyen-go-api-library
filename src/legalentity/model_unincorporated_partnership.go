/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the UnincorporatedPartnership type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UnincorporatedPartnership{}

// UnincorporatedPartnership struct for UnincorporatedPartnership
type UnincorporatedPartnership struct {
	// The two-character [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the governing country.
	CountryOfGoverningLaw string `json:"countryOfGoverningLaw"`
	// The date when the legal arrangement was incorporated in YYYY-MM-DD format.
	DateOfIncorporation *string `json:"dateOfIncorporation,omitempty"`
	// Short description about the Legal Arrangement.
	Description *string `json:"description,omitempty"`
	// The registered name, if different from the `name`.
	DoingBusinessAs *string `json:"doingBusinessAs,omitempty"`
	// The legal name.
	Name                     string   `json:"name"`
	PrincipalPlaceOfBusiness *Address `json:"principalPlaceOfBusiness,omitempty"`
	RegisteredAddress        Address  `json:"registeredAddress"`
	// The registration number.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	// The tax information of the entity.
	TaxInformation []TaxInformation `json:"taxInformation,omitempty"`
	// Type of Partnership.  Possible values: *  **limitedPartnership** *  **generalPartnership** *  **familyPartnership** *  **commercialPartnership** *  **publicPartnership** *  **otherPartnership** *  **gbr** *  **gmbh** *  **kgaa** *  **cv** *  **vof** *  **maatschap** *  **privateFundLimitedPartnership** *  **businessTrustEntity** *  **businessPartnership**
	Type string `json:"type"`
	// The reason for not providing a VAT number.  Possible values: **industryExemption**, **belowTaxThreshold**.
	VatAbsenceReason *string `json:"vatAbsenceReason,omitempty"`
	// The VAT number.
	VatNumber *string `json:"vatNumber,omitempty"`
}

// NewUnincorporatedPartnership instantiates a new UnincorporatedPartnership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnincorporatedPartnership(countryOfGoverningLaw string, name string, registeredAddress Address, type_ string) *UnincorporatedPartnership {
	this := UnincorporatedPartnership{}
	this.CountryOfGoverningLaw = countryOfGoverningLaw
	this.Name = name
	this.RegisteredAddress = registeredAddress
	this.Type = type_
	return &this
}

// NewUnincorporatedPartnershipWithDefaults instantiates a new UnincorporatedPartnership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnincorporatedPartnershipWithDefaults() *UnincorporatedPartnership {
	this := UnincorporatedPartnership{}
	return &this
}

// GetCountryOfGoverningLaw returns the CountryOfGoverningLaw field value
func (o *UnincorporatedPartnership) GetCountryOfGoverningLaw() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryOfGoverningLaw
}

// GetCountryOfGoverningLawOk returns a tuple with the CountryOfGoverningLaw field value
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetCountryOfGoverningLawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryOfGoverningLaw, true
}

// SetCountryOfGoverningLaw sets field value
func (o *UnincorporatedPartnership) SetCountryOfGoverningLaw(v string) {
	o.CountryOfGoverningLaw = v
}

// GetDateOfIncorporation returns the DateOfIncorporation field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetDateOfIncorporation() string {
	if o == nil || common.IsNil(o.DateOfIncorporation) {
		var ret string
		return ret
	}
	return *o.DateOfIncorporation
}

// GetDateOfIncorporationOk returns a tuple with the DateOfIncorporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetDateOfIncorporationOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfIncorporation) {
		return nil, false
	}
	return o.DateOfIncorporation, true
}

// HasDateOfIncorporation returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasDateOfIncorporation() bool {
	if o != nil && !common.IsNil(o.DateOfIncorporation) {
		return true
	}

	return false
}

// SetDateOfIncorporation gets a reference to the given string and assigns it to the DateOfIncorporation field.
func (o *UnincorporatedPartnership) SetDateOfIncorporation(v string) {
	o.DateOfIncorporation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnincorporatedPartnership) SetDescription(v string) {
	o.Description = &v
}

// GetDoingBusinessAs returns the DoingBusinessAs field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetDoingBusinessAs() string {
	if o == nil || common.IsNil(o.DoingBusinessAs) {
		var ret string
		return ret
	}
	return *o.DoingBusinessAs
}

// GetDoingBusinessAsOk returns a tuple with the DoingBusinessAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetDoingBusinessAsOk() (*string, bool) {
	if o == nil || common.IsNil(o.DoingBusinessAs) {
		return nil, false
	}
	return o.DoingBusinessAs, true
}

// HasDoingBusinessAs returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasDoingBusinessAs() bool {
	if o != nil && !common.IsNil(o.DoingBusinessAs) {
		return true
	}

	return false
}

// SetDoingBusinessAs gets a reference to the given string and assigns it to the DoingBusinessAs field.
func (o *UnincorporatedPartnership) SetDoingBusinessAs(v string) {
	o.DoingBusinessAs = &v
}

// GetName returns the Name field value
func (o *UnincorporatedPartnership) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UnincorporatedPartnership) SetName(v string) {
	o.Name = v
}

// GetPrincipalPlaceOfBusiness returns the PrincipalPlaceOfBusiness field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetPrincipalPlaceOfBusiness() Address {
	if o == nil || common.IsNil(o.PrincipalPlaceOfBusiness) {
		var ret Address
		return ret
	}
	return *o.PrincipalPlaceOfBusiness
}

// GetPrincipalPlaceOfBusinessOk returns a tuple with the PrincipalPlaceOfBusiness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetPrincipalPlaceOfBusinessOk() (*Address, bool) {
	if o == nil || common.IsNil(o.PrincipalPlaceOfBusiness) {
		return nil, false
	}
	return o.PrincipalPlaceOfBusiness, true
}

// HasPrincipalPlaceOfBusiness returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasPrincipalPlaceOfBusiness() bool {
	if o != nil && !common.IsNil(o.PrincipalPlaceOfBusiness) {
		return true
	}

	return false
}

// SetPrincipalPlaceOfBusiness gets a reference to the given Address and assigns it to the PrincipalPlaceOfBusiness field.
func (o *UnincorporatedPartnership) SetPrincipalPlaceOfBusiness(v Address) {
	o.PrincipalPlaceOfBusiness = &v
}

// GetRegisteredAddress returns the RegisteredAddress field value
func (o *UnincorporatedPartnership) GetRegisteredAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.RegisteredAddress
}

// GetRegisteredAddressOk returns a tuple with the RegisteredAddress field value
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetRegisteredAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegisteredAddress, true
}

// SetRegisteredAddress sets field value
func (o *UnincorporatedPartnership) SetRegisteredAddress(v Address) {
	o.RegisteredAddress = v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetRegistrationNumber() string {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasRegistrationNumber() bool {
	if o != nil && !common.IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *UnincorporatedPartnership) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetTaxInformation returns the TaxInformation field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetTaxInformation() []TaxInformation {
	if o == nil || common.IsNil(o.TaxInformation) {
		var ret []TaxInformation
		return ret
	}
	return o.TaxInformation
}

// GetTaxInformationOk returns a tuple with the TaxInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetTaxInformationOk() ([]TaxInformation, bool) {
	if o == nil || common.IsNil(o.TaxInformation) {
		return nil, false
	}
	return o.TaxInformation, true
}

// HasTaxInformation returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasTaxInformation() bool {
	if o != nil && !common.IsNil(o.TaxInformation) {
		return true
	}

	return false
}

// SetTaxInformation gets a reference to the given []TaxInformation and assigns it to the TaxInformation field.
func (o *UnincorporatedPartnership) SetTaxInformation(v []TaxInformation) {
	o.TaxInformation = v
}

// GetType returns the Type field value
func (o *UnincorporatedPartnership) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UnincorporatedPartnership) SetType(v string) {
	o.Type = v
}

// GetVatAbsenceReason returns the VatAbsenceReason field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetVatAbsenceReason() string {
	if o == nil || common.IsNil(o.VatAbsenceReason) {
		var ret string
		return ret
	}
	return *o.VatAbsenceReason
}

// GetVatAbsenceReasonOk returns a tuple with the VatAbsenceReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetVatAbsenceReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.VatAbsenceReason) {
		return nil, false
	}
	return o.VatAbsenceReason, true
}

// HasVatAbsenceReason returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasVatAbsenceReason() bool {
	if o != nil && !common.IsNil(o.VatAbsenceReason) {
		return true
	}

	return false
}

// SetVatAbsenceReason gets a reference to the given string and assigns it to the VatAbsenceReason field.
func (o *UnincorporatedPartnership) SetVatAbsenceReason(v string) {
	o.VatAbsenceReason = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UnincorporatedPartnership) GetVatNumber() string {
	if o == nil || common.IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnincorporatedPartnership) GetVatNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UnincorporatedPartnership) HasVatNumber() bool {
	if o != nil && !common.IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UnincorporatedPartnership) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o UnincorporatedPartnership) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnincorporatedPartnership) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryOfGoverningLaw"] = o.CountryOfGoverningLaw
	if !common.IsNil(o.DateOfIncorporation) {
		toSerialize["dateOfIncorporation"] = o.DateOfIncorporation
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.DoingBusinessAs) {
		toSerialize["doingBusinessAs"] = o.DoingBusinessAs
	}
	toSerialize["name"] = o.Name
	if !common.IsNil(o.PrincipalPlaceOfBusiness) {
		toSerialize["principalPlaceOfBusiness"] = o.PrincipalPlaceOfBusiness
	}
	toSerialize["registeredAddress"] = o.RegisteredAddress
	if !common.IsNil(o.RegistrationNumber) {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if !common.IsNil(o.TaxInformation) {
		toSerialize["taxInformation"] = o.TaxInformation
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.VatAbsenceReason) {
		toSerialize["vatAbsenceReason"] = o.VatAbsenceReason
	}
	if !common.IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	return toSerialize, nil
}

type NullableUnincorporatedPartnership struct {
	value *UnincorporatedPartnership
	isSet bool
}

func (v NullableUnincorporatedPartnership) Get() *UnincorporatedPartnership {
	return v.value
}

func (v *NullableUnincorporatedPartnership) Set(val *UnincorporatedPartnership) {
	v.value = val
	v.isSet = true
}

func (v NullableUnincorporatedPartnership) IsSet() bool {
	return v.isSet
}

func (v *NullableUnincorporatedPartnership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnincorporatedPartnership(val *UnincorporatedPartnership) *NullableUnincorporatedPartnership {
	return &NullableUnincorporatedPartnership{value: val, isSet: true}
}

func (v NullableUnincorporatedPartnership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnincorporatedPartnership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UnincorporatedPartnership) isValidType() bool {
	var allowedEnumValues = []string{"limitedPartnership", "generalPartnership", "familyPartnership", "commercialPartnership", "publicPartnership", "otherPartnership", "gbr", "gmbh", "kgaa", "cv", "vof", "maatschap", "privateFundLimitedPartnership", "businessTrustEntity", "businessPartnership"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
func (o *UnincorporatedPartnership) isValidVatAbsenceReason() bool {
	var allowedEnumValues = []string{"industryExemption", "belowTaxThreshold"}
	for _, allowed := range allowedEnumValues {
		if o.GetVatAbsenceReason() == allowed {
			return true
		}
	}
	return false
}
