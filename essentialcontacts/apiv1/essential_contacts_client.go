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

package essentialcontacts

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	essentialcontactspb "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
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

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	CreateContact   []gax.CallOption
	UpdateContact   []gax.CallOption
	ListContacts    []gax.CallOption
	GetContact      []gax.CallOption
	DeleteContact   []gax.CallOption
	ComputeContacts []gax.CallOption
	SendTestMessage []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("essentialcontacts.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("essentialcontacts.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://essentialcontacts.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		CreateContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListContacts: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ComputeContacts: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SendTestMessage: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		CreateContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		UpdateContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ListContacts: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeleteContact: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		ComputeContacts: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		SendTestMessage: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalClient is an interface that defines the methods available from Essential Contacts API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateContact(context.Context, *essentialcontactspb.CreateContactRequest, ...gax.CallOption) (*essentialcontactspb.Contact, error)
	UpdateContact(context.Context, *essentialcontactspb.UpdateContactRequest, ...gax.CallOption) (*essentialcontactspb.Contact, error)
	ListContacts(context.Context, *essentialcontactspb.ListContactsRequest, ...gax.CallOption) *ContactIterator
	GetContact(context.Context, *essentialcontactspb.GetContactRequest, ...gax.CallOption) (*essentialcontactspb.Contact, error)
	DeleteContact(context.Context, *essentialcontactspb.DeleteContactRequest, ...gax.CallOption) error
	ComputeContacts(context.Context, *essentialcontactspb.ComputeContactsRequest, ...gax.CallOption) *ContactIterator
	SendTestMessage(context.Context, *essentialcontactspb.SendTestMessageRequest, ...gax.CallOption) error
}

// Client is a client for interacting with Essential Contacts API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Manages contacts for important Google Cloud notifications.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateContact adds a new contact for a resource.
func (c *Client) CreateContact(ctx context.Context, req *essentialcontactspb.CreateContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	return c.internalClient.CreateContact(ctx, req, opts...)
}

// UpdateContact updates a contact.
// Note: A contact’s email address cannot be changed.
func (c *Client) UpdateContact(ctx context.Context, req *essentialcontactspb.UpdateContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	return c.internalClient.UpdateContact(ctx, req, opts...)
}

// ListContacts lists the contacts that have been set on a resource.
func (c *Client) ListContacts(ctx context.Context, req *essentialcontactspb.ListContactsRequest, opts ...gax.CallOption) *ContactIterator {
	return c.internalClient.ListContacts(ctx, req, opts...)
}

// GetContact gets a single contact.
func (c *Client) GetContact(ctx context.Context, req *essentialcontactspb.GetContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	return c.internalClient.GetContact(ctx, req, opts...)
}

// DeleteContact deletes a contact.
func (c *Client) DeleteContact(ctx context.Context, req *essentialcontactspb.DeleteContactRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteContact(ctx, req, opts...)
}

// ComputeContacts lists all contacts for the resource that are subscribed to the
// specified notification categories, including contacts inherited from
// any parent resources.
func (c *Client) ComputeContacts(ctx context.Context, req *essentialcontactspb.ComputeContactsRequest, opts ...gax.CallOption) *ContactIterator {
	return c.internalClient.ComputeContacts(ctx, req, opts...)
}

