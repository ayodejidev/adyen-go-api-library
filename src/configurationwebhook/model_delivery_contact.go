/*
Configuration webhooks

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the DeliveryContact type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeliveryContact{}

// DeliveryContact struct for DeliveryContact
type DeliveryContact struct {
	Address DeliveryAddress `json:"address"`
	// The company name of the contact.
	Company *string `json:"company,omitempty"`
	// The email address of the contact.
	Email *string `json:"email,omitempty"`
	// The full phone number of the contact provided as a single string. It will be handled as a landline phone. **Examples:** \"0031 6 11 22 33 44\", \"+316/1122-3344\", \"(0031) 611223344\"
	FullPhoneNumber *string      `json:"fullPhoneNumber,omitempty"`
	Name            Name         `json:"name"`
	PhoneNumber     *PhoneNumber `json:"phoneNumber,omitempty"`
	// The URL of the contact's website.
	WebAddress *string `json:"webAddress,omitempty"`
}

// NewDeliveryContact instantiates a new DeliveryContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryContact(address DeliveryAddress, name Name) *DeliveryContact {
	this := DeliveryContact{}
	this.Address = address
	this.Name = name
	return &this
}

// NewDeliveryContactWithDefaults instantiates a new DeliveryContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryContactWithDefaults() *DeliveryContact {
	this := DeliveryContact{}
	return &this
}

// GetAddress returns the Address field value
func (o *DeliveryContact) GetAddress() DeliveryAddress {
	if o == nil {
		var ret DeliveryAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetAddressOk() (*DeliveryAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *DeliveryContact) SetAddress(v DeliveryAddress) {
	o.Address = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *DeliveryContact) GetCompany() string {
	if o == nil || common.IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetCompanyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *DeliveryContact) HasCompany() bool {
	if o != nil && !common.IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *DeliveryContact) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DeliveryContact) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DeliveryContact) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DeliveryContact) SetEmail(v string) {
	o.Email = &v
}

// GetFullPhoneNumber returns the FullPhoneNumber field value if set, zero value otherwise.
func (o *DeliveryContact) GetFullPhoneNumber() string {
	if o == nil || common.IsNil(o.FullPhoneNumber) {
		var ret string
		return ret
	}
	return *o.FullPhoneNumber
}

// GetFullPhoneNumberOk returns a tuple with the FullPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetFullPhoneNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.FullPhoneNumber) {
		return nil, false
	}
	return o.FullPhoneNumber, true
}

// HasFullPhoneNumber returns a boolean if a field has been set.
func (o *DeliveryContact) HasFullPhoneNumber() bool {
	if o != nil && !common.IsNil(o.FullPhoneNumber) {
		return true
	}

	return false
}

// SetFullPhoneNumber gets a reference to the given string and assigns it to the FullPhoneNumber field.
func (o *DeliveryContact) SetFullPhoneNumber(v string) {
	o.FullPhoneNumber = &v
}

// GetName returns the Name field value
func (o *DeliveryContact) GetName() Name {
	if o == nil {
		var ret Name
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetNameOk() (*Name, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeliveryContact) SetName(v Name) {
	o.Name = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *DeliveryContact) GetPhoneNumber() PhoneNumber {
	if o == nil || common.IsNil(o.PhoneNumber) {
		var ret PhoneNumber
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetPhoneNumberOk() (*PhoneNumber, bool) {
	if o == nil || common.IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *DeliveryContact) HasPhoneNumber() bool {
	if o != nil && !common.IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given PhoneNumber and assigns it to the PhoneNumber field.
func (o *DeliveryContact) SetPhoneNumber(v PhoneNumber) {
	o.PhoneNumber = &v
}

// GetWebAddress returns the WebAddress field value if set, zero value otherwise.
func (o *DeliveryContact) GetWebAddress() string {
	if o == nil || common.IsNil(o.WebAddress) {
		var ret string
		return ret
	}
	return *o.WebAddress
}

// GetWebAddressOk returns a tuple with the WebAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryContact) GetWebAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WebAddress) {
		return nil, false
	}
	return o.WebAddress, true
}

// HasWebAddress returns a boolean if a field has been set.
func (o *DeliveryContact) HasWebAddress() bool {
	if o != nil && !common.IsNil(o.WebAddress) {
		return true
	}

	return false
}

// SetWebAddress gets a reference to the given string and assigns it to the WebAddress field.
func (o *DeliveryContact) SetWebAddress(v string) {
	o.WebAddress = &v
}

func (o DeliveryContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !common.IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.FullPhoneNumber) {
		toSerialize["fullPhoneNumber"] = o.FullPhoneNumber
	}
	toSerialize["name"] = o.Name
	if !common.IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !common.IsNil(o.WebAddress) {
		toSerialize["webAddress"] = o.WebAddress
	}
	return toSerialize, nil
}

type NullableDeliveryContact struct {
	value *DeliveryContact
	isSet bool
}

func (v NullableDeliveryContact) Get() *DeliveryContact {
	return v.value
}

func (v *NullableDeliveryContact) Set(val *DeliveryContact) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryContact) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryContact(val *DeliveryContact) *NullableDeliveryContact {
	return &NullableDeliveryContact{value: val, isSet: true}
}

func (v NullableDeliveryContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
