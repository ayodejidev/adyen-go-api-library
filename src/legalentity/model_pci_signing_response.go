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

// checks if the PciSigningResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PciSigningResponse{}

// PciSigningResponse struct for PciSigningResponse
type PciSigningResponse struct {
	// The unique identifiers of the signed PCI documents.
	PciQuestionnaireIds []string `json:"pciQuestionnaireIds,omitempty"`
	// The [legal entity ID](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities__resParam_id) of the individual who signed the PCI questionnaire.
	SignedBy *string `json:"signedBy,omitempty"`
}

// NewPciSigningResponse instantiates a new PciSigningResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciSigningResponse() *PciSigningResponse {
	this := PciSigningResponse{}
	return &this
}

// NewPciSigningResponseWithDefaults instantiates a new PciSigningResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciSigningResponseWithDefaults() *PciSigningResponse {
	this := PciSigningResponse{}
	return &this
}

// GetPciQuestionnaireIds returns the PciQuestionnaireIds field value if set, zero value otherwise.
func (o *PciSigningResponse) GetPciQuestionnaireIds() []string {
	if o == nil || common.IsNil(o.PciQuestionnaireIds) {
		var ret []string
		return ret
	}
	return o.PciQuestionnaireIds
}

// GetPciQuestionnaireIdsOk returns a tuple with the PciQuestionnaireIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSigningResponse) GetPciQuestionnaireIdsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.PciQuestionnaireIds) {
		return nil, false
	}
	return o.PciQuestionnaireIds, true
}

// HasPciQuestionnaireIds returns a boolean if a field has been set.
func (o *PciSigningResponse) HasPciQuestionnaireIds() bool {
	if o != nil && !common.IsNil(o.PciQuestionnaireIds) {
		return true
	}

	return false
}

// SetPciQuestionnaireIds gets a reference to the given []string and assigns it to the PciQuestionnaireIds field.
func (o *PciSigningResponse) SetPciQuestionnaireIds(v []string) {
	o.PciQuestionnaireIds = v
}

// GetSignedBy returns the SignedBy field value if set, zero value otherwise.
func (o *PciSigningResponse) GetSignedBy() string {
	if o == nil || common.IsNil(o.SignedBy) {
		var ret string
		return ret
	}
	return *o.SignedBy
}

// GetSignedByOk returns a tuple with the SignedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSigningResponse) GetSignedByOk() (*string, bool) {
	if o == nil || common.IsNil(o.SignedBy) {
		return nil, false
	}
	return o.SignedBy, true
}

// HasSignedBy returns a boolean if a field has been set.
func (o *PciSigningResponse) HasSignedBy() bool {
	if o != nil && !common.IsNil(o.SignedBy) {
		return true
	}

	return false
}

// SetSignedBy gets a reference to the given string and assigns it to the SignedBy field.
func (o *PciSigningResponse) SetSignedBy(v string) {
	o.SignedBy = &v
}

func (o PciSigningResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PciSigningResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.PciQuestionnaireIds) {
		toSerialize["pciQuestionnaireIds"] = o.PciQuestionnaireIds
	}
	if !common.IsNil(o.SignedBy) {
		toSerialize["signedBy"] = o.SignedBy
	}
	return toSerialize, nil
}

type NullablePciSigningResponse struct {
	value *PciSigningResponse
	isSet bool
}

func (v NullablePciSigningResponse) Get() *PciSigningResponse {
	return v.value
}

func (v *NullablePciSigningResponse) Set(val *PciSigningResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePciSigningResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePciSigningResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciSigningResponse(val *PciSigningResponse) *NullablePciSigningResponse {
	return &NullablePciSigningResponse{value: val, isSet: true}
}

func (v NullablePciSigningResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciSigningResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
