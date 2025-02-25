/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the AmexInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AmexInfo{}

// AmexInfo struct for AmexInfo
type AmexInfo struct {
	// Merchant ID (MID) number. Format: 10 numeric characters.  You must provide this field when you request `gatewayContract` or `paymentDesignatorContract` service levels.
	MidNumber *string `json:"midNumber,omitempty"`
	// Indicates whether the Amex Merchant ID is reused from a previously setup Amex payment method.  This is only applicable for `gatewayContract` and `paymentDesignatorContract` service levels.  The default value is **false**.
	ReuseMidNumber *bool `json:"reuseMidNumber,omitempty"`
	// Specifies the service level (settlement type) of this payment method. Possible values: * **noContract**: Adyen holds the contract with American Express. * **gatewayContract**: American Express receives the settlement and handles disputes, then pays out to you or your sub-merchant directly. * **paymentDesignatorContract**: Adyen receives the settlement, and handles disputes and payouts.
	ServiceLevel string `json:"serviceLevel"`
}

// NewAmexInfo instantiates a new AmexInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmexInfo(serviceLevel string) *AmexInfo {
	this := AmexInfo{}
	var reuseMidNumber bool = false
	this.ReuseMidNumber = &reuseMidNumber
	this.ServiceLevel = serviceLevel
	return &this
}

// NewAmexInfoWithDefaults instantiates a new AmexInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmexInfoWithDefaults() *AmexInfo {
	this := AmexInfo{}
	var reuseMidNumber bool = false
	this.ReuseMidNumber = &reuseMidNumber
	return &this
}

// GetMidNumber returns the MidNumber field value if set, zero value otherwise.
func (o *AmexInfo) GetMidNumber() string {
	if o == nil || common.IsNil(o.MidNumber) {
		var ret string
		return ret
	}
	return *o.MidNumber
}

// GetMidNumberOk returns a tuple with the MidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmexInfo) GetMidNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.MidNumber) {
		return nil, false
	}
	return o.MidNumber, true
}

// HasMidNumber returns a boolean if a field has been set.
func (o *AmexInfo) HasMidNumber() bool {
	if o != nil && !common.IsNil(o.MidNumber) {
		return true
	}

	return false
}

// SetMidNumber gets a reference to the given string and assigns it to the MidNumber field.
func (o *AmexInfo) SetMidNumber(v string) {
	o.MidNumber = &v
}

// GetReuseMidNumber returns the ReuseMidNumber field value if set, zero value otherwise.
func (o *AmexInfo) GetReuseMidNumber() bool {
	if o == nil || common.IsNil(o.ReuseMidNumber) {
		var ret bool
		return ret
	}
	return *o.ReuseMidNumber
}

// GetReuseMidNumberOk returns a tuple with the ReuseMidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmexInfo) GetReuseMidNumberOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReuseMidNumber) {
		return nil, false
	}
	return o.ReuseMidNumber, true
}

// HasReuseMidNumber returns a boolean if a field has been set.
func (o *AmexInfo) HasReuseMidNumber() bool {
	if o != nil && !common.IsNil(o.ReuseMidNumber) {
		return true
	}

	return false
}

// SetReuseMidNumber gets a reference to the given bool and assigns it to the ReuseMidNumber field.
func (o *AmexInfo) SetReuseMidNumber(v bool) {
	o.ReuseMidNumber = &v
}

// GetServiceLevel returns the ServiceLevel field value
func (o *AmexInfo) GetServiceLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value
// and a boolean to check if the value has been set.
func (o *AmexInfo) GetServiceLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceLevel, true
}

// SetServiceLevel sets field value
func (o *AmexInfo) SetServiceLevel(v string) {
	o.ServiceLevel = v
}

func (o AmexInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmexInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.MidNumber) {
		toSerialize["midNumber"] = o.MidNumber
	}
	if !common.IsNil(o.ReuseMidNumber) {
		toSerialize["reuseMidNumber"] = o.ReuseMidNumber
	}
	toSerialize["serviceLevel"] = o.ServiceLevel
	return toSerialize, nil
}

type NullableAmexInfo struct {
	value *AmexInfo
	isSet bool
}

func (v NullableAmexInfo) Get() *AmexInfo {
	return v.value
}

func (v *NullableAmexInfo) Set(val *AmexInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAmexInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAmexInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmexInfo(val *AmexInfo) *NullableAmexInfo {
	return &NullableAmexInfo{value: val, isSet: true}
}

func (v NullableAmexInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmexInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AmexInfo) isValidServiceLevel() bool {
	var allowedEnumValues = []string{"noContract", "gatewayContract", "paymentDesignatorContract"}
	for _, allowed := range allowedEnumValues {
		if o.GetServiceLevel() == allowed {
			return true
		}
	}
	return false
}
