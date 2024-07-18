/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the UninstallAndroidCertificateDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UninstallAndroidCertificateDetails{}

// UninstallAndroidCertificateDetails struct for UninstallAndroidCertificateDetails
type UninstallAndroidCertificateDetails struct {
	// The unique identifier of the certificate to be uninstalled.
	CertificateId *string `json:"certificateId,omitempty"`
	// Type of terminal action: Uninstall an Android certificate.
	Type *string `json:"type,omitempty"`
}

// NewUninstallAndroidCertificateDetails instantiates a new UninstallAndroidCertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUninstallAndroidCertificateDetails() *UninstallAndroidCertificateDetails {
	this := UninstallAndroidCertificateDetails{}
	var type_ string = "UninstallAndroidCertificate"
	this.Type = &type_
	return &this
}

// NewUninstallAndroidCertificateDetailsWithDefaults instantiates a new UninstallAndroidCertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUninstallAndroidCertificateDetailsWithDefaults() *UninstallAndroidCertificateDetails {
	this := UninstallAndroidCertificateDetails{}
	var type_ string = "UninstallAndroidCertificate"
	this.Type = &type_
	return &this
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *UninstallAndroidCertificateDetails) GetCertificateId() string {
	if o == nil || common.IsNil(o.CertificateId) {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UninstallAndroidCertificateDetails) GetCertificateIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CertificateId) {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *UninstallAndroidCertificateDetails) HasCertificateId() bool {
	if o != nil && !common.IsNil(o.CertificateId) {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *UninstallAndroidCertificateDetails) SetCertificateId(v string) {
	o.CertificateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UninstallAndroidCertificateDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UninstallAndroidCertificateDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UninstallAndroidCertificateDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UninstallAndroidCertificateDetails) SetType(v string) {
	o.Type = &v
}

func (o UninstallAndroidCertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UninstallAndroidCertificateDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CertificateId) {
		toSerialize["certificateId"] = o.CertificateId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUninstallAndroidCertificateDetails struct {
	value *UninstallAndroidCertificateDetails
	isSet bool
}

func (v NullableUninstallAndroidCertificateDetails) Get() *UninstallAndroidCertificateDetails {
	return v.value
}

func (v *NullableUninstallAndroidCertificateDetails) Set(val *UninstallAndroidCertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUninstallAndroidCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUninstallAndroidCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUninstallAndroidCertificateDetails(val *UninstallAndroidCertificateDetails) *NullableUninstallAndroidCertificateDetails {
	return &NullableUninstallAndroidCertificateDetails{value: val, isSet: true}
}

func (v NullableUninstallAndroidCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUninstallAndroidCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *UninstallAndroidCertificateDetails) isValidType() bool {
	var allowedEnumValues = []string{"UninstallAndroidCertificate"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
