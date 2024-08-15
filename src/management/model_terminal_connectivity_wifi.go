/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the TerminalConnectivityWifi type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalConnectivityWifi{}

// TerminalConnectivityWifi struct for TerminalConnectivityWifi
type TerminalConnectivityWifi struct {
	// The terminal's IP address in the Wi-Fi network.
	IpAddress *string `json:"ipAddress,omitempty"`
	// The terminal's MAC address in the Wi-Fi network.
	MacAddress *string `json:"macAddress,omitempty"`
	// The SSID of the Wi-Fi network that the terminal is connected to.
	Ssid *string `json:"ssid,omitempty"`
}

// NewTerminalConnectivityWifi instantiates a new TerminalConnectivityWifi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalConnectivityWifi() *TerminalConnectivityWifi {
	this := TerminalConnectivityWifi{}
	return &this
}

// NewTerminalConnectivityWifiWithDefaults instantiates a new TerminalConnectivityWifi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalConnectivityWifiWithDefaults() *TerminalConnectivityWifi {
	this := TerminalConnectivityWifi{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *TerminalConnectivityWifi) GetIpAddress() string {
	if o == nil || common.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivityWifi) GetIpAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *TerminalConnectivityWifi) HasIpAddress() bool {
	if o != nil && !common.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *TerminalConnectivityWifi) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *TerminalConnectivityWifi) GetMacAddress() string {
	if o == nil || common.IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivityWifi) GetMacAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *TerminalConnectivityWifi) HasMacAddress() bool {
	if o != nil && !common.IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *TerminalConnectivityWifi) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *TerminalConnectivityWifi) GetSsid() string {
	if o == nil || common.IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivityWifi) GetSsidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *TerminalConnectivityWifi) HasSsid() bool {
	if o != nil && !common.IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *TerminalConnectivityWifi) SetSsid(v string) {
	o.Ssid = &v
}

func (o TerminalConnectivityWifi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalConnectivityWifi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !common.IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !common.IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	return toSerialize, nil
}

type NullableTerminalConnectivityWifi struct {
	value *TerminalConnectivityWifi
	isSet bool
}

func (v NullableTerminalConnectivityWifi) Get() *TerminalConnectivityWifi {
	return v.value
}

func (v *NullableTerminalConnectivityWifi) Set(val *TerminalConnectivityWifi) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalConnectivityWifi) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalConnectivityWifi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalConnectivityWifi(val *TerminalConnectivityWifi) *NullableTerminalConnectivityWifi {
	return &NullableTerminalConnectivityWifi{value: val, isSet: true}
}

func (v NullableTerminalConnectivityWifi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalConnectivityWifi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
