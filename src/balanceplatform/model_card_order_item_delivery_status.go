/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the CardOrderItemDeliveryStatus type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CardOrderItemDeliveryStatus{}

// CardOrderItemDeliveryStatus struct for CardOrderItemDeliveryStatus
type CardOrderItemDeliveryStatus struct {
	// An error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The status of the PIN delivery.
	Status *string `json:"status,omitempty"`
	// The tracking number of the PIN delivery.
	TrackingNumber *string `json:"trackingNumber,omitempty"`
}

// NewCardOrderItemDeliveryStatus instantiates a new CardOrderItemDeliveryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardOrderItemDeliveryStatus() *CardOrderItemDeliveryStatus {
	this := CardOrderItemDeliveryStatus{}
	return &this
}

// NewCardOrderItemDeliveryStatusWithDefaults instantiates a new CardOrderItemDeliveryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardOrderItemDeliveryStatusWithDefaults() *CardOrderItemDeliveryStatus {
	this := CardOrderItemDeliveryStatus{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetErrorMessage() string {
	if o == nil || common.IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasErrorMessage() bool {
	if o != nil && !common.IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CardOrderItemDeliveryStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CardOrderItemDeliveryStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *CardOrderItemDeliveryStatus) GetTrackingNumber() string {
	if o == nil || common.IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardOrderItemDeliveryStatus) GetTrackingNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *CardOrderItemDeliveryStatus) HasTrackingNumber() bool {
	if o != nil && !common.IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *CardOrderItemDeliveryStatus) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

func (o CardOrderItemDeliveryStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardOrderItemDeliveryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	return toSerialize, nil
}

type NullableCardOrderItemDeliveryStatus struct {
	value *CardOrderItemDeliveryStatus
	isSet bool
}

func (v NullableCardOrderItemDeliveryStatus) Get() *CardOrderItemDeliveryStatus {
	return v.value
}

func (v *NullableCardOrderItemDeliveryStatus) Set(val *CardOrderItemDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCardOrderItemDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCardOrderItemDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardOrderItemDeliveryStatus(val *CardOrderItemDeliveryStatus) *NullableCardOrderItemDeliveryStatus {
	return &NullableCardOrderItemDeliveryStatus{value: val, isSet: true}
}

func (v NullableCardOrderItemDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardOrderItemDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CardOrderItemDeliveryStatus) isValidStatus() bool {
	var allowedEnumValues = []string{"created", "delivered", "notApplicable", "processing", "produced", "rejected", "shipped", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
