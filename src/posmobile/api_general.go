/*
POS Mobile API

API version: 68
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package posmobile

import (
	"context"
	"net/http"
	"net/url"

	"github.com/adyen/adyen-go-api-library/v11/src/common"
)

// GeneralApi service
type GeneralApi common.Service

// All parameters accepted by GeneralApi.CreateCommunicationSession
type GeneralApiCreateCommunicationSessionInput struct {
	createSessionRequest *CreateSessionRequest
}

func (r GeneralApiCreateCommunicationSessionInput) CreateSessionRequest(createSessionRequest CreateSessionRequest) GeneralApiCreateCommunicationSessionInput {
	r.createSessionRequest = &createSessionRequest
	return r
}

/*
Prepare a request for CreateCommunicationSession

@return GeneralApiCreateCommunicationSessionInput
*/
func (a *GeneralApi) CreateCommunicationSessionInput() GeneralApiCreateCommunicationSessionInput {
	return GeneralApiCreateCommunicationSessionInput{}
}

/*
CreateCommunicationSession Create a communication session

Establishes a secure communications session between the POS Mobile SDK and the Adyen payments platform, through mutual authentication.
The request sends a setup token that identifies the SDK and the device. The response returns a session token that the SDK can use to authenticate responses received from the Adyen payments platform.
>This request applies to **mobile in-person** transactions. You cannot use this request to create online payments sessions.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r GeneralApiCreateCommunicationSessionInput - Request parameters, see CreateCommunicationSessionInput
@return CreateSessionResponse, *http.Response, error
*/
func (a *GeneralApi) CreateCommunicationSession(ctx context.Context, r GeneralApiCreateCommunicationSessionInput) (CreateSessionResponse, *http.Response, error) {
	res := &CreateSessionResponse{}
	path := "/sessions"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createSessionRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
