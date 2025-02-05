//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolReplicationLinksClient contains the methods for the SQLPoolReplicationLinks group.
// Don't use this type directly, use NewSQLPoolReplicationLinksClient() instead.
type SQLPoolReplicationLinksClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSQLPoolReplicationLinksClient creates a new instance of SQLPoolReplicationLinksClient with the specified values.
func NewSQLPoolReplicationLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLPoolReplicationLinksClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SQLPoolReplicationLinksClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetByName - Get SQL pool replication link by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolReplicationLinksClient) GetByName(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, linkID string, options *SQLPoolReplicationLinksGetByNameOptions) (SQLPoolReplicationLinksGetByNameResponse, error) {
	req, err := client.getByNameCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, linkID, options)
	if err != nil {
		return SQLPoolReplicationLinksGetByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolReplicationLinksGetByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolReplicationLinksGetByNameResponse{}, client.getByNameHandleError(resp)
	}
	return client.getByNameHandleResponse(resp)
}

// getByNameCreateRequest creates the GetByName request.
func (client *SQLPoolReplicationLinksClient) getByNameCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, linkID string, options *SQLPoolReplicationLinksGetByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/replicationLinks/{linkId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if linkID == "" {
		return nil, errors.New("parameter linkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkId}", url.PathEscape(linkID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByNameHandleResponse handles the GetByName response.
func (client *SQLPoolReplicationLinksClient) getByNameHandleResponse(resp *http.Response) (SQLPoolReplicationLinksGetByNameResponse, error) {
	result := SQLPoolReplicationLinksGetByNameResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationLink); err != nil {
		return SQLPoolReplicationLinksGetByNameResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByNameHandleError handles the GetByName error response.
func (client *SQLPoolReplicationLinksClient) getByNameHandleError(resp *http.Response) error {
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

// List - Lists a Sql pool's replication links.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLPoolReplicationLinksClient) List(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolReplicationLinksListOptions) *SQLPoolReplicationLinksListPager {
	return &SQLPoolReplicationLinksListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolReplicationLinksListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ReplicationLinkListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SQLPoolReplicationLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolReplicationLinksListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/replicationLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLPoolReplicationLinksClient) listHandleResponse(resp *http.Response) (SQLPoolReplicationLinksListResponse, error) {
	result := SQLPoolReplicationLinksListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationLinkListResult); err != nil {
		return SQLPoolReplicationLinksListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SQLPoolReplicationLinksClient) listHandleError(resp *http.Response) error {
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
