/*
Transfer webhooks

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transferwebhook

import (
	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// APIClient manages communication with the Transfer webhooks API v3
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	common common.Service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
}

// NewAPIClient creates a new API client.
func NewAPIClient(client *common.Client) *APIClient {
	c := &APIClient{}
	c.common.Client = client
	c.common.BasePath = func() string {
		return client.Cfg.Endpoint
	}

	// API Services

	return c
}
