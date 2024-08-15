/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the TerminalSettingsNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalSettingsNotificationRequest{}

// TerminalSettingsNotificationRequest struct for TerminalSettingsNotificationRequest
type TerminalSettingsNotificationRequest struct {
	// Timestamp for when the webhook was created.
	CreatedAt time.Time            `json:"createdAt"`
	Data      TerminalSettingsData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewTerminalSettingsNotificationRequest instantiates a new TerminalSettingsNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalSettingsNotificationRequest(createdAt time.Time, data TerminalSettingsData, environment string, type_ string) *TerminalSettingsNotificationRequest {
	this := TerminalSettingsNotificationRequest{}
	this.CreatedAt = createdAt
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewTerminalSettingsNotificationRequestWithDefaults instantiates a new TerminalSettingsNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalSettingsNotificationRequestWithDefaults() *TerminalSettingsNotificationRequest {
	this := TerminalSettingsNotificationRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TerminalSettingsNotificationRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TerminalSettingsNotificationRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TerminalSettingsNotificationRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *TerminalSettingsNotificationRequest) GetData() TerminalSettingsData {
	if o == nil {
		var ret TerminalSettingsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TerminalSettingsNotificationRequest) GetDataOk() (*TerminalSettingsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TerminalSettingsNotificationRequest) SetData(v TerminalSettingsData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *TerminalSettingsNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *TerminalSettingsNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *TerminalSettingsNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *TerminalSettingsNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerminalSettingsNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerminalSettingsNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o TerminalSettingsNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalSettingsNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTerminalSettingsNotificationRequest struct {
	value *TerminalSettingsNotificationRequest
	isSet bool
}

func (v NullableTerminalSettingsNotificationRequest) Get() *TerminalSettingsNotificationRequest {
	return v.value
}

func (v *NullableTerminalSettingsNotificationRequest) Set(val *TerminalSettingsNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalSettingsNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalSettingsNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalSettingsNotificationRequest(val *TerminalSettingsNotificationRequest) *NullableTerminalSettingsNotificationRequest {
	return &NullableTerminalSettingsNotificationRequest{value: val, isSet: true}
}

func (v NullableTerminalSettingsNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalSettingsNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TerminalSettingsNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"terminalSettings.modified"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
