/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// checks if the EcontextVoucherDetails type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EcontextVoucherDetails{}

// EcontextVoucherDetails struct for EcontextVoucherDetails
type EcontextVoucherDetails struct {
	// The checkout attempt identifier.
	CheckoutAttemptId *string `json:"checkoutAttemptId,omitempty"`
	// The shopper's first name.
	FirstName string `json:"firstName"`
	// The shopper's last name.
	LastName string `json:"lastName"`
	// The shopper's email.
	ShopperEmail string `json:"shopperEmail"`
	// The shopper's contact number. It must have an international number format, for example **+31 20 779 1846**. Formats like **+31 (0)20 779 1846** or **0031 20 779 1846** are not accepted.
	TelephoneNumber string `json:"telephoneNumber"`
	// **econtextvoucher**
	Type string `json:"type"`
}

// NewEcontextVoucherDetails instantiates a new EcontextVoucherDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcontextVoucherDetails(firstName string, lastName string, shopperEmail string, telephoneNumber string, type_ string) *EcontextVoucherDetails {
	this := EcontextVoucherDetails{}
	this.FirstName = firstName
	this.LastName = lastName
	this.ShopperEmail = shopperEmail
	this.TelephoneNumber = telephoneNumber
	this.Type = type_
	return &this
}

// NewEcontextVoucherDetailsWithDefaults instantiates a new EcontextVoucherDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcontextVoucherDetailsWithDefaults() *EcontextVoucherDetails {
	this := EcontextVoucherDetails{}
	return &this
}

// GetCheckoutAttemptId returns the CheckoutAttemptId field value if set, zero value otherwise.
func (o *EcontextVoucherDetails) GetCheckoutAttemptId() string {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		var ret string
		return ret
	}
	return *o.CheckoutAttemptId
}

// GetCheckoutAttemptIdOk returns a tuple with the CheckoutAttemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcontextVoucherDetails) GetCheckoutAttemptIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.CheckoutAttemptId) {
		return nil, false
	}
	return o.CheckoutAttemptId, true
}

// HasCheckoutAttemptId returns a boolean if a field has been set.
func (o *EcontextVoucherDetails) HasCheckoutAttemptId() bool {
	if o != nil && !common.IsNil(o.CheckoutAttemptId) {
		return true
	}

	return false
}

// SetCheckoutAttemptId gets a reference to the given string and assigns it to the CheckoutAttemptId field.
func (o *EcontextVoucherDetails) SetCheckoutAttemptId(v string) {
	o.CheckoutAttemptId = &v
}

// GetFirstName returns the FirstName field value
func (o *EcontextVoucherDetails) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *EcontextVoucherDetails) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *EcontextVoucherDetails) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *EcontextVoucherDetails) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *EcontextVoucherDetails) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *EcontextVoucherDetails) SetLastName(v string) {
	o.LastName = v
}

// GetShopperEmail returns the ShopperEmail field value
func (o *EcontextVoucherDetails) GetShopperEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShopperEmail
}

// GetShopperEmailOk returns a tuple with the ShopperEmail field value
// and a boolean to check if the value has been set.
func (o *EcontextVoucherDetails) GetShopperEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShopperEmail, true
}

// SetShopperEmail sets field value
func (o *EcontextVoucherDetails) SetShopperEmail(v string) {
	o.ShopperEmail = v
}

// GetTelephoneNumber returns the TelephoneNumber field value
func (o *EcontextVoucherDetails) GetTelephoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TelephoneNumber
}

// GetTelephoneNumberOk returns a tuple with the TelephoneNumber field value
// and a boolean to check if the value has been set.
func (o *EcontextVoucherDetails) GetTelephoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TelephoneNumber, true
}

// SetTelephoneNumber sets field value
func (o *EcontextVoucherDetails) SetTelephoneNumber(v string) {
	o.TelephoneNumber = v
}

// GetType returns the Type field value
func (o *EcontextVoucherDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EcontextVoucherDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EcontextVoucherDetails) SetType(v string) {
	o.Type = v
}

func (o EcontextVoucherDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcontextVoucherDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.CheckoutAttemptId) {
		toSerialize["checkoutAttemptId"] = o.CheckoutAttemptId
	}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["shopperEmail"] = o.ShopperEmail
	toSerialize["telephoneNumber"] = o.TelephoneNumber
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableEcontextVoucherDetails struct {
	value *EcontextVoucherDetails
	isSet bool
}

func (v NullableEcontextVoucherDetails) Get() *EcontextVoucherDetails {
	return v.value
}

func (v *NullableEcontextVoucherDetails) Set(val *EcontextVoucherDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEcontextVoucherDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEcontextVoucherDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcontextVoucherDetails(val *EcontextVoucherDetails) *NullableEcontextVoucherDetails {
	return &NullableEcontextVoucherDetails{value: val, isSet: true}
}

func (v NullableEcontextVoucherDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcontextVoucherDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *EcontextVoucherDetails) isValidType() bool {
	var allowedEnumValues = []string{"econtext_seven_eleven", "econtext_online", "econtext", "econtext_stores", "econtext_atm"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
