/*
Management API

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// APIKeyCompanyLevelApi service
type APIKeyCompanyLevelApi common.Service

// All parameters accepted by APIKeyCompanyLevelApi.GenerateNewApiKey
type APIKeyCompanyLevelApiGenerateNewApiKeyInput struct {
	companyId       string
	apiCredentialId string
}

/*
Prepare a request for GenerateNewApiKey
@param companyId The unique identifier of the company account.@param apiCredentialId Unique identifier of the API credential.
@return APIKeyCompanyLevelApiGenerateNewApiKeyInput
*/
func (a *APIKeyCompanyLevelApi) GenerateNewApiKeyInput(companyId string, apiCredentialId string) APIKeyCompanyLevelApiGenerateNewApiKeyInput {
	return APIKeyCompanyLevelApiGenerateNewApiKeyInput{
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
	}
}

/*
GenerateNewApiKey Generate new API key

Returns a new API key for the API credential. You can use the new API key a few minutes after generating it. The old API key stops working 24 hours after generating a new one.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r APIKeyCompanyLevelApiGenerateNewApiKeyInput - Request parameters, see GenerateNewApiKeyInput
@return GenerateApiKeyResponse, *http.Response, error
*/
func (a *APIKeyCompanyLevelApi) GenerateNewApiKey(ctx context.Context, r APIKeyCompanyLevelApiGenerateNewApiKeyInput) (GenerateApiKeyResponse, *http.Response, error) {
	res := &GenerateApiKeyResponse{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/generateApiKey"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return *res, httpRes, err
	}

	var serviceError common.RestServiceError
	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}
	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return *res, httpRes, decodeError
		}
		return *res, httpRes, serviceError
	}

	return *res, httpRes, err
}
