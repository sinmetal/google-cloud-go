// Copyright 2023 Google LLC
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

package dashboard

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	dashboardpb "cloud.google.com/go/monitoring/dashboard/apiv1/dashboardpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newDashboardsClientHook clientHook

// DashboardsCallOptions contains the retry settings for each method of DashboardsClient.
type DashboardsCallOptions struct {
	CreateDashboard []gax.CallOption
	ListDashboards  []gax.CallOption
	GetDashboard    []gax.CallOption
	DeleteDashboard []gax.CallOption
	UpdateDashboard []gax.CallOption
}

func defaultDashboardsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultDashboardsCallOptions() *DashboardsCallOptions {
	return &DashboardsCallOptions{
		CreateDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		ListDashboards: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		UpdateDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
	}
}

func defaultDashboardsRESTCallOptions() *DashboardsCallOptions {
	return &DashboardsCallOptions{
		CreateDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		ListDashboards: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusInternalServerError)
			}),
		},
		GetDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable,
					http.StatusInternalServerError)
			}),
		},
		DeleteDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
		UpdateDashboard: []gax.CallOption{
			gax.WithTimeout(30000 * time.Millisecond),
		},
	}
}

// internalDashboardsClient is an interface that defines the methods available from Cloud Monitoring API.
type internalDashboardsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateDashboard(context.Context, *dashboardpb.CreateDashboardRequest, ...gax.CallOption) (*dashboardpb.Dashboard, error)
	ListDashboards(context.Context, *dashboardpb.ListDashboardsRequest, ...gax.CallOption) *DashboardIterator
	GetDashboard(context.Context, *dashboardpb.GetDashboardRequest, ...gax.CallOption) (*dashboardpb.Dashboard, error)
	DeleteDashboard(context.Context, *dashboardpb.DeleteDashboardRequest, ...gax.CallOption) error
	UpdateDashboard(context.Context, *dashboardpb.UpdateDashboardRequest, ...gax.CallOption) (*dashboardpb.Dashboard, error)
}

// DashboardsClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages Stackdriver dashboards. A dashboard is an arrangement of data display
// widgets in a specific layout.
type DashboardsClient struct {
	// The internal transport-dependent client.
	internalClient internalDashboardsClient

	// The call options for this service.
	CallOptions *DashboardsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DashboardsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DashboardsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *DashboardsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateDashboard creates a new custom dashboard. For examples on how you can use this API to
// create dashboards, see Managing dashboards by
// API (at https://cloud.google.com/monitoring/dashboards/api-dashboard). This
// method requires the monitoring.dashboards.create permission on the
// specified project. For more information about permissions, see Cloud
// Identity and Access Management (at https://cloud.google.com/iam).
func (c *DashboardsClient) CreateDashboard(ctx context.Context, req *dashboardpb.CreateDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	return c.internalClient.CreateDashboard(ctx, req, opts...)
}

// ListDashboards lists the existing dashboards.
//
// This method requires the monitoring.dashboards.list permission
// on the specified project. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *DashboardsClient) ListDashboards(ctx context.Context, req *dashboardpb.ListDashboardsRequest, opts ...gax.CallOption) *DashboardIterator {
	return c.internalClient.ListDashboards(ctx, req, opts...)
}

// GetDashboard fetches a specific dashboard.
//
// This method requires the monitoring.dashboards.get permission
// on the specified dashboard. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *DashboardsClient) GetDashboard(ctx context.Context, req *dashboardpb.GetDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	return c.internalClient.GetDashboard(ctx, req, opts...)
}

// DeleteDashboard deletes an existing custom dashboard.
//
// This method requires the monitoring.dashboards.delete permission
// on the specified dashboard. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *DashboardsClient) DeleteDashboard(ctx context.Context, req *dashboardpb.DeleteDashboardRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteDashboard(ctx, req, opts...)
}

// UpdateDashboard replaces an existing custom dashboard with a new definition.
//
// This method requires the monitoring.dashboards.update permission
// on the specified dashboard. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *DashboardsClient) UpdateDashboard(ctx context.Context, req *dashboardpb.UpdateDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	return c.internalClient.UpdateDashboard(ctx, req, opts...)
}

// dashboardsGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type dashboardsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing DashboardsClient
	CallOptions **DashboardsCallOptions

	// The gRPC API client.
	dashboardsClient dashboardpb.DashboardsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDashboardsClient creates a new dashboards service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages Stackdriver dashboards. A dashboard is an arrangement of data display
