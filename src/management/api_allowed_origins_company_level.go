/*
Management API

API version: 1
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

	"github.com/adyen/adyen-go-api-library/v7/src/common"
)

// AllowedOriginsCompanyLevelApi service
type AllowedOriginsCompanyLevelApi common.Service

// All parameters accepted by AllowedOriginsCompanyLevelApi.CreateAllowedOrigin
type AllowedOriginsCompanyLevelApiCreateAllowedOriginInput struct {
	companyId       string
	apiCredentialId string
	allowedOrigin   *AllowedOrigin
}

func (r AllowedOriginsCompanyLevelApiCreateAllowedOriginInput) AllowedOrigin(allowedOrigin AllowedOrigin) AllowedOriginsCompanyLevelApiCreateAllowedOriginInput {
	r.allowedOrigin = &allowedOrigin
	return r
}

/*
Prepare a request for CreateAllowedOrigin
@param companyId The unique identifier of the company account.@param apiCredentialId Unique identifier of the API credential.
@return AllowedOriginsCompanyLevelApiCreateAllowedOriginInput
*/
func (a *AllowedOriginsCompanyLevelApi) CreateAllowedOriginInput(companyId string, apiCredentialId string) AllowedOriginsCompanyLevelApiCreateAllowedOriginInput {
	return AllowedOriginsCompanyLevelApiCreateAllowedOriginInput{
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
	}
}

/*
CreateAllowedOrigin Create an allowed origin

Adds a new [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) to the API credential's list of allowed origins.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsCompanyLevelApiCreateAllowedOriginInput - Request parameters, see CreateAllowedOriginInput
@return AllowedOriginsResponse, *http.Response, error
*/
func (a *AllowedOriginsCompanyLevelApi) CreateAllowedOrigin(ctx context.Context, r AllowedOriginsCompanyLevelApiCreateAllowedOriginInput) (AllowedOriginsResponse, *http.Response, error) {
	res := &AllowedOriginsResponse{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.allowedOrigin,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

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

// All parameters accepted by AllowedOriginsCompanyLevelApi.DeleteAllowedOrigin
type AllowedOriginsCompanyLevelApiDeleteAllowedOriginInput struct {
	companyId       string
	apiCredentialId string
	originId        string
}

/*
Prepare a request for DeleteAllowedOrigin
@param companyId The unique identifier of the company account.@param apiCredentialId Unique identifier of the API credential.@param originId Unique identifier of the allowed origin.
@return AllowedOriginsCompanyLevelApiDeleteAllowedOriginInput
*/
func (a *AllowedOriginsCompanyLevelApi) DeleteAllowedOriginInput(companyId string, apiCredentialId string, originId string) AllowedOriginsCompanyLevelApiDeleteAllowedOriginInput {
	return AllowedOriginsCompanyLevelApiDeleteAllowedOriginInput{
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
		originId:        originId,
	}
}

/*
DeleteAllowedOrigin Delete an allowed origin

Removes the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path. As soon as an allowed origin is removed, we no longer accept client-side requests from that domain.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsCompanyLevelApiDeleteAllowedOriginInput - Request parameters, see DeleteAllowedOriginInput
@return *http.Response, error
*/
func (a *AllowedOriginsCompanyLevelApi) DeleteAllowedOrigin(ctx context.Context, r AllowedOriginsCompanyLevelApiDeleteAllowedOriginInput) (*http.Response, error) {
	var res interface{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	path = strings.Replace(path, "{"+"originId"+"}", url.PathEscape(common.ParameterValueToString(r.originId, "originId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodDelete,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	var serviceError common.RestServiceError

	if httpRes.StatusCode == 400 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 401 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 403 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 422 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	if httpRes.StatusCode == 500 {
		body, _ := ioutil.ReadAll(httpRes.Body)
		decodeError := json.Unmarshal([]byte(body), &serviceError)
		if decodeError != nil {
			return httpRes, decodeError
		}
		return httpRes, serviceError
	}

	return httpRes, err
}

// All parameters accepted by AllowedOriginsCompanyLevelApi.GetAllowedOrigin
type AllowedOriginsCompanyLevelApiGetAllowedOriginInput struct {
	companyId       string
	apiCredentialId string
	originId        string
}

/*
Prepare a request for GetAllowedOrigin
@param companyId The unique identifier of the company account.@param apiCredentialId Unique identifier of the API credential.@param originId Unique identifier of the allowed origin.
@return AllowedOriginsCompanyLevelApiGetAllowedOriginInput
*/
func (a *AllowedOriginsCompanyLevelApi) GetAllowedOriginInput(companyId string, apiCredentialId string, originId string) AllowedOriginsCompanyLevelApiGetAllowedOriginInput {
	return AllowedOriginsCompanyLevelApiGetAllowedOriginInput{
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
		originId:        originId,
	}
}

/*
GetAllowedOrigin Get an allowed origin

Returns the [allowed origin](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsCompanyLevelApiGetAllowedOriginInput - Request parameters, see GetAllowedOriginInput
@return AllowedOrigin, *http.Response, error
*/
func (a *AllowedOriginsCompanyLevelApi) GetAllowedOrigin(ctx context.Context, r AllowedOriginsCompanyLevelApiGetAllowedOriginInput) (AllowedOrigin, *http.Response, error) {
	res := &AllowedOrigin{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins/{originId}"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	path = strings.Replace(path, "{"+"originId"+"}", url.PathEscape(common.ParameterValueToString(r.originId, "originId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

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

// All parameters accepted by AllowedOriginsCompanyLevelApi.ListAllowedOrigins
type AllowedOriginsCompanyLevelApiListAllowedOriginsInput struct {
	companyId       string
	apiCredentialId string
}

/*
Prepare a request for ListAllowedOrigins
@param companyId The unique identifier of the company account.@param apiCredentialId Unique identifier of the API credential.
@return AllowedOriginsCompanyLevelApiListAllowedOriginsInput
*/
func (a *AllowedOriginsCompanyLevelApi) ListAllowedOriginsInput(companyId string, apiCredentialId string) AllowedOriginsCompanyLevelApiListAllowedOriginsInput {
	return AllowedOriginsCompanyLevelApiListAllowedOriginsInput{
		companyId:       companyId,
		apiCredentialId: apiCredentialId,
	}
}

/*
ListAllowedOrigins Get a list of allowed origins

Returns the list of [allowed origins](https://docs.adyen.com/development-resources/client-side-authentication#allowed-origins) for the API credential identified in the path.

To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—API credentials read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AllowedOriginsCompanyLevelApiListAllowedOriginsInput - Request parameters, see ListAllowedOriginsInput
@return AllowedOriginsResponse, *http.Response, error
*/
func (a *AllowedOriginsCompanyLevelApi) ListAllowedOrigins(ctx context.Context, r AllowedOriginsCompanyLevelApiListAllowedOriginsInput) (AllowedOriginsResponse, *http.Response, error) {
	res := &AllowedOriginsResponse{}
	path := "/companies/{companyId}/apiCredentials/{apiCredentialId}/allowedOrigins"
	path = strings.Replace(path, "{"+"companyId"+"}", url.PathEscape(common.ParameterValueToString(r.companyId, "companyId")), -1)
	path = strings.Replace(path, "{"+"apiCredentialId"+"}", url.PathEscape(common.ParameterValueToString(r.apiCredentialId, "apiCredentialId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		nil,
		res,
		http.MethodGet,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

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
