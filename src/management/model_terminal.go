/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the Terminal type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Terminal{}

// Terminal struct for Terminal
type Terminal struct {
	Assignment   *TerminalAssignment   `json:"assignment,omitempty"`
	Connectivity *TerminalConnectivity `json:"connectivity,omitempty"`
	// The software release currently in use on the terminal.
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`
	// The unique identifier of the terminal.
	Id *string `json:"id,omitempty"`
	// Date and time of the last activity on the terminal. Not included when the last activity was more than 14 days ago.
	LastActivityAt *time.Time `json:"lastActivityAt,omitempty"`
	// Date and time of the last transaction on the terminal. Not included when the last transaction was more than 14 days ago.
	LastTransactionAt *time.Time `json:"lastTransactionAt,omitempty"`
	// The model name of the terminal.
	Model *string `json:"model,omitempty"`
	// The exact time of the terminal reboot, in the timezone of the terminal in **HH:mm** format.
	RestartLocalTime *string `json:"restartLocalTime,omitempty"`
	// The serial number of the terminal.
	SerialNumber *string `json:"serialNumber,omitempty"`
}

// NewTerminal instantiates a new Terminal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerminal() *Terminal {
	this := Terminal{}
	return &this
}

// NewTerminalWithDefaults instantiates a new Terminal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerminalWithDefaults() *Terminal {
	this := Terminal{}
	return &this
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *Terminal) GetAssignment() TerminalAssignment {
	if o == nil || common.IsNil(o.Assignment) {
		var ret TerminalAssignment
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetAssignmentOk() (*TerminalAssignment, bool) {
	if o == nil || common.IsNil(o.Assignment) {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *Terminal) HasAssignment() bool {
	if o != nil && !common.IsNil(o.Assignment) {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given TerminalAssignment and assigns it to the Assignment field.
func (o *Terminal) SetAssignment(v TerminalAssignment) {
	o.Assignment = &v
}

// GetConnectivity returns the Connectivity field value if set, zero value otherwise.
func (o *Terminal) GetConnectivity() TerminalConnectivity {
	if o == nil || common.IsNil(o.Connectivity) {
		var ret TerminalConnectivity
		return ret
	}
	return *o.Connectivity
}

// GetConnectivityOk returns a tuple with the Connectivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetConnectivityOk() (*TerminalConnectivity, bool) {
	if o == nil || common.IsNil(o.Connectivity) {
		return nil, false
	}
	return o.Connectivity, true
}

// HasConnectivity returns a boolean if a field has been set.
func (o *Terminal) HasConnectivity() bool {
	if o != nil && !common.IsNil(o.Connectivity) {
		return true
	}

	return false
}

// SetConnectivity gets a reference to the given TerminalConnectivity and assigns it to the Connectivity field.
func (o *Terminal) SetConnectivity(v TerminalConnectivity) {
	o.Connectivity = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *Terminal) GetFirmwareVersion() string {
	if o == nil || common.IsNil(o.FirmwareVersion) {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirmwareVersion) {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *Terminal) HasFirmwareVersion() bool {
	if o != nil && !common.IsNil(o.FirmwareVersion) {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *Terminal) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Terminal) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Terminal) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Terminal) SetId(v string) {
	o.Id = &v
}

// GetLastActivityAt returns the LastActivityAt field value if set, zero value otherwise.
func (o *Terminal) GetLastActivityAt() time.Time {
	if o == nil || common.IsNil(o.LastActivityAt) {
		var ret time.Time
		return ret
	}
	return *o.LastActivityAt
}

// GetLastActivityAtOk returns a tuple with the LastActivityAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetLastActivityAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LastActivityAt) {
		return nil, false
	}
	return o.LastActivityAt, true
}

// HasLastActivityAt returns a boolean if a field has been set.
func (o *Terminal) HasLastActivityAt() bool {
	if o != nil && !common.IsNil(o.LastActivityAt) {
		return true
	}

	return false
}

// SetLastActivityAt gets a reference to the given time.Time and assigns it to the LastActivityAt field.
func (o *Terminal) SetLastActivityAt(v time.Time) {
	o.LastActivityAt = &v
}

// GetLastTransactionAt returns the LastTransactionAt field value if set, zero value otherwise.
func (o *Terminal) GetLastTransactionAt() time.Time {
	if o == nil || common.IsNil(o.LastTransactionAt) {
		var ret time.Time
		return ret
	}
	return *o.LastTransactionAt
}

// GetLastTransactionAtOk returns a tuple with the LastTransactionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetLastTransactionAtOk() (*time.Time, bool) {
	if o == nil || common.IsNil(o.LastTransactionAt) {
		return nil, false
	}
	return o.LastTransactionAt, true
}

// HasLastTransactionAt returns a boolean if a field has been set.
func (o *Terminal) HasLastTransactionAt() bool {
	if o != nil && !common.IsNil(o.LastTransactionAt) {
		return true
	}

	return false
}

// SetLastTransactionAt gets a reference to the given time.Time and assigns it to the LastTransactionAt field.
func (o *Terminal) SetLastTransactionAt(v time.Time) {
	o.LastTransactionAt = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *Terminal) GetModel() string {
	if o == nil || common.IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetModelOk() (*string, bool) {
	if o == nil || common.IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *Terminal) HasModel() bool {
	if o != nil && !common.IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *Terminal) SetModel(v string) {
	o.Model = &v
}

// GetRestartLocalTime returns the RestartLocalTime field value if set, zero value otherwise.
func (o *Terminal) GetRestartLocalTime() string {
	if o == nil || common.IsNil(o.RestartLocalTime) {
		var ret string
		return ret
	}
	return *o.RestartLocalTime
}

// GetRestartLocalTimeOk returns a tuple with the RestartLocalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetRestartLocalTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RestartLocalTime) {
		return nil, false
	}
	return o.RestartLocalTime, true
}

// HasRestartLocalTime returns a boolean if a field has been set.
func (o *Terminal) HasRestartLocalTime() bool {
	if o != nil && !common.IsNil(o.RestartLocalTime) {
		return true
	}

	return false
}

// SetRestartLocalTime gets a reference to the given string and assigns it to the RestartLocalTime field.
func (o *Terminal) SetRestartLocalTime(v string) {
	o.RestartLocalTime = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Terminal) GetSerialNumber() string {
	if o == nil || common.IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Terminal) GetSerialNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Terminal) HasSerialNumber() bool {
	if o != nil && !common.IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Terminal) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o Terminal) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Terminal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Assignment) {
		toSerialize["assignment"] = o.Assignment
	}
	if !common.IsNil(o.Connectivity) {
		toSerialize["connectivity"] = o.Connectivity
	}
	if !common.IsNil(o.FirmwareVersion) {
		toSerialize["firmwareVersion"] = o.FirmwareVersion
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.LastActivityAt) {
		toSerialize["lastActivityAt"] = o.LastActivityAt
	}
	if !common.IsNil(o.LastTransactionAt) {
		toSerialize["lastTransactionAt"] = o.LastTransactionAt
	}
	if !common.IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !common.IsNil(o.RestartLocalTime) {
		toSerialize["restartLocalTime"] = o.RestartLocalTime
	}
	if !common.IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return toSerialize, nil
}

type NullableTerminal struct {
	value *Terminal
	isSet bool
}

func (v NullableTerminal) Get() *Terminal {
	return v.value
}

func (v *NullableTerminal) Set(val *Terminal) {
	v.value = val
	v.isSet = true
}

func (v NullableTerminal) IsSet() bool {
	return v.isSet
}

func (v *NullableTerminal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerminal(val *Terminal) *NullableTerminal {
	return &NullableTerminal{value: val, isSet: true}
}

func (v NullableTerminal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerminal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