// SendTestMessage allows a contact admin to send a test message to contact to verify that it
// has been configured correctly.
func (c *Client) SendTestMessage(ctx context.Context, req *essentialcontactspb.SendTestMessageRequest, opts ...gax.CallOption) error {
	return c.internalClient.SendTestMessage(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Essential Contacts API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client essentialcontactspb.EssentialContactsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new essential contacts service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Manages contacts for important Google Cloud notifications.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:    connPool,
		client:      essentialcontactspb.NewEssentialContactsServiceClient(connPool),
		CallOptions: &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions
}

// NewRESTClient creates a new essential contacts service rest client.
//
// Manages contacts for important Google Cloud notifications.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://essentialcontacts.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://essentialcontacts.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://essentialcontacts.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) CreateContact(ctx context.Context, req *essentialcontactspb.CreateContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateContact[0:len((*c.CallOptions).CreateContact):len((*c.CallOptions).CreateContact)], opts...)
	var resp *essentialcontactspb.Contact
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateContact(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) UpdateContact(ctx context.Context, req *essentialcontactspb.UpdateContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "contact.name", url.QueryEscape(req.GetContact().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateContact[0:len((*c.CallOptions).UpdateContact):len((*c.CallOptions).UpdateContact)], opts...)
	var resp *essentialcontactspb.Contact
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.UpdateContact(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListContacts(ctx context.Context, req *essentialcontactspb.ListContactsRequest, opts ...gax.CallOption) *ContactIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListContacts[0:len((*c.CallOptions).ListContacts):len((*c.CallOptions).ListContacts)], opts...)
	it := &ContactIterator{}
	req = proto.Clone(req).(*essentialcontactspb.ListContactsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*essentialcontactspb.Contact, string, error) {
		resp := &essentialcontactspb.ListContactsResponse{}
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
			resp, err = c.client.ListContacts(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetContacts(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) GetContact(ctx context.Context, req *essentialcontactspb.GetContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetContact[0:len((*c.CallOptions).GetContact):len((*c.CallOptions).GetContact)], opts...)
	var resp *essentialcontactspb.Contact
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetContact(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) DeleteContact(ctx context.Context, req *essentialcontactspb.DeleteContactRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteContact[0:len((*c.CallOptions).DeleteContact):len((*c.CallOptions).DeleteContact)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteContact(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) ComputeContacts(ctx context.Context, req *essentialcontactspb.ComputeContactsRequest, opts ...gax.CallOption) *ContactIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ComputeContacts[0:len((*c.CallOptions).ComputeContacts):len((*c.CallOptions).ComputeContacts)], opts...)
	it := &ContactIterator{}
	req = proto.Clone(req).(*essentialcontactspb.ComputeContactsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*essentialcontactspb.Contact, string, error) {
		resp := &essentialcontactspb.ComputeContactsResponse{}
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
			resp, err = c.client.ComputeContacts(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetContacts(), resp.GetNextPageToken(), nil
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

func (c *gRPCClient) SendTestMessage(ctx context.Context, req *essentialcontactspb.SendTestMessageRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SendTestMessage[0:len((*c.CallOptions).SendTestMessage):len((*c.CallOptions).SendTestMessage)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.SendTestMessage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateContact adds a new contact for a resource.
func (c *restClient) CreateContact(ctx context.Context, req *essentialcontactspb.CreateContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetContact()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/contacts", req.GetParent())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateContact[0:len((*c.CallOptions).CreateContact):len((*c.CallOptions).CreateContact)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &essentialcontactspb.Contact{}
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

// UpdateContact updates a contact.
// Note: A contact’s email address cannot be changed.
func (c *restClient) UpdateContact(ctx context.Context, req *essentialcontactspb.UpdateContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetContact()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetContact().GetName())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "contact.name", url.QueryEscape(req.GetContact().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).UpdateContact[0:len((*c.CallOptions).UpdateContact):len((*c.CallOptions).UpdateContact)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &essentialcontactspb.Contact{}
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

// ListContacts lists the contacts that have been set on a resource.
func (c *restClient) ListContacts(ctx context.Context, req *essentialcontactspb.ListContactsRequest, opts ...gax.CallOption) *ContactIterator {
	it := &ContactIterator{}
	req = proto.Clone(req).(*essentialcontactspb.ListContactsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*essentialcontactspb.Contact, string, error) {
		resp := &essentialcontactspb.ListContactsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/contacts", req.GetParent())

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
		return resp.GetContacts(), resp.GetNextPageToken(), nil
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

// GetContact gets a single contact.
func (c *restClient) GetContact(ctx context.Context, req *essentialcontactspb.GetContactRequest, opts ...gax.CallOption) (*essentialcontactspb.Contact, error) {
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
	opts = append((*c.CallOptions).GetContact[0:len((*c.CallOptions).GetContact):len((*c.CallOptions).GetContact)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &essentialcontactspb.Contact{}
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

// DeleteContact deletes a contact.
func (c *restClient) DeleteContact(ctx context.Context, req *essentialcontactspb.DeleteContactRequest, opts ...gax.CallOption) error {
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

// ComputeContacts lists all contacts for the resource that are subscribed to the
// specified notification categories, including contacts inherited from
// any parent resources.
func (c *restClient) ComputeContacts(ctx context.Context, req *essentialcontactspb.ComputeContactsRequest, opts ...gax.CallOption) *ContactIterator {
	it := &ContactIterator{}
	req = proto.Clone(req).(*essentialcontactspb.ComputeContactsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*essentialcontactspb.Contact, string, error) {
		resp := &essentialcontactspb.ComputeContactsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/%v/contacts:compute", req.GetParent())

		params := url.Values{}
		params.Add("$alt", "json;enum-encoding=int")
		if items := req.GetNotificationCategories(); len(items) > 0 {
			for _, item := range items {
				params.Add("notificationCategories", fmt.Sprintf("%v", item))
			}
		}
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
		return resp.GetContacts(), resp.GetNextPageToken(), nil
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

// SendTestMessage allows a contact admin to send a test message to contact to verify that it
// has been configured correctly.
func (c *restClient) SendTestMessage(ctx context.Context, req *essentialcontactspb.SendTestMessageRequest, opts ...gax.CallOption) error {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v/contacts:sendTestMessage", req.GetResource())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// ContactIterator manages a stream of *essentialcontactspb.Contact.
type ContactIterator struct {
	items    []*essentialcontactspb.Contact
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
	InternalFetch func(pageSize int, pageToken string) (results []*essentialcontactspb.Contact, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ContactIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ContactIterator) Next() (*essentialcontactspb.Contact, error) {
	var item *essentialcontactspb.Contact
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ContactIterator) bufLen() int {
	return len(it.items)
}

func (it *ContactIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
