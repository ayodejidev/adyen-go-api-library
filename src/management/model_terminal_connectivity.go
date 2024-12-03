/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the TerminalConnectivity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TerminalConnectivity{}

// TerminalConnectivity struct for TerminalConnectivity
type TerminalConnectivity struct {
	Bluetooth *TerminalConnectivityBluetooth `json:"bluetooth,omitempty"`
	Cellular  *TerminalConnectivityCellular  `json:"cellular,omitempty"`
	Ethernet  *TerminalConnectivityEthernet  `json:"ethernet,omitempty"`
	Wifi      *TerminalConnectivityWifi      `json:"wifi,omitempty"`
}

// NewTerminalConnectivity instantiates a new TerminalConnectivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminalConnectivity() *TerminalConnectivity {
	this := TerminalConnectivity{}
	return &this
}

// NewTerminalConnectivityWithDefaults instantiates a new TerminalConnectivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalConnectivityWithDefaults() *TerminalConnectivity {
	this := TerminalConnectivity{}
	return &this
}

// GetBluetooth returns the Bluetooth field value if set, zero value otherwise.
func (o *TerminalConnectivity) GetBluetooth() TerminalConnectivityBluetooth {
	if o == nil || common.IsNil(o.Bluetooth) {
		var ret TerminalConnectivityBluetooth
		return ret
	}
	return *o.Bluetooth
}

// GetBluetoothOk returns a tuple with the Bluetooth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivity) GetBluetoothOk() (*TerminalConnectivityBluetooth, bool) {
	if o == nil || common.IsNil(o.Bluetooth) {
		return nil, false
	}
	return o.Bluetooth, true
}

// HasBluetooth returns a boolean if a field has been set.
func (o *TerminalConnectivity) HasBluetooth() bool {
	if o != nil && !common.IsNil(o.Bluetooth) {
		return true
	}

	return false
}

// SetBluetooth gets a reference to the given TerminalConnectivityBluetooth and assigns it to the Bluetooth field.
func (o *TerminalConnectivity) SetBluetooth(v TerminalConnectivityBluetooth) {
	o.Bluetooth = &v
}

// GetCellular returns the Cellular field value if set, zero value otherwise.
func (o *TerminalConnectivity) GetCellular() TerminalConnectivityCellular {
	if o == nil || common.IsNil(o.Cellular) {
		var ret TerminalConnectivityCellular
		return ret
	}
	return *o.Cellular
}

// GetCellularOk returns a tuple with the Cellular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivity) GetCellularOk() (*TerminalConnectivityCellular, bool) {
	if o == nil || common.IsNil(o.Cellular) {
		return nil, false
	}
	return o.Cellular, true
}

// HasCellular returns a boolean if a field has been set.
func (o *TerminalConnectivity) HasCellular() bool {
	if o != nil && !common.IsNil(o.Cellular) {
		return true
	}

	return false
}

// SetCellular gets a reference to the given TerminalConnectivityCellular and assigns it to the Cellular field.
func (o *TerminalConnectivity) SetCellular(v TerminalConnectivityCellular) {
	o.Cellular = &v
}

// GetEthernet returns the Ethernet field value if set, zero value otherwise.
func (o *TerminalConnectivity) GetEthernet() TerminalConnectivityEthernet {
	if o == nil || common.IsNil(o.Ethernet) {
		var ret TerminalConnectivityEthernet
		return ret
	}
	return *o.Ethernet
}

// GetEthernetOk returns a tuple with the Ethernet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivity) GetEthernetOk() (*TerminalConnectivityEthernet, bool) {
	if o == nil || common.IsNil(o.Ethernet) {
		return nil, false
	}
	return o.Ethernet, true
}

// HasEthernet returns a boolean if a field has been set.
func (o *TerminalConnectivity) HasEthernet() bool {
	if o != nil && !common.IsNil(o.Ethernet) {
		return true
	}

	return false
}

// SetEthernet gets a reference to the given TerminalConnectivityEthernet and assigns it to the Ethernet field.
func (o *TerminalConnectivity) SetEthernet(v TerminalConnectivityEthernet) {
	o.Ethernet = &v
}

// GetWifi returns the Wifi field value if set, zero value otherwise.
func (o *TerminalConnectivity) GetWifi() TerminalConnectivityWifi {
	if o == nil || common.IsNil(o.Wifi) {
		var ret TerminalConnectivityWifi
		return ret
	}
	return *o.Wifi
}

// GetWifiOk returns a tuple with the Wifi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerminalConnectivity) GetWifiOk() (*TerminalConnectivityWifi, bool) {
	if o == nil || common.IsNil(o.Wifi) {
		return nil, false
	}
	return o.Wifi, true
}

// HasWifi returns a boolean if a field has been set.
func (o *TerminalConnectivity) HasWifi() bool {
	if o != nil && !common.IsNil(o.Wifi) {
		return true
	}

	return false
}

// SetWifi gets a reference to the given TerminalConnectivityWifi and assigns it to the Wifi field.
func (o *TerminalConnectivity) SetWifi(v TerminalConnectivityWifi) {
	o.Wifi = &v
}

func (o TerminalConnectivity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerminalConnectivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Bluetooth) {
		toSerialize["bluetooth"] = o.Bluetooth
	}
	if !common.IsNil(o.Cellular) {
		toSerialize["cellular"] = o.Cellular
	}
	if !common.IsNil(o.Ethernet) {
		toSerialize["ethernet"] = o.Ethernet
	}
	if !common.IsNil(o.Wifi) {
		toSerialize["wifi"] = o.Wifi
	}
	return toSerialize, nil
}

type NullableTerminalConnectivity struct {
	value *TerminalConnectivity
	isSet bool
}

func (v NullableTerminalConnectivity) Get() *TerminalConnectivity {
	return v.value
}

func (v *NullableTerminalConnectivity) Set(val *TerminalConnectivity) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminalConnectivity) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminalConnectivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminalConnectivity(val *TerminalConnectivity) *NullableTerminalConnectivity {
	return &NullableTerminalConnectivity{value: val, isSet: true}
}

func (v NullableTerminalConnectivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminalConnectivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
