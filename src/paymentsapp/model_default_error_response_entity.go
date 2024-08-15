/*
Payments App API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentsapp

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the DefaultErrorResponseEntity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DefaultErrorResponseEntity{}

// DefaultErrorResponseEntity Standardized error response following RFC-7807 format
type DefaultErrorResponseEntity struct {
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// Unique business error code.
	ErrorCode *string `json:"errorCode,omitempty"`
	// A URI that identifies the specific occurrence of the problem if applicable.
	Instance *string `json:"instance,omitempty"`
	// Array of fields with validation errors when applicable.
	InvalidFields []InvalidField `json:"invalidFields,omitempty"`
	// The unique reference for the request.
	RequestId *string `json:"requestId,omitempty"`
	// The HTTP status code.
	Status *int32 `json:"status,omitempty"`
	// A short, human-readable summary of the problem type.
	Title *string `json:"title,omitempty"`
	// A URI that identifies the validation error type. It points to human-readable documentation for the problem type.
	Type *string `json:"type,omitempty"`
}

// NewDefaultErrorResponseEntity instantiates a new DefaultErrorResponseEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultErrorResponseEntity() *DefaultErrorResponseEntity {
	this := DefaultErrorResponseEntity{}
	return &this
}

// NewDefaultErrorResponseEntityWithDefaults instantiates a new DefaultErrorResponseEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultErrorResponseEntityWithDefaults() *DefaultErrorResponseEntity {
	this := DefaultErrorResponseEntity{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetDetail() string {
	if o == nil || common.IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasDetail() bool {
	if o != nil && !common.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *DefaultErrorResponseEntity) SetDetail(v string) {
	o.Detail = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetErrorCode() string {
	if o == nil || common.IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetErrorCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasErrorCode() bool {
	if o != nil && !common.IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *DefaultErrorResponseEntity) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetInstance() string {
	if o == nil || common.IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetInstanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasInstance() bool {
	if o != nil && !common.IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *DefaultErrorResponseEntity) SetInstance(v string) {
	o.Instance = &v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetInvalidFields() []InvalidField {
	if o == nil || common.IsNil(o.InvalidFields) {
		var ret []InvalidField
		return ret
	}
	return o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetInvalidFieldsOk() ([]InvalidField, bool) {
	if o == nil || common.IsNil(o.InvalidFields) {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasInvalidFields() bool {
	if o != nil && !common.IsNil(o.InvalidFields) {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []InvalidField and assigns it to the InvalidFields field.
func (o *DefaultErrorResponseEntity) SetInvalidFields(v []InvalidField) {
	o.InvalidFields = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetRequestId() string {
	if o == nil || common.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetRequestIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasRequestId() bool {
	if o != nil && !common.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DefaultErrorResponseEntity) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetStatus() int32 {
	if o == nil || common.IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetStatusOk() (*int32, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *DefaultErrorResponseEntity) SetStatus(v int32) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetTitle() string {
	if o == nil || common.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasTitle() bool {
	if o != nil && !common.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *DefaultErrorResponseEntity) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DefaultErrorResponseEntity) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultErrorResponseEntity) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DefaultErrorResponseEntity) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DefaultErrorResponseEntity) SetType(v string) {
	o.Type = &v
}

func (o DefaultErrorResponseEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultErrorResponseEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !common.IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !common.IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !common.IsNil(o.InvalidFields) {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if !common.IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDefaultErrorResponseEntity struct {
	value *DefaultErrorResponseEntity
	isSet bool
}

func (v NullableDefaultErrorResponseEntity) Get() *DefaultErrorResponseEntity {
	return v.value
}

func (v *NullableDefaultErrorResponseEntity) Set(val *DefaultErrorResponseEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultErrorResponseEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultErrorResponseEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultErrorResponseEntity(val *DefaultErrorResponseEntity) *NullableDefaultErrorResponseEntity {
	return &NullableDefaultErrorResponseEntity{value: val, isSet: true}
}

func (v NullableDefaultErrorResponseEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultErrorResponseEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
