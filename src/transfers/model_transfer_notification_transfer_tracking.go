/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the TransferNotificationTransferTracking type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TransferNotificationTransferTracking{}

// TransferNotificationTransferTracking struct for TransferNotificationTransferTracking
type TransferNotificationTransferTracking struct {
	// The estimated time the beneficiary should have access to the funds.
	EstimatedArrivalTime *time.Time `json:"estimatedArrivalTime,omitempty"`
	// The tracking status of the transfer.
	Status *string `json:"status,omitempty"`
}

// NewTransferNotificationTransferTracking instantiates a new TransferNotificationTransferTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferNotificationTransferTracking() *TransferNotificationTransferTracking {
	this := TransferNotificationTransferTracking{}
	return &this
}

// NewTransferNotificationTransferTrackingWithDefaults instantiates a new TransferNotificationTransferTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferNotificationTransferTrackingWithDefaults() *TransferNotificationTransferTracking {
	this := TransferNotificationTransferTracking{}
	return &this
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise.
func (o *TransferNotificationTransferTracking) GetEstimatedArrivalTime() time.Time {
	if o == nil || common.IsNil(o.EstimatedArrivalTime) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedArrivalTime
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationTransferTracking) GetEstimatedArrivalTimeOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.EstimatedArrivalTime) {
		return nil, false
	}
	return o.EstimatedArrivalTime, true
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *TransferNotificationTransferTracking) HasEstimatedArrivalTime() bool {
	if o != nil && !common.IsNil(o.EstimatedArrivalTime) {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given time.Time and assigns it to the EstimatedArrivalTime field.
func (o *TransferNotificationTransferTracking) SetEstimatedArrivalTime(v time.Time) {
	o.EstimatedArrivalTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransferNotificationTransferTracking) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferNotificationTransferTracking) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransferNotificationTransferTracking) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TransferNotificationTransferTracking) SetStatus(v string) {
	o.Status = &v
}

func (o TransferNotificationTransferTracking) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferNotificationTransferTracking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EstimatedArrivalTime) {
		toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableTransferNotificationTransferTracking struct {
	value *TransferNotificationTransferTracking
	isSet bool
}

func (v NullableTransferNotificationTransferTracking) Get() *TransferNotificationTransferTracking {
	return v.value
}

func (v *NullableTransferNotificationTransferTracking) Set(val *TransferNotificationTransferTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferNotificationTransferTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferNotificationTransferTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferNotificationTransferTracking(val *TransferNotificationTransferTracking) *NullableTransferNotificationTransferTracking {
	return &NullableTransferNotificationTransferTracking{value: val, isSet: true}
}

func (v NullableTransferNotificationTransferTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferNotificationTransferTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *TransferNotificationTransferTracking) isValidStatus() bool {
	var allowedEnumValues = []string{"credited"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
