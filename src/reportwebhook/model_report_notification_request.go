/*
Report webhooks

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package reportwebhook

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the ReportNotificationRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ReportNotificationRequest{}

// ReportNotificationRequest struct for ReportNotificationRequest
type ReportNotificationRequest struct {
	Data ReportNotificationData `json:"data"`
	// The environment from which the webhook originated.  Possible values: **test**, **live**.
	Environment string `json:"environment"`
	// When the event was queued.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of webhook.
	Type string `json:"type"`
}

// NewReportNotificationRequest instantiates a new ReportNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportNotificationRequest(data ReportNotificationData, environment string, type_ string) *ReportNotificationRequest {
	this := ReportNotificationRequest{}
	this.Data = data
	this.Environment = environment
	this.Type = type_
	return &this
}

// NewReportNotificationRequestWithDefaults instantiates a new ReportNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportNotificationRequestWithDefaults() *ReportNotificationRequest {
	this := ReportNotificationRequest{}
	return &this
}

// GetData returns the Data field value
func (o *ReportNotificationRequest) GetData() ReportNotificationData {
	if o == nil {
		var ret ReportNotificationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReportNotificationRequest) GetDataOk() (*ReportNotificationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReportNotificationRequest) SetData(v ReportNotificationData) {
	o.Data = v
}

// GetEnvironment returns the Environment field value
func (o *ReportNotificationRequest) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *ReportNotificationRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *ReportNotificationRequest) SetEnvironment(v string) {
	o.Environment = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ReportNotificationRequest) GetTimestamp() time.Time {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportNotificationRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ReportNotificationRequest) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *ReportNotificationRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetType returns the Type field value
func (o *ReportNotificationRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReportNotificationRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReportNotificationRequest) SetType(v string) {
	o.Type = v
}

func (o ReportNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["environment"] = o.Environment
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableReportNotificationRequest struct {
	value *ReportNotificationRequest
	isSet bool
}

func (v NullableReportNotificationRequest) Get() *ReportNotificationRequest {
	return v.value
}

func (v *NullableReportNotificationRequest) Set(val *ReportNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportNotificationRequest(val *ReportNotificationRequest) *NullableReportNotificationRequest {
	return &NullableReportNotificationRequest{value: val, isSet: true}
}

func (v NullableReportNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *ReportNotificationRequest) isValidType() bool {
	var allowedEnumValues = []string{"balancePlatform.report.created"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
