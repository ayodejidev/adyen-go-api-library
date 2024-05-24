/*
Transfer webhooks

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v9/src/common"
)

// checks if the ConfirmationTrackingData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ConfirmationTrackingData{}

// ConfirmationTrackingData struct for ConfirmationTrackingData
type ConfirmationTrackingData struct {
	// The status of the transfer.  Possible values:    - **credited**: the funds are credited to your user's transfer instrument or bank account.
	Status string `json:"status"`
	// The type of the tracking event.  Possible values:   - **confirmation**: the transfer passed Adyen's internal review.
	Type string `json:"type"`
}

// NewConfirmationTrackingData instantiates a new ConfirmationTrackingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmationTrackingData(status string, type_ string) *ConfirmationTrackingData {
	this := ConfirmationTrackingData{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewConfirmationTrackingDataWithDefaults instantiates a new ConfirmationTrackingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmationTrackingDataWithDefaults() *ConfirmationTrackingData {
	this := ConfirmationTrackingData{}
	var type_ string = "confirmation"
	this.Type = type_
	return &this
}

// GetStatus returns the Status field value
func (o *ConfirmationTrackingData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ConfirmationTrackingData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ConfirmationTrackingData) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ConfirmationTrackingData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfirmationTrackingData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfirmationTrackingData) SetType(v string) {
	o.Type = v
}

func (o ConfirmationTrackingData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfirmationTrackingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableConfirmationTrackingData struct {
	value *ConfirmationTrackingData
	isSet bool
}

func (v NullableConfirmationTrackingData) Get() *ConfirmationTrackingData {
	return v.value
}

func (v *NullableConfirmationTrackingData) Set(val *ConfirmationTrackingData) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmationTrackingData) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmationTrackingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmationTrackingData(val *ConfirmationTrackingData) *NullableConfirmationTrackingData {
	return &NullableConfirmationTrackingData{value: val, isSet: true}
}

func (v NullableConfirmationTrackingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfirmationTrackingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ConfirmationTrackingData) isValidStatus() bool {
	var allowedEnumValues = []string{"credited"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *ConfirmationTrackingData) isValidType() bool {
	var allowedEnumValues = []string{"confirmation"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
