/*
Adyen BinLookup API

API version: 54
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package binlookup

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the CostEstimateAssumptions type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CostEstimateAssumptions{}

// CostEstimateAssumptions struct for CostEstimateAssumptions
type CostEstimateAssumptions struct {
	// If true, the cardholder is expected to successfully authorise via 3D Secure.
	Assume3DSecureAuthenticated *bool `json:"assume3DSecureAuthenticated,omitempty"`
	// If true, the transaction is expected to have valid Level 3 data.
	AssumeLevel3Data *bool `json:"assumeLevel3Data,omitempty"`
	// If not zero, the number of installments.
	Installments *int32 `json:"installments,omitempty"`
}

// NewCostEstimateAssumptions instantiates a new CostEstimateAssumptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCostEstimateAssumptions() *CostEstimateAssumptions {
	this := CostEstimateAssumptions{}
	return &this
}

// NewCostEstimateAssumptionsWithDefaults instantiates a new CostEstimateAssumptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCostEstimateAssumptionsWithDefaults() *CostEstimateAssumptions {
	this := CostEstimateAssumptions{}
	return &this
}

// GetAssume3DSecureAuthenticated returns the Assume3DSecureAuthenticated field value if set, zero value otherwise.
func (o *CostEstimateAssumptions) GetAssume3DSecureAuthenticated() bool {
	if o == nil || common.IsNil(o.Assume3DSecureAuthenticated) {
		var ret bool
		return ret
	}
	return *o.Assume3DSecureAuthenticated
}

// GetAssume3DSecureAuthenticatedOk returns a tuple with the Assume3DSecureAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateAssumptions) GetAssume3DSecureAuthenticatedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Assume3DSecureAuthenticated) {
		return nil, false
	}
	return o.Assume3DSecureAuthenticated, true
}

// HasAssume3DSecureAuthenticated returns a boolean if a field has been set.
func (o *CostEstimateAssumptions) HasAssume3DSecureAuthenticated() bool {
	if o != nil && !common.IsNil(o.Assume3DSecureAuthenticated) {
		return true
	}

	return false
}

// SetAssume3DSecureAuthenticated gets a reference to the given bool and assigns it to the Assume3DSecureAuthenticated field.
func (o *CostEstimateAssumptions) SetAssume3DSecureAuthenticated(v bool) {
	o.Assume3DSecureAuthenticated = &v
}

// GetAssumeLevel3Data returns the AssumeLevel3Data field value if set, zero value otherwise.
func (o *CostEstimateAssumptions) GetAssumeLevel3Data() bool {
	if o == nil || common.IsNil(o.AssumeLevel3Data) {
		var ret bool
		return ret
	}
	return *o.AssumeLevel3Data
}

// GetAssumeLevel3DataOk returns a tuple with the AssumeLevel3Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateAssumptions) GetAssumeLevel3DataOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AssumeLevel3Data) {
		return nil, false
	}
	return o.AssumeLevel3Data, true
}

// HasAssumeLevel3Data returns a boolean if a field has been set.
func (o *CostEstimateAssumptions) HasAssumeLevel3Data() bool {
	if o != nil && !common.IsNil(o.AssumeLevel3Data) {
		return true
	}

	return false
}

// SetAssumeLevel3Data gets a reference to the given bool and assigns it to the AssumeLevel3Data field.
func (o *CostEstimateAssumptions) SetAssumeLevel3Data(v bool) {
	o.AssumeLevel3Data = &v
}

// GetInstallments returns the Installments field value if set, zero value otherwise.
func (o *CostEstimateAssumptions) GetInstallments() int32 {
	if o == nil || common.IsNil(o.Installments) {
		var ret int32
		return ret
	}
	return *o.Installments
}

// GetInstallmentsOk returns a tuple with the Installments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostEstimateAssumptions) GetInstallmentsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Installments) {
		return nil, false
	}
	return o.Installments, true
}

// HasInstallments returns a boolean if a field has been set.
func (o *CostEstimateAssumptions) HasInstallments() bool {
	if o != nil && !common.IsNil(o.Installments) {
		return true
	}

	return false
}

// SetInstallments gets a reference to the given int32 and assigns it to the Installments field.
func (o *CostEstimateAssumptions) SetInstallments(v int32) {
	o.Installments = &v
}

func (o CostEstimateAssumptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CostEstimateAssumptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assume3DSecureAuthenticated) {
		toSerialize["assume3DSecureAuthenticated"] = o.Assume3DSecureAuthenticated
	}
	if !common.IsNil(o.AssumeLevel3Data) {
		toSerialize["assumeLevel3Data"] = o.AssumeLevel3Data
	}
	if !common.IsNil(o.Installments) {
		toSerialize["installments"] = o.Installments
	}
	return toSerialize, nil
}

type NullableCostEstimateAssumptions struct {
	value *CostEstimateAssumptions
	isSet bool
}

func (v NullableCostEstimateAssumptions) Get() *CostEstimateAssumptions {
	return v.value
}

func (v *NullableCostEstimateAssumptions) Set(val *CostEstimateAssumptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCostEstimateAssumptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCostEstimateAssumptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCostEstimateAssumptions(val *CostEstimateAssumptions) *NullableCostEstimateAssumptions {
	return &NullableCostEstimateAssumptions{value: val, isSet: true}
}

func (v NullableCostEstimateAssumptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCostEstimateAssumptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
