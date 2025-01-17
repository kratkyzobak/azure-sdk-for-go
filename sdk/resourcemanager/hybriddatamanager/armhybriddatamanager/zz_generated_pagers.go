//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// DataServicesClientListByDataManagerPager provides operations for iterating over paged responses.
type DataServicesClientListByDataManagerPager struct {
	client    *DataServicesClient
	current   DataServicesClientListByDataManagerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataServicesClientListByDataManagerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataServicesClientListByDataManagerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataServicesClientListByDataManagerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataServiceList.NextLink == nil || len(*p.current.DataServiceList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataManagerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataServicesClientListByDataManagerResponse page.
func (p *DataServicesClientListByDataManagerPager) PageResponse() DataServicesClientListByDataManagerResponse {
	return p.current
}

// DataStoreTypesClientListByDataManagerPager provides operations for iterating over paged responses.
type DataStoreTypesClientListByDataManagerPager struct {
	client    *DataStoreTypesClient
	current   DataStoreTypesClientListByDataManagerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataStoreTypesClientListByDataManagerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataStoreTypesClientListByDataManagerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataStoreTypesClientListByDataManagerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataStoreTypeList.NextLink == nil || len(*p.current.DataStoreTypeList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataManagerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataStoreTypesClientListByDataManagerResponse page.
func (p *DataStoreTypesClientListByDataManagerPager) PageResponse() DataStoreTypesClientListByDataManagerResponse {
	return p.current
}

// DataStoresClientListByDataManagerPager provides operations for iterating over paged responses.
type DataStoresClientListByDataManagerPager struct {
	client    *DataStoresClient
	current   DataStoresClientListByDataManagerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, DataStoresClientListByDataManagerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *DataStoresClientListByDataManagerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *DataStoresClientListByDataManagerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataStoreList.NextLink == nil || len(*p.current.DataStoreList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataManagerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current DataStoresClientListByDataManagerResponse page.
func (p *DataStoresClientListByDataManagerPager) PageResponse() DataStoresClientListByDataManagerResponse {
	return p.current
}

// JobDefinitionsClientListByDataManagerPager provides operations for iterating over paged responses.
type JobDefinitionsClientListByDataManagerPager struct {
	client    *JobDefinitionsClient
	current   JobDefinitionsClientListByDataManagerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobDefinitionsClientListByDataManagerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobDefinitionsClientListByDataManagerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobDefinitionsClientListByDataManagerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobDefinitionList.NextLink == nil || len(*p.current.JobDefinitionList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataManagerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current JobDefinitionsClientListByDataManagerResponse page.
func (p *JobDefinitionsClientListByDataManagerPager) PageResponse() JobDefinitionsClientListByDataManagerResponse {
	return p.current
}

// JobDefinitionsClientListByDataServicePager provides operations for iterating over paged responses.
type JobDefinitionsClientListByDataServicePager struct {
	client    *JobDefinitionsClient
	current   JobDefinitionsClientListByDataServiceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobDefinitionsClientListByDataServiceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobDefinitionsClientListByDataServicePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobDefinitionsClientListByDataServicePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobDefinitionList.NextLink == nil || len(*p.current.JobDefinitionList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataServiceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current JobDefinitionsClientListByDataServiceResponse page.
func (p *JobDefinitionsClientListByDataServicePager) PageResponse() JobDefinitionsClientListByDataServiceResponse {
	return p.current
}

// JobsClientListByDataManagerPager provides operations for iterating over paged responses.
type JobsClientListByDataManagerPager struct {
	client    *JobsClient
	current   JobsClientListByDataManagerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListByDataManagerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsClientListByDataManagerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsClientListByDataManagerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobList.NextLink == nil || len(*p.current.JobList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataManagerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current JobsClientListByDataManagerResponse page.
func (p *JobsClientListByDataManagerPager) PageResponse() JobsClientListByDataManagerResponse {
	return p.current
}

// JobsClientListByDataServicePager provides operations for iterating over paged responses.
type JobsClientListByDataServicePager struct {
	client    *JobsClient
	current   JobsClientListByDataServiceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListByDataServiceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsClientListByDataServicePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsClientListByDataServicePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobList.NextLink == nil || len(*p.current.JobList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataServiceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current JobsClientListByDataServiceResponse page.
func (p *JobsClientListByDataServicePager) PageResponse() JobsClientListByDataServiceResponse {
	return p.current
}

// JobsClientListByJobDefinitionPager provides operations for iterating over paged responses.
type JobsClientListByJobDefinitionPager struct {
	client    *JobsClient
	current   JobsClientListByJobDefinitionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, JobsClientListByJobDefinitionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *JobsClientListByJobDefinitionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *JobsClientListByJobDefinitionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.JobList.NextLink == nil || len(*p.current.JobList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByJobDefinitionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current JobsClientListByJobDefinitionResponse page.
func (p *JobsClientListByJobDefinitionPager) PageResponse() JobsClientListByJobDefinitionResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailableProviderOperations.NextLink == nil || len(*p.current.AvailableProviderOperations.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
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

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// PublicKeysClientListByDataManagerPager provides operations for iterating over paged responses.
type PublicKeysClientListByDataManagerPager struct {
	client    *PublicKeysClient
	current   PublicKeysClientListByDataManagerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PublicKeysClientListByDataManagerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PublicKeysClientListByDataManagerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PublicKeysClientListByDataManagerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PublicKeyList.NextLink == nil || len(*p.current.PublicKeyList.NextLink) == 0 {
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
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByDataManagerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PublicKeysClientListByDataManagerResponse page.
func (p *PublicKeysClientListByDataManagerPager) PageResponse() PublicKeysClientListByDataManagerResponse {
	return p.current
}
