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

// checks if the BusinessLineInfoUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BusinessLineInfoUpdate{}

// BusinessLineInfoUpdate struct for BusinessLineInfoUpdate
type BusinessLineInfoUpdate struct {
	// The capability for which you are creating the business line. For example, **receivePayments**.
	// Deprecated
	Capability *string `json:"capability,omitempty"`
	// A code that represents the industry of your legal entity. For example, **4431A** for computer software stores.
	IndustryCode *string `json:"industryCode,omitempty"`
	// Unique identifier of the [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) that owns the business line.
	LegalEntityId *string `json:"legalEntityId,omitempty"`
	// A list of channels where goods or services are sold.  Possible values: **pos**, **posMoto**, **eCommerce**, **ecomMoto**, **payByLink**.  Required only in combination with the `service` **paymentProcessing**.
	SalesChannels []string `json:"salesChannels,omitempty"`
	// The service for which you are creating the business line.    Possible values: *  **paymentProcessing** *  **banking**
	Service       *string        `json:"service,omitempty"`
	SourceOfFunds *SourceOfFunds `json:"sourceOfFunds,omitempty"`
	// List of website URLs where your user's goods or services are sold. When this is required for a service but your user does not have an online presence, provide the reason in the `webDataExemption` object.
	WebData          []WebData         `json:"webData,omitempty"`
	WebDataExemption *WebDataExemption `json:"webDataExemption,omitempty"`
}

// NewBusinessLineInfoUpdate instantiates a new BusinessLineInfoUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessLineInfoUpdate() *BusinessLineInfoUpdate {
	this := BusinessLineInfoUpdate{}
	return &this
}

// NewBusinessLineInfoUpdateWithDefaults instantiates a new BusinessLineInfoUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessLineInfoUpdateWithDefaults() *BusinessLineInfoUpdate {
	this := BusinessLineInfoUpdate{}
	return &this
}

// GetCapability returns the Capability field value if set, zero value otherwise.
// Deprecated
func (o *BusinessLineInfoUpdate) GetCapability() string {
	if o == nil || common.IsNil(o.Capability) {
		var ret string
		return ret
	}
	return *o.Capability
}

// GetCapabilityOk returns a tuple with the Capability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BusinessLineInfoUpdate) GetCapabilityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Capability) {
		return nil, false
	}
	return o.Capability, true
}

// HasCapability returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasCapability() bool {
	if o != nil && !common.IsNil(o.Capability) {
		return true
	}

	return false
}

// SetCapability gets a reference to the given string and assigns it to the Capability field.
// Deprecated
func (o *BusinessLineInfoUpdate) SetCapability(v string) {
	o.Capability = &v
}

// GetIndustryCode returns the IndustryCode field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetIndustryCode() string {
	if o == nil || common.IsNil(o.IndustryCode) {
		var ret string
		return ret
	}
	return *o.IndustryCode
}

// GetIndustryCodeOk returns a tuple with the IndustryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetIndustryCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.IndustryCode) {
		return nil, false
	}
	return o.IndustryCode, true
}

// HasIndustryCode returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasIndustryCode() bool {
	if o != nil && !common.IsNil(o.IndustryCode) {
		return true
	}

	return false
}

// SetIndustryCode gets a reference to the given string and assigns it to the IndustryCode field.
func (o *BusinessLineInfoUpdate) SetIndustryCode(v string) {
	o.IndustryCode = &v
}

// GetLegalEntityId returns the LegalEntityId field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetLegalEntityId() string {
	if o == nil || common.IsNil(o.LegalEntityId) {
		var ret string
		return ret
	}
	return *o.LegalEntityId
}

// GetLegalEntityIdOk returns a tuple with the LegalEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetLegalEntityIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.LegalEntityId) {
		return nil, false
	}
	return o.LegalEntityId, true
}

// HasLegalEntityId returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasLegalEntityId() bool {
	if o != nil && !common.IsNil(o.LegalEntityId) {
		return true
	}

	return false
}

// SetLegalEntityId gets a reference to the given string and assigns it to the LegalEntityId field.
func (o *BusinessLineInfoUpdate) SetLegalEntityId(v string) {
	o.LegalEntityId = &v
}

// GetSalesChannels returns the SalesChannels field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetSalesChannels() []string {
	if o == nil || common.IsNil(o.SalesChannels) {
		var ret []string
		return ret
	}
	return o.SalesChannels
}

// GetSalesChannelsOk returns a tuple with the SalesChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetSalesChannelsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.SalesChannels) {
		return nil, false
	}
	return o.SalesChannels, true
}

// HasSalesChannels returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasSalesChannels() bool {
	if o != nil && !common.IsNil(o.SalesChannels) {
		return true
	}

	return false
}