// widgets in a specific layout.
func NewDashboardsClient(ctx context.Context, opts ...option.ClientOption) (*DashboardsClient, error) {
	clientOpts := defaultDashboardsGRPCClientOptions()
	if newDashboardsClientHook != nil {
		hookOpts, err := newDashboardsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := DashboardsClient{CallOptions: defaultDashboardsCallOptions()}

	c := &dashboardsGRPCClient{
		connPool:         connPool,
		dashboardsClient: dashboardpb.NewDashboardsServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *dashboardsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *dashboardsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *dashboardsGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type dashboardsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing DashboardsClient
	CallOptions **DashboardsCallOptions
}

// NewDashboardsRESTClient creates a new dashboards service rest client.
//
// Manages Stackdriver dashboards. A dashboard is an arrangement of data display
// widgets in a specific layout.
func NewDashboardsRESTClient(ctx context.Context, opts ...option.ClientOption) (*DashboardsClient, error) {
	clientOpts := append(defaultDashboardsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultDashboardsRESTCallOptions()
	c := &dashboardsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &DashboardsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultDashboardsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://monitoring.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://monitoring.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *dashboardsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *dashboardsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *dashboardsRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *dashboardsGRPCClient) CreateDashboard(ctx context.Context, req *dashboardpb.CreateDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateDashboard[0:len((*c.CallOptions).CreateDashboard):len((*c.CallOptions).CreateDashboard)], opts...)
	var resp *dashboardpb.Dashboard
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dashboardsClient.CreateDashboard(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *dashboardsGRPCClient) ListDashboards(ctx context.Context, req *dashboardpb.ListDashboardsRequest, opts ...gax.CallOption) *DashboardIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListDashboards[0:len((*c.CallOptions).ListDashboards):len((*c.CallOptions).ListDashboards)], opts...)
	it := &DashboardIterator{}
	req = proto.Clone(req).(*dashboardpb.ListDashboardsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dashboardpb.Dashboard, string, error) {
		resp := &dashboardpb.ListDashboardsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.dashboardsClient.ListDashboards(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDashboards(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *dashboardsGRPCClient) GetDashboard(ctx context.Context, req *dashboardpb.GetDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetDashboard[0:len((*c.CallOptions).GetDashboard):len((*c.CallOptions).GetDashboard)], opts...)
	var resp *dashboardpb.Dashboard
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dashboardsClient.GetDashboard(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *dashboardsGRPCClient) DeleteDashboard(ctx context.Context, req *dashboardpb.DeleteDashboardRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteDashboard[0:len((*c.CallOptions).DeleteDashboard):len((*c.CallOptions).DeleteDashboard)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.dashboardsClient.DeleteDashboard(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *dashboardsGRPCClient) UpdateDashboard(ctx context.Context, req *dashboardpb.UpdateDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "dashboard.name", url.QueryEscape(req.GetDashboard().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateDashboard[0:len((*c.CallOptions).UpdateDashboard):len((*c.CallOptions).UpdateDashboard)], opts...)
	var resp *dashboardpb.Dashboard
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.dashboardsClient.UpdateDashboard(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateDashboard creates a new custom dashboard. For examples on how you can use this API to
// create dashboards, see Managing dashboards by
// API (at https://cloud.google.com/monitoring/dashboards/api-dashboard). This
// method requires the monitoring.dashboards.create permission on the
// specified project. For more information about permissions, see Cloud
// Identity and Access Management (at https://cloud.google.com/iam).
func (c *dashboardsRESTClient) CreateDashboard(ctx context.Context, req *dashboardpb.CreateDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetDashboard()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/dashboards", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetValidateOnly() {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateDashboard[0:len((*c.CallOptions).CreateDashboard):len((*c.CallOptions).CreateDashboard)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dashboardpb.Dashboard{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListDashboards lists the existing dashboards.
//
// This method requires the monitoring.dashboards.list permission
// on the specified project. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *dashboardsRESTClient) ListDashboards(ctx context.Context, req *dashboardpb.ListDashboardsRequest, opts ...gax.CallOption) *DashboardIterator {
	it := &DashboardIterator{}
	req = proto.Clone(req).(*dashboardpb.ListDashboardsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dashboardpb.Dashboard, string, error) {
		resp := &dashboardpb.ListDashboardsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/dashboards", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetDashboards(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// GetDashboard fetches a specific dashboard.
//
// This method requires the monitoring.dashboards.get permission
// on the specified dashboard. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *dashboardsRESTClient) GetDashboard(ctx context.Context, req *dashboardpb.GetDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetDashboard[0:len((*c.CallOptions).GetDashboard):len((*c.CallOptions).GetDashboard)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dashboardpb.Dashboard{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// DeleteDashboard deletes an existing custom dashboard.
//
// This method requires the monitoring.dashboards.delete permission
// on the specified dashboard. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *dashboardsRESTClient) DeleteDashboard(ctx context.Context, req *dashboardpb.DeleteDashboardRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// UpdateDashboard replaces an existing custom dashboard with a new definition.
//
// This method requires the monitoring.dashboards.update permission
// on the specified dashboard. For more information, see
// Cloud Identity and Access Management (at https://cloud.google.com/iam).
func (c *dashboardsRESTClient) UpdateDashboard(ctx context.Context, req *dashboardpb.UpdateDashboardRequest, opts ...gax.CallOption) (*dashboardpb.Dashboard, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetDashboard()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetDashboard().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetValidateOnly() {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "dashboard.name", url.QueryEscape(req.GetDashboard().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).UpdateDashboard[0:len((*c.CallOptions).UpdateDashboard):len((*c.CallOptions).UpdateDashboard)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dashboardpb.Dashboard{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// DashboardIterator manages a stream of *dashboardpb.Dashboard.
type DashboardIterator struct {
	items    []*dashboardpb.Dashboard
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
	InternalFetch func(pageSize int, pageToken string) (results []*dashboardpb.Dashboard, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DashboardIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DashboardIterator) Next() (*dashboardpb.Dashboard, error) {
	var item *dashboardpb.Dashboard
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DashboardIterator) bufLen() int {
	return len(it.items)
}

func (it *DashboardIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
