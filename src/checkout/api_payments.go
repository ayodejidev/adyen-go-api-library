/*
Adyen Checkout API

API version: 71
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package checkout

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/adyen/adyen-go-api-library/v14/src/common"
)

// PaymentsApi service
type PaymentsApi common.Service

// All parameters accepted by PaymentsApi.CardDetails
type PaymentsApiCardDetailsInput struct {
	idempotencyKey     *string
	cardDetailsRequest *CardDetailsRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r PaymentsApiCardDetailsInput) IdempotencyKey(idempotencyKey string) PaymentsApiCardDetailsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r PaymentsApiCardDetailsInput) CardDetailsRequest(cardDetailsRequest CardDetailsRequest) PaymentsApiCardDetailsInput {
	r.cardDetailsRequest = &cardDetailsRequest
	return r
}

/*
Prepare a request for CardDetails

@return PaymentsApiCardDetailsInput
*/
func (a *PaymentsApi) CardDetailsInput() PaymentsApiCardDetailsInput {
	return PaymentsApiCardDetailsInput{}
}

/*
CardDetails Get the list of brands on the card

Send a request with at least the first 6 digits of the card number to get a response with an array of brands on the card. If you include [your supported brands](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/cardDetails__reqParam_supportedBrands) in the request, the response also tells you if you support each [brand that was identified](https://docs.adyen.com/api-explorer/#/CheckoutService/latest/post/cardDetails__resParam_details).

If you have an API-only integration and collect card data, use this endpoint to find out if the shopper's card is co-branded. For co-branded cards, you must let the shopper choose the brand to pay with  if you support both brands.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsApiCardDetailsInput - Request parameters, see CardDetailsInput
@return CardDetailsResponse, *http.Response, error
*/
func (a *PaymentsApi) CardDetails(ctx context.Context, r PaymentsApiCardDetailsInput) (CardDetailsResponse, *http.Response, error) {
	res := &CardDetailsResponse{}
	path := "/cardDetails"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.cardDetailsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentsApi.GetResultOfPaymentSession
type PaymentsApiGetResultOfPaymentSessionInput struct {
	sessionId     string
	sessionResult *string
}

// The &#x60;sessionResult&#x60; value from the Drop-in or Component.
func (r PaymentsApiGetResultOfPaymentSessionInput) SessionResult(sessionResult string) PaymentsApiGetResultOfPaymentSessionInput {
	r.sessionResult = &sessionResult
	return r
}

/*
Prepare a request for GetResultOfPaymentSession
@param sessionId A unique identifier of the session.
@return PaymentsApiGetResultOfPaymentSessionInput
*/
func (a *PaymentsApi) GetResultOfPaymentSessionInput(sessionId string) PaymentsApiGetResultOfPaymentSessionInput {
	return PaymentsApiGetResultOfPaymentSessionInput{
		sessionId: sessionId,
	}
}

/*
GetResultOfPaymentSession Get the result of a payment session

Returns the status of the payment session with the `sessionId` and `sessionResult` specified in the path.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsApiGetResultOfPaymentSessionInput - Request parameters, see GetResultOfPaymentSessionInput
@return SessionResultResponse, *http.Response, error
*/
func (a *PaymentsApi) GetResultOfPaymentSession(ctx context.Context, r PaymentsApiGetResultOfPaymentSessionInput) (SessionResultResponse, *http.Response, error) {
	res := &SessionResultResponse{}
	path := "/sessions/{sessionId}"
	path = strings.Replace(path, "{"+"sessionId"+"}", url.PathEscape(common.ParameterValueToString(r.sessionId, "sessionId")), -1)
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.sessionResult != nil {
		common.ParameterAddToQuery(queryParams, "sessionResult", r.sessionResult, "")
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

	return *res, httpRes, err
}

// All parameters accepted by PaymentsApi.PaymentMethods
type PaymentsApiPaymentMethodsInput struct {
	idempotencyKey        *string
	paymentMethodsRequest *PaymentMethodsRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r PaymentsApiPaymentMethodsInput) IdempotencyKey(idempotencyKey string) PaymentsApiPaymentMethodsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r PaymentsApiPaymentMethodsInput) PaymentMethodsRequest(paymentMethodsRequest PaymentMethodsRequest) PaymentsApiPaymentMethodsInput {
	r.paymentMethodsRequest = &paymentMethodsRequest
	return r
}

/*
Prepare a request for PaymentMethods

@return PaymentsApiPaymentMethodsInput
*/
func (a *PaymentsApi) PaymentMethodsInput() PaymentsApiPaymentMethodsInput {
	return PaymentsApiPaymentMethodsInput{}
}

/*
PaymentMethods Get a list of available payment methods

Queries the available payment methods for a transaction based on the transaction context (like amount, country, and currency). Besides giving back a list of the available payment methods, the response also returns which input details you need to collect from the shopper (to be submitted to `/payments`).

Although we highly recommend using this endpoint to ensure you are always offering the most up-to-date list of payment methods, its usage is optional. You can, for example, also cache the `/paymentMethods` response and update it once a week.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsApiPaymentMethodsInput - Request parameters, see PaymentMethodsInput
@return PaymentMethodsResponse, *http.Response, error
*/
func (a *PaymentsApi) PaymentMethods(ctx context.Context, r PaymentsApiPaymentMethodsInput) (PaymentMethodsResponse, *http.Response, error) {
	res := &PaymentMethodsResponse{}
	path := "/paymentMethods"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentMethodsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentsApi.Payments
type PaymentsApiPaymentsInput struct {
	idempotencyKey *string
	paymentRequest *PaymentRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r PaymentsApiPaymentsInput) IdempotencyKey(idempotencyKey string) PaymentsApiPaymentsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r PaymentsApiPaymentsInput) PaymentRequest(paymentRequest PaymentRequest) PaymentsApiPaymentsInput {
	r.paymentRequest = &paymentRequest
	return r
}

/*
Prepare a request for Payments

@return PaymentsApiPaymentsInput
*/
func (a *PaymentsApi) PaymentsInput() PaymentsApiPaymentsInput {
	return PaymentsApiPaymentsInput{}
}

/*
Payments Start a transaction

Sends payment parameters (like amount, country, and currency) together with other required input details collected from the shopper. To know more about required parameters for specific payment methods, refer to our [payment method guides](https://docs.adyen.com/payment-methods).
The response depends on the [payment flow](https://docs.adyen.com/payment-methods#payment-flow):
* For a direct flow, the response includes a `pspReference` and a `resultCode` with the payment result, for example **Authorised** or **Refused**.
* For a redirect or additional action, the response contains an `action` object.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsApiPaymentsInput - Request parameters, see PaymentsInput
@return PaymentResponse, *http.Response, error
*/
func (a *PaymentsApi) Payments(ctx context.Context, r PaymentsApiPaymentsInput) (PaymentResponse, *http.Response, error) {
	res := &PaymentResponse{}
	path := "/payments"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentsApi.PaymentsDetails
type PaymentsApiPaymentsDetailsInput struct {
	idempotencyKey        *string
	paymentDetailsRequest *PaymentDetailsRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r PaymentsApiPaymentsDetailsInput) IdempotencyKey(idempotencyKey string) PaymentsApiPaymentsDetailsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r PaymentsApiPaymentsDetailsInput) PaymentDetailsRequest(paymentDetailsRequest PaymentDetailsRequest) PaymentsApiPaymentsDetailsInput {
	r.paymentDetailsRequest = &paymentDetailsRequest
	return r
}

/*
Prepare a request for PaymentsDetails

@return PaymentsApiPaymentsDetailsInput
*/
func (a *PaymentsApi) PaymentsDetailsInput() PaymentsApiPaymentsDetailsInput {
	return PaymentsApiPaymentsDetailsInput{}
}

/*
PaymentsDetails Submit details for a payment

Submits details for a payment created using `/payments`. This step is only needed when no final state has been reached on the `/payments` request, for example when the shopper was redirected to another page to complete the payment.



@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsApiPaymentsDetailsInput - Request parameters, see PaymentsDetailsInput
@return PaymentDetailsResponse, *http.Response, error
*/
func (a *PaymentsApi) PaymentsDetails(ctx context.Context, r PaymentsApiPaymentsDetailsInput) (PaymentDetailsResponse, *http.Response, error) {
	res := &PaymentDetailsResponse{}
	path := "/payments/details"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.paymentDetailsRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}

// All parameters accepted by PaymentsApi.Sessions
type PaymentsApiSessionsInput struct {
	idempotencyKey               *string
	createCheckoutSessionRequest *CreateCheckoutSessionRequest
}

// A unique identifier for the message with a maximum of 64 characters (we recommend a UUID).
func (r PaymentsApiSessionsInput) IdempotencyKey(idempotencyKey string) PaymentsApiSessionsInput {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r PaymentsApiSessionsInput) CreateCheckoutSessionRequest(createCheckoutSessionRequest CreateCheckoutSessionRequest) PaymentsApiSessionsInput {
	r.createCheckoutSessionRequest = &createCheckoutSessionRequest
	return r
}

/*
Prepare a request for Sessions

@return PaymentsApiSessionsInput
*/
func (a *PaymentsApi) SessionsInput() PaymentsApiSessionsInput {
	return PaymentsApiSessionsInput{}
}

/*
Sessions Create a payment session

Creates a payment session for [Web Drop-in](https://docs.adyen.com/online-payments/web-drop-in) and [Web Components](https://docs.adyen.com/online-payments/web-components) integrations.

The response contains encrypted payment session data. The front end then uses the session data to make any required server-side calls for the payment flow.

You get the payment outcome asynchronously, in an [AUTHORISATION](https://docs.adyen.com/api-explorer/#/Webhooks/latest/post/AUTHORISATION) webhook.

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param r PaymentsApiSessionsInput - Request parameters, see SessionsInput
@return CreateCheckoutSessionResponse, *http.Response, error
*/
func (a *PaymentsApi) Sessions(ctx context.Context, r PaymentsApiSessionsInput) (CreateCheckoutSessionResponse, *http.Response, error) {
	res := &CreateCheckoutSessionResponse{}
	path := "/sessions"
	queryParams := url.Values{}
	headerParams := make(map[string]string)
	if r.idempotencyKey != nil {
		common.ParameterAddToHeaderOrQuery(headerParams, "Idempotency-Key", r.idempotencyKey, "")
	}
	httpRes, err := common.SendAPIRequest(
		ctx,
		a.Client,
		r.createCheckoutSessionRequest,
		res,
		http.MethodPost,
		a.BasePath()+path,
		queryParams,
		headerParams,
	)

	return *res, httpRes, err
}
