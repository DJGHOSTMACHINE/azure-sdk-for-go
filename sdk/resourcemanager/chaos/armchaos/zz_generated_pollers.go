//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ExperimentsCancelPoller provides polling facilities until the operation reaches a terminal state.
type ExperimentsCancelPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ExperimentsCancelPoller) Done() bool {
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
func (p *ExperimentsCancelPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ExperimentsCancelResponse will be returned.
func (p *ExperimentsCancelPoller) FinalResponse(ctx context.Context) (ExperimentsCancelResponse, error) {
	respType := ExperimentsCancelResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.ExperimentCancelOperationResult)
	if err != nil {
		return ExperimentsCancelResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ExperimentsCancelPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ExperimentsCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type ExperimentsCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ExperimentsCreateOrUpdatePoller) Done() bool {
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
func (p *ExperimentsCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ExperimentsCreateOrUpdateResponse will be returned.
func (p *ExperimentsCreateOrUpdatePoller) FinalResponse(ctx context.Context) (ExperimentsCreateOrUpdateResponse, error) {
	respType := ExperimentsCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Experiment)
	if err != nil {
		return ExperimentsCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ExperimentsCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
