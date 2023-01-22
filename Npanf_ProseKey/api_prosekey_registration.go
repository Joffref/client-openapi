/*
Npanf_ProseKey

PAnF ProseKey Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npanf_ProseKey

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// ProsekeyRegistrationApiService ProsekeyRegistrationApi service
type ProsekeyRegistrationApiService service

type ApiProseKeyRegistrationRequest struct {
	ctx context.Context
	ApiService *ProsekeyRegistrationApiService
	proseContextInfo *ProseContextInfo
}

func (r ApiProseKeyRegistrationRequest) ProseContextInfo(proseContextInfo ProseContextInfo) ApiProseKeyRegistrationRequest {
	r.proseContextInfo = &proseContextInfo
	return r
}

func (r ApiProseKeyRegistrationRequest) Execute() (*http.Response, error) {
	return r.ApiService.ProseKeyRegistrationExecute(r)
}

/*
ProseKeyRegistration Register the Prose Key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProseKeyRegistrationRequest
*/
func (a *ProsekeyRegistrationApiService) ProseKeyRegistration(ctx context.Context) ApiProseKeyRegistrationRequest {
	return ApiProseKeyRegistrationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ProsekeyRegistrationApiService) ProseKeyRegistrationExecute(r ApiProseKeyRegistrationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProsekeyRegistrationApiService.ProseKeyRegistration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/prose-keys/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.proseContextInfo == nil {
		return nil, reportError("proseContextInfo is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.proseContextInfo
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}