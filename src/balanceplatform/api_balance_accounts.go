/*
Configuration API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package balanceplatform

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

// BalanceAccountsApi service
type BalanceAccountsApi common.Service

// All parameters accepted by BalanceAccountsApi.CreateBalanceAccount
type BalanceAccountsApiCreateBalanceAccountInput struct {
	balanceAccountInfo *BalanceAccountInfo
}

func (r BalanceAccountsApiCreateBalanceAccountInput) BalanceAccountInfo(balanceAccountInfo BalanceAccountInfo) BalanceAccountsApiCreateBalanceAccountInput {
	r.balanceAccountInfo = &balanceAccountInfo
	return r
}

/*
Prepare a request for CreateBalanceAccount

@return BalanceAccountsApiCreateBalanceAccountInput
*/
func (a *BalanceAccountsApi) CreateBalanceAccountInput() BalanceAccountsApiCreateBalanceAccountInput {
	return BalanceAccountsApiCreateBalanceAccountInput{}
}

/*
CreateBalanceAccount Create a balance account

Creates a balance account that holds the funds of the associated account holder.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiCreateBalanceAccountInput - Request parameters, see CreateBalanceAccountInput
@return BalanceAccount, *http.Response, error
*/
func (a *BalanceAccountsApi) CreateBalanceAccount(ctx context.Context, r BalanceAccountsApiCreateBalanceAccountInput) (BalanceAccount, *http.Response, error) {
	res := &BalanceAccount{}
	path := "/balanceAccounts"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.balanceAccountInfo,
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

// All parameters accepted by BalanceAccountsApi.CreateSweep
type BalanceAccountsApiCreateSweepInput struct {
	balanceAccountId     string
	sweepConfigurationV2 *SweepConfigurationV2
}

func (r BalanceAccountsApiCreateSweepInput) SweepConfigurationV2(sweepConfigurationV2 SweepConfigurationV2) BalanceAccountsApiCreateSweepInput {
	r.sweepConfigurationV2 = &sweepConfigurationV2
	return r
}

/*
Prepare a request for CreateSweep
@param balanceAccountId The unique identifier of the balance account.
@return BalanceAccountsApiCreateSweepInput
*/
func (a *BalanceAccountsApi) CreateSweepInput(balanceAccountId string) BalanceAccountsApiCreateSweepInput {
	return BalanceAccountsApiCreateSweepInput{
		balanceAccountId: balanceAccountId,
	}
}

/*
CreateSweep Create a sweep

Creates a sweep that results in moving funds from or to a balance account.

A sweep pulls in or pushes out funds based on a defined schedule, amount, currency, and a source or a destination.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiCreateSweepInput - Request parameters, see CreateSweepInput
@return SweepConfigurationV2, *http.Response, error
*/
func (a *BalanceAccountsApi) CreateSweep(ctx context.Context, r BalanceAccountsApiCreateSweepInput) (SweepConfigurationV2, *http.Response, error) {
	res := &SweepConfigurationV2{}
	path := "/balanceAccounts/{balanceAccountId}/sweeps"
	path = strings.Replace(path, "{"+"balanceAccountId"+"}", url.PathEscape(common.ParameterValueToString(r.balanceAccountId, "balanceAccountId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.sweepConfigurationV2,
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

// All parameters accepted by BalanceAccountsApi.DeleteSweep
type BalanceAccountsApiDeleteSweepInput struct {
	balanceAccountId string
	sweepId          string
}

/*
Prepare a request for DeleteSweep
@param balanceAccountId The unique identifier of the balance account.@param sweepId The unique identifier of the sweep.
@return BalanceAccountsApiDeleteSweepInput
*/
func (a *BalanceAccountsApi) DeleteSweepInput(balanceAccountId string, sweepId string) BalanceAccountsApiDeleteSweepInput {
	return BalanceAccountsApiDeleteSweepInput{
		balanceAccountId: balanceAccountId,
		sweepId:          sweepId,
	}
}

/*
DeleteSweep Delete a sweep

Deletes a sweep for a balance account.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiDeleteSweepInput - Request parameters, see DeleteSweepInput
@return *http.Response, error
*/
func (a *BalanceAccountsApi) DeleteSweep(ctx context.Context, r BalanceAccountsApiDeleteSweepInput) (*http.Response, error) {
	var res interface{}
	path := "/balanceAccounts/{balanceAccountId}/sweeps/{sweepId}"
	path = strings.Replace(path, "{"+"balanceAccountId"+"}", url.PathEscape(common.ParameterValueToString(r.balanceAccountId, "balanceAccountId")), -1)
	path = strings.Replace(path, "{"+"sweepId"+"}", url.PathEscape(common.ParameterValueToString(r.sweepId, "sweepId")), -1)
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

// All parameters accepted by BalanceAccountsApi.GetAllPaymentInstrumentsForBalanceAccount
type BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput struct {
	id     string
	offset *int32
	limit  *int32
}

// The number of items that you want to skip.
func (r BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput) Offset(offset int32) BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput {
	r.offset = &offset
	return r
}

// The number of items returned per page, maximum 100 items. By default, the response returns 10 items per page.
func (r BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput) Limit(limit int32) BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetAllPaymentInstrumentsForBalanceAccount
@param id The unique identifier of the balance account.
@return BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput
*/
func (a *BalanceAccountsApi) GetAllPaymentInstrumentsForBalanceAccountInput(id string) BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput {
	return BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput{
		id: id,
	}
}

/*
GetAllPaymentInstrumentsForBalanceAccount Get all payment instruments for a balance account

Returns a paginated list of the payment instruments associated with a balance account.

To fetch multiple pages, use the query parameters.For example, to limit the page to 3 payment instruments and to skip the first 6, use `/balanceAccounts/{id}/paymentInstruments?limit=3&offset=6`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput - Request parameters, see GetAllPaymentInstrumentsForBalanceAccountInput
@return PaginatedPaymentInstrumentsResponse, *http.Response, error
*/
func (a *BalanceAccountsApi) GetAllPaymentInstrumentsForBalanceAccount(ctx context.Context, r BalanceAccountsApiGetAllPaymentInstrumentsForBalanceAccountInput) (PaginatedPaymentInstrumentsResponse, *http.Response, error) {
	res := &PaginatedPaymentInstrumentsResponse{}
	path := "/balanceAccounts/{id}/paymentInstruments"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
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

// All parameters accepted by BalanceAccountsApi.GetAllSweepsForBalanceAccount
type BalanceAccountsApiGetAllSweepsForBalanceAccountInput struct {
	balanceAccountId string
	offset           *int32
	limit            *int32
}

// The number of items that you want to skip.
func (r BalanceAccountsApiGetAllSweepsForBalanceAccountInput) Offset(offset int32) BalanceAccountsApiGetAllSweepsForBalanceAccountInput {
	r.offset = &offset
	return r
}

// The number of items returned per page, maximum 100 items. By default, the response returns 10 items per page.
func (r BalanceAccountsApiGetAllSweepsForBalanceAccountInput) Limit(limit int32) BalanceAccountsApiGetAllSweepsForBalanceAccountInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetAllSweepsForBalanceAccount
@param balanceAccountId The unique identifier of the balance account.
@return BalanceAccountsApiGetAllSweepsForBalanceAccountInput
*/
func (a *BalanceAccountsApi) GetAllSweepsForBalanceAccountInput(balanceAccountId string) BalanceAccountsApiGetAllSweepsForBalanceAccountInput {
	return BalanceAccountsApiGetAllSweepsForBalanceAccountInput{
		balanceAccountId: balanceAccountId,
	}
}

/*
GetAllSweepsForBalanceAccount Get all sweeps for a balance account

Returns a list of the sweeps configured for a balance account.

To fetch multiple pages, use the query parameters. For example, to limit the page to 5 sweeps and to skip the first 10, use `/balanceAccounts/{balanceAccountId}/sweeps?limit=5&offset=10`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiGetAllSweepsForBalanceAccountInput - Request parameters, see GetAllSweepsForBalanceAccountInput
@return BalanceSweepConfigurationsResponse, *http.Response, error
*/
func (a *BalanceAccountsApi) GetAllSweepsForBalanceAccount(ctx context.Context, r BalanceAccountsApiGetAllSweepsForBalanceAccountInput) (BalanceSweepConfigurationsResponse, *http.Response, error) {
	res := &BalanceSweepConfigurationsResponse{}
	path := "/balanceAccounts/{balanceAccountId}/sweeps"
	path = strings.Replace(path, "{"+"balanceAccountId"+"}", url.PathEscape(common.ParameterValueToString(r.balanceAccountId, "balanceAccountId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.offset != nil {
		common.ParameterAddToQuery(queryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		common.ParameterAddToQuery(queryParams, "limit", r.limit, "")
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

// All parameters accepted by BalanceAccountsApi.GetBalanceAccount
type BalanceAccountsApiGetBalanceAccountInput struct {
	id string
}

/*
Prepare a request for GetBalanceAccount
@param id The unique identifier of the balance account.
@return BalanceAccountsApiGetBalanceAccountInput
*/
func (a *BalanceAccountsApi) GetBalanceAccountInput(id string) BalanceAccountsApiGetBalanceAccountInput {
	return BalanceAccountsApiGetBalanceAccountInput{
		id: id,
	}
}

/*
GetBalanceAccount Get a balance account

Returns a balance account and its balances for the default currency and other currencies with a non-zero balance.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiGetBalanceAccountInput - Request parameters, see GetBalanceAccountInput
@return BalanceAccount, *http.Response, error
*/
func (a *BalanceAccountsApi) GetBalanceAccount(ctx context.Context, r BalanceAccountsApiGetBalanceAccountInput) (BalanceAccount, *http.Response, error) {
	res := &BalanceAccount{}
	path := "/balanceAccounts/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
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

// All parameters accepted by BalanceAccountsApi.GetSweep
type BalanceAccountsApiGetSweepInput struct {
	balanceAccountId string
	sweepId          string
}

/*
Prepare a request for GetSweep
@param balanceAccountId The unique identifier of the balance account.@param sweepId The unique identifier of the sweep.
@return BalanceAccountsApiGetSweepInput
*/
func (a *BalanceAccountsApi) GetSweepInput(balanceAccountId string, sweepId string) BalanceAccountsApiGetSweepInput {
	return BalanceAccountsApiGetSweepInput{
		balanceAccountId: balanceAccountId,
		sweepId:          sweepId,
	}
}

/*
GetSweep Get a sweep

Returns a sweep.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiGetSweepInput - Request parameters, see GetSweepInput
@return SweepConfigurationV2, *http.Response, error
*/
func (a *BalanceAccountsApi) GetSweep(ctx context.Context, r BalanceAccountsApiGetSweepInput) (SweepConfigurationV2, *http.Response, error) {
	res := &SweepConfigurationV2{}
	path := "/balanceAccounts/{balanceAccountId}/sweeps/{sweepId}"
	path = strings.Replace(path, "{"+"balanceAccountId"+"}", url.PathEscape(common.ParameterValueToString(r.balanceAccountId, "balanceAccountId")), -1)
	path = strings.Replace(path, "{"+"sweepId"+"}", url.PathEscape(common.ParameterValueToString(r.sweepId, "sweepId")), -1)
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

// All parameters accepted by BalanceAccountsApi.UpdateBalanceAccount
type BalanceAccountsApiUpdateBalanceAccountInput struct {
	id                          string
	balanceAccountUpdateRequest *BalanceAccountUpdateRequest
}

func (r BalanceAccountsApiUpdateBalanceAccountInput) BalanceAccountUpdateRequest(balanceAccountUpdateRequest BalanceAccountUpdateRequest) BalanceAccountsApiUpdateBalanceAccountInput {
	r.balanceAccountUpdateRequest = &balanceAccountUpdateRequest
	return r
}

/*
Prepare a request for UpdateBalanceAccount
@param id The unique identifier of the balance account.
@return BalanceAccountsApiUpdateBalanceAccountInput
*/
func (a *BalanceAccountsApi) UpdateBalanceAccountInput(id string) BalanceAccountsApiUpdateBalanceAccountInput {
	return BalanceAccountsApiUpdateBalanceAccountInput{
		id: id,
	}
}

/*
UpdateBalanceAccount Update a balance account

Updates a balance account.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiUpdateBalanceAccountInput - Request parameters, see UpdateBalanceAccountInput
@return BalanceAccount, *http.Response, error
*/
func (a *BalanceAccountsApi) UpdateBalanceAccount(ctx context.Context, r BalanceAccountsApiUpdateBalanceAccountInput) (BalanceAccount, *http.Response, error) {
	res := &BalanceAccount{}
	path := "/balanceAccounts/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.balanceAccountUpdateRequest,
		res,
		http.MethodPatch,
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

// All parameters accepted by BalanceAccountsApi.UpdateSweep
type BalanceAccountsApiUpdateSweepInput struct {
	balanceAccountId           string
	sweepId                    string
	updateSweepConfigurationV2 *UpdateSweepConfigurationV2
}

func (r BalanceAccountsApiUpdateSweepInput) UpdateSweepConfigurationV2(updateSweepConfigurationV2 UpdateSweepConfigurationV2) BalanceAccountsApiUpdateSweepInput {
	r.updateSweepConfigurationV2 = &updateSweepConfigurationV2
	return r
}

/*
Prepare a request for UpdateSweep
@param balanceAccountId The unique identifier of the balance account.@param sweepId The unique identifier of the sweep.
@return BalanceAccountsApiUpdateSweepInput
*/
func (a *BalanceAccountsApi) UpdateSweepInput(balanceAccountId string, sweepId string) BalanceAccountsApiUpdateSweepInput {
	return BalanceAccountsApiUpdateSweepInput{
		balanceAccountId: balanceAccountId,
		sweepId:          sweepId,
	}
}

/*
UpdateSweep Update a sweep

Updates a sweep. When updating a sweep resource, note that if a request parameter is not provided, the parameter is left unchanged.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r BalanceAccountsApiUpdateSweepInput - Request parameters, see UpdateSweepInput
@return SweepConfigurationV2, *http.Response, error
*/
func (a *BalanceAccountsApi) UpdateSweep(ctx context.Context, r BalanceAccountsApiUpdateSweepInput) (SweepConfigurationV2, *http.Response, error) {
	res := &SweepConfigurationV2{}
	path := "/balanceAccounts/{balanceAccountId}/sweeps/{sweepId}"
	path = strings.Replace(path, "{"+"balanceAccountId"+"}", url.PathEscape(common.ParameterValueToString(r.balanceAccountId, "balanceAccountId")), -1)
	path = strings.Replace(path, "{"+"sweepId"+"}", url.PathEscape(common.ParameterValueToString(r.sweepId, "sweepId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.updateSweepConfigurationV2,
		res,
		http.MethodPatch,
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
