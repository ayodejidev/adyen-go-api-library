/*
Legal Entity Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package legalentity

import (
	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// APIClient manages communication with the Legal Entity Management API API v3
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	BusinessLinesApi *BusinessLinesApi

	DocumentsApi *DocumentsApi

	HostedOnboardingApi *HostedOnboardingApi

	LegalEntitiesApi *LegalEntitiesApi

	PCIQuestionnairesApi *PCIQuestionnairesApi

	TaxEDeliveryConsentApi *TaxEDeliveryConsentApi

	TermsOfServiceApi *TermsOfServiceApi

	TransferInstrumentsApi *TransferInstrumentsApi
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
	c.common.Client = client
	c.common.BasePath = func() string {
		return client.Cfg.LegalEntityEndpoint
	}

	// API Services
	c.BusinessLinesApi = (*BusinessLinesApi)(&c.common)
	c.DocumentsApi = (*DocumentsApi)(&c.common)
	c.HostedOnboardingApi = (*HostedOnboardingApi)(&c.common)
	c.LegalEntitiesApi = (*LegalEntitiesApi)(&c.common)
	c.PCIQuestionnairesApi = (*PCIQuestionnairesApi)(&c.common)
	c.TaxEDeliveryConsentApi = (*TaxEDeliveryConsentApi)(&c.common)
	c.TermsOfServiceApi = (*TermsOfServiceApi)(&c.common)
	c.TransferInstrumentsApi = (*TransferInstrumentsApi)(&c.common)

	return c
}
