/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the ApplicationInfo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ApplicationInfo{}

// ApplicationInfo struct for ApplicationInfo
type ApplicationInfo struct {
	AdyenLibrary             *CommonField              `json:"adyenLibrary,omitempty"`
	AdyenPaymentSource       *CommonField              `json:"adyenPaymentSource,omitempty"`
	ExternalPlatform         *ExternalPlatform         `json:"externalPlatform,omitempty"`
	MerchantApplication      *CommonField              `json:"merchantApplication,omitempty"`
	MerchantDevice           *MerchantDevice           `json:"merchantDevice,omitempty"`
	ShopperInteractionDevice *ShopperInteractionDevice `json:"shopperInteractionDevice,omitempty"`
}

// NewApplicationInfo instantiates a new ApplicationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationInfo() *ApplicationInfo {
	this := ApplicationInfo{}
	return &this
}

// NewApplicationInfoWithDefaults instantiates a new ApplicationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationInfoWithDefaults() *ApplicationInfo {
	this := ApplicationInfo{}
	return &this
}

// GetAdyenLibrary returns the AdyenLibrary field value if set, zero value otherwise.
func (o *ApplicationInfo) GetAdyenLibrary() CommonField {
	if o == nil || common.IsNil(o.AdyenLibrary) {
		var ret CommonField
		return ret
	}
	return *o.AdyenLibrary
}

// GetAdyenLibraryOk returns a tuple with the AdyenLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInfo) GetAdyenLibraryOk() (*CommonField, bool) {
	if o == nil || common.IsNil(o.AdyenLibrary) {
		return nil, false
	}
	return o.AdyenLibrary, true
}

// HasAdyenLibrary returns a boolean if a field has been set.
func (o *ApplicationInfo) HasAdyenLibrary() bool {
	if o != nil && !common.IsNil(o.AdyenLibrary) {
		return true
	}

	return false
}

// SetAdyenLibrary gets a reference to the given CommonField and assigns it to the AdyenLibrary field.
func (o *ApplicationInfo) SetAdyenLibrary(v CommonField) {
	o.AdyenLibrary = &v
}

// GetAdyenPaymentSource returns the AdyenPaymentSource field value if set, zero value otherwise.
func (o *ApplicationInfo) GetAdyenPaymentSource() CommonField {
	if o == nil || common.IsNil(o.AdyenPaymentSource) {
		var ret CommonField
		return ret
	}
	return *o.AdyenPaymentSource
}

// GetAdyenPaymentSourceOk returns a tuple with the AdyenPaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInfo) GetAdyenPaymentSourceOk() (*CommonField, bool) {
	if o == nil || common.IsNil(o.AdyenPaymentSource) {
		return nil, false
	}
	return o.AdyenPaymentSource, true
}

// HasAdyenPaymentSource returns a boolean if a field has been set.
func (o *ApplicationInfo) HasAdyenPaymentSource() bool {
	if o != nil && !common.IsNil(o.AdyenPaymentSource) {
		return true
	}

	return false
}

// SetAdyenPaymentSource gets a reference to the given CommonField and assigns it to the AdyenPaymentSource field.
func (o *ApplicationInfo) SetAdyenPaymentSource(v CommonField) {
	o.AdyenPaymentSource = &v
}

// GetExternalPlatform returns the ExternalPlatform field value if set, zero value otherwise.
func (o *ApplicationInfo) GetExternalPlatform() ExternalPlatform {
	if o == nil || common.IsNil(o.ExternalPlatform) {
		var ret ExternalPlatform
		return ret
	}
	return *o.ExternalPlatform
}

// GetExternalPlatformOk returns a tuple with the ExternalPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInfo) GetExternalPlatformOk() (*ExternalPlatform, bool) {
	if o == nil || common.IsNil(o.ExternalPlatform) {
		return nil, false
	}
	return o.ExternalPlatform, true
}

// HasExternalPlatform returns a boolean if a field has been set.
func (o *ApplicationInfo) HasExternalPlatform() bool {
	if o != nil && !common.IsNil(o.ExternalPlatform) {
		return true
	}

	return false
}

// SetExternalPlatform gets a reference to the given ExternalPlatform and assigns it to the ExternalPlatform field.
func (o *ApplicationInfo) SetExternalPlatform(v ExternalPlatform) {
	o.ExternalPlatform = &v
}

