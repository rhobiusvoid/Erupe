// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newTargetHttpProxiesClientHook clientHook

// TargetHttpProxiesCallOptions contains the retry settings for each method of TargetHttpProxiesClient.
type TargetHttpProxiesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete         []gax.CallOption
	Get            []gax.CallOption
	Insert         []gax.CallOption
	List           []gax.CallOption
	Patch          []gax.CallOption
	SetUrlMap      []gax.CallOption
}

// internalTargetHttpProxiesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalTargetHttpProxiesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListTargetHttpProxiesRequest, ...gax.CallOption) *TargetHttpProxiesScopedListPairIterator
	Delete(context.Context, *computepb.DeleteTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetTargetHttpProxyRequest, ...gax.CallOption) (*computepb.TargetHttpProxy, error)
	Insert(context.Context, *computepb.InsertTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListTargetHttpProxiesRequest, ...gax.CallOption) *TargetHttpProxyIterator
	Patch(context.Context, *computepb.PatchTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
	SetUrlMap(context.Context, *computepb.SetUrlMapTargetHttpProxyRequest, ...gax.CallOption) (*Operation, error)
}

// TargetHttpProxiesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetHttpProxies API.
type TargetHttpProxiesClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetHttpProxiesClient

	// The call options for this service.
	CallOptions *TargetHttpProxiesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetHttpProxiesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetHttpProxiesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TargetHttpProxiesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves the list of all TargetHttpProxy resources, regional and global, available to the specified project.
func (c *TargetHttpProxiesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxiesScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified TargetHttpProxy resource.
func (c *TargetHttpProxiesClient) Delete(ctx context.Context, req *computepb.DeleteTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified TargetHttpProxy resource. Gets a list of available target HTTP proxies by making a list() request.
func (c *TargetHttpProxiesClient) Get(ctx context.Context, req *computepb.GetTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpProxy, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a TargetHttpProxy resource in the specified project using the data included in the request.
func (c *TargetHttpProxiesClient) Insert(ctx context.Context, req *computepb.InsertTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of TargetHttpProxy resources available to the specified project.
func (c *TargetHttpProxiesClient) List(ctx context.Context, req *computepb.ListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxyIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified TargetHttpProxy resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *TargetHttpProxiesClient) Patch(ctx context.Context, req *computepb.PatchTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetUrlMap changes the URL map for TargetHttpProxy.
func (c *TargetHttpProxiesClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.SetUrlMap(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetHttpProxiesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTargetHttpProxiesRESTClient creates a new target http proxies rest client.
//
// The TargetHttpProxies API.
func NewTargetHttpProxiesRESTClient(ctx context.Context, opts ...option.ClientOption) (*TargetHttpProxiesClient, error) {
	clientOpts := append(defaultTargetHttpProxiesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &targetHttpProxiesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &TargetHttpProxiesClient{internalClient: c, CallOptions: &TargetHttpProxiesCallOptions{}}, nil
}

func defaultTargetHttpProxiesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *targetHttpProxiesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetHttpProxiesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *targetHttpProxiesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves the list of all TargetHttpProxy resources, regional and global, available to the specified project.
func (c *targetHttpProxiesRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxiesScopedListPairIterator {
	it := &TargetHttpProxiesScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListTargetHttpProxiesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]TargetHttpProxiesScopedListPair, string, error) {
		resp := &computepb.TargetHttpProxyAggregatedList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/targetHttpProxies", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp

		elems := make([]TargetHttpProxiesScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, TargetHttpProxiesScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Delete deletes the specified TargetHttpProxy resource.
func (c *targetHttpProxiesRESTClient) Delete(ctx context.Context, req *computepb.DeleteTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetHttpProxies/%v", req.GetProject(), req.GetTargetHttpProxy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified TargetHttpProxy resource. Gets a list of available target HTTP proxies by making a list() request.
func (c *targetHttpProxiesRESTClient) Get(ctx context.Context, req *computepb.GetTargetHttpProxyRequest, opts ...gax.CallOption) (*computepb.TargetHttpProxy, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetHttpProxies/%v", req.GetProject(), req.GetTargetHttpProxy())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.TargetHttpProxy{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// Insert creates a TargetHttpProxy resource in the specified project using the data included in the request.
func (c *targetHttpProxiesRESTClient) Insert(ctx context.Context, req *computepb.InsertTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetHttpProxyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetHttpProxies", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves the list of TargetHttpProxy resources available to the specified project.
func (c *targetHttpProxiesRESTClient) List(ctx context.Context, req *computepb.ListTargetHttpProxiesRequest, opts ...gax.CallOption) *TargetHttpProxyIterator {
	it := &TargetHttpProxyIterator{}
	req = proto.Clone(req).(*computepb.ListTargetHttpProxiesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.TargetHttpProxy, string, error) {
		resp := &computepb.TargetHttpProxyList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetHttpProxies", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// Patch patches the specified TargetHttpProxy resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *targetHttpProxiesRESTClient) Patch(ctx context.Context, req *computepb.PatchTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTargetHttpProxyResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/targetHttpProxies/%v", req.GetProject(), req.GetTargetHttpProxy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// SetUrlMap changes the URL map for TargetHttpProxy.
func (c *targetHttpProxiesRESTClient) SetUrlMap(ctx context.Context, req *computepb.SetUrlMapTargetHttpProxyRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetUrlMapReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/targetHttpProxies/%v/setUrlMap", req.GetProject(), req.GetTargetHttpProxy())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// TargetHttpProxiesScopedListPair is a holder type for string/*computepb.TargetHttpProxiesScopedList map entries
type TargetHttpProxiesScopedListPair struct {
	Key   string
	Value *computepb.TargetHttpProxiesScopedList
}

// TargetHttpProxiesScopedListPairIterator manages a stream of TargetHttpProxiesScopedListPair.
type TargetHttpProxiesScopedListPairIterator struct {
	items    []TargetHttpProxiesScopedListPair
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []TargetHttpProxiesScopedListPair, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TargetHttpProxiesScopedListPairIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TargetHttpProxiesScopedListPairIterator) Next() (TargetHttpProxiesScopedListPair, error) {
	var item TargetHttpProxiesScopedListPair
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TargetHttpProxiesScopedListPairIterator) bufLen() int {
	return len(it.items)
}

func (it *TargetHttpProxiesScopedListPairIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
