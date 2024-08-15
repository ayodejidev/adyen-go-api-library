/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// checks if the PartyIdentification type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartyIdentification{}

// PartyIdentification struct for PartyIdentification
type PartyIdentification struct {
	Address *Address `json:"address,omitempty"`
	// The date of birth of the individual in [ISO-8601](https://www.w3.org/TR/NOTE-datetime) format. For example, **YYYY-MM-DD**.  Allowed only when `type` is **individual**.
	DateOfBirth *string `json:"dateOfBirth,omitempty"`
	// The first name of the individual.  Supported characters: [a-z] [A-Z] - . / — and space.  This parameter is: - Allowed only when `type` is **individual**. - Required when `category` is **card**.
	FirstName *string `json:"firstName,omitempty"`
	// The full name of the entity that owns the bank account or card.  Supported characters: [a-z] [A-Z] [0-9] , . ; : - — / \\ + & ! ? @ ( ) \" ' and space.  Required when `category` is **bank**.
	FullName *string `json:"fullName,omitempty"`
	// The last name of the individual.  Supported characters: [a-z] [A-Z] - . / — and space.  This parameter is: - Allowed only when `type` is **individual**. - Required when `category` is **card**.
	LastName *string `json:"lastName,omitempty"`
	// A unique reference to identify the party or counterparty involved in the transfer. For example, your client's unique wallet or payee ID.  Required when you include `cardIdentification.storedPaymentMethodId`.
	Reference *string `json:"reference,omitempty"`
	// The type of entity that owns the bank account or card.  Possible values: **individual**, **organization**, or **unknown**.  Required when `category` is **card**. In this case, the value must be **individual**.
	Type *string `json:"type,omitempty"`
}

// NewPartyIdentification instantiates a new PartyIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyIdentification() *PartyIdentification {
	this := PartyIdentification{}
	var type_ string = "unknown"
	this.Type = &type_
	return &this
}

// NewPartyIdentificationWithDefaults instantiates a new PartyIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyIdentificationWithDefaults() *PartyIdentification {
	this := PartyIdentification{}
	var type_ string = "unknown"
	this.Type = &type_
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PartyIdentification) GetAddress() Address {
	if o == nil || common.IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetAddressOk() (*Address, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PartyIdentification) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *PartyIdentification) SetAddress(v Address) {
	o.Address = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *PartyIdentification) GetDateOfBirth() string {
	if o == nil || common.IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetDateOfBirthOk() (*string, bool) {
	if o == nil || common.IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *PartyIdentification) HasDateOfBirth() bool {
	if o != nil && !common.IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *PartyIdentification) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PartyIdentification) GetFirstName() string {
	if o == nil || common.IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetFirstNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PartyIdentification) HasFirstName() bool {
	if o != nil && !common.IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PartyIdentification) SetFirstName(v string) {
	o.FirstName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *PartyIdentification) GetFullName() string {
	if o == nil || common.IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetFullNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *PartyIdentification) HasFullName() bool {
	if o != nil && !common.IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *PartyIdentification) SetFullName(v string) {
	o.FullName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PartyIdentification) GetLastName() string {
	if o == nil || common.IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetLastNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PartyIdentification) HasLastName() bool {
	if o != nil && !common.IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PartyIdentification) SetLastName(v string) {
	o.LastName = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PartyIdentification) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PartyIdentification) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PartyIdentification) SetReference(v string) {
	o.Reference = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PartyIdentification) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PartyIdentification) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PartyIdentification) SetType(v string) {
	o.Type = &v
}

func (o PartyIdentification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartyIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !common.IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !common.IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !common.IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullablePartyIdentification struct {
	value *PartyIdentification
	isSet bool
}

func (v NullablePartyIdentification) Get() *PartyIdentification {
	return v.value
}

func (v *NullablePartyIdentification) Set(val *PartyIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyIdentification(val *PartyIdentification) *NullablePartyIdentification {
	return &NullablePartyIdentification{value: val, isSet: true}
}

func (v NullablePartyIdentification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *PartyIdentification) isValidType() bool {
	var allowedEnumValues = []string{"individual", "organization", "unknown"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
