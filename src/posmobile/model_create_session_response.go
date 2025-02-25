/*
POS Mobile API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posmobile

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the CreateSessionResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateSessionResponse{}

// CreateSessionResponse struct for CreateSessionResponse
type CreateSessionResponse struct {
	// The unique identifier of the session.
	Id *string `json:"id,omitempty"`
	// The unique identifier of the SDK installation. If you create the [Terminal API](https://docs.adyen.com/point-of-sale/design-your-integration/terminal-api/) transaction request on your backend, use this as the `POIID` in the `MessageHeader` of the request.
	InstallationId *string `json:"installationId,omitempty"`
	// The unique identifier of your merchant account.
	MerchantAccount *string `json:"merchantAccount,omitempty"`
	// The data that the SDK uses to authenticate responses from the Adyen payments platform. Pass this value to your POS app.
	SdkData *string `json:"sdkData,omitempty"`
	// The unique identifier of the store that you want to process transactions for.
	Store *string `json:"store,omitempty"`
}

// NewCreateSessionResponse instantiates a new CreateSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSessionResponse() *CreateSessionResponse {
	this := CreateSessionResponse{}
	return &this
}

// NewCreateSessionResponseWithDefaults instantiates a new CreateSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSessionResponseWithDefaults() *CreateSessionResponse {
	this := CreateSessionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateSessionResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateSessionResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateSessionResponse) SetId(v string) {
	o.Id = &v
}

// GetInstallationId returns the InstallationId field value if set, zero value otherwise.
func (o *CreateSessionResponse) GetInstallationId() string {
	if o == nil || common.IsNil(o.InstallationId) {
		var ret string
		return ret
	}
	return *o.InstallationId
}

// GetInstallationIdOk returns a tuple with the InstallationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionResponse) GetInstallationIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.InstallationId) {
		return nil, false
	}
	return o.InstallationId, true
}

// HasInstallationId returns a boolean if a field has been set.
func (o *CreateSessionResponse) HasInstallationId() bool {
	if o != nil && !common.IsNil(o.InstallationId) {
		return true
	}

	return false
}

// SetInstallationId gets a reference to the given string and assigns it to the InstallationId field.
func (o *CreateSessionResponse) SetInstallationId(v string) {
	o.InstallationId = &v
}

// GetMerchantAccount returns the MerchantAccount field value if set, zero value otherwise.
func (o *CreateSessionResponse) GetMerchantAccount() string {
	if o == nil || common.IsNil(o.MerchantAccount) {
		var ret string
		return ret
	}
	return *o.MerchantAccount
}

// GetMerchantAccountOk returns a tuple with the MerchantAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionResponse) GetMerchantAccountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccount) {
		return nil, false
	}
	return o.MerchantAccount, true
}

// HasMerchantAccount returns a boolean if a field has been set.
func (o *CreateSessionResponse) HasMerchantAccount() bool {
	if o != nil && !common.IsNil(o.MerchantAccount) {
		return true
	}

	return false
}

// SetMerchantAccount gets a reference to the given string and assigns it to the MerchantAccount field.
func (o *CreateSessionResponse) SetMerchantAccount(v string) {
	o.MerchantAccount = &v
}

// GetSdkData returns the SdkData field value if set, zero value otherwise.
func (o *CreateSessionResponse) GetSdkData() string {
	if o == nil || common.IsNil(o.SdkData) {
		var ret string
		return ret
	}
	return *o.SdkData
}

// GetSdkDataOk returns a tuple with the SdkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionResponse) GetSdkDataOk() (*string, bool) {
	if o == nil || common.IsNil(o.SdkData) {
		return nil, false
	}
	return o.SdkData, true
}

// HasSdkData returns a boolean if a field has been set.
func (o *CreateSessionResponse) HasSdkData() bool {
	if o != nil && !common.IsNil(o.SdkData) {
		return true
	}

	return false
}

// SetSdkData gets a reference to the given string and assigns it to the SdkData field.
func (o *CreateSessionResponse) SetSdkData(v string) {
	o.SdkData = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *CreateSessionResponse) GetStore() string {
	if o == nil || common.IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionResponse) GetStoreOk() (*string, bool) {
	if o == nil || common.IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *CreateSessionResponse) HasStore() bool {
	if o != nil && !common.IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *CreateSessionResponse) SetStore(v string) {
	o.Store = &v
}

func (o CreateSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.InstallationId) {
		toSerialize["installationId"] = o.InstallationId
	}
	if !common.IsNil(o.MerchantAccount) {
		toSerialize["merchantAccount"] = o.MerchantAccount
	}
	if !common.IsNil(o.SdkData) {
		toSerialize["sdkData"] = o.SdkData
	}
	if !common.IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	return toSerialize, nil
}

type NullableCreateSessionResponse struct {
	value *CreateSessionResponse
	isSet bool
}

func (v NullableCreateSessionResponse) Get() *CreateSessionResponse {
	return v.value
}

func (v *NullableCreateSessionResponse) Set(val *CreateSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSessionResponse(val *CreateSessionResponse) *NullableCreateSessionResponse {
	return &NullableCreateSessionResponse{value: val, isSet: true}
}

func (v NullableCreateSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
