/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"encoding/json"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// checks if the LegalEntity type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LegalEntity{}

// LegalEntity struct for LegalEntity
type LegalEntity struct {
	// Contains key-value pairs that specify the actions that the legal entity can do in your platform.The key is a capability required for your integration. For example, **issueCard** for Issuing.The value is an object containing the settings for the capability.
	Capabilities *map[string]LegalEntityCapability `json:"capabilities,omitempty"`
	// List of documents uploaded for the legal entity.
	DocumentDetails []DocumentReference `json:"documentDetails,omitempty"`
	// List of documents uploaded for the legal entity.
	// Deprecated
	Documents []EntityReference `json:"documents,omitempty"`
	// List of legal entities associated with the current legal entity. For example, ultimate beneficial owners associated with an organization through ownership or control, or as signatories.
	EntityAssociations []LegalEntityAssociation `json:"entityAssociations,omitempty"`
	// The unique identifier of the legal entity.
	Id           string        `json:"id"`
	Individual   *Individual   `json:"individual,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	// List of verification errors related to capabilities for the legal entity.
	Problems []CapabilityProblem `json:"problems,omitempty"`
	// Your reference for the legal entity, maximum 150 characters.
	Reference          *string             `json:"reference,omitempty"`
	SoleProprietorship *SoleProprietorship `json:"soleProprietorship,omitempty"`
	// List of transfer instruments that the legal entity owns.
	TransferInstruments []TransferInstrumentReference `json:"transferInstruments,omitempty"`
	Trust               *Trust                        `json:"trust,omitempty"`
	// The type of legal entity.  Possible values: **individual**, **organization**, **soleProprietorship**, or **trust**.
	Type *string `json:"type,omitempty"`
	// List of verification deadlines and the capabilities that will be disallowed if verification errors are not resolved.
	VerificationDeadlines []VerificationDeadline `json:"verificationDeadlines,omitempty"`
	// A key-value pair that specifies the verification process for a legal entity. Set to **upfront** for upfront verification for [marketplaces](https://docs.adyen.com/marketplaces/onboard-users#upfront).
	VerificationPlan *string `json:"verificationPlan,omitempty"`
}

// NewLegalEntity instantiates a new LegalEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalEntity(id string) *LegalEntity {
	this := LegalEntity{}
	this.Id = id
	return &this
}

// NewLegalEntityWithDefaults instantiates a new LegalEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalEntityWithDefaults() *LegalEntity {
	this := LegalEntity{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *LegalEntity) GetCapabilities() map[string]LegalEntityCapability {
	if o == nil || common.IsNil(o.Capabilities) {
		var ret map[string]LegalEntityCapability
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetCapabilitiesOk() (*map[string]LegalEntityCapability, bool) {
	if o == nil || common.IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *LegalEntity) HasCapabilities() bool {
	if o != nil && !common.IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]LegalEntityCapability and assigns it to the Capabilities field.
func (o *LegalEntity) SetCapabilities(v map[string]LegalEntityCapability) {
	o.Capabilities = &v
}

// GetDocumentDetails returns the DocumentDetails field value if set, zero value otherwise.
func (o *LegalEntity) GetDocumentDetails() []DocumentReference {
	if o == nil || common.IsNil(o.DocumentDetails) {
		var ret []DocumentReference
		return ret
	}
	return o.DocumentDetails
}

// GetDocumentDetailsOk returns a tuple with the DocumentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetDocumentDetailsOk() ([]DocumentReference, bool) {
	if o == nil || common.IsNil(o.DocumentDetails) {
		return nil, false
	}
	return o.DocumentDetails, true
}

// HasDocumentDetails returns a boolean if a field has been set.
func (o *LegalEntity) HasDocumentDetails() bool {
	if o != nil && !common.IsNil(o.DocumentDetails) {
		return true
	}

	return false
}

// SetDocumentDetails gets a reference to the given []DocumentReference and assigns it to the DocumentDetails field.
func (o *LegalEntity) SetDocumentDetails(v []DocumentReference) {
	o.DocumentDetails = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
// Deprecated
func (o *LegalEntity) GetDocuments() []EntityReference {
	if o == nil || common.IsNil(o.Documents) {
		var ret []EntityReference
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LegalEntity) GetDocumentsOk() ([]EntityReference, bool) {
	if o == nil || common.IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *LegalEntity) HasDocuments() bool {
	if o != nil && !common.IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []EntityReference and assigns it to the Documents field.
// Deprecated
func (o *LegalEntity) SetDocuments(v []EntityReference) {
	o.Documents = v
}

// GetEntityAssociations returns the EntityAssociations field value if set, zero value otherwise.
func (o *LegalEntity) GetEntityAssociations() []LegalEntityAssociation {
	if o == nil || common.IsNil(o.EntityAssociations) {
		var ret []LegalEntityAssociation
		return ret
	}
	return o.EntityAssociations
}

// GetEntityAssociationsOk returns a tuple with the EntityAssociations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetEntityAssociationsOk() ([]LegalEntityAssociation, bool) {
	if o == nil || common.IsNil(o.EntityAssociations) {
		return nil, false
	}
	return o.EntityAssociations, true
}

// HasEntityAssociations returns a boolean if a field has been set.
func (o *LegalEntity) HasEntityAssociations() bool {
	if o != nil && !common.IsNil(o.EntityAssociations) {
		return true
	}

	return false
}

// SetEntityAssociations gets a reference to the given []LegalEntityAssociation and assigns it to the EntityAssociations field.
func (o *LegalEntity) SetEntityAssociations(v []LegalEntityAssociation) {
	o.EntityAssociations = v
}

// GetId returns the Id field value
func (o *LegalEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LegalEntity) SetId(v string) {
	o.Id = v
}

// GetIndividual returns the Individual field value if set, zero value otherwise.
func (o *LegalEntity) GetIndividual() Individual {
	if o == nil || common.IsNil(o.Individual) {
		var ret Individual
		return ret
	}
	return *o.Individual
}

// GetIndividualOk returns a tuple with the Individual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetIndividualOk() (*Individual, bool) {
	if o == nil || common.IsNil(o.Individual) {
		return nil, false
	}
	return o.Individual, true
}

// HasIndividual returns a boolean if a field has been set.
func (o *LegalEntity) HasIndividual() bool {
	if o != nil && !common.IsNil(o.Individual) {
		return true
	}

	return false
}

// SetIndividual gets a reference to the given Individual and assigns it to the Individual field.
func (o *LegalEntity) SetIndividual(v Individual) {
	o.Individual = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *LegalEntity) GetOrganization() Organization {
	if o == nil || common.IsNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetOrganizationOk() (*Organization, bool) {
	if o == nil || common.IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *LegalEntity) HasOrganization() bool {
	if o != nil && !common.IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *LegalEntity) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *LegalEntity) GetProblems() []CapabilityProblem {
	if o == nil || common.IsNil(o.Problems) {
		var ret []CapabilityProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetProblemsOk() ([]CapabilityProblem, bool) {
	if o == nil || common.IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *LegalEntity) HasProblems() bool {
	if o != nil && !common.IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []CapabilityProblem and assigns it to the Problems field.
func (o *LegalEntity) SetProblems(v []CapabilityProblem) {
	o.Problems = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LegalEntity) GetReference() string {
	if o == nil || common.IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetReferenceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LegalEntity) HasReference() bool {
	if o != nil && !common.IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LegalEntity) SetReference(v string) {
	o.Reference = &v
}

// GetSoleProprietorship returns the SoleProprietorship field value if set, zero value otherwise.
func (o *LegalEntity) GetSoleProprietorship() SoleProprietorship {
	if o == nil || common.IsNil(o.SoleProprietorship) {
		var ret SoleProprietorship
		return ret
	}
	return *o.SoleProprietorship
}

// GetSoleProprietorshipOk returns a tuple with the SoleProprietorship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetSoleProprietorshipOk() (*SoleProprietorship, bool) {
	if o == nil || common.IsNil(o.SoleProprietorship) {
		return nil, false
	}
	return o.SoleProprietorship, true
}

// HasSoleProprietorship returns a boolean if a field has been set.
func (o *LegalEntity) HasSoleProprietorship() bool {
	if o != nil && !common.IsNil(o.SoleProprietorship) {
		return true
	}

	return false
}

// SetSoleProprietorship gets a reference to the given SoleProprietorship and assigns it to the SoleProprietorship field.
func (o *LegalEntity) SetSoleProprietorship(v SoleProprietorship) {
	o.SoleProprietorship = &v
}

// GetTransferInstruments returns the TransferInstruments field value if set, zero value otherwise.
func (o *LegalEntity) GetTransferInstruments() []TransferInstrumentReference {
	if o == nil || common.IsNil(o.TransferInstruments) {
		var ret []TransferInstrumentReference
		return ret
	}
	return o.TransferInstruments
}

// GetTransferInstrumentsOk returns a tuple with the TransferInstruments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetTransferInstrumentsOk() ([]TransferInstrumentReference, bool) {
	if o == nil || common.IsNil(o.TransferInstruments) {
		return nil, false
	}
	return o.TransferInstruments, true
}

// HasTransferInstruments returns a boolean if a field has been set.
func (o *LegalEntity) HasTransferInstruments() bool {
	if o != nil && !common.IsNil(o.TransferInstruments) {
		return true
	}

	return false
}

// SetTransferInstruments gets a reference to the given []TransferInstrumentReference and assigns it to the TransferInstruments field.
func (o *LegalEntity) SetTransferInstruments(v []TransferInstrumentReference) {
	o.TransferInstruments = v
}

// GetTrust returns the Trust field value if set, zero value otherwise.
func (o *LegalEntity) GetTrust() Trust {
	if o == nil || common.IsNil(o.Trust) {
		var ret Trust
		return ret
	}
	return *o.Trust
}

// GetTrustOk returns a tuple with the Trust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetTrustOk() (*Trust, bool) {
	if o == nil || common.IsNil(o.Trust) {
		return nil, false
	}
	return o.Trust, true
}

// HasTrust returns a boolean if a field has been set.
func (o *LegalEntity) HasTrust() bool {
	if o != nil && !common.IsNil(o.Trust) {
		return true
	}

	return false
}

// SetTrust gets a reference to the given Trust and assigns it to the Trust field.
func (o *LegalEntity) SetTrust(v Trust) {
	o.Trust = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LegalEntity) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LegalEntity) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LegalEntity) SetType(v string) {
	o.Type = &v
}

// GetVerificationDeadlines returns the VerificationDeadlines field value if set, zero value otherwise.
func (o *LegalEntity) GetVerificationDeadlines() []VerificationDeadline {
	if o == nil || common.IsNil(o.VerificationDeadlines) {
		var ret []VerificationDeadline
		return ret
	}
	return o.VerificationDeadlines
}

// GetVerificationDeadlinesOk returns a tuple with the VerificationDeadlines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetVerificationDeadlinesOk() ([]VerificationDeadline, bool) {
	if o == nil || common.IsNil(o.VerificationDeadlines) {
		return nil, false
	}
	return o.VerificationDeadlines, true
}

// HasVerificationDeadlines returns a boolean if a field has been set.
func (o *LegalEntity) HasVerificationDeadlines() bool {
	if o != nil && !common.IsNil(o.VerificationDeadlines) {
		return true
	}

	return false
}

// SetVerificationDeadlines gets a reference to the given []VerificationDeadline and assigns it to the VerificationDeadlines field.
func (o *LegalEntity) SetVerificationDeadlines(v []VerificationDeadline) {
	o.VerificationDeadlines = v
}

// GetVerificationPlan returns the VerificationPlan field value if set, zero value otherwise.
func (o *LegalEntity) GetVerificationPlan() string {
	if o == nil || common.IsNil(o.VerificationPlan) {
		var ret string
		return ret
	}
	return *o.VerificationPlan
}

// GetVerificationPlanOk returns a tuple with the VerificationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalEntity) GetVerificationPlanOk() (*string, bool) {
	if o == nil || common.IsNil(o.VerificationPlan) {
		return nil, false
	}
	return o.VerificationPlan, true
}

// HasVerificationPlan returns a boolean if a field has been set.
func (o *LegalEntity) HasVerificationPlan() bool {
	if o != nil && !common.IsNil(o.VerificationPlan) {
		return true
	}

	return false
}

// SetVerificationPlan gets a reference to the given string and assigns it to the VerificationPlan field.
func (o *LegalEntity) SetVerificationPlan(v string) {
	o.VerificationPlan = &v
}

func (o LegalEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !common.IsNil(o.DocumentDetails) {
		toSerialize["documentDetails"] = o.DocumentDetails
	}
	if !common.IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !common.IsNil(o.EntityAssociations) {
		toSerialize["entityAssociations"] = o.EntityAssociations
	}
	toSerialize["id"] = o.Id
	if !common.IsNil(o.Individual) {
		toSerialize["individual"] = o.Individual
	}
	if !common.IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !common.IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	if !common.IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !common.IsNil(o.SoleProprietorship) {
		toSerialize["soleProprietorship"] = o.SoleProprietorship
	}
	if !common.IsNil(o.TransferInstruments) {
		toSerialize["transferInstruments"] = o.TransferInstruments
	}
	if !common.IsNil(o.Trust) {
		toSerialize["trust"] = o.Trust
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.VerificationDeadlines) {
		toSerialize["verificationDeadlines"] = o.VerificationDeadlines
	}
	if !common.IsNil(o.VerificationPlan) {
		toSerialize["verificationPlan"] = o.VerificationPlan
	}
	return toSerialize, nil
}

type NullableLegalEntity struct {
	value *LegalEntity
	isSet bool
}

func (v NullableLegalEntity) Get() *LegalEntity {
	return v.value
}

func (v *NullableLegalEntity) Set(val *LegalEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalEntity(val *LegalEntity) *NullableLegalEntity {
	return &NullableLegalEntity{value: val, isSet: true}
}

func (v NullableLegalEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

func (o *LegalEntity) isValidType() bool {
	var allowedEnumValues = []string{"individual", "organization", "soleProprietorship", "trust", "unincorporatedPartnership"}
	for _, allowed := range allowedEnumValues {
		if o.GetType() == allowed {
			return true
		}
	}
	return false
}
