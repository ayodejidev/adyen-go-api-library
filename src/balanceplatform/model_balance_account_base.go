/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the BalanceAccountBase type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceAccountBase{}

// BalanceAccountBase struct for BalanceAccountBase
type BalanceAccountBase struct {
	// The unique identifier of the [account holder](https://docs.adyen.com/api-explorer/#/balanceplatform/latest/post/accountHolders__resParam_id) associated with the balance account.
	AccountHolderId string `json:"accountHolderId"`
	// The default three-character [ISO currency code](https://docs.adyen.com/development-resources/currency-codes) of the balance account. The default value is **EUR**. > After a balance account is created, you cannot change its default currency.
	DefaultCurrencyCode *string `json:"defaultCurrencyCode,omitempty"`
	// A human-readable description of the balance account, maximum 300 characters. You can use this parameter to distinguish between multiple balance accounts under an account holder.
	Description *string `json:"description,omitempty"`
	// The unique identifier of the balance account.
	Id string `json:"id"`
	// A set of key and value pairs for general use. The keys do not have specific names and may be used for storing miscellaneous data as desired. > Note that during an update of metadata, the omission of existing key-value pairs will result in the deletion of those key-value pairs.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The unique identifier of the account of the migrated account holder in the classic integration.
	MigratedAccountCode          *string                       `json:"migratedAccountCode,omitempty"`
	PlatformPaymentConfiguration *PlatformPaymentConfiguration `json:"platformPaymentConfiguration,omitempty"`
	// Your reference for the balance account, maximum 150 characters.
	Reference *string `json:"reference,omitempty"`
	// The status of the balance account, set to **active** by default.
	Status *string `json:"status,omitempty"`
	// The time zone of the balance account. For example, **Europe/Amsterdam**. Defaults to the time zone of the account holder if no time zone is set. For possible values, see the [list of time zone codes](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewBalanceAccountBase instantiates a new BalanceAccountBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceAccountBase(accountHolderId string, id string) *BalanceAccountBase {
	this := BalanceAccountBase{}
	this.AccountHolderId = accountHolderId
	this.Id = id
	return &this
}

// NewBalanceAccountBaseWithDefaults instantiates a new BalanceAccountBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceAccountBaseWithDefaults() *BalanceAccountBase {
	this := BalanceAccountBase{}
	return &this
}

// GetAccountHolderId returns the AccountHolderId field value
func (o *BalanceAccountBase) GetAccountHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolderId
}

// GetAccountHolderIdOk returns a tuple with the AccountHolderId field value
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetAccountHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolderId, true
}

// SetAccountHolderId sets field value
func (o *BalanceAccountBase) SetAccountHolderId(v string) {
	o.AccountHolderId = v
}

// GetDefaultCurrencyCode returns the DefaultCurrencyCode field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetDefaultCurrencyCode() string {
	if o == nil || common.IsNil(o.DefaultCurrencyCode) {
		var ret string
		return ret
	}
	return *o.DefaultCurrencyCode
}

// GetDefaultCurrencyCodeOk returns a tuple with the DefaultCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetDefaultCurrencyCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.DefaultCurrencyCode) {
		return nil, false
	}
	return o.DefaultCurrencyCode, true
}

// HasDefaultCurrencyCode returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasDefaultCurrencyCode() bool {
	if o != nil && !common.IsNil(o.DefaultCurrencyCode) {
		return true
	}

	return false
}

// SetDefaultCurrencyCode gets a reference to the given string and assigns it to the DefaultCurrencyCode field.
func (o *BalanceAccountBase) SetDefaultCurrencyCode(v string) {
	o.DefaultCurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BalanceAccountBase) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *BalanceAccountBase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BalanceAccountBase) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetMetadata() map[string]string {
	if o == nil || common.IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || common.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasMetadata() bool {
	if o != nil && !common.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *BalanceAccountBase) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMigratedAccountCode returns the MigratedAccountCode field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetMigratedAccountCode() string {
	if o == nil || common.IsNil(o.MigratedAccountCode) {
		var ret string
		return ret
	}
	return *o.MigratedAccountCode
}

// GetMigratedAccountCodeOk returns a tuple with the MigratedAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetMigratedAccountCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MigratedAccountCode) {
		return nil, false
	}
	return o.MigratedAccountCode, true
}

// HasMigratedAccountCode returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasMigratedAccountCode() bool {
	if o != nil && !common.IsNil(o.MigratedAccountCode) {
		return true
	}

	return false
}

// SetMigratedAccountCode gets a reference to the given string and assigns it to the MigratedAccountCode field.
func (o *BalanceAccountBase) SetMigratedAccountCode(v string) {
	o.MigratedAccountCode = &v
}

// GetPlatformPaymentConfiguration returns the PlatformPaymentConfiguration field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetPlatformPaymentConfiguration() PlatformPaymentConfiguration {
	if o == nil || common.IsNil(o.PlatformPaymentConfiguration) {
		var ret PlatformPaymentConfiguration
		return ret
	}
	return *o.PlatformPaymentConfiguration
}

// GetPlatformPaymentConfigurationOk returns a tuple with the PlatformPaymentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetPlatformPaymentConfigurationOk() (*PlatformPaymentConfiguration, bool) {
	if o == nil || common.IsNil(o.PlatformPaymentConfiguration) {
		return nil, false
	}
	return o.PlatformPaymentConfiguration, true
}

// HasPlatformPaymentConfiguration returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasPlatformPaymentConfiguration() bool {
	if o != nil && !common.IsNil(o.PlatformPaymentConfiguration) {
		return true
	}

	return false
}

// SetPlatformPaymentConfiguration gets a reference to the given PlatformPaymentConfiguration and assigns it to the PlatformPaymentConfiguration field.
func (o *BalanceAccountBase) SetPlatformPaymentConfiguration(v PlatformPaymentConfiguration) {
	o.PlatformPaymentConfiguration = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *BalanceAccountBase) SetReference(v string) {
	o.Reference = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BalanceAccountBase) SetStatus(v string) {
	o.Status = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *BalanceAccountBase) GetTimeZone() string {
	if o == nil || common.IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceAccountBase) GetTimeZoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *BalanceAccountBase) HasTimeZone() bool {
	if o != nil && !common.IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *BalanceAccountBase) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o BalanceAccountBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceAccountBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountHolderId"] = o.AccountHolderId
	if !common.IsNil(o.DefaultCurrencyCode) {
		toSerialize["defaultCurrencyCode"] = o.DefaultCurrencyCode
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !common.IsNil(o.MigratedAccountCode) {
		toSerialize["migratedAccountCode"] = o.MigratedAccountCode
	}
	if !common.IsNil(o.PlatformPaymentConfiguration) {
		toSerialize["platformPaymentConfiguration"] = o.PlatformPaymentConfiguration
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return toSerialize, nil
}

type NullableBalanceAccountBase struct {
	value *BalanceAccountBase
	isSet bool
}

func (v NullableBalanceAccountBase) Get() *BalanceAccountBase {
	return v.value
}

func (v *NullableBalanceAccountBase) Set(val *BalanceAccountBase) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceAccountBase) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceAccountBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceAccountBase(val *BalanceAccountBase) *NullableBalanceAccountBase {
	return &NullableBalanceAccountBase{value: val, isSet: true}
}

func (v NullableBalanceAccountBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceAccountBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *BalanceAccountBase) isValidStatus() bool {
	var allowedEnumValues = []string{"active", "closed", "inactive", "suspended"}
	for _, allowed := range allowedEnumValues {
		if o.GetStatus() == allowed {
			return true
		}
	}
	return false
}
