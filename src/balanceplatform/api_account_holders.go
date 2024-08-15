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

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// AccountHoldersApi service
type AccountHoldersApi common.Service

// All parameters accepted by AccountHoldersApi.CreateAccountHolder
type AccountHoldersApiCreateAccountHolderInput struct {
	accountHolderInfo *AccountHolderInfo
}

func (r AccountHoldersApiCreateAccountHolderInput) AccountHolderInfo(accountHolderInfo AccountHolderInfo) AccountHoldersApiCreateAccountHolderInput {
	r.accountHolderInfo = &accountHolderInfo
	return r
}

/*
Prepare a request for CreateAccountHolder

@return AccountHoldersApiCreateAccountHolderInput
*/
func (a *AccountHoldersApi) CreateAccountHolderInput() AccountHoldersApiCreateAccountHolderInput {
	return AccountHoldersApiCreateAccountHolderInput{}
}

/*
CreateAccountHolder Create an account holder

Creates an account holder linked to a [legal entity](https://docs.adyen.com/api-explorer/#/legalentity/latest/post/legalEntities).



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountHoldersApiCreateAccountHolderInput - Request parameters, see CreateAccountHolderInput
@return AccountHolder, *http.Response, error
*/
func (a *AccountHoldersApi) CreateAccountHolder(ctx context.Context, r AccountHoldersApiCreateAccountHolderInput) (AccountHolder, *http.Response, error) {
	res := &AccountHolder{}
	path := "/accountHolders"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.accountHolderInfo,
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

// All parameters accepted by AccountHoldersApi.GetAccountHolder
type AccountHoldersApiGetAccountHolderInput struct {
	id string
}

/*
Prepare a request for GetAccountHolder
@param id The unique identifier of the account holder.
@return AccountHoldersApiGetAccountHolderInput
*/
func (a *AccountHoldersApi) GetAccountHolderInput(id string) AccountHoldersApiGetAccountHolderInput {
	return AccountHoldersApiGetAccountHolderInput{
		id: id,
	}
}

/*
GetAccountHolder Get an account holder

Returns an account holder.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountHoldersApiGetAccountHolderInput - Request parameters, see GetAccountHolderInput
@return AccountHolder, *http.Response, error
*/
func (a *AccountHoldersApi) GetAccountHolder(ctx context.Context, r AccountHoldersApiGetAccountHolderInput) (AccountHolder, *http.Response, error) {
	res := &AccountHolder{}
	path := "/accountHolders/{id}"
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

// All parameters accepted by AccountHoldersApi.GetAllBalanceAccountsOfAccountHolder
type AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput struct {
	id     string
	offset *int32
	limit  *int32
}

// The number of items that you want to skip.
func (r AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput) Offset(offset int32) AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput {
	r.offset = &offset
	return r
}

// The number of items returned per page, maximum 100 items. By default, the response returns 10 items per page.
func (r AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput) Limit(limit int32) AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput {
	r.limit = &limit
	return r
}

/*
Prepare a request for GetAllBalanceAccountsOfAccountHolder
@param id The unique identifier of the account holder.
@return AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput
*/
func (a *AccountHoldersApi) GetAllBalanceAccountsOfAccountHolderInput(id string) AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput {
	return AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput{
		id: id,
	}
}

/*
GetAllBalanceAccountsOfAccountHolder Get all balance accounts of an account holder

Returns a paginated list of the balance accounts associated with an account holder. To fetch multiple pages, use the query parameters.

For example, to limit the page to 5 balance accounts and skip the first 10, use `/accountHolders/{id}/balanceAccounts?limit=5&offset=10`.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput - Request parameters, see GetAllBalanceAccountsOfAccountHolderInput
@return PaginatedBalanceAccountsResponse, *http.Response, error
*/
func (a *AccountHoldersApi) GetAllBalanceAccountsOfAccountHolder(ctx context.Context, r AccountHoldersApiGetAllBalanceAccountsOfAccountHolderInput) (PaginatedBalanceAccountsResponse, *http.Response, error) {
	res := &PaginatedBalanceAccountsResponse{}
	path := "/accountHolders/{id}/balanceAccounts"
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

// All parameters accepted by AccountHoldersApi.GetTaxForm
type AccountHoldersApiGetTaxFormInput struct {
	id       string
	formType *string
	year     *int32
}

// The type of tax form you want to retrieve. Accepted values are **US1099k** and **US1099nec**
func (r AccountHoldersApiGetTaxFormInput) FormType(formType string) AccountHoldersApiGetTaxFormInput {
	r.formType = &formType
	return r
}

// The tax year in YYYY format for the tax form you want to retrieve
func (r AccountHoldersApiGetTaxFormInput) Year(year int32) AccountHoldersApiGetTaxFormInput {
	r.year = &year
	return r
}

/*
Prepare a request for GetTaxForm
@param id The unique identifier of the account holder.
@return AccountHoldersApiGetTaxFormInput
*/
func (a *AccountHoldersApi) GetTaxFormInput(id string) AccountHoldersApiGetTaxFormInput {
	return AccountHoldersApiGetTaxFormInput{
		id: id,
	}
}

/*
GetTaxForm Get a tax form

Generates a tax form for account holders operating in the US. For more information, refer to US tax forms for [marketplaces](https://docs.adyen.com/marketplaces/us-tax-forms/) or [platforms](https://docs.adyen.com/platforms/us-tax-forms/) .

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountHoldersApiGetTaxFormInput - Request parameters, see GetTaxFormInput
@return GetTaxFormResponse, *http.Response, error
*/
func (a *AccountHoldersApi) GetTaxForm(ctx context.Context, r AccountHoldersApiGetTaxFormInput) (GetTaxFormResponse, *http.Response, error) {
	res := &GetTaxFormResponse{}
	path := "/accountHolders/{id}/taxForms"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.formType != nil {
		common.ParameterAddToQuery(queryParams, "formType", r.formType, "")
	}
	if r.year != nil {
		common.ParameterAddToQuery(queryParams, "year", r.year, "")
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

// All parameters accepted by AccountHoldersApi.UpdateAccountHolder
type AccountHoldersApiUpdateAccountHolderInput struct {
	id                         string
	accountHolderUpdateRequest *AccountHolderUpdateRequest
}

func (r AccountHoldersApiUpdateAccountHolderInput) AccountHolderUpdateRequest(accountHolderUpdateRequest AccountHolderUpdateRequest) AccountHoldersApiUpdateAccountHolderInput {
	r.accountHolderUpdateRequest = &accountHolderUpdateRequest
	return r
}

/*
Prepare a request for UpdateAccountHolder
@param id The unique identifier of the account holder.
@return AccountHoldersApiUpdateAccountHolderInput
*/
func (a *AccountHoldersApi) UpdateAccountHolderInput(id string) AccountHoldersApiUpdateAccountHolderInput {
	return AccountHoldersApiUpdateAccountHolderInput{
		id: id,
	}
}

/*
UpdateAccountHolder Update an account holder

Updates an account holder. When updating an account holder resource, if a parameter is not provided in the request, it is left unchanged.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r AccountHoldersApiUpdateAccountHolderInput - Request parameters, see UpdateAccountHolderInput
@return AccountHolder, *http.Response, error
*/
func (a *AccountHoldersApi) UpdateAccountHolder(ctx context.Context, r AccountHoldersApiUpdateAccountHolderInput) (AccountHolder, *http.Response, error) {
	res := &AccountHolder{}
	path := "/accountHolders/{id}"
	path = strings.Replace(path, "{"+"id"+"}", url.PathEscape(common.ParameterValueToString(r.id, "id")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.accountHolderUpdateRequest,
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
