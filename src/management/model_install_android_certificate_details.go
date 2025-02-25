/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the InstallAndroidCertificateDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &InstallAndroidCertificateDetails{}

// InstallAndroidCertificateDetails struct for InstallAndroidCertificateDetails
type InstallAndroidCertificateDetails struct {
	// The unique identifier of the certificate to be installed.
	CertificateId *string `json:"certificateId,omitempty"`
	// Type of terminal action: Install an Android certificate.
	Type *string `json:"type,omitempty"`
}

// NewInstallAndroidCertificateDetails instantiates a new InstallAndroidCertificateDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallAndroidCertificateDetails() *InstallAndroidCertificateDetails {
	this := InstallAndroidCertificateDetails{}
	var type_ string = "InstallAndroidCertificate"
	this.Type = &type_
	return &this
}

// NewInstallAndroidCertificateDetailsWithDefaults instantiates a new InstallAndroidCertificateDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallAndroidCertificateDetailsWithDefaults() *InstallAndroidCertificateDetails {
	this := InstallAndroidCertificateDetails{}
	var type_ string = "InstallAndroidCertificate"
	this.Type = &type_
	return &this
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise.
func (o *InstallAndroidCertificateDetails) GetCertificateId() string {
	if o == nil || common.IsNil(o.CertificateId) {
		var ret string
		return ret
	}
	return *o.CertificateId
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallAndroidCertificateDetails) GetCertificateIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CertificateId) {
		return nil, false
	}
	return o.CertificateId, true
}

// HasCertificateId returns a boolean if a field has been set.
func (o *InstallAndroidCertificateDetails) HasCertificateId() bool {
	if o != nil && !common.IsNil(o.CertificateId) {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given string and assigns it to the CertificateId field.
func (o *InstallAndroidCertificateDetails) SetCertificateId(v string) {
	o.CertificateId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InstallAndroidCertificateDetails) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallAndroidCertificateDetails) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InstallAndroidCertificateDetails) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InstallAndroidCertificateDetails) SetType(v string) {
	o.Type = &v
}

func (o InstallAndroidCertificateDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallAndroidCertificateDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CertificateId) {
		toSerialize["certificateId"] = o.CertificateId
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInstallAndroidCertificateDetails struct {
	value *InstallAndroidCertificateDetails
	isSet bool
}

func (v NullableInstallAndroidCertificateDetails) Get() *InstallAndroidCertificateDetails {
	return v.value
}

func (v *NullableInstallAndroidCertificateDetails) Set(val *InstallAndroidCertificateDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallAndroidCertificateDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallAndroidCertificateDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallAndroidCertificateDetails(val *InstallAndroidCertificateDetails) *NullableInstallAndroidCertificateDetails {
	return &NullableInstallAndroidCertificateDetails{value: val, isSet: true}
}

func (v NullableInstallAndroidCertificateDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallAndroidCertificateDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *InstallAndroidCertificateDetails) isValidType() bool {
	var allowedEnumValues = []string{"InstallAndroidCertificate"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
