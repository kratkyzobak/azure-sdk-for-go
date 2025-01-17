//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfidentialledger

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

// LedgerClient contains the methods for the Ledger group.
// Don't use this type directly, use NewLedgerClient() instead.
type LedgerClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLedgerClient creates a new instance of LedgerClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLedgerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LedgerClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &LedgerClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Creates a Confidential Ledger with the specified ledger parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ledgerName - Name of the Confidential Ledger
// confidentialLedger - Confidential Ledger Create Request Body
// options - LedgerClientBeginCreateOptions contains the optional parameters for the LedgerClient.BeginCreate method.
func (client *LedgerClient) BeginCreate(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginCreateOptions) (LedgerClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
	if err != nil {
		return LedgerClientCreatePollerResponse{}, err
	}
	result := LedgerClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LedgerClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return LedgerClientCreatePollerResponse{}, err
	}
	result.Poller = &LedgerClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a Confidential Ledger with the specified ledger parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LedgerClient) create(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LedgerClient) createCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, confidentialLedger)
}

// BeginDelete - Deletes an existing Confidential Ledger.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ledgerName - Name of the Confidential Ledger
// options - LedgerClientBeginDeleteOptions contains the optional parameters for the LedgerClient.BeginDelete method.
func (client *LedgerClient) BeginDelete(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientBeginDeleteOptions) (LedgerClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, ledgerName, options)
	if err != nil {
		return LedgerClientDeletePollerResponse{}, err
	}
	result := LedgerClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LedgerClient.Delete", "", resp, client.pl)
	if err != nil {
		return LedgerClientDeletePollerResponse{}, err
	}
	result.Poller = &LedgerClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes an existing Confidential Ledger.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LedgerClient) deleteOperation(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ledgerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LedgerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the properties of a Confidential Ledger.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ledgerName - Name of the Confidential Ledger
// options - LedgerClientGetOptions contains the optional parameters for the LedgerClient.Get method.
func (client *LedgerClient) Get(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientGetOptions) (LedgerClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ledgerName, options)
	if err != nil {
		return LedgerClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LedgerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LedgerClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LedgerClient) getCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, options *LedgerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LedgerClient) getHandleResponse(resp *http.Response) (LedgerClientGetResponse, error) {
	result := LedgerClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfidentialLedger); err != nil {
		return LedgerClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Retrieves the properties of all Confidential Ledgers.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - LedgerClientListByResourceGroupOptions contains the optional parameters for the LedgerClient.ListByResourceGroup
// method.
func (client *LedgerClient) ListByResourceGroup(resourceGroupName string, options *LedgerClientListByResourceGroupOptions) *LedgerClientListByResourceGroupPager {
	return &LedgerClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp LedgerClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.List.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *LedgerClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *LedgerClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *LedgerClient) listByResourceGroupHandleResponse(resp *http.Response) (LedgerClientListByResourceGroupResponse, error) {
	result := LedgerClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return LedgerClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Retrieves the properties of all Confidential Ledgers.
// If the operation fails it returns an *azcore.ResponseError type.
// options - LedgerClientListBySubscriptionOptions contains the optional parameters for the LedgerClient.ListBySubscription
// method.
func (client *LedgerClient) ListBySubscription(options *LedgerClientListBySubscriptionOptions) *LedgerClientListBySubscriptionPager {
	return &LedgerClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp LedgerClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.List.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *LedgerClient) listBySubscriptionCreateRequest(ctx context.Context, options *LedgerClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConfidentialLedger/ledgers/"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *LedgerClient) listBySubscriptionHandleResponse(resp *http.Response) (LedgerClientListBySubscriptionResponse, error) {
	result := LedgerClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.List); err != nil {
		return LedgerClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates properties of Confidential Ledger
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// ledgerName - Name of the Confidential Ledger
// confidentialLedger - Confidential Ledger request body for Updating Ledger
// options - LedgerClientBeginUpdateOptions contains the optional parameters for the LedgerClient.BeginUpdate method.
func (client *LedgerClient) BeginUpdate(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginUpdateOptions) (LedgerClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
	if err != nil {
		return LedgerClientUpdatePollerResponse{}, err
	}
	result := LedgerClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("LedgerClient.Update", "", resp, client.pl)
	if err != nil {
		return LedgerClientUpdatePollerResponse{}, err
	}
	result.Poller = &LedgerClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates properties of Confidential Ledger
// If the operation fails it returns an *azcore.ResponseError type.
func (client *LedgerClient) update(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, ledgerName, confidentialLedger, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *LedgerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, ledgerName string, confidentialLedger ConfidentialLedger, options *LedgerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConfidentialLedger/ledgers/{ledgerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ledgerName == "" {
		return nil, errors.New("parameter ledgerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ledgerName}", url.PathEscape(ledgerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, confidentialLedger)
}
