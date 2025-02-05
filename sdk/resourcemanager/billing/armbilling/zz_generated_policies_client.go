//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PoliciesClient contains the methods for the Policies group.
// Don't use this type directly, use NewPoliciesClient() instead.
type PoliciesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewPoliciesClient creates a new instance of PoliciesClient with the specified values.
func NewPoliciesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *PoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PoliciesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetByBillingProfile - Lists the policies for a billing profile. This operation is supported only for billing accounts with agreement type Microsoft Customer
// Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PoliciesClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, options *PoliciesGetByBillingProfileOptions) (PoliciesGetByBillingProfileResponse, error) {
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return PoliciesGetByBillingProfileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoliciesGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoliciesGetByBillingProfileResponse{}, client.getByBillingProfileHandleError(resp)
	}
	return client.getByBillingProfileHandleResponse(resp)
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *PoliciesClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *PoliciesGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *PoliciesClient) getByBillingProfileHandleResponse(resp *http.Response) (PoliciesGetByBillingProfileResponse, error) {
	result := PoliciesGetByBillingProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Policy); err != nil {
		return PoliciesGetByBillingProfileResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByBillingProfileHandleError handles the GetByBillingProfile error response.
func (client *PoliciesClient) getByBillingProfileHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetByCustomer - Lists the policies for a customer. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PoliciesClient) GetByCustomer(ctx context.Context, billingAccountName string, customerName string, options *PoliciesGetByCustomerOptions) (PoliciesGetByCustomerResponse, error) {
	req, err := client.getByCustomerCreateRequest(ctx, billingAccountName, customerName, options)
	if err != nil {
		return PoliciesGetByCustomerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoliciesGetByCustomerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoliciesGetByCustomerResponse{}, client.getByCustomerHandleError(resp)
	}
	return client.getByCustomerHandleResponse(resp)
}

// getByCustomerCreateRequest creates the GetByCustomer request.
func (client *PoliciesClient) getByCustomerCreateRequest(ctx context.Context, billingAccountName string, customerName string, options *PoliciesGetByCustomerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByCustomerHandleResponse handles the GetByCustomer response.
func (client *PoliciesClient) getByCustomerHandleResponse(resp *http.Response) (PoliciesGetByCustomerResponse, error) {
	result := PoliciesGetByCustomerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerPolicy); err != nil {
		return PoliciesGetByCustomerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByCustomerHandleError handles the GetByCustomer error response.
func (client *PoliciesClient) getByCustomerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates the policies for a billing profile. This operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PoliciesClient) Update(ctx context.Context, billingAccountName string, billingProfileName string, parameters Policy, options *PoliciesUpdateOptions) (PoliciesUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, billingAccountName, billingProfileName, parameters, options)
	if err != nil {
		return PoliciesUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoliciesUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoliciesUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PoliciesClient) updateCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, parameters Policy, options *PoliciesUpdateOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/policies/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *PoliciesClient) updateHandleResponse(resp *http.Response) (PoliciesUpdateResponse, error) {
	result := PoliciesUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Policy); err != nil {
		return PoliciesUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *PoliciesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateCustomer - Updates the policies for a customer. This operation is supported only for billing accounts with agreement type Microsoft Partner Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *PoliciesClient) UpdateCustomer(ctx context.Context, billingAccountName string, customerName string, parameters CustomerPolicy, options *PoliciesUpdateCustomerOptions) (PoliciesUpdateCustomerResponse, error) {
	req, err := client.updateCustomerCreateRequest(ctx, billingAccountName, customerName, parameters, options)
	if err != nil {
		return PoliciesUpdateCustomerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PoliciesUpdateCustomerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PoliciesUpdateCustomerResponse{}, client.updateCustomerHandleError(resp)
	}
	return client.updateCustomerHandleResponse(resp)
}

// updateCustomerCreateRequest creates the UpdateCustomer request.
func (client *PoliciesClient) updateCustomerCreateRequest(ctx context.Context, billingAccountName string, customerName string, parameters CustomerPolicy, options *PoliciesUpdateCustomerOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/policies/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if customerName == "" {
		return nil, errors.New("parameter customerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customerName}", url.PathEscape(customerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateCustomerHandleResponse handles the UpdateCustomer response.
func (client *PoliciesClient) updateCustomerHandleResponse(resp *http.Response) (PoliciesUpdateCustomerResponse, error) {
	result := PoliciesUpdateCustomerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomerPolicy); err != nil {
		return PoliciesUpdateCustomerResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateCustomerHandleError handles the UpdateCustomer error response.
func (client *PoliciesClient) updateCustomerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
