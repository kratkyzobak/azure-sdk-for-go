//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// UsageClient contains the methods for the Usage group.
// Don't use this type directly, use NewUsageClient() instead.
type UsageClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewUsageClient creates a new instance of UsageClient with the specified values.
func NewUsageClient(con *arm.Connection, subscriptionID string) *UsageClient {
	return &UsageClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
// If the operation fails it returns a generic error.
func (client *UsageClient) List(location string, options *UsageListOptions) *UsageListPager {
	return &UsageListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp UsageListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListUsagesResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UsageClient) listCreateRequest(ctx context.Context, location string, options *UsageListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageClient) listHandleResponse(resp *http.Response) (UsageListResponse, error) {
	result := UsageListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListUsagesResult); err != nil {
		return UsageListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *UsageClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}