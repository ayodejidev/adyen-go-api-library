/*
Management Webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package managementwebhook

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// checks if the CapabilityProblemEntity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CapabilityProblemEntity{}

// CapabilityProblemEntity struct for CapabilityProblemEntity
type CapabilityProblemEntity struct {
	// List of document IDs to which the verification errors related to the capabilities correspond to.
	Documents []string `json:"documents,omitempty"`
	// The ID of the entity.
	Id    *string                           `json:"id,omitempty"`
	Owner *CapabilityProblemEntityRecursive `json:"owner,omitempty"`
	// The type of entity.  Possible values: **LegalEntity**, **BankAccount**, or **Document**.
	Type *string `json:"type,omitempty"`
}

// NewCapabilityProblemEntity instantiates a new CapabilityProblemEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapabilityProblemEntity() *CapabilityProblemEntity {
	this := CapabilityProblemEntity{}
	return &this
}

// NewCapabilityProblemEntityWithDefaults instantiates a new CapabilityProblemEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapabilityProblemEntityWithDefaults() *CapabilityProblemEntity {
	this := CapabilityProblemEntity{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *CapabilityProblemEntity) GetDocuments() []string {
	if o == nil || common.IsNil(o.Documents) {
		var ret []string
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntity) GetDocumentsOk() ([]string, bool) {
	if o == nil || common.IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *CapabilityProblemEntity) HasDocuments() bool {
	if o != nil && !common.IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []string and assigns it to the Documents field.
func (o *CapabilityProblemEntity) SetDocuments(v []string) {
	o.Documents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CapabilityProblemEntity) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntity) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CapabilityProblemEntity) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CapabilityProblemEntity) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CapabilityProblemEntity) GetOwner() CapabilityProblemEntityRecursive {
	if o == nil || common.IsNil(o.Owner) {
		var ret CapabilityProblemEntityRecursive
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntity) GetOwnerOk() (*CapabilityProblemEntityRecursive, bool) {
	if o == nil || common.IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CapabilityProblemEntity) HasOwner() bool {
	if o != nil && !common.IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given CapabilityProblemEntityRecursive and assigns it to the Owner field.
func (o *CapabilityProblemEntity) SetOwner(v CapabilityProblemEntityRecursive) {
	o.Owner = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CapabilityProblemEntity) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapabilityProblemEntity) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CapabilityProblemEntity) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CapabilityProblemEntity) SetType(v string) {
	o.Type = &v
}

func (o CapabilityProblemEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CapabilityProblemEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCapabilityProblemEntity struct {
	value *CapabilityProblemEntity
	isSet bool
}

func (v NullableCapabilityProblemEntity) Get() *CapabilityProblemEntity {
	return v.value
}

func (v *NullableCapabilityProblemEntity) Set(val *CapabilityProblemEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableCapabilityProblemEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableCapabilityProblemEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapabilityProblemEntity(val *CapabilityProblemEntity) *NullableCapabilityProblemEntity {
	return &NullableCapabilityProblemEntity{value: val, isSet: true}
}

func (v NullableCapabilityProblemEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapabilityProblemEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *CapabilityProblemEntity) isValidType() bool {
	var allowedEnumValues = []string{"BankAccount", "Document", "LegalEntity"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