// SetSalesChannels gets a reference to the given []string and assigns it to the SalesChannels field.
func (o *BusinessLineInfoUpdate) SetSalesChannels(v []string) {
	o.SalesChannels = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetService() string {
	if o == nil || common.IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetServiceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasService() bool {
	if o != nil && !common.IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *BusinessLineInfoUpdate) SetService(v string) {
	o.Service = &v
}

// GetSourceOfFunds returns the SourceOfFunds field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetSourceOfFunds() SourceOfFunds {
	if o == nil || common.IsNil(o.SourceOfFunds) {
		var ret SourceOfFunds
		return ret
	}
	return *o.SourceOfFunds
}

// GetSourceOfFundsOk returns a tuple with the SourceOfFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetSourceOfFundsOk() (*SourceOfFunds, bool) {
	if o == nil || common.IsNil(o.SourceOfFunds) {
		return nil, false
	}
	return o.SourceOfFunds, true
}

// HasSourceOfFunds returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasSourceOfFunds() bool {
	if o != nil && !common.IsNil(o.SourceOfFunds) {
		return true
	}

	return false
}

// SetSourceOfFunds gets a reference to the given SourceOfFunds and assigns it to the SourceOfFunds field.
func (o *BusinessLineInfoUpdate) SetSourceOfFunds(v SourceOfFunds) {
	o.SourceOfFunds = &v
}

// GetWebData returns the WebData field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetWebData() []WebData {
	if o == nil || common.IsNil(o.WebData) {
		var ret []WebData
		return ret
	}
	return o.WebData
}

// GetWebDataOk returns a tuple with the WebData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetWebDataOk() ([]WebData, bool) {
	if o == nil || common.IsNil(o.WebData) {
		return nil, false
	}
	return o.WebData, true
}

// HasWebData returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasWebData() bool {
	if o != nil && !common.IsNil(o.WebData) {
		return true
	}

	return false
}

// SetWebData gets a reference to the given []WebData and assigns it to the WebData field.
func (o *BusinessLineInfoUpdate) SetWebData(v []WebData) {
	o.WebData = v
}

// GetWebDataExemption returns the WebDataExemption field value if set, zero value otherwise.
func (o *BusinessLineInfoUpdate) GetWebDataExemption() WebDataExemption {
	if o == nil || common.IsNil(o.WebDataExemption) {
		var ret WebDataExemption
		return ret
	}
	return *o.WebDataExemption
}

// GetWebDataExemptionOk returns a tuple with the WebDataExemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessLineInfoUpdate) GetWebDataExemptionOk() (*WebDataExemption, bool) {
	if o == nil || common.IsNil(o.WebDataExemption) {
		return nil, false
	}
	return o.WebDataExemption, true
}

// HasWebDataExemption returns a boolean if a field has been set.
func (o *BusinessLineInfoUpdate) HasWebDataExemption() bool {
	if o != nil && !common.IsNil(o.WebDataExemption) {
		return true
	}

	return false
}

// SetWebDataExemption gets a reference to the given WebDataExemption and assigns it to the WebDataExemption field.
func (o *BusinessLineInfoUpdate) SetWebDataExemption(v WebDataExemption) {
	o.WebDataExemption = &v
}

func (o BusinessLineInfoUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessLineInfoUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capability) {
		toSerialize["capability"] = o.Capability
	}
	if !common.IsNil(o.IndustryCode) {
		toSerialize["industryCode"] = o.IndustryCode
	}
	if !common.IsNil(o.LegalEntityId) {
		toSerialize["legalEntityId"] = o.LegalEntityId
	}
	if !common.IsNil(o.SalesChannels) {
		toSerialize["salesChannels"] = o.SalesChannels
	}
	if !common.IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
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

type NullableBusinessLineInfoUpdate struct {
	value *BusinessLineInfoUpdate
	isSet bool
}

func (v NullableBusinessLineInfoUpdate) Get() *BusinessLineInfoUpdate {
	return v.value
}

func (v *NullableBusinessLineInfoUpdate) Set(val *BusinessLineInfoUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessLineInfoUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessLineInfoUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessLineInfoUpdate(val *BusinessLineInfoUpdate) *NullableBusinessLineInfoUpdate {
	return &NullableBusinessLineInfoUpdate{value: val, isSet: true}
}

func (v NullableBusinessLineInfoUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessLineInfoUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BusinessLineInfoUpdate) isValidCapability() bool {
	var allowedEnumValues = []string{"receivePayments", "receiveFromPlatformPayments", "issueBankAccount"}
	for _, allowed := range allowedEnumValues {
		if o.GetCapability() == allowed {
			return true
		}
	}
	return false
}
func (o *BusinessLineInfoUpdate) isValidService() bool {
	var allowedEnumValues = []string{"paymentProcessing", "banking"}
	for _, allowed := range allowedEnumValues {
		if o.GetService() == allowed {
			return true
		}
	}
	return false
}
