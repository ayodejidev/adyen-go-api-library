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

// checks if the MerchantUpdatedNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MerchantUpdatedNotificationRequest{}

// MerchantUpdatedNotificationRequest struct for MerchantUpdatedNotificationRequest
type MerchantUpdatedNotificationRequest struct {
	// Timestamp for when the webhook was created.
	CreatedAt time.Time                     `json:"createdAt"`
	Data      AccountUpdateNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// Type of notification.
	Type string `json:"type"`
}

// NewMerchantUpdatedNotificationRequest instantiates a new MerchantUpdatedNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUpdatedNotificationRequest(createdAt time.Time, data AccountUpdateNotificationData, environment string, type_ string) *MerchantUpdatedNotificationRequest {
	this := MerchantUpdatedNotificationRequest{}
	this.CreatedAt = createdAt
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewMerchantUpdatedNotificationRequestWithDefaults instantiates a new MerchantUpdatedNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUpdatedNotificationRequestWithDefaults() *MerchantUpdatedNotificationRequest {
	this := MerchantUpdatedNotificationRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *MerchantUpdatedNotificationRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MerchantUpdatedNotificationRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MerchantUpdatedNotificationRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *MerchantUpdatedNotificationRequest) GetData() AccountUpdateNotificationData {
	if o == nil {
		var ret AccountUpdateNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MerchantUpdatedNotificationRequest) GetDataOk() (*AccountUpdateNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MerchantUpdatedNotificationRequest) SetData(v AccountUpdateNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *MerchantUpdatedNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *MerchantUpdatedNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *MerchantUpdatedNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetType returns the Type field value
func (o *MerchantUpdatedNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MerchantUpdatedNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MerchantUpdatedNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o MerchantUpdatedNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantUpdatedNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMerchantUpdatedNotificationRequest struct {
	value *MerchantUpdatedNotificationRequest
	isSet bool
}

func (v NullableMerchantUpdatedNotificationRequest) Get() *MerchantUpdatedNotificationRequest {
	return v.value
}

func (v *NullableMerchantUpdatedNotificationRequest) Set(val *MerchantUpdatedNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUpdatedNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUpdatedNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUpdatedNotificationRequest(val *MerchantUpdatedNotificationRequest) *NullableMerchantUpdatedNotificationRequest {
	return &NullableMerchantUpdatedNotificationRequest{value: val, isSet: true}
}

func (v NullableMerchantUpdatedNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUpdatedNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *MerchantUpdatedNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"merchant.updated"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
