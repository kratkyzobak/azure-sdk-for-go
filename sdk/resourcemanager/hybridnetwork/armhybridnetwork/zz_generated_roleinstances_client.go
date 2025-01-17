//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// RoleInstancesClient contains the methods for the RoleInstances group.
// Don't use this type directly, use NewRoleInstancesClient() instead.
type RoleInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRoleInstancesClient creates a new instance of RoleInstancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRoleInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RoleInstancesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RoleInstancesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets the information of role instance of vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// roleInstanceName - The name of the role instance of the vendor network function.
// options - RoleInstancesClientGetOptions contains the optional parameters for the RoleInstancesClient.Get method.
func (client *RoleInstancesClient) Get(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientGetOptions) (RoleInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return RoleInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleInstancesClient) getCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleInstancesClient) getHandleResponse(resp *http.Response) (RoleInstancesClientGetResponse, error) {
	result := RoleInstancesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleInstance); err != nil {
		return RoleInstancesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the information of role instances of vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// options - RoleInstancesClientListOptions contains the optional parameters for the RoleInstancesClient.List method.
func (client *RoleInstancesClient) List(locationName string, vendorName string, serviceKey string, options *RoleInstancesClientListOptions) *RoleInstancesClientListPager {
	return &RoleInstancesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, locationName, vendorName, serviceKey, options)
		},
		advancer: func(ctx context.Context, resp RoleInstancesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NetworkFunctionRoleInstanceListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *RoleInstancesClient) listCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, options *RoleInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RoleInstancesClient) listHandleResponse(resp *http.Response) (RoleInstancesClientListResponse, error) {
	result := RoleInstancesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkFunctionRoleInstanceListResult); err != nil {
		return RoleInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginRestart - Restarts a role instance of a vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// roleInstanceName - The name of the role instance of the vendor network function.
// options - RoleInstancesClientBeginRestartOptions contains the optional parameters for the RoleInstancesClient.BeginRestart
// method.
func (client *RoleInstancesClient) BeginRestart(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginRestartOptions) (RoleInstancesClientRestartPollerResponse, error) {
	resp, err := client.restart(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return RoleInstancesClientRestartPollerResponse{}, err
	}
	result := RoleInstancesClientRestartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoleInstancesClient.Restart", "location", resp, client.pl)
	if err != nil {
		return RoleInstancesClientRestartPollerResponse{}, err
	}
	result.Poller = &RoleInstancesClientRestartPoller{
		pt: pt,
	}
	return result, nil
}

// Restart - Restarts a role instance of a vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoleInstancesClient) restart(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginRestartOptions) (*http.Response, error) {
	req, err := client.restartCreateRequest(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// restartCreateRequest creates the Restart request.
func (client *RoleInstancesClient) restartCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}/restart"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStart - Starts a role instance of a vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// roleInstanceName - The name of the role instance of the vendor network function.
// options - RoleInstancesClientBeginStartOptions contains the optional parameters for the RoleInstancesClient.BeginStart
// method.
func (client *RoleInstancesClient) BeginStart(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginStartOptions) (RoleInstancesClientStartPollerResponse, error) {
	resp, err := client.start(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return RoleInstancesClientStartPollerResponse{}, err
	}
	result := RoleInstancesClientStartPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoleInstancesClient.Start", "location", resp, client.pl)
	if err != nil {
		return RoleInstancesClientStartPollerResponse{}, err
	}
	result.Poller = &RoleInstancesClientStartPoller{
		pt: pt,
	}
	return result, nil
}

// Start - Starts a role instance of a vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoleInstancesClient) start(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginStartOptions) (*http.Response, error) {
	req, err := client.startCreateRequest(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startCreateRequest creates the Start request.
func (client *RoleInstancesClient) startCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}/start"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStop - Powers off (stop) a role instance of a vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
// locationName - The Azure region where the network function resource was created by customer.
// vendorName - The name of the vendor.
// serviceKey - The GUID for the vendor network function.
// roleInstanceName - The name of the role instance of the vendor network function.
// options - RoleInstancesClientBeginStopOptions contains the optional parameters for the RoleInstancesClient.BeginStop method.
func (client *RoleInstancesClient) BeginStop(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginStopOptions) (RoleInstancesClientStopPollerResponse, error) {
	resp, err := client.stop(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return RoleInstancesClientStopPollerResponse{}, err
	}
	result := RoleInstancesClientStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("RoleInstancesClient.Stop", "location", resp, client.pl)
	if err != nil {
		return RoleInstancesClientStopPollerResponse{}, err
	}
	result.Poller = &RoleInstancesClientStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Powers off (stop) a role instance of a vendor network function.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *RoleInstancesClient) stop(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, locationName, vendorName, serviceKey, roleInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *RoleInstancesClient) stopCreateRequest(ctx context.Context, locationName string, vendorName string, serviceKey string, roleInstanceName string, options *RoleInstancesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}/roleInstances/{roleInstanceName}/stop"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if serviceKey == "" {
		return nil, errors.New("parameter serviceKey cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceKey}", url.PathEscape(serviceKey))
	if roleInstanceName == "" {
		return nil, errors.New("parameter roleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleInstanceName}", url.PathEscape(roleInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
