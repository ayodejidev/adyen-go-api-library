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

// checks if the UpdateMerchantUserRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UpdateMerchantUserRequest{}

// UpdateMerchantUserRequest struct for UpdateMerchantUserRequest
type UpdateMerchantUserRequest struct {
	// The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user.
	AccountGroups []string `json:"accountGroups,omitempty"`
	// Sets the status of the user to active (**true**) or inactive (**false**).
	Active *bool `json:"active,omitempty"`
	// The email address of the user.
	Email *string `json:"email,omitempty"`
	// The requested login method for the user. To use SSO, you must already have SSO configured with Adyen before creating the user.  Possible values: **Username & account**, **Email**, or **SSO**
	LoginMethod *string `json:"loginMethod,omitempty"`
	Name        *Name2  `json:"name,omitempty"`
	// The list of [roles](https://docs.adyen.com/account/user-roles) for this user.
	Roles []string `json:"roles,omitempty"`
	// The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**.
	TimeZoneCode *string `json:"timeZoneCode,omitempty"`
}

// NewUpdateMerchantUserRequest instantiates a new UpdateMerchantUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMerchantUserRequest() *UpdateMerchantUserRequest {
	this := UpdateMerchantUserRequest{}
	return &this
}

// NewUpdateMerchantUserRequestWithDefaults instantiates a new UpdateMerchantUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMerchantUserRequestWithDefaults() *UpdateMerchantUserRequest {
	this := UpdateMerchantUserRequest{}
	return &this
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetAccountGroups() []string {
	if o == nil || common.IsNil(o.AccountGroups) {
		var ret []string
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetAccountGroupsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasAccountGroups() bool {
	if o != nil && !common.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []string and assigns it to the AccountGroups field.
func (o *UpdateMerchantUserRequest) SetAccountGroups(v []string) {
	o.AccountGroups = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetActive() bool {
	if o == nil || common.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetActiveOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasActive() bool {
	if o != nil && !common.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateMerchantUserRequest) SetActive(v bool) {
	o.Active = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateMerchantUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetLoginMethod returns the LoginMethod field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetLoginMethod() string {
	if o == nil || common.IsNil(o.LoginMethod) {
		var ret string
		return ret
	}
	return *o.LoginMethod
}

// GetLoginMethodOk returns a tuple with the LoginMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetLoginMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoginMethod) {
		return nil, false
	}
	return o.LoginMethod, true
}

// HasLoginMethod returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasLoginMethod() bool {
	if o != nil && !common.IsNil(o.LoginMethod) {
		return true
	}

	return false
}

// SetLoginMethod gets a reference to the given string and assigns it to the LoginMethod field.
func (o *UpdateMerchantUserRequest) SetLoginMethod(v string) {
	o.LoginMethod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetName() Name2 {
	if o == nil || common.IsNil(o.Name) {
		var ret Name2
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetNameOk() (*Name2, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name2 and assigns it to the Name field.
func (o *UpdateMerchantUserRequest) SetName(v Name2) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UpdateMerchantUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTimeZoneCode returns the TimeZoneCode field value if set, zero value otherwise.
func (o *UpdateMerchantUserRequest) GetTimeZoneCode() string {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		var ret string
		return ret
	}
	return *o.TimeZoneCode
}

// GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMerchantUserRequest) GetTimeZoneCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		return nil, false
	}
	return o.TimeZoneCode, true
}

// HasTimeZoneCode returns a boolean if a field has been set.
func (o *UpdateMerchantUserRequest) HasTimeZoneCode() bool {
	if o != nil && !common.IsNil(o.TimeZoneCode) {
		return true
	}

	return false
}

// SetTimeZoneCode gets a reference to the given string and assigns it to the TimeZoneCode field.
func (o *UpdateMerchantUserRequest) SetTimeZoneCode(v string) {
	o.TimeZoneCode = &v
}

func (o UpdateMerchantUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMerchantUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !common.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.LoginMethod) {
		toSerialize["loginMethod"] = o.LoginMethod
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !common.IsNil(o.TimeZoneCode) {
		toSerialize["timeZoneCode"] = o.TimeZoneCode
	}
	return toSerialize, nil
}

type NullableUpdateMerchantUserRequest struct {
	value *UpdateMerchantUserRequest
	isSet bool
}

func (v NullableUpdateMerchantUserRequest) Get() *UpdateMerchantUserRequest {
	return v.value
}

func (v *NullableUpdateMerchantUserRequest) Set(val *UpdateMerchantUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMerchantUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMerchantUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMerchantUserRequest(val *UpdateMerchantUserRequest) *NullableUpdateMerchantUserRequest {
	return &NullableUpdateMerchantUserRequest{value: val, isSet: true}
}

func (v NullableUpdateMerchantUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMerchantUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
