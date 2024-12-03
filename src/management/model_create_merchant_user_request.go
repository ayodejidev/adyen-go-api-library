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

// checks if the CreateMerchantUserRequest type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CreateMerchantUserRequest{}

// CreateMerchantUserRequest struct for CreateMerchantUserRequest
type CreateMerchantUserRequest struct {
	// The list of [account groups](https://docs.adyen.com/account/account-structure#account-groups) associated with this user.
	AccountGroups []string `json:"accountGroups,omitempty"`
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

// NewCreateMerchantUserRequest instantiates a new CreateMerchantUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMerchantUserRequest(email string, name Name, username string) *CreateMerchantUserRequest {
	this := CreateMerchantUserRequest{}
	this.Email = email
	this.Name = name
	this.Username = username
	return &this
}

// NewCreateMerchantUserRequestWithDefaults instantiates a new CreateMerchantUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMerchantUserRequestWithDefaults() *CreateMerchantUserRequest {
	this := CreateMerchantUserRequest{}
	return &this
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetAccountGroups() []string {
	if o == nil || common.IsNil(o.AccountGroups) {
		var ret []string
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetAccountGroupsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasAccountGroups() bool {
	if o != nil && !common.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []string and assigns it to the AccountGroups field.
func (o *CreateMerchantUserRequest) SetAccountGroups(v []string) {
	o.AccountGroups = v
}

// GetEmail returns the Email field value
func (o *CreateMerchantUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateMerchantUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetLoginMethod returns the LoginMethod field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetLoginMethod() string {
	if o == nil || common.IsNil(o.LoginMethod) {
		var ret string
		return ret
	}
	return *o.LoginMethod
}

// GetLoginMethodOk returns a tuple with the LoginMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetLoginMethodOk() (*string, bool) {
	if o == nil || common.IsNil(o.LoginMethod) {
		return nil, false
	}
	return o.LoginMethod, true
}

// HasLoginMethod returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasLoginMethod() bool {
	if o != nil && !common.IsNil(o.LoginMethod) {
		return true
	}

	return false
}

// SetLoginMethod gets a reference to the given string and assigns it to the LoginMethod field.
func (o *CreateMerchantUserRequest) SetLoginMethod(v string) {
	o.LoginMethod = &v
}

// GetName returns the Name field value
func (o *CreateMerchantUserRequest) GetName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateMerchantUserRequest) SetName(v Name) {
	o.Name = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetRoles() []string {
	if o == nil || common.IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasRoles() bool {
	if o != nil && !common.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *CreateMerchantUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetTimeZoneCode returns the TimeZoneCode field value if set, zero value otherwise.
func (o *CreateMerchantUserRequest) GetTimeZoneCode() string {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		var ret string
		return ret
	}
	return *o.TimeZoneCode
}

// GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetTimeZoneCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeZoneCode) {
		return nil, false
	}
	return o.TimeZoneCode, true
}

// HasTimeZoneCode returns a boolean if a field has been set.
func (o *CreateMerchantUserRequest) HasTimeZoneCode() bool {
	if o != nil && !common.IsNil(o.TimeZoneCode) {
		return true
	}

	return false
}

// SetTimeZoneCode gets a reference to the given string and assigns it to the TimeZoneCode field.
func (o *CreateMerchantUserRequest) SetTimeZoneCode(v string) {
	o.TimeZoneCode = &v
}

// GetUsername returns the Username field value
func (o *CreateMerchantUserRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateMerchantUserRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateMerchantUserRequest) SetUsername(v string) {
	o.Username = v
}

func (o CreateMerchantUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMerchantUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
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

type NullableCreateMerchantUserRequest struct {
	value *CreateMerchantUserRequest
	isSet bool
}

func (v NullableCreateMerchantUserRequest) Get() *CreateMerchantUserRequest {
	return v.value
}

func (v *NullableCreateMerchantUserRequest) Set(val *CreateMerchantUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMerchantUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMerchantUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMerchantUserRequest(val *CreateMerchantUserRequest) *NullableCreateMerchantUserRequest {
	return &NullableCreateMerchantUserRequest{value: val, isSet: true}
}

func (v NullableCreateMerchantUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMerchantUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
