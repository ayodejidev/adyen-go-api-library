/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the CapabilitySettings type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapabilitySettings{}

// CapabilitySettings struct for CapabilitySettings
type CapabilitySettings struct {
	//
	AmountPerIndustry *map[string]Amount `json:"amountPerIndustry,omitempty"`
	//
	AuthorizedCardUsers *bool `json:"authorizedCardUsers,omitempty"`
	//
	FundingSource []string `json:"fundingSource,omitempty"`
	//
	Interval  *string `json:"interval,omitempty"`
	MaxAmount *Amount `json:"maxAmount,omitempty"`
}

// NewCapabilitySettings instantiates a new CapabilitySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilitySettings() *CapabilitySettings {
	this := CapabilitySettings{}
	return &this
}

// NewCapabilitySettingsWithDefaults instantiates a new CapabilitySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilitySettingsWithDefaults() *CapabilitySettings {
	this := CapabilitySettings{}
	return &this
}

// GetAmountPerIndustry returns the AmountPerIndustry field value if set, zero value otherwise.
func (o *CapabilitySettings) GetAmountPerIndustry() map[string]Amount {
	if o == nil || common.IsNil(o.AmountPerIndustry) {
		var ret map[string]Amount
		return ret
	}
	return *o.AmountPerIndustry
}

// GetAmountPerIndustryOk returns a tuple with the AmountPerIndustry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySettings) GetAmountPerIndustryOk() (*map[string]Amount, bool) {
	if o == nil || common.IsNil(o.AmountPerIndustry) {
		return nil, false
	}
	return o.AmountPerIndustry, true
}

// HasAmountPerIndustry returns a boolean if a field has been set.
func (o *CapabilitySettings) HasAmountPerIndustry() bool {
	if o != nil && !common.IsNil(o.AmountPerIndustry) {
		return true
	}

	return false
}

// SetAmountPerIndustry gets a reference to the given map[string]Amount and assigns it to the AmountPerIndustry field.
func (o *CapabilitySettings) SetAmountPerIndustry(v map[string]Amount) {
	o.AmountPerIndustry = &v
}

// GetAuthorizedCardUsers returns the AuthorizedCardUsers field value if set, zero value otherwise.
func (o *CapabilitySettings) GetAuthorizedCardUsers() bool {
	if o == nil || common.IsNil(o.AuthorizedCardUsers) {
		var ret bool
		return ret
	}
	return *o.AuthorizedCardUsers
}

// GetAuthorizedCardUsersOk returns a tuple with the AuthorizedCardUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySettings) GetAuthorizedCardUsersOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AuthorizedCardUsers) {
		return nil, false
	}
	return o.AuthorizedCardUsers, true
}

// HasAuthorizedCardUsers returns a boolean if a field has been set.
func (o *CapabilitySettings) HasAuthorizedCardUsers() bool {
	if o != nil && !common.IsNil(o.AuthorizedCardUsers) {
		return true
	}

	return false
}

// SetAuthorizedCardUsers gets a reference to the given bool and assigns it to the AuthorizedCardUsers field.
func (o *CapabilitySettings) SetAuthorizedCardUsers(v bool) {
	o.AuthorizedCardUsers = &v
}

// GetFundingSource returns the FundingSource field value if set, zero value otherwise.
func (o *CapabilitySettings) GetFundingSource() []string {
	if o == nil || common.IsNil(o.FundingSource) {
		var ret []string
		return ret
	}
	return o.FundingSource
}

// GetFundingSourceOk returns a tuple with the FundingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySettings) GetFundingSourceOk() ([]string, bool) {
	if o == nil || common.IsNil(o.FundingSource) {
		return nil, false
	}
	return o.FundingSource, true
}

// HasFundingSource returns a boolean if a field has been set.
func (o *CapabilitySettings) HasFundingSource() bool {
	if o != nil && !common.IsNil(o.FundingSource) {
		return true
	}

	return false
}

// SetFundingSource gets a reference to the given []string and assigns it to the FundingSource field.
func (o *CapabilitySettings) SetFundingSource(v []string) {
	o.FundingSource = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CapabilitySettings) GetInterval() string {
	if o == nil || common.IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySettings) GetIntervalOk() (*string, bool) {
	if o == nil || common.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CapabilitySettings) HasInterval() bool {
	if o != nil && !common.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *CapabilitySettings) SetInterval(v string) {
	o.Interval = &v
}

// GetMaxAmount returns the MaxAmount field value if set, zero value otherwise.
func (o *CapabilitySettings) GetMaxAmount() Amount {
	if o == nil || common.IsNil(o.MaxAmount) {
		var ret Amount
		return ret
	}
	return *o.MaxAmount
}

// GetMaxAmountOk returns a tuple with the MaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilitySettings) GetMaxAmountOk() (*Amount, bool) {
	if o == nil || common.IsNil(o.MaxAmount) {
		return nil, false
	}
	return o.MaxAmount, true
}

// HasMaxAmount returns a boolean if a field has been set.
func (o *CapabilitySettings) HasMaxAmount() bool {
	if o != nil && !common.IsNil(o.MaxAmount) {
		return true
	}

	return false
}

// SetMaxAmount gets a reference to the given Amount and assigns it to the MaxAmount field.
func (o *CapabilitySettings) SetMaxAmount(v Amount) {
	o.MaxAmount = &v
}

func (o CapabilitySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilitySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AmountPerIndustry) {
		toSerialize["amountPerIndustry"] = o.AmountPerIndustry
	}
	if !common.IsNil(o.AuthorizedCardUsers) {
		toSerialize["authorizedCardUsers"] = o.AuthorizedCardUsers
	}
	if !common.IsNil(o.FundingSource) {
		toSerialize["fundingSource"] = o.FundingSource
	}
	if !common.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !common.IsNil(o.MaxAmount) {
		toSerialize["maxAmount"] = o.MaxAmount
	}
	return toSerialize, nil
}

type NullableCapabilitySettings struct {
	value *CapabilitySettings
	isSet bool
}

func (v NullableCapabilitySettings) Get() *CapabilitySettings {
	return v.value
}

func (v *NullableCapabilitySettings) Set(val *CapabilitySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilitySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilitySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilitySettings(val *CapabilitySettings) *NullableCapabilitySettings {
	return &NullableCapabilitySettings{value: val, isSet: true}
}

func (v NullableCapabilitySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilitySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CapabilitySettings) isValidInterval() bool {
	var allowedEnumValues = []string{"daily", "monthly", "weekly"}
	for _, allowed := range allowedEnumValues {
		if o.GetInterval() == allowed {
			return true
		}
	}
	return false
}
