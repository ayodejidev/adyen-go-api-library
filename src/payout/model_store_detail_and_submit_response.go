/*
Adyen Payout API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the StoreDetailAndSubmitResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StoreDetailAndSubmitResponse{}

// StoreDetailAndSubmitResponse struct for StoreDetailAndSubmitResponse
type StoreDetailAndSubmitResponse struct {
	// This field contains additional data, which may be returned in a particular response.
	AdditionalData *map[string]string `json:"additionalData,omitempty"`
	// A new reference to uniquely identify this request.
	PspReference string `json:"pspReference"`
	// In case of refusal, an informational message for the reason.
	RefusalReason *string `json:"refusalReason,omitempty"`
	// The response:  * In case of success is payout-submit-received. * In case of an error, an informational message is returned.
	ResultCode string `json:"resultCode"`
}

// NewStoreDetailAndSubmitResponse instantiates a new StoreDetailAndSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDetailAndSubmitResponse(pspReference string, resultCode string) *StoreDetailAndSubmitResponse {
	this := StoreDetailAndSubmitResponse{}
	this.PspReference = pspReference
	this.ResultCode = resultCode
	return &this
}

// NewStoreDetailAndSubmitResponseWithDefaults instantiates a new StoreDetailAndSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDetailAndSubmitResponseWithDefaults() *StoreDetailAndSubmitResponse {
	this := StoreDetailAndSubmitResponse{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitResponse) GetAdditionalData() map[string]string {
	if o == nil || common.IsNil(o.AdditionalData) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitResponse) GetAdditionalDataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitResponse) HasAdditionalData() bool {
	if o != nil && !common.IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]string and assigns it to the AdditionalData field.
func (o *StoreDetailAndSubmitResponse) SetAdditionalData(v map[string]string) {
	o.AdditionalData = &v
}

// GetPspReference returns the PspReference field value
func (o *StoreDetailAndSubmitResponse) GetPspReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PspReference, true
}

// SetPspReference sets field value
func (o *StoreDetailAndSubmitResponse) SetPspReference(v string) {
	o.PspReference = v
}

// GetRefusalReason returns the RefusalReason field value if set, zero value otherwise.
func (o *StoreDetailAndSubmitResponse) GetRefusalReason() string {
	if o == nil || common.IsNil(o.RefusalReason) {
		var ret string
		return ret
	}
	return *o.RefusalReason
}

// GetRefusalReasonOk returns a tuple with the RefusalReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitResponse) GetRefusalReasonOk() (*string, bool) {
	if o == nil || common.IsNil(o.RefusalReason) {
		return nil, false
	}
	return o.RefusalReason, true
}

// HasRefusalReason returns a boolean if a field has been set.
func (o *StoreDetailAndSubmitResponse) HasRefusalReason() bool {
	if o != nil && !common.IsNil(o.RefusalReason) {
		return true
	}

	return false
}

// SetRefusalReason gets a reference to the given string and assigns it to the RefusalReason field.
func (o *StoreDetailAndSubmitResponse) SetRefusalReason(v string) {
	o.RefusalReason = &v
}

// GetResultCode returns the ResultCode field value
func (o *StoreDetailAndSubmitResponse) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *StoreDetailAndSubmitResponse) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *StoreDetailAndSubmitResponse) SetResultCode(v string) {
	o.ResultCode = v
}

func (o StoreDetailAndSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDetailAndSubmitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdditionalData) {
		toSerialize["additionalData"] = o.AdditionalData
	}
	toSerialize["pspReference"] = o.PspReference
	if !common.IsNil(o.RefusalReason) {
		toSerialize["refusalReason"] = o.RefusalReason
	}
	toSerialize["resultCode"] = o.ResultCode
	return toSerialize, nil
}

type NullableStoreDetailAndSubmitResponse struct {
	value *StoreDetailAndSubmitResponse
	isSet bool
}

func (v NullableStoreDetailAndSubmitResponse) Get() *StoreDetailAndSubmitResponse {
	return v.value
}

func (v *NullableStoreDetailAndSubmitResponse) Set(val *StoreDetailAndSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDetailAndSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDetailAndSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDetailAndSubmitResponse(val *StoreDetailAndSubmitResponse) *NullableStoreDetailAndSubmitResponse {
	return &NullableStoreDetailAndSubmitResponse{value: val, isSet: true}
}

func (v NullableStoreDetailAndSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDetailAndSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
