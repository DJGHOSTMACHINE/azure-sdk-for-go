//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MigrationRecoveryPointsClient contains the methods for the MigrationRecoveryPoints group.
// Don't use this type directly, use NewMigrationRecoveryPointsClient() instead.
type MigrationRecoveryPointsClient struct {
	ep                string
	pl                runtime.Pipeline
	resourceName      string
	resourceGroupName string
	subscriptionID    string
}

// NewMigrationRecoveryPointsClient creates a new instance of MigrationRecoveryPointsClient with the specified values.
func NewMigrationRecoveryPointsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MigrationRecoveryPointsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &MigrationRecoveryPointsClient{resourceName: resourceName, resourceGroupName: resourceGroupName, subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a recovery point for a migration item.
// If the operation fails it returns a generic error.
func (client *MigrationRecoveryPointsClient) Get(ctx context.Context, fabricName string, protectionContainerName string, migrationItemName string, migrationRecoveryPointName string, options *MigrationRecoveryPointsGetOptions) (MigrationRecoveryPointsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, protectionContainerName, migrationItemName, migrationRecoveryPointName, options)
	if err != nil {
		return MigrationRecoveryPointsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationRecoveryPointsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationRecoveryPointsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MigrationRecoveryPointsClient) getCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, migrationItemName string, migrationRecoveryPointName string, options *MigrationRecoveryPointsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints/{migrationRecoveryPointName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if migrationItemName == "" {
		return nil, errors.New("parameter migrationItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationItemName}", url.PathEscape(migrationItemName))
	if migrationRecoveryPointName == "" {
		return nil, errors.New("parameter migrationRecoveryPointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationRecoveryPointName}", url.PathEscape(migrationRecoveryPointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MigrationRecoveryPointsClient) getHandleResponse(resp *http.Response) (MigrationRecoveryPointsGetResponse, error) {
	result := MigrationRecoveryPointsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationRecoveryPoint); err != nil {
		return MigrationRecoveryPointsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MigrationRecoveryPointsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByReplicationMigrationItems - Gets the recovery points for a migration item.
// If the operation fails it returns a generic error.
func (client *MigrationRecoveryPointsClient) ListByReplicationMigrationItems(fabricName string, protectionContainerName string, migrationItemName string, options *MigrationRecoveryPointsListByReplicationMigrationItemsOptions) *MigrationRecoveryPointsListByReplicationMigrationItemsPager {
	return &MigrationRecoveryPointsListByReplicationMigrationItemsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByReplicationMigrationItemsCreateRequest(ctx, fabricName, protectionContainerName, migrationItemName, options)
		},
		advancer: func(ctx context.Context, resp MigrationRecoveryPointsListByReplicationMigrationItemsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MigrationRecoveryPointCollection.NextLink)
		},
	}
}

// listByReplicationMigrationItemsCreateRequest creates the ListByReplicationMigrationItems request.
func (client *MigrationRecoveryPointsClient) listByReplicationMigrationItemsCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, migrationItemName string, options *MigrationRecoveryPointsListByReplicationMigrationItemsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if migrationItemName == "" {
		return nil, errors.New("parameter migrationItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationItemName}", url.PathEscape(migrationItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByReplicationMigrationItemsHandleResponse handles the ListByReplicationMigrationItems response.
func (client *MigrationRecoveryPointsClient) listByReplicationMigrationItemsHandleResponse(resp *http.Response) (MigrationRecoveryPointsListByReplicationMigrationItemsResponse, error) {
	result := MigrationRecoveryPointsListByReplicationMigrationItemsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationRecoveryPointCollection); err != nil {
		return MigrationRecoveryPointsListByReplicationMigrationItemsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByReplicationMigrationItemsHandleError handles the ListByReplicationMigrationItems error response.
func (client *MigrationRecoveryPointsClient) listByReplicationMigrationItemsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
