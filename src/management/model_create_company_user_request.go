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

// checks if the CreateCompanyUserRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateCompanyUserRequest{}

// CreateCompanyUserRequest struct for CreateCompanyUserRequest
type CreateCompanyUserRequest struct {
	// The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user.
	AccountGroups []string `json:"accountGroups,omitempty"`
	// The list of [merchant accounts](https://docs.adyen.com/account/account-structure#merchant-accounts) associated with this user.
	AssociatedMerchantAccounts []string `json:"associatedMerchantAccounts,omitempty"`
	// The email address of the user.
	Email string `json:"email"`
	// The requested login method for the user. To use SSO, you must already have SSO configured with Adyen before creating the user.  Possible values: **Username & account**, **Email**, or **SSO**
	LoginMethod *string `json:"loginMethod,omitempty"`
	Name        Name    `json:"name"`
	// The list of [roles](https://docs.adyen.com/account/user-roles) for this user.
	Roles []string `json:"roles,omitempty"`
	// The [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) of the time zone of the user. For example, **Europe/Amsterdam**.
	TimeZoneCode *string `json:"timeZoneCode,omitempty"`
	// The user's email address that will be their username. Must be the same as the one in the `email` field.
	Username string `json:"username"`
}

// NewCreateCompanyUserRequest instantiates a new CreateCompanyUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCompanyUserRequest(email string, name Name, username string) *CreateCompanyUserRequest {
	this := CreateCompanyUserRequest{}
	this.Email = email
	this.Name = name
	this.Username = username
	return &this
}

// NewCreateCompanyUserRequestWithDefaults instantiates a new CreateCompanyUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCompanyUserRequestWithDefaults() *CreateCompanyUserRequest {
	this := CreateCompanyUserRequest{}
	return &this
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *CreateCompanyUserRequest) GetAccountGroups() []string {
	if o == nil || common.IsNil(o.AccountGroups) {
		var ret []string
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetAccountGroupsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *CreateCompanyUserRequest) HasAccountGroups() bool {
	if o != nil && !common.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []string and assigns it to the AccountGroups field.
func (o *CreateCompanyUserRequest) SetAccountGroups(v []string) {
	o.AccountGroups = v
}

// GetAssociatedMerchantAccounts returns the AssociatedMerchantAccounts field value if set, zero value otherwise.
func (o *CreateCompanyUserRequest) GetAssociatedMerchantAccounts() []string {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		var ret []string
		return ret
	}
	return o.AssociatedMerchantAccounts
}

// GetAssociatedMerchantAccountsOk returns a tuple with the AssociatedMerchantAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetAssociatedMerchantAccountsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AssociatedMerchantAccounts) {
		return nil, false
	}
	return o.AssociatedMerchantAccounts, true
}

// HasAssociatedMerchantAccounts returns a boolean if a field has been set.
func (o *CreateCompanyUserRequest) HasAssociatedMerchantAccounts() bool {
	if o != nil && !common.IsNil(o.AssociatedMerchantAccounts) {
		return true
	}

	return false
}

// SetAssociatedMerchantAccounts gets a reference to the given []string and assigns it to the AssociatedMerchantAccounts field.
func (o *CreateCompanyUserRequest) SetAssociatedMerchantAccounts(v []string) {
	o.AssociatedMerchantAccounts = v
}

// GetEmail returns the Email field value
func (o *CreateCompanyUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateCompanyUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetLoginMethod returns the LoginMethod field value if set, zero value otherwise.
func (o *CreateCompanyUserRequest) GetLoginMethod() string {
	if o == nil || common.IsNil(o.LoginMethod) {
		var ret string
		return ret
	}
	return *o.LoginMethod
}

// GetLoginMethodOk returns a tuple with the LoginMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetLoginMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoginMethod) {
		return nil, false
	}
	return o.LoginMethod, true
}

// HasLoginMethod returns a boolean if a field has been set.
func (o *CreateCompanyUserRequest) HasLoginMethod() bool {
	if o != nil && !common.IsNil(o.LoginMethod) {
		return true
	}

	return false
}

// SetLoginMethod gets a reference to the given string and assigns it to the LoginMethod field.
func (o *CreateCompanyUserRequest) SetLoginMethod(v string) {
	o.LoginMethod = &v
}

// GetName returns the Name field value
func (o *CreateCompanyUserRequest) GetName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCompanyUserRequest) SetName(v Name) {
	o.Name = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateCompanyUserRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateCompanyUserRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateCompanyUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTimeZoneCode returns the TimeZoneCode field value if set, zero value otherwise.
func (o *CreateCompanyUserRequest) GetTimeZoneCode() string {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		var ret string
		return ret
	}
	return *o.TimeZoneCode
}

// GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetTimeZoneCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		return nil, false
	}
	return o.TimeZoneCode, true
}

// HasTimeZoneCode returns a boolean if a field has been set.
func (o *CreateCompanyUserRequest) HasTimeZoneCode() bool {
	if o != nil && !common.IsNil(o.TimeZoneCode) {
		return true
	}

	return false
}

// SetTimeZoneCode gets a reference to the given string and assigns it to the TimeZoneCode field.
func (o *CreateCompanyUserRequest) SetTimeZoneCode(v string) {
	o.TimeZoneCode = &v
}

// GetUsername returns the Username field value
func (o *CreateCompanyUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateCompanyUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateCompanyUserRequest) SetUsername(v string) {
	o.Username = v
}

func (o CreateCompanyUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCompanyUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !common.IsNil(o.AssociatedMerchantAccounts) {
		toSerialize["associatedMerchantAccounts"] = o.AssociatedMerchantAccounts
	}
	toSerialize["email"] = o.Email
	if !common.IsNil(o.LoginMethod) {
		toSerialize["loginMethod"] = o.LoginMethod
	}
	toSerialize["name"] = o.Name
	if !common.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !common.IsNil(o.TimeZoneCode) {
		toSerialize["timeZoneCode"] = o.TimeZoneCode
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableCreateCompanyUserRequest struct {
	value *CreateCompanyUserRequest
	isSet bool
}

func (v NullableCreateCompanyUserRequest) Get() *CreateCompanyUserRequest {
	return v.value
}

func (v *NullableCreateCompanyUserRequest) Set(val *CreateCompanyUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCompanyUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCompanyUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCompanyUserRequest(val *CreateCompanyUserRequest) *NullableCreateCompanyUserRequest {
	return &NullableCreateCompanyUserRequest{value: val, isSet: true}
}

func (v NullableCreateCompanyUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCompanyUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
