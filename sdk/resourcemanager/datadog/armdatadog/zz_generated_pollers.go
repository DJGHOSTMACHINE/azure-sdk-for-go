//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// MonitorsCreatePoller provides polling facilities until the operation reaches a terminal state.
type MonitorsCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MonitorsCreatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MonitorsCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MonitorsCreateResponse will be returned.
func (p *MonitorsCreatePoller) FinalResponse(ctx context.Context) (MonitorsCreateResponse, error) {
	respType := MonitorsCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DatadogMonitorResource)
	if err != nil {
		return MonitorsCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MonitorsCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MonitorsDeletePoller provides polling facilities until the operation reaches a terminal state.
type MonitorsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MonitorsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MonitorsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MonitorsDeleteResponse will be returned.
func (p *MonitorsDeletePoller) FinalResponse(ctx context.Context) (MonitorsDeleteResponse, error) {
	respType := MonitorsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return MonitorsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MonitorsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MonitorsUpdatePoller provides polling facilities until the operation reaches a terminal state.
type MonitorsUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MonitorsUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *MonitorsUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MonitorsUpdateResponse will be returned.
func (p *MonitorsUpdatePoller) FinalResponse(ctx context.Context) (MonitorsUpdateResponse, error) {
	respType := MonitorsUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DatadogMonitorResource)
	if err != nil {
		return MonitorsUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MonitorsUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SingleSignOnConfigurationsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type SingleSignOnConfigurationsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SingleSignOnConfigurationsCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *SingleSignOnConfigurationsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SingleSignOnConfigurationsCreateOrUpdateResponse will be returned.
func (p *SingleSignOnConfigurationsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (SingleSignOnConfigurationsCreateOrUpdateResponse, error) {
	respType := SingleSignOnConfigurationsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DatadogSingleSignOnResource)
	if err != nil {
		return SingleSignOnConfigurationsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SingleSignOnConfigurationsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
