/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the AccountSupportingEntityCapability type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountSupportingEntityCapability{}

// AccountSupportingEntityCapability struct for AccountSupportingEntityCapability
type AccountSupportingEntityCapability struct {
	// Indicates whether the supporting entity capability is allowed. Adyen sets this to **true** if the verification is successful and the account holder is permitted to use the capability.
	Allowed *bool `json:"allowed,omitempty"`
	// The capability level that is allowed for the account holder.  Possible values: **notApplicable**, **low**, **medium**, **high**.
	AllowedLevel *string `json:"allowedLevel,omitempty"`
	// Indicates whether the capability is enabled. If **false**, the capability is temporarily disabled for the account holder.
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the supporting entity.
	Id *string `json:"id,omitempty"`
	// Indicates whether the capability is requested. To check whether the account holder is permitted to use the capability, refer to the `allowed` field.
	Requested *bool `json:"requested,omitempty"`
	// The requested level of the capability. Some capabilities, such as those used in [card issuing](https://docs.adyen.com/issuing/add-capabilities#capability-levels), have different levels. Levels increase the capability, but also require additional checks and increased monitoring.  Possible values: **notApplicable**, **low**, **medium**, **high**.
	RequestedLevel *string `json:"requestedLevel,omitempty"`
	// The status of the verification checks for the supporting entity capability.  Possible values:  * **pending**: Adyen is running the verification.  * **invalid**: The verification failed. Check if the `errors` array contains more information.  * **valid**: The verification has been successfully completed.  * **rejected**: Adyen has verified the information, but found reasons to not allow the capability.
	VerificationStatus *string `json:"verificationStatus,omitempty"`
}

// NewAccountSupportingEntityCapability instantiates a new AccountSupportingEntityCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSupportingEntityCapability() *AccountSupportingEntityCapability {
	this := AccountSupportingEntityCapability{}
	return &this
}

// NewAccountSupportingEntityCapabilityWithDefaults instantiates a new AccountSupportingEntityCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSupportingEntityCapabilityWithDefaults() *AccountSupportingEntityCapability {
	this := AccountSupportingEntityCapability{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetAllowed() bool {
	if o == nil || common.IsNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetAllowedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasAllowed() bool {
	if o != nil && !common.IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *AccountSupportingEntityCapability) SetAllowed(v bool) {
	o.Allowed = &v
}

// GetAllowedLevel returns the AllowedLevel field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetAllowedLevel() string {
	if o == nil || common.IsNil(o.AllowedLevel) {
		var ret string
		return ret
	}
	return *o.AllowedLevel
}

// GetAllowedLevelOk returns a tuple with the AllowedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetAllowedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.AllowedLevel) {
		return nil, false
	}
	return o.AllowedLevel, true
}

// HasAllowedLevel returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasAllowedLevel() bool {
	if o != nil && !common.IsNil(o.AllowedLevel) {
		return true
	}

	return false
}

// SetAllowedLevel gets a reference to the given string and assigns it to the AllowedLevel field.
func (o *AccountSupportingEntityCapability) SetAllowedLevel(v string) {
	o.AllowedLevel = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetEnabled() bool {
	if o == nil || common.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetEnabledOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasEnabled() bool {
	if o != nil && !common.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AccountSupportingEntityCapability) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountSupportingEntityCapability) SetId(v string) {
	o.Id = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetRequested() bool {
	if o == nil || common.IsNil(o.Requested) {
		var ret bool
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetRequestedOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasRequested() bool {
	if o != nil && !common.IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given bool and assigns it to the Requested field.
func (o *AccountSupportingEntityCapability) SetRequested(v bool) {
	o.Requested = &v
}

// GetRequestedLevel returns the RequestedLevel field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetRequestedLevel() string {
	if o == nil || common.IsNil(o.RequestedLevel) {
		var ret string
		return ret
	}
	return *o.RequestedLevel
}

// GetRequestedLevelOk returns a tuple with the RequestedLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetRequestedLevelOk() (*string, bool) {
	if o == nil || common.IsNil(o.RequestedLevel) {
		return nil, false
	}
	return o.RequestedLevel, true
}

// HasRequestedLevel returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasRequestedLevel() bool {
	if o != nil && !common.IsNil(o.RequestedLevel) {
		return true
	}

	return false
}

// SetRequestedLevel gets a reference to the given string and assigns it to the RequestedLevel field.
func (o *AccountSupportingEntityCapability) SetRequestedLevel(v string) {
	o.RequestedLevel = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *AccountSupportingEntityCapability) GetVerificationStatus() string {
	if o == nil || common.IsNil(o.VerificationStatus) {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSupportingEntityCapability) GetVerificationStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationStatus) {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountSupportingEntityCapability) HasVerificationStatus() bool {
	if o != nil && !common.IsNil(o.VerificationStatus) {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *AccountSupportingEntityCapability) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

func (o AccountSupportingEntityCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSupportingEntityCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	if !common.IsNil(o.AllowedLevel) {
		toSerialize["allowedLevel"] = o.AllowedLevel
	}
	if !common.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !common.IsNil(o.RequestedLevel) {
		toSerialize["requestedLevel"] = o.RequestedLevel
	}
	if !common.IsNil(o.VerificationStatus) {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	return toSerialize, nil
}

type NullableAccountSupportingEntityCapability struct {
	value *AccountSupportingEntityCapability
	isSet bool
}

func (v NullableAccountSupportingEntityCapability) Get() *AccountSupportingEntityCapability {
	return v.value
}

func (v *NullableAccountSupportingEntityCapability) Set(val *AccountSupportingEntityCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSupportingEntityCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSupportingEntityCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSupportingEntityCapability(val *AccountSupportingEntityCapability) *NullableAccountSupportingEntityCapability {
	return &NullableAccountSupportingEntityCapability{value: val, isSet: true}
}

func (v NullableAccountSupportingEntityCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSupportingEntityCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *AccountSupportingEntityCapability) isValidAllowedLevel() bool {
	var allowedEnumValues = []string{"high", "low", "medium", "notApplicable"}
	for _, allowed := range allowedEnumValues {
		if o.GetAllowedLevel() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountSupportingEntityCapability) isValidRequestedLevel() bool {
	var allowedEnumValues = []string{"high", "low", "medium", "notApplicable"}
	for _, allowed := range allowedEnumValues {
		if o.GetRequestedLevel() == allowed {
			return true
		}
	}
	return false
}
func (o *AccountSupportingEntityCapability) isValidVerificationStatus() bool {
	var allowedEnumValues = []string{"invalid", "pending", "rejected", "valid"}
	for _, allowed := range allowedEnumValues {
		if o.GetVerificationStatus() == allowed {
			return true
		}
	}
	return false
}
