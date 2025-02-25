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

	"github.com/adyen/adyen-go-api-library/v19/src/common"
)

// TerminalsTerminalLevelApi service
type TerminalsTerminalLevelApi common.Service

// All parameters accepted by TerminalsTerminalLevelApi.ListTerminals
type TerminalsTerminalLevelApiListTerminalsInput struct {
	searchQuery *string
	otpQuery    *string
	countries   *string
	merchantIds *string
	storeIds    *string
	brandModels *string
	pageNumber  *int32
	pageSize    *int32
}

// Returns terminals with an ID that contains the specified string. If present, other query parameters are ignored.
func (r TerminalsTerminalLevelApiListTerminalsInput) SearchQuery(searchQuery string) TerminalsTerminalLevelApiListTerminalsInput {
	r.searchQuery = &searchQuery
	return r
}

// Returns one or more terminals associated with the one-time passwords specified in the request. If this query parameter is used, other query parameters are ignored.
func (r TerminalsTerminalLevelApiListTerminalsInput) OtpQuery(otpQuery string) TerminalsTerminalLevelApiListTerminalsInput {
	r.otpQuery = &otpQuery
	return r
}

// Returns terminals located in the countries specified by their [two-letter country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
func (r TerminalsTerminalLevelApiListTerminalsInput) Countries(countries string) TerminalsTerminalLevelApiListTerminalsInput {
	r.countries = &countries
	return r
}

// Returns terminals that belong to the merchant accounts specified by their unique merchant account ID.
func (r TerminalsTerminalLevelApiListTerminalsInput) MerchantIds(merchantIds string) TerminalsTerminalLevelApiListTerminalsInput {
	r.merchantIds = &merchantIds
	return r
}

// Returns terminals that are assigned to the [stores](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/stores) specified by their unique store ID.
func (r TerminalsTerminalLevelApiListTerminalsInput) StoreIds(storeIds string) TerminalsTerminalLevelApiListTerminalsInput {
	r.storeIds = &storeIds
	return r
}

// Returns terminals of the [models](https://docs.adyen.com/api-explorer/#/ManagementService/latest/get/companies/{companyId}/terminalModels) specified in the format *brand.model*.
func (r TerminalsTerminalLevelApiListTerminalsInput) BrandModels(brandModels string) TerminalsTerminalLevelApiListTerminalsInput {
	r.brandModels = &brandModels
	return r
}

// The number of the page to fetch.
func (r TerminalsTerminalLevelApiListTerminalsInput) PageNumber(pageNumber int32) TerminalsTerminalLevelApiListTerminalsInput {
	r.pageNumber = &pageNumber
	return r
}

// The number of items to have on a page, maximum 100. The default is 20 items on a page.
func (r TerminalsTerminalLevelApiListTerminalsInput) PageSize(pageSize int32) TerminalsTerminalLevelApiListTerminalsInput {
	r.pageSize = &pageSize
	return r
}

/*
Prepare a request for ListTerminals

@return TerminalsTerminalLevelApiListTerminalsInput
*/
func (a *TerminalsTerminalLevelApi) ListTerminalsInput() TerminalsTerminalLevelApiListTerminalsInput {
	return TerminalsTerminalLevelApiListTerminalsInput{}
}

/*
ListTerminals Get a list of terminals

Returns the payment terminals that the API credential has access to and that match the query parameters.
To make this request, your API credential must have the following [roles](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API — Terminal actions read

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalsTerminalLevelApiListTerminalsInput - Request parameters, see ListTerminalsInput
@return ListTerminalsResponse, *http.Response, error
*/
func (a *TerminalsTerminalLevelApi) ListTerminals(ctx context.Context, r TerminalsTerminalLevelApiListTerminalsInput) (ListTerminalsResponse, *http.Response, error) {
	res := &ListTerminalsResponse{}
	path := "/terminals"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.searchQuery != nil {
		common.ParameterAddToQuery(queryParams, "searchQuery", r.searchQuery, "")
	}
	if r.otpQuery != nil {
		common.ParameterAddToQuery(queryParams, "otpQuery", r.otpQuery, "")
	}
	if r.countries != nil {
		common.ParameterAddToQuery(queryParams, "countries", r.countries, "")
	}
	if r.merchantIds != nil {
		common.ParameterAddToQuery(queryParams, "merchantIds", r.merchantIds, "")
	}
	if r.storeIds != nil {
		common.ParameterAddToQuery(queryParams, "storeIds", r.storeIds, "")
	}
	if r.brandModels != nil {
		common.ParameterAddToQuery(queryParams, "brandModels", r.brandModels, "")
	}
	if r.pageNumber != nil {
		common.ParameterAddToQuery(queryParams, "pageNumber", r.pageNumber, "")
	}
	if r.pageSize != nil {
		common.ParameterAddToQuery(queryParams, "pageSize", r.pageSize, "")
	}
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

// All parameters accepted by TerminalsTerminalLevelApi.ReassignTerminal
type TerminalsTerminalLevelApiReassignTerminalInput struct {
	terminalId                  string
	terminalReassignmentRequest *TerminalReassignmentRequest
}

func (r TerminalsTerminalLevelApiReassignTerminalInput) TerminalReassignmentRequest(terminalReassignmentRequest TerminalReassignmentRequest) TerminalsTerminalLevelApiReassignTerminalInput {
	r.terminalReassignmentRequest = &terminalReassignmentRequest
	return r
}

/*
Prepare a request for ReassignTerminal
@param terminalId The unique identifier of the payment terminal.
@return TerminalsTerminalLevelApiReassignTerminalInput
*/
func (a *TerminalsTerminalLevelApi) ReassignTerminalInput(terminalId string) TerminalsTerminalLevelApiReassignTerminalInput {
	return TerminalsTerminalLevelApiReassignTerminalInput{
		terminalId: terminalId,
	}
}

/*
ReassignTerminal Reassign a terminal

Reassigns a payment terminal to a company account, merchant account, merchant account inventory, or a store.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Assign Terminal

In the live environment, requests to this endpoint are subject to [rate limits](https://docs.adyen.com/point-of-sale/automating-terminal-management#rate-limits-in-the-live-environment).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalsTerminalLevelApiReassignTerminalInput - Request parameters, see ReassignTerminalInput
@return *http.Response, error
*/
func (a *TerminalsTerminalLevelApi) ReassignTerminal(ctx context.Context, r TerminalsTerminalLevelApiReassignTerminalInput) (*http.Response, error) {
	var res interface{}
	path := "/terminals/{terminalId}/reassign"
	path = strings.Replace(path, "{"+"terminalId"+"}", url.PathEscape(common.ParameterValueToString(r.terminalId, "terminalId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.terminalReassignmentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	if httpRes == nil {
		return httpRes, err
	}

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
