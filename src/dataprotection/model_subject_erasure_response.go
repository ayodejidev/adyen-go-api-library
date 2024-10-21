/*
Adyen Data Protection API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dataprotection

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the SubjectErasureResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SubjectErasureResponse{}

// SubjectErasureResponse struct for SubjectErasureResponse
type SubjectErasureResponse struct {
	// The result of this operation.
	Result *string `json:"result,omitempty"`
}

// NewSubjectErasureResponse instantiates a new SubjectErasureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectErasureResponse() *SubjectErasureResponse {
	this := SubjectErasureResponse{}
	return &this
}

// NewSubjectErasureResponseWithDefaults instantiates a new SubjectErasureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectErasureResponseWithDefaults() *SubjectErasureResponse {
	this := SubjectErasureResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SubjectErasureResponse) GetResult() string {
	if o == nil || common.IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectErasureResponse) GetResultOk() (*string, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SubjectErasureResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *SubjectErasureResponse) SetResult(v string) {
	o.Result = &v
}

func (o SubjectErasureResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectErasureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableSubjectErasureResponse struct {
	value *SubjectErasureResponse
	isSet bool
}

func (v NullableSubjectErasureResponse) Get() *SubjectErasureResponse {
	return v.value
}

func (v *NullableSubjectErasureResponse) Set(val *SubjectErasureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectErasureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectErasureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectErasureResponse(val *SubjectErasureResponse) *NullableSubjectErasureResponse {
	return &NullableSubjectErasureResponse{value: val, isSet: true}
}

func (v NullableSubjectErasureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectErasureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *SubjectErasureResponse) isValidResult() bool {
	var allowedEnumValues = []string{"ACTIVE_RECURRING_TOKEN_EXISTS", "ALREADY_PROCESSED", "PAYMENT_NOT_FOUND", "SUCCESS"}
	for _, allowed := range allowedEnumValues {
		if o.GetResult() == allowed {
			return true
		}
	}
	return false
}
