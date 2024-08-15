/*
Transfers API

API version: 4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transfers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v12/src/common"
)

// CapitalApi service
type CapitalApi common.Service

// All parameters accepted by CapitalApi.GetCapitalAccount
type CapitalApiGetCapitalAccountInput struct {
	counterpartyAccountHolderId *string
}

// The counterparty account holder id.
func (r CapitalApiGetCapitalAccountInput) CounterpartyAccountHolderId(counterpartyAccountHolderId string) CapitalApiGetCapitalAccountInput {
	r.counterpartyAccountHolderId = &counterpartyAccountHolderId
	return r
}

/*
Prepare a request for GetCapitalAccount

@return CapitalApiGetCapitalAccountInput
*/
func (a *CapitalApi) GetCapitalAccountInput() CapitalApiGetCapitalAccountInput {
	return CapitalApiGetCapitalAccountInput{}
}

/*
GetCapitalAccount Get a capital account

Returns a list of grants with status and outstanding balances.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r CapitalApiGetCapitalAccountInput - Request parameters, see GetCapitalAccountInput
@return CapitalGrants, *http.Response, error
*/
func (a *CapitalApi) GetCapitalAccount(ctx context.Context, r CapitalApiGetCapitalAccountInput) (CapitalGrants, *http.Response, error) {
	res := &CapitalGrants{}
	path := "/grants"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.counterpartyAccountHolderId != nil {
		common.ParameterAddToQuery(queryParams, "counterpartyAccountHolderId", r.counterpartyAccountHolderId, "")
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
	if httpRes.StatusCode == 404 {
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

// All parameters accepted by CapitalApi.GetGrantReferenceDetails
type CapitalApiGetGrantReferenceDetailsInput struct {
	id string
}

/*
Prepare a request for GetGrantReferenceDetails
@param id The unique identifier of the grant.
@return CapitalApiGetGrantReferenceDetailsInput
*/
func (a *CapitalApi) GetGrantReferenceDetailsInput(id string) CapitalApiGetGrantReferenceDetailsInput {
	return CapitalApiGetGrantReferenceDetailsInput{
		id: id,
	}
}

/*
GetGrantReferenceDetails Get grant reference details

Returns the details of a capital account specified in the path.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r CapitalApiGetGrantReferenceDetailsInput - Request parameters, see GetGrantReferenceDetailsInput
@return CapitalGrant, *http.Response, error
*/
func (a *CapitalApi) GetGrantReferenceDetails(ctx context.Context, r CapitalApiGetGrantReferenceDetailsInput) (CapitalGrant, *http.Response, error) {
	res := &CapitalGrant{}
	path := "/grants/{id}"
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
	if httpRes.StatusCode == 404 {
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

// All parameters accepted by CapitalApi.RequestGrantPayout
type CapitalApiRequestGrantPayoutInput struct {
	capitalGrantInfo *CapitalGrantInfo
}

func (r CapitalApiRequestGrantPayoutInput) CapitalGrantInfo(capitalGrantInfo CapitalGrantInfo) CapitalApiRequestGrantPayoutInput {
	r.capitalGrantInfo = &capitalGrantInfo
	return r
}

/*
Prepare a request for RequestGrantPayout

@return CapitalApiRequestGrantPayoutInput
*/
func (a *CapitalApi) RequestGrantPayoutInput() CapitalApiRequestGrantPayoutInput {
	return CapitalApiRequestGrantPayoutInput{}
}

/*
RequestGrantPayout Request a grant payout

Requests the payout of the selected grant offer.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r CapitalApiRequestGrantPayoutInput - Request parameters, see RequestGrantPayoutInput
@return CapitalGrant, *http.Response, error
*/
func (a *CapitalApi) RequestGrantPayout(ctx context.Context, r CapitalApiRequestGrantPayoutInput) (CapitalGrant, *http.Response, error) {
	res := &CapitalGrant{}
	path := "/grants"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.capitalGrantInfo,
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
	if httpRes.StatusCode == 404 {
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
