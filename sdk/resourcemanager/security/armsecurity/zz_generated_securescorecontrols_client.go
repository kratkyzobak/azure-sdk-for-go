//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// SecureScoreControlsClient contains the methods for the SecureScoreControls group.
// Don't use this type directly, use NewSecureScoreControlsClient() instead.
type SecureScoreControlsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecureScoreControlsClient creates a new instance of SecureScoreControlsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecureScoreControlsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecureScoreControlsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &SecureScoreControlsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// List - Get all security controls within a scope
// If the operation fails it returns an *azcore.ResponseError type.
// options - SecureScoreControlsClientListOptions contains the optional parameters for the SecureScoreControlsClient.List
// method.
func (client *SecureScoreControlsClient) List(options *SecureScoreControlsClientListOptions) *SecureScoreControlsClientListPager {
	return &SecureScoreControlsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SecureScoreControlsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecureScoreControlList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecureScoreControlsClient) listCreateRequest(ctx context.Context, options *SecureScoreControlsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScoreControls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecureScoreControlsClient) listHandleResponse(resp *http.Response) (SecureScoreControlsClientListResponse, error) {
	result := SecureScoreControlsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlList); err != nil {
		return SecureScoreControlsClientListResponse{}, err
	}
	return result, nil
}

// ListBySecureScore - Get all security controls for a specific initiative within a scope
// If the operation fails it returns an *azcore.ResponseError type.
// secureScoreName - The initiative name. For the ASC Default initiative, use 'ascScore' as in the sample request below.
// options - SecureScoreControlsClientListBySecureScoreOptions contains the optional parameters for the SecureScoreControlsClient.ListBySecureScore
// method.
func (client *SecureScoreControlsClient) ListBySecureScore(secureScoreName string, options *SecureScoreControlsClientListBySecureScoreOptions) *SecureScoreControlsClientListBySecureScorePager {
	return &SecureScoreControlsClientListBySecureScorePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySecureScoreCreateRequest(ctx, secureScoreName, options)
		},
		advancer: func(ctx context.Context, resp SecureScoreControlsClientListBySecureScoreResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecureScoreControlList.NextLink)
		},
	}
}

// listBySecureScoreCreateRequest creates the ListBySecureScore request.
func (client *SecureScoreControlsClient) listBySecureScoreCreateRequest(ctx context.Context, secureScoreName string, options *SecureScoreControlsClientListBySecureScoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/secureScores/{secureScoreName}/secureScoreControls"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if secureScoreName == "" {
		return nil, errors.New("parameter secureScoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{secureScoreName}", url.PathEscape(secureScoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySecureScoreHandleResponse handles the ListBySecureScore response.
func (client *SecureScoreControlsClient) listBySecureScoreHandleResponse(resp *http.Response) (SecureScoreControlsClientListBySecureScoreResponse, error) {
	result := SecureScoreControlsClientListBySecureScoreResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecureScoreControlList); err != nil {
		return SecureScoreControlsClientListBySecureScoreResponse{}, err
	}
	return result, nil
}
