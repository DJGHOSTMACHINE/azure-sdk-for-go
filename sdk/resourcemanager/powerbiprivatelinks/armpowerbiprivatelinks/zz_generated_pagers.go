//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// PrivateEndpointConnectionsListByResourcePager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsListByResourcePager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsListByResourceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsListByResourceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsListByResourcePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsListByResourcePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsListByResourceResponse page.
func (p *PrivateEndpointConnectionsListByResourcePager) PageResponse() PrivateEndpointConnectionsListByResourceResponse {
	return p.current
}

// PrivateLinkResourcesListByResourcePager provides operations for iterating over paged responses.
type PrivateLinkResourcesListByResourcePager struct {
	client    *PrivateLinkResourcesClient
	current   PrivateLinkResourcesListByResourceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkResourcesListByResourceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateLinkResourcesListByResourcePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateLinkResourcesListByResourcePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkResourcesListResult.NextLink == nil || len(*p.current.PrivateLinkResourcesListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateLinkResourcesListByResourceResponse page.
func (p *PrivateLinkResourcesListByResourcePager) PageResponse() PrivateLinkResourcesListByResourceResponse {
	return p.current
}
