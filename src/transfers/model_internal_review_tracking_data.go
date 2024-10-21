/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the InternalReviewTrackingData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InternalReviewTrackingData{}

// InternalReviewTrackingData struct for InternalReviewTrackingData
type InternalReviewTrackingData struct {
	// The reason why the transfer failed Adyen's internal review.   Possible values:  - **refusedForRegulatoryReasons**: the transfer does not comply with Adyen's risk policy. For more information, [contact the Support Team](https://www.adyen.help/hc/en-us/requests/new).
	Reason *string `json:"reason,omitempty"`
	// The status of the transfer.  Possible values:   - **pending**: the transfer is under internal review.  - **failed**: the transfer failed Adyen's internal review. For details, see `reason`.
	Status string `json:"status"`
	// The type of tracking event.   Possible values:    - **internalReview**: the transfer was flagged because it does not comply with Adyen's risk policy.
	Type string `json:"type"`
}

// NewInternalReviewTrackingData instantiates a new InternalReviewTrackingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalReviewTrackingData(status string, type_ string) *InternalReviewTrackingData {
	this := InternalReviewTrackingData{}
	this.Status = status
	this.Type = type_
	return &this
}

// NewInternalReviewTrackingDataWithDefaults instantiates a new InternalReviewTrackingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalReviewTrackingDataWithDefaults() *InternalReviewTrackingData {
	this := InternalReviewTrackingData{}
	var type_ string = "internalReview"
	this.Type = type_
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InternalReviewTrackingData) GetReason() string {
	if o == nil || common.IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalReviewTrackingData) GetReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InternalReviewTrackingData) HasReason() bool {
	if o != nil && !common.IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InternalReviewTrackingData) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *InternalReviewTrackingData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InternalReviewTrackingData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InternalReviewTrackingData) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *InternalReviewTrackingData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InternalReviewTrackingData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InternalReviewTrackingData) SetType(v string) {
	o.Type = v
}

func (o InternalReviewTrackingData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalReviewTrackingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableInternalReviewTrackingData struct {
	value *InternalReviewTrackingData
	isSet bool
}

func (v NullableInternalReviewTrackingData) Get() *InternalReviewTrackingData {
	return v.value
}

func (v *NullableInternalReviewTrackingData) Set(val *InternalReviewTrackingData) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalReviewTrackingData) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalReviewTrackingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalReviewTrackingData(val *InternalReviewTrackingData) *NullableInternalReviewTrackingData {
	return &NullableInternalReviewTrackingData{value: val, isSet: true}
}

func (v NullableInternalReviewTrackingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalReviewTrackingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *InternalReviewTrackingData) isValidReason() bool {
	var allowedEnumValues = []string{"refusedForRegulatoryReasons"}
	for _, allowed := range allowedEnumValues {
		if o.GetReason() == allowed {
			return true
		}
	}
	return false
}
func (o *InternalReviewTrackingData) isValidStatus() bool {
	var allowedEnumValues = []string{"pending", "failed"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
func (o *InternalReviewTrackingData) isValidType() bool {
	var allowedEnumValues = []string{"internalReview"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
