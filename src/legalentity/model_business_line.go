/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the BusinessLine type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BusinessLine{}

// BusinessLine struct for BusinessLine
type BusinessLine struct {
	// The capability for which you are creating the business line.  Possible values: **receivePayments**, **receiveFromPlatformPayments**, **issueBankAccount**
	// Deprecated
	Capability *string `json:"capability,omitempty"`
	// The unique identifier of the business line.
	Id string `json:"id"`
	// A code that represents the industry of the legal entity for [marketplaces](https://docs.adyen.com/marketplaces/verification-requirements/reference-additional-products/#list-industry-codes) or [platforms](https://docs.adyen.com/platforms/verification-requirements/reference-additional-products/#list-industry-codes). For example, **4431A** for computer software stores.
	IndustryCode string `json:"industryCode"`
	// Unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the business line.
	LegalEntityId string `json:"legalEntityId"`
	// The verification errors related to capabilities for this supporting entity.
	Problems []CapabilityProblem `json:"problems,omitempty"`
	// A list of channels where goods or services are sold.  Possible values: **pos**, **posMoto**, **eCommerce**, **ecomMoto**, **payByLink**.  Required only in combination with the `service` **paymentProcessing**.
	SalesChannels []string `json:"salesChannels,omitempty"`
	// The service for which you are creating the business line.    Possible values: *  **paymentProcessing** *  **banking**
	Service       string         `json:"service"`
	SourceOfFunds *SourceOfFunds `json:"sourceOfFunds,omitempty"`
	// List of website URLs where your user's goods or services are sold. When this is required for a service but your user does not have an online presence, provide the reason in the `webDataExemption` object.
	WebData          []WebData         `json:"webData,omitempty"`
	WebDataExemption *WebDataExemption `json:"webDataExemption,omitempty"`
}

// NewBusinessLine instantiates a new BusinessLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessLine(id string, industryCode string, legalEntityId string, service string) *BusinessLine {
	this := BusinessLine{}
	this.Id = id
	this.IndustryCode = industryCode
	this.LegalEntityId = legalEntityId
	this.Service = service
	return &this
}

// NewBusinessLineWithDefaults instantiates a new BusinessLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessLineWithDefaults() *BusinessLine {
	this := BusinessLine{}
	return &this
}

// GetCapability returns the Capability field value if set, zero value otherwise.
// Deprecated
func (o *BusinessLine) GetCapability() string {
	if o == nil || common.IsNil(o.Capability) {
		var ret string
		return ret
	}
	return *o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BusinessLine) GetCapabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Capability) {
		return nil, false
	}
	return o.Capability, true
}

// HasCapability returns a boolean if a field has been set.
func (o *BusinessLine) HasCapability() bool {
	if o != nil && !common.IsNil(o.Capability) {
		return true
	}

	return false
}

// SetCapability gets a reference to the given string and assigns it to the Capability field.
// Deprecated
func (o *BusinessLine) SetCapability(v string) {
	o.Capability = &v
}

// GetId returns the Id field value
func (o *BusinessLine) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BusinessLine) SetId(v string) {
	o.Id = v
}

// GetIndustryCode returns the IndustryCode field value
func (o *BusinessLine) GetIndustryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndustryCode
}

// GetIndustryCodeOk returns a tuple with the IndustryCode field value
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetIndustryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndustryCode, true
}

// SetIndustryCode sets field value
func (o *BusinessLine) SetIndustryCode(v string) {
	o.IndustryCode = v
}

// GetLegalEntityId returns the LegalEntityId field value
func (o *BusinessLine) GetLegalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetLegalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegalEntityId, true
}

// SetLegalEntityId sets field value
func (o *BusinessLine) SetLegalEntityId(v string) {
	o.LegalEntityId = v
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *BusinessLine) GetProblems() []CapabilityProblem {
	if o == nil || common.IsNil(o.Problems) {
		var ret []CapabilityProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetProblemsOk() ([]CapabilityProblem, bool) {
	if o == nil || common.IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *BusinessLine) HasProblems() bool {
	if o != nil && !common.IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []CapabilityProblem and assigns it to the Problems field.
func (o *BusinessLine) SetProblems(v []CapabilityProblem) {
	o.Problems = v
}

// GetSalesChannels returns the SalesChannels field value if set, zero value otherwise.
func (o *BusinessLine) GetSalesChannels() []string {
	if o == nil || common.IsNil(o.SalesChannels) {
		var ret []string
		return ret
	}
	return o.SalesChannels
}

// GetSalesChannelsOk returns a tuple with the SalesChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetSalesChannelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SalesChannels) {
		return nil, false
	}
	return o.SalesChannels, true
}

// HasSalesChannels returns a boolean if a field has been set.
func (o *BusinessLine) HasSalesChannels() bool {
	if o != nil && !common.IsNil(o.SalesChannels) {
		return true
	}

	return false
}

// SetSalesChannels gets a reference to the given []string and assigns it to the SalesChannels field.
func (o *BusinessLine) SetSalesChannels(v []string) {
	o.SalesChannels = v
}

// GetService returns the Service field value
func (o *BusinessLine) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *BusinessLine) SetService(v string) {
	o.Service = v
}