// GetMerchantApplication returns the MerchantApplication field value if set, zero value otherwise.
func (o *ApplicationInfo) GetMerchantApplication() CommonField {
	if o == nil || common.IsNil(o.MerchantApplication) {
		var ret CommonField
		return ret
	}
	return *o.MerchantApplication
}

// GetMerchantApplicationOk returns a tuple with the MerchantApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInfo) GetMerchantApplicationOk() (*CommonField, bool) {
	if o == nil || common.IsNil(o.MerchantApplication) {
		return nil, false
	}
	return o.MerchantApplication, true
}

// HasMerchantApplication returns a boolean if a field has been set.
func (o *ApplicationInfo) HasMerchantApplication() bool {
	if o != nil && !common.IsNil(o.MerchantApplication) {
		return true
	}

	return false
}

// SetMerchantApplication gets a reference to the given CommonField and assigns it to the MerchantApplication field.
func (o *ApplicationInfo) SetMerchantApplication(v CommonField) {
	o.MerchantApplication = &v
}

// GetMerchantDevice returns the MerchantDevice field value if set, zero value otherwise.
func (o *ApplicationInfo) GetMerchantDevice() MerchantDevice {
	if o == nil || common.IsNil(o.MerchantDevice) {
		var ret MerchantDevice
		return ret
	}
	return *o.MerchantDevice
}

// GetMerchantDeviceOk returns a tuple with the MerchantDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInfo) GetMerchantDeviceOk() (*MerchantDevice, bool) {
	if o == nil || common.IsNil(o.MerchantDevice) {
		return nil, false
	}
	return o.MerchantDevice, true
}

// HasMerchantDevice returns a boolean if a field has been set.
func (o *ApplicationInfo) HasMerchantDevice() bool {
	if o != nil && !common.IsNil(o.MerchantDevice) {
		return true
	}

	return false
}

// SetMerchantDevice gets a reference to the given MerchantDevice and assigns it to the MerchantDevice field.
func (o *ApplicationInfo) SetMerchantDevice(v MerchantDevice) {
	o.MerchantDevice = &v
}

// GetShopperInteractionDevice returns the ShopperInteractionDevice field value if set, zero value otherwise.
func (o *ApplicationInfo) GetShopperInteractionDevice() ShopperInteractionDevice {
	if o == nil || common.IsNil(o.ShopperInteractionDevice) {
		var ret ShopperInteractionDevice
		return ret
	}
	return *o.ShopperInteractionDevice
}

// GetShopperInteractionDeviceOk returns a tuple with the ShopperInteractionDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationInfo) GetShopperInteractionDeviceOk() (*ShopperInteractionDevice, bool) {
	if o == nil || common.IsNil(o.ShopperInteractionDevice) {
		return nil, false
	}
	return o.ShopperInteractionDevice, true
}

// HasShopperInteractionDevice returns a boolean if a field has been set.
func (o *ApplicationInfo) HasShopperInteractionDevice() bool {
	if o != nil && !common.IsNil(o.ShopperInteractionDevice) {
		return true
	}

	return false
}

// SetShopperInteractionDevice gets a reference to the given ShopperInteractionDevice and assigns it to the ShopperInteractionDevice field.
func (o *ApplicationInfo) SetShopperInteractionDevice(v ShopperInteractionDevice) {
	o.ShopperInteractionDevice = &v
}

func (o ApplicationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AdyenLibrary) {
		toSerialize["adyenLibrary"] = o.AdyenLibrary
	}
	if !common.IsNil(o.AdyenPaymentSource) {
		toSerialize["adyenPaymentSource"] = o.AdyenPaymentSource
	}
	if !common.IsNil(o.ExternalPlatform) {
		toSerialize["externalPlatform"] = o.ExternalPlatform
	}
	if !common.IsNil(o.MerchantApplication) {
		toSerialize["merchantApplication"] = o.MerchantApplication
	}
	if !common.IsNil(o.MerchantDevice) {
		toSerialize["merchantDevice"] = o.MerchantDevice
	}
	if !common.IsNil(o.ShopperInteractionDevice) {
		toSerialize["shopperInteractionDevice"] = o.ShopperInteractionDevice
	}
	return toSerialize, nil
}

type NullableApplicationInfo struct {
	value *ApplicationInfo
	isSet bool
}

func (v NullableApplicationInfo) Get() *ApplicationInfo {
	return v.value
}

func (v *NullableApplicationInfo) Set(val *ApplicationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationInfo(val *ApplicationInfo) *NullableApplicationInfo {
	return &NullableApplicationInfo{value: val, isSet: true}
}

func (v NullableApplicationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
