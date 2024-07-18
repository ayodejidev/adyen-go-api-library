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

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// GrantAccountsApi service
type GrantAccountsApi common.Service

// All parameters accepted by GrantAccountsApi.GetGrantAccount
type GrantAccountsApiGetGrantAccountInput struct {
	id string
}

/*
Prepare a request for GetGrantAccount
@param id The unique identifier of the grant account.
@return GrantAccountsApiGetGrantAccountInput
*/
func (a *GrantAccountsApi) GetGrantAccountInput(id string) GrantAccountsApiGetGrantAccountInput {
	return GrantAccountsApiGetGrantAccountInput{
		id: id,
	}
}

/*
GetGrantAccount Get a grant account

Returns the details of the [grant account](https://docs.adyen.com/platforms/capital#grant-account).

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GrantAccountsApiGetGrantAccountInput - Request parameters, see GetGrantAccountInput
@return CapitalGrantAccount, *http.Response, error
*/
func (a *GrantAccountsApi) GetGrantAccount(ctx context.Context, r GrantAccountsApiGetGrantAccountInput) (CapitalGrantAccount, *http.Response, error) {
	res := &CapitalGrantAccount{}
	path := "/grantAccounts/{id}"
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
