// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Taxes Service Adapter SPI
 *
 * ## Overview Tax Service Adapter is a Service Provider Interface (SPI) that helps Chargebee communicate with API platforms providing the following services: - Address validation - Estimating taxes for sales - Submitting invoices for tax filing  If you have an API that provides the services mentioned above and want Chargebee to be able to integrate with it, you can build an adapter service according to this specification.  ## Terminology Here's a list of terms we've used to describe this specification. ### Merchant    Chargebee’s customer who is carrying out the business of selling products. Also known as the Seller. The Seller has one or more Chargebee accounts. ### Chargebee   The SaaS which manages subscriptions and revenue operations on behalf of the Merchant. Chargebee makes API calls to the Tax Service Adapter for tax estimation and tax filing of transactions between the Merchant and the Customer. ### Customer  Merchant’s customer who purchases products from the Merchant. Their customer record(s) and other details are stored within Chargebee under the Merchant’s account. ### Tax Authority The institution, usually a government, that is responsible for collecting taxes for commercial transactions between Merchants and their Customers.   ### Tax Service Provider The service which provides tax estimation and tax filing services for merchants. ### Tax Service Adapter  The SPI for which the spec is defined in this document. This adapter serves the following purposes: - Translates API requests (as defined in this document) from Chargebee to the request format specified by the Tax Service Provider's API. - Translates API responses from the Tax Service Provider APIs to the format understood by Chargebee (as defined in this document). ### Authorization Chargebee uses HTTP header-based authorization for all the API endpoints associated with Tax Service Adapter. We dynamically pass this authorization key in the HTTP header. The parameter that holds this key is found in the JSON object `api_configuration` required for configuring your onboarding on Chargebee's marketplace. In the `api_configuration` object, our [Taxes Service Adapter SPI](https://chargebee.atlassian.net/l/cp/ca5aZ1mA) checks the authorization key parameter from `credential_configuration.id` and creates the HTTP header-based input query parameter for authorization. The `credential_configuration` is an array of objects with an `id` attribute, and the value of `id` is the parameter containing the authorization key.    Following are the JSON snippets for your reference.     ```json       \"api_configuration\": {           \"api_base_url\": \"https://xyz.abc.com/chargebee\",           \"credential_configuration\": [              {               \"id\":\"authorization_key\",               \"name\": \"Authorization Key\",                \"type\": \"text\",               \"is_sensitive\": true              },             {               \"id\": \"client_secret\",               \"name\": \"Client Secret\",               \"type\": \"text\",               \"is_sensitive\": true             }           ]         }    ```     `\"Authorization\":\"{\"<authorization_key>\":\"test_GykCvTykZFTBqHQAFzjMN9R1Thgdho0Q\"}\"` 
 *
 * API version: 0.3.7
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CreditNoteAPIController binds http requests to an api service and writes the service results to the http response
type CreditNoteAPIController struct {
	service CreditNoteAPIServicer
	errorHandler ErrorHandler
}

// CreditNoteAPIOption for how the controller is set up.
type CreditNoteAPIOption func(*CreditNoteAPIController)

// WithCreditNoteAPIErrorHandler inject ErrorHandler into controller
func WithCreditNoteAPIErrorHandler(h ErrorHandler) CreditNoteAPIOption {
	return func(c *CreditNoteAPIController) {
		c.errorHandler = h
	}
}

// NewCreditNoteAPIController creates a default api controller
func NewCreditNoteAPIController(s CreditNoteAPIServicer, opts ...CreditNoteAPIOption) *CreditNoteAPIController {
	controller := &CreditNoteAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the CreditNoteAPIController
func (c *CreditNoteAPIController) Routes() Routes {
	return Routes{
		"CommitCreditNote": Route{
			strings.ToUpper("Post"),
			"/api/v1/credit-notes/{creditNoteId}/commit",
			c.CommitCreditNote,
		},
		"CreateCreditNote": Route{
			strings.ToUpper("Post"),
			"/api/v1/credit-notes",
			c.CreateCreditNote,
		},
		"FetchCreditNote": Route{
			strings.ToUpper("Get"),
			"/api/v1/credit-notes/{creditNoteId}",
			c.FetchCreditNote,
		},
		"VoidCreditNote": Route{
			strings.ToUpper("Post"),
			"/api/v1/credit-notes/{creditNoteId}/void",
			c.VoidCreditNote,
		},
	}
}

// CommitCreditNote - Commit credit note
func (c *CreditNoteAPIController) CommitCreditNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	creditNoteIdParam := params["creditNoteId"]
	if creditNoteIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"creditNoteId"}, nil)
		return
	}
	var invoiceIdParam string
	if query.Has("invoiceId") {
		param := query.Get("invoiceId")

		invoiceIdParam = param
	} else {
	}
	result, err := c.service.CommitCreditNote(r.Context(), creditNoteIdParam, invoiceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateCreditNote - Create credit note
func (c *CreditNoteAPIController) CreateCreditNote(w http.ResponseWriter, r *http.Request) {
	creditNoteRequestParam := CreditNoteRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&creditNoteRequestParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertCreditNoteRequestRequired(creditNoteRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertCreditNoteRequestConstraints(creditNoteRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateCreditNote(r.Context(), creditNoteRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// FetchCreditNote - Retrieve credit note
func (c *CreditNoteAPIController) FetchCreditNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	creditNoteIdParam := params["creditNoteId"]
	if creditNoteIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"creditNoteId"}, nil)
		return
	}
	var invoiceIdParam string
	if query.Has("invoiceId") {
		param := query.Get("invoiceId")

		invoiceIdParam = param
	} else {
	}
	result, err := c.service.FetchCreditNote(r.Context(), creditNoteIdParam, invoiceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// VoidCreditNote - Void credit note
func (c *CreditNoteAPIController) VoidCreditNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	creditNoteIdParam := params["creditNoteId"]
	if creditNoteIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"creditNoteId"}, nil)
		return
	}
	var invoiceIdParam string
	if query.Has("invoiceId") {
		param := query.Get("invoiceId")

		invoiceIdParam = param
	} else {
	}
	result, err := c.service.VoidCreditNote(r.Context(), creditNoteIdParam, invoiceIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
