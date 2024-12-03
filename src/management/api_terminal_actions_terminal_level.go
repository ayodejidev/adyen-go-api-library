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

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// TerminalActionsTerminalLevelApi service
type TerminalActionsTerminalLevelApi common.Service

// All parameters accepted by TerminalActionsTerminalLevelApi.CreateTerminalAction
type TerminalActionsTerminalLevelApiCreateTerminalActionInput struct {
	scheduleTerminalActionsRequest *ScheduleTerminalActionsRequest
}

func (r TerminalActionsTerminalLevelApiCreateTerminalActionInput) ScheduleTerminalActionsRequest(scheduleTerminalActionsRequest ScheduleTerminalActionsRequest) TerminalActionsTerminalLevelApiCreateTerminalActionInput {
	r.scheduleTerminalActionsRequest = &scheduleTerminalActionsRequest
	return r
}

/*
Prepare a request for CreateTerminalAction

@return TerminalActionsTerminalLevelApiCreateTerminalActionInput
*/
func (a *TerminalActionsTerminalLevelApi) CreateTerminalActionInput() TerminalActionsTerminalLevelApiCreateTerminalActionInput {
	return TerminalActionsTerminalLevelApiCreateTerminalActionInput{}
}

/*
CreateTerminalAction Create a terminal action

Schedules a [terminal action](https://docs.adyen.com/point-of-sale/automating-terminal-management/terminal-actions-api) by specifying the action and the terminals that the action must be applied to.

The following restrictions apply:
* You can schedule only one action at a time. For example, to install a new app version and remove an old app version, you have to make two API requests.
* The maximum number of terminals in a request is **100**. For example, to apply an action to 250 terminals, you have to divide the terminals over three API requests.
* If there is an error with one or more terminal IDs in the request, the action is scheduled for none of the terminals. You need to fix the error and try again.

To make this request, your API credential must have the following [role](https://docs.adyen.com/development-resources/api-credentials#api-permissions):
* Management API—Terminal actions read and write

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r TerminalActionsTerminalLevelApiCreateTerminalActionInput - Request parameters, see CreateTerminalActionInput
@return ScheduleTerminalActionsResponse, *http.Response, error
*/
func (a *TerminalActionsTerminalLevelApi) CreateTerminalAction(ctx context.Context, r TerminalActionsTerminalLevelApiCreateTerminalActionInput) (ScheduleTerminalActionsResponse, *http.Response, error) {
	res := &ScheduleTerminalActionsResponse{}
	path := "/terminals/scheduleActions"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.scheduleTerminalActionsRequest,
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
