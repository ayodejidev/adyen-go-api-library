/*
POS Terminal Management API

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posterminalmanagement

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v15/src/common"
)

// checks if the Store type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Store{}

// Store struct for Store
type Store struct {
	Address *Address `json:"address,omitempty"`
	// The description of the store.
	Description *string `json:"description,omitempty"`
	// The list of terminals assigned to the store.
	InStoreTerminals []string `json:"inStoreTerminals,omitempty"`
	// The code of the merchant account.
	MerchantAccountCode *string `json:"merchantAccountCode,omitempty"`
	// The status of the store:  - `PreActive`: the store has been created, but not yet activated.   - `Active`: the store has been activated. This means you can process payments for this store.   - `Inactive`: the store is currently not active.   - `InactiveWithModifications`: the store is currently not active, but payment modifications such as refunds are possible.   - `Closed`: the store has been closed.
	Status *string `json:"status,omitempty"`
	// The code of the store.
	Store string `json:"store"`
}

// NewStore instantiates a new Store object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStore(store string) *Store {
	this := Store{}
	this.Store = store
	return &this
}

// NewStoreWithDefaults instantiates a new Store object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreWithDefaults() *Store {
	this := Store{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Store) GetAddress() Address {
	if o == nil || common.IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Store) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Store) SetAddress(v Address) {
	o.Address = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Store) GetDescription() string {
	if o == nil || common.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetDescriptionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Store) HasDescription() bool {
	if o != nil && !common.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Store) SetDescription(v string) {
	o.Description = &v
}

// GetInStoreTerminals returns the InStoreTerminals field value if set, zero value otherwise.
func (o *Store) GetInStoreTerminals() []string {
	if o == nil || common.IsNil(o.InStoreTerminals) {
		var ret []string
		return ret
	}
	return o.InStoreTerminals
}

// GetInStoreTerminalsOk returns a tuple with the InStoreTerminals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetInStoreTerminalsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.InStoreTerminals) {
		return nil, false
	}
	return o.InStoreTerminals, true
}

// HasInStoreTerminals returns a boolean if a field has been set.
func (o *Store) HasInStoreTerminals() bool {
	if o != nil && !common.IsNil(o.InStoreTerminals) {
		return true
	}

	return false
}

// SetInStoreTerminals gets a reference to the given []string and assigns it to the InStoreTerminals field.
func (o *Store) SetInStoreTerminals(v []string) {
	o.InStoreTerminals = v
}

// GetMerchantAccountCode returns the MerchantAccountCode field value if set, zero value otherwise.
func (o *Store) GetMerchantAccountCode() string {
	if o == nil || common.IsNil(o.MerchantAccountCode) {
		var ret string
		return ret
	}
	return *o.MerchantAccountCode
}

// GetMerchantAccountCodeOk returns a tuple with the MerchantAccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetMerchantAccountCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.MerchantAccountCode) {
		return nil, false
	}
	return o.MerchantAccountCode, true
}

// HasMerchantAccountCode returns a boolean if a field has been set.
func (o *Store) HasMerchantAccountCode() bool {
	if o != nil && !common.IsNil(o.MerchantAccountCode) {
		return true
	}

	return false
}

// SetMerchantAccountCode gets a reference to the given string and assigns it to the MerchantAccountCode field.
func (o *Store) SetMerchantAccountCode(v string) {
	o.MerchantAccountCode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Store) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Store) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Store) SetStatus(v string) {
	o.Status = &v
}

// GetStore returns the Store field value
func (o *Store) GetStore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Store
}

// GetStoreOk returns a tuple with the Store field value
// and a boolean to check if the value has been set.
func (o *Store) GetStoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Store, true
}

// SetStore sets field value
func (o *Store) SetStore(v string) {
	o.Store = v
}

func (o Store) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Store) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !common.IsNil(o.InStoreTerminals) {
		toSerialize["inStoreTerminals"] = o.InStoreTerminals
	}
	if !common.IsNil(o.MerchantAccountCode) {
		toSerialize["merchantAccountCode"] = o.MerchantAccountCode
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["store"] = o.Store
	return toSerialize, nil
}

type NullableStore struct {
	value *Store
	isSet bool
}

func (v NullableStore) Get() *Store {
	return v.value
}

func (v *NullableStore) Set(val *Store) {
	v.value = val
	v.isSet = true
}

func (v NullableStore) IsSet() bool {
	return v.isSet
}

func (v *NullableStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStore(val *Store) *NullableStore {
	return &NullableStore{value: val, isSet: true}
}

func (v NullableStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
