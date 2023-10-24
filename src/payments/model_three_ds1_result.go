/*
Adyen Payment API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package payments

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// checks if the ThreeDS1Result type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ThreeDS1Result{}

// ThreeDS1Result struct for ThreeDS1Result
type ThreeDS1Result struct {
	// The cardholder authentication value (base64 encoded).
	Cavv *string `json:"cavv,omitempty"`
	// The CAVV algorithm used.
	CavvAlgorithm *string `json:"cavvAlgorithm,omitempty"`
	// 3D Secure Electronic Commerce Indicator (ECI).
	Eci *string `json:"eci,omitempty"`
	// The authentication response from the ACS.
	ThreeDAuthenticatedResponse *string `json:"threeDAuthenticatedResponse,omitempty"`
	// Whether 3D Secure was offered or not.
	ThreeDOfferedResponse *string `json:"threeDOfferedResponse,omitempty"`
	// A unique transaction identifier generated by the MPI on behalf of the merchant to identify the 3D Secure transaction, in `Base64` encoding.
	Xid *string `json:"xid,omitempty"`
}

// NewThreeDS1Result instantiates a new ThreeDS1Result object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreeDS1Result() *ThreeDS1Result {
	this := ThreeDS1Result{}
	return &this
}

// NewThreeDS1ResultWithDefaults instantiates a new ThreeDS1Result object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreeDS1ResultWithDefaults() *ThreeDS1Result {
	this := ThreeDS1Result{}
	return &this
}

// GetCavv returns the Cavv field value if set, zero value otherwise.
func (o *ThreeDS1Result) GetCavv() string {
	if o == nil || common.IsNil(o.Cavv) {
		var ret string
		return ret
	}
	return *o.Cavv
}

// GetCavvOk returns a tuple with the Cavv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS1Result) GetCavvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cavv) {
		return nil, false
	}
	return o.Cavv, true
}

// HasCavv returns a boolean if a field has been set.
func (o *ThreeDS1Result) HasCavv() bool {
	if o != nil && !common.IsNil(o.Cavv) {
		return true
	}

	return false
}

// SetCavv gets a reference to the given string and assigns it to the Cavv field.
func (o *ThreeDS1Result) SetCavv(v string) {
	o.Cavv = &v
}

// GetCavvAlgorithm returns the CavvAlgorithm field value if set, zero value otherwise.
func (o *ThreeDS1Result) GetCavvAlgorithm() string {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		var ret string
		return ret
	}
	return *o.CavvAlgorithm
}

// GetCavvAlgorithmOk returns a tuple with the CavvAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS1Result) GetCavvAlgorithmOk() (*string, bool) {
	if o == nil || common.IsNil(o.CavvAlgorithm) {
		return nil, false
	}
	return o.CavvAlgorithm, true
}

// HasCavvAlgorithm returns a boolean if a field has been set.
func (o *ThreeDS1Result) HasCavvAlgorithm() bool {
	if o != nil && !common.IsNil(o.CavvAlgorithm) {
		return true
	}

	return false
}

// SetCavvAlgorithm gets a reference to the given string and assigns it to the CavvAlgorithm field.
func (o *ThreeDS1Result) SetCavvAlgorithm(v string) {
	o.CavvAlgorithm = &v
}

// GetEci returns the Eci field value if set, zero value otherwise.
func (o *ThreeDS1Result) GetEci() string {
	if o == nil || common.IsNil(o.Eci) {
		var ret string
		return ret
	}
	return *o.Eci
}

// GetEciOk returns a tuple with the Eci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS1Result) GetEciOk() (*string, bool) {
	if o == nil || common.IsNil(o.Eci) {
		return nil, false
	}
	return o.Eci, true
}

// HasEci returns a boolean if a field has been set.
func (o *ThreeDS1Result) HasEci() bool {
	if o != nil && !common.IsNil(o.Eci) {
		return true
	}

	return false
}

// SetEci gets a reference to the given string and assigns it to the Eci field.
func (o *ThreeDS1Result) SetEci(v string) {
	o.Eci = &v
}

// GetThreeDAuthenticatedResponse returns the ThreeDAuthenticatedResponse field value if set, zero value otherwise.
func (o *ThreeDS1Result) GetThreeDAuthenticatedResponse() string {
	if o == nil || common.IsNil(o.ThreeDAuthenticatedResponse) {
		var ret string
		return ret
	}
	return *o.ThreeDAuthenticatedResponse
}

// GetThreeDAuthenticatedResponseOk returns a tuple with the ThreeDAuthenticatedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS1Result) GetThreeDAuthenticatedResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDAuthenticatedResponse) {
		return nil, false
	}
	return o.ThreeDAuthenticatedResponse, true
}

// HasThreeDAuthenticatedResponse returns a boolean if a field has been set.
func (o *ThreeDS1Result) HasThreeDAuthenticatedResponse() bool {
	if o != nil && !common.IsNil(o.ThreeDAuthenticatedResponse) {
		return true
	}

	return false
}

// SetThreeDAuthenticatedResponse gets a reference to the given string and assigns it to the ThreeDAuthenticatedResponse field.
func (o *ThreeDS1Result) SetThreeDAuthenticatedResponse(v string) {
	o.ThreeDAuthenticatedResponse = &v
}

// GetThreeDOfferedResponse returns the ThreeDOfferedResponse field value if set, zero value otherwise.
func (o *ThreeDS1Result) GetThreeDOfferedResponse() string {
	if o == nil || common.IsNil(o.ThreeDOfferedResponse) {
		var ret string
		return ret
	}
	return *o.ThreeDOfferedResponse
}

// GetThreeDOfferedResponseOk returns a tuple with the ThreeDOfferedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS1Result) GetThreeDOfferedResponseOk() (*string, bool) {
	if o == nil || common.IsNil(o.ThreeDOfferedResponse) {
		return nil, false
	}
	return o.ThreeDOfferedResponse, true
}

// HasThreeDOfferedResponse returns a boolean if a field has been set.
func (o *ThreeDS1Result) HasThreeDOfferedResponse() bool {
	if o != nil && !common.IsNil(o.ThreeDOfferedResponse) {
		return true
	}

	return false
}

// SetThreeDOfferedResponse gets a reference to the given string and assigns it to the ThreeDOfferedResponse field.
func (o *ThreeDS1Result) SetThreeDOfferedResponse(v string) {
	o.ThreeDOfferedResponse = &v
}

// GetXid returns the Xid field value if set, zero value otherwise.
func (o *ThreeDS1Result) GetXid() string {
	if o == nil || common.IsNil(o.Xid) {
		var ret string
		return ret
	}
	return *o.Xid
}

// GetXidOk returns a tuple with the Xid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreeDS1Result) GetXidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Xid) {
		return nil, false
	}
	return o.Xid, true
}

// HasXid returns a boolean if a field has been set.
func (o *ThreeDS1Result) HasXid() bool {
	if o != nil && !common.IsNil(o.Xid) {
		return true
	}

	return false
}

// SetXid gets a reference to the given string and assigns it to the Xid field.
func (o *ThreeDS1Result) SetXid(v string) {
	o.Xid = &v
}

func (o ThreeDS1Result) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreeDS1Result) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Cavv) {
		toSerialize["cavv"] = o.Cavv
	}
	if !common.IsNil(o.CavvAlgorithm) {
		toSerialize["cavvAlgorithm"] = o.CavvAlgorithm
	}
	if !common.IsNil(o.Eci) {
		toSerialize["eci"] = o.Eci
	}
	if !common.IsNil(o.ThreeDAuthenticatedResponse) {
		toSerialize["threeDAuthenticatedResponse"] = o.ThreeDAuthenticatedResponse
	}
	if !common.IsNil(o.ThreeDOfferedResponse) {
		toSerialize["threeDOfferedResponse"] = o.ThreeDOfferedResponse
	}
	if !common.IsNil(o.Xid) {
		toSerialize["xid"] = o.Xid
	}
	return toSerialize, nil
}

type NullableThreeDS1Result struct {
	value *ThreeDS1Result
	isSet bool
}

func (v NullableThreeDS1Result) Get() *ThreeDS1Result {
	return v.value
}

func (v *NullableThreeDS1Result) Set(val *ThreeDS1Result) {
	v.value = val
	v.isSet = true
}

func (v NullableThreeDS1Result) IsSet() bool {
	return v.isSet
}

func (v *NullableThreeDS1Result) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreeDS1Result(val *ThreeDS1Result) *NullableThreeDS1Result {
	return &NullableThreeDS1Result{value: val, isSet: true}
}

func (v NullableThreeDS1Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreeDS1Result) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