// GetSourceOfFunds returns the SourceOfFunds field value if set, zero value otherwise.
func (o *BusinessLine) GetSourceOfFunds() SourceOfFunds {
	if o == nil || common.IsNil(o.SourceOfFunds) {
		var ret SourceOfFunds
		return ret
	}
	return *o.SourceOfFunds
}

// GetSourceOfFundsOk returns a tuple with the SourceOfFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetSourceOfFundsOk() (*SourceOfFunds, bool) {
	if o == nil || common.IsNil(o.SourceOfFunds) {
		return nil, false
	}
	return o.SourceOfFunds, true
}

// HasSourceOfFunds returns a boolean if a field has been set.
func (o *BusinessLine) HasSourceOfFunds() bool {
	if o != nil && !common.IsNil(o.SourceOfFunds) {
		return true
	}

	return false
}

// SetSourceOfFunds gets a reference to the given SourceOfFunds and assigns it to the SourceOfFunds field.
func (o *BusinessLine) SetSourceOfFunds(v SourceOfFunds) {
	o.SourceOfFunds = &v
}

// GetWebData returns the WebData field value if set, zero value otherwise.
func (o *BusinessLine) GetWebData() []WebData {
	if o == nil || common.IsNil(o.WebData) {
		var ret []WebData
		return ret
	}
	return o.WebData
}

// GetWebDataOk returns a tuple with the WebData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetWebDataOk() ([]WebData, bool) {
	if o == nil || common.IsNil(o.WebData) {
		return nil, false
	}
	return o.WebData, true
}

// HasWebData returns a boolean if a field has been set.
func (o *BusinessLine) HasWebData() bool {
	if o != nil && !common.IsNil(o.WebData) {
		return true
	}

	return false
}

// SetWebData gets a reference to the given []WebData and assigns it to the WebData field.
func (o *BusinessLine) SetWebData(v []WebData) {
	o.WebData = v
}

// GetWebDataExemption returns the WebDataExemption field value if set, zero value otherwise.
func (o *BusinessLine) GetWebDataExemption() WebDataExemption {
	if o == nil || common.IsNil(o.WebDataExemption) {
		var ret WebDataExemption
		return ret
	}
	return *o.WebDataExemption
}

// GetWebDataExemptionOk returns a tuple with the WebDataExemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLine) GetWebDataExemptionOk() (*WebDataExemption, bool) {
	if o == nil || common.IsNil(o.WebDataExemption) {
		return nil, false
	}
	return o.WebDataExemption, true
}

// HasWebDataExemption returns a boolean if a field has been set.
func (o *BusinessLine) HasWebDataExemption() bool {
	if o != nil && !common.IsNil(o.WebDataExemption) {
		return true
	}

	return false
}

// SetWebDataExemption gets a reference to the given WebDataExemption and assigns it to the WebDataExemption field.
func (o *BusinessLine) SetWebDataExemption(v WebDataExemption) {
	o.WebDataExemption = &v
}

func (o BusinessLine) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capability) {
		toSerialize["capability"] = o.Capability
	}
	toSerialize["id"] = o.Id
	toSerialize["industryCode"] = o.IndustryCode
	toSerialize["legalEntityId"] = o.LegalEntityId
	if !common.IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	if !common.IsNil(o.SalesChannels) {
		toSerialize["salesChannels"] = o.SalesChannels
	}
	toSerialize["service"] = o.Service
	if !common.IsNil(o.SourceOfFunds) {
		toSerialize["sourceOfFunds"] = o.SourceOfFunds
	}
	if !common.IsNil(o.WebData) {
		toSerialize["webData"] = o.WebData
	}
	if !common.IsNil(o.WebDataExemption) {
		toSerialize["webDataExemption"] = o.WebDataExemption
	}
	return toSerialize, nil
}

type NullableBusinessLine struct {
	value *BusinessLine
	isSet bool
}

func (v NullableBusinessLine) Get() *BusinessLine {
	return v.value
}

func (v *NullableBusinessLine) Set(val *BusinessLine) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessLine) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessLine(val *BusinessLine) *NullableBusinessLine {
	return &NullableBusinessLine{value: val, isSet: true}
}

func (v NullableBusinessLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BusinessLine) isValidCapability() bool {
	var allowedEnumValues = []string{"receivePayments", "receiveFromPlatformPayments", "issueBankAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetCapability() == allowed {
			return true
		}
	}
	return false
}
func (o *BusinessLine) isValidService() bool {
	var allowedEnumValues = []string{"paymentProcessing", "banking"}
	for _, allowed := range allowedEnumValues {
		if o.GetService() == allowed {
			return true
		}
	}
	return false
}
