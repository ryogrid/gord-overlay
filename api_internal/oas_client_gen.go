// Code generated by ogen, DO NOT EDIT.

package api_internal

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// InternalServiceDeleteValueInner invokes InternalService_DeleteValueInner operation.
	//
	// POST /server.InternalService/DeleteValueInner
	InternalServiceDeleteValueInner(ctx context.Context, params InternalServiceDeleteValueInnerParams) error
	// InternalServiceFindClosestPrecedingNode invokes InternalService_FindClosestPrecedingNode operation.
	//
	// POST /server.InternalService/FindClosestPrecedingNode
	InternalServiceFindClosestPrecedingNode(ctx context.Context, params InternalServiceFindClosestPrecedingNodeParams) error
	// InternalServiceFindSuccessorByList invokes InternalService_FindSuccessorByList operation.
	//
	// POST /server.InternalService/FindSuccessorByList
	InternalServiceFindSuccessorByList(ctx context.Context, params InternalServiceFindSuccessorByListParams) error
	// InternalServiceFindSuccessorByTable invokes InternalService_FindSuccessorByTable operation.
	//
	// POST /server.InternalService/FindSuccessorByTable
	InternalServiceFindSuccessorByTable(ctx context.Context, params InternalServiceFindSuccessorByTableParams) error
	// InternalServiceGetValueInner invokes InternalService_GetValueInner operation.
	//
	// POST /server.InternalService/GetValueInner
	InternalServiceGetValueInner(ctx context.Context, params InternalServiceGetValueInnerParams) error
	// InternalServiceNotify invokes InternalService_Notify operation.
	//
	// POST /server.InternalService/Notify
	InternalServiceNotify(ctx context.Context, params InternalServiceNotifyParams) error
	// InternalServicePing invokes InternalService_Ping operation.
	//
	// POST /server.InternalService/Ping
	InternalServicePing(ctx context.Context) error
	// InternalServicePredecessor invokes InternalService_Predecessor operation.
	//
	// POST /server.InternalService/Predecessor
	InternalServicePredecessor(ctx context.Context) error
	// InternalServicePutValueInner invokes InternalService_PutValueInner operation.
	//
	// POST /server.InternalService/PutValueInner
	InternalServicePutValueInner(ctx context.Context, params InternalServicePutValueInnerParams) error
	// InternalServiceSuccessors invokes InternalService_Successors operation.
	//
	// POST /server.InternalService/Successors
	InternalServiceSuccessors(ctx context.Context) error
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// InternalServiceDeleteValueInner invokes InternalService_DeleteValueInner operation.
//
// POST /server.InternalService/DeleteValueInner
func (c *Client) InternalServiceDeleteValueInner(ctx context.Context, params InternalServiceDeleteValueInnerParams) error {
	_, err := c.sendInternalServiceDeleteValueInner(ctx, params)
	return err
}

func (c *Client) sendInternalServiceDeleteValueInner(ctx context.Context, params InternalServiceDeleteValueInnerParams) (res *InternalServiceDeleteValueInnerOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_DeleteValueInner"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/DeleteValueInner"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceDeleteValueInner",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/DeleteValueInner"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "key" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "key",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Key.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceDeleteValueInnerResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServiceFindClosestPrecedingNode invokes InternalService_FindClosestPrecedingNode operation.
//
// POST /server.InternalService/FindClosestPrecedingNode
func (c *Client) InternalServiceFindClosestPrecedingNode(ctx context.Context, params InternalServiceFindClosestPrecedingNodeParams) error {
	_, err := c.sendInternalServiceFindClosestPrecedingNode(ctx, params)
	return err
}

func (c *Client) sendInternalServiceFindClosestPrecedingNode(ctx context.Context, params InternalServiceFindClosestPrecedingNodeParams) (res *InternalServiceFindClosestPrecedingNodeOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_FindClosestPrecedingNode"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/FindClosestPrecedingNode"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceFindClosestPrecedingNode",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/FindClosestPrecedingNode"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.BytesToString(params.ID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceFindClosestPrecedingNodeResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServiceFindSuccessorByList invokes InternalService_FindSuccessorByList operation.
//
// POST /server.InternalService/FindSuccessorByList
func (c *Client) InternalServiceFindSuccessorByList(ctx context.Context, params InternalServiceFindSuccessorByListParams) error {
	_, err := c.sendInternalServiceFindSuccessorByList(ctx, params)
	return err
}

func (c *Client) sendInternalServiceFindSuccessorByList(ctx context.Context, params InternalServiceFindSuccessorByListParams) (res *InternalServiceFindSuccessorByListOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_FindSuccessorByList"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/FindSuccessorByList"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceFindSuccessorByList",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/FindSuccessorByList"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.BytesToString(params.ID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceFindSuccessorByListResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServiceFindSuccessorByTable invokes InternalService_FindSuccessorByTable operation.
//
// POST /server.InternalService/FindSuccessorByTable
func (c *Client) InternalServiceFindSuccessorByTable(ctx context.Context, params InternalServiceFindSuccessorByTableParams) error {
	_, err := c.sendInternalServiceFindSuccessorByTable(ctx, params)
	return err
}

func (c *Client) sendInternalServiceFindSuccessorByTable(ctx context.Context, params InternalServiceFindSuccessorByTableParams) (res *InternalServiceFindSuccessorByTableOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_FindSuccessorByTable"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/FindSuccessorByTable"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceFindSuccessorByTable",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/FindSuccessorByTable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.BytesToString(params.ID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceFindSuccessorByTableResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServiceGetValueInner invokes InternalService_GetValueInner operation.
//
// POST /server.InternalService/GetValueInner
func (c *Client) InternalServiceGetValueInner(ctx context.Context, params InternalServiceGetValueInnerParams) error {
	_, err := c.sendInternalServiceGetValueInner(ctx, params)
	return err
}

func (c *Client) sendInternalServiceGetValueInner(ctx context.Context, params InternalServiceGetValueInnerParams) (res *InternalServiceGetValueInnerOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_GetValueInner"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/GetValueInner"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceGetValueInner",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/GetValueInner"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "key" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "key",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Key.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceGetValueInnerResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServiceNotify invokes InternalService_Notify operation.
//
// POST /server.InternalService/Notify
func (c *Client) InternalServiceNotify(ctx context.Context, params InternalServiceNotifyParams) error {
	_, err := c.sendInternalServiceNotify(ctx, params)
	return err
}

func (c *Client) sendInternalServiceNotify(ctx context.Context, params InternalServiceNotifyParams) (res *InternalServiceNotifyOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Notify"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Notify"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceNotify",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/Notify"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "host" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "host",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Host.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceNotifyResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServicePing invokes InternalService_Ping operation.
//
// POST /server.InternalService/Ping
func (c *Client) InternalServicePing(ctx context.Context) error {
	_, err := c.sendInternalServicePing(ctx)
	return err
}

func (c *Client) sendInternalServicePing(ctx context.Context) (res *InternalServicePingOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Ping"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Ping"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServicePing",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/Ping"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServicePingResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServicePredecessor invokes InternalService_Predecessor operation.
//
// POST /server.InternalService/Predecessor
func (c *Client) InternalServicePredecessor(ctx context.Context) error {
	_, err := c.sendInternalServicePredecessor(ctx)
	return err
}

func (c *Client) sendInternalServicePredecessor(ctx context.Context) (res *InternalServicePredecessorOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Predecessor"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Predecessor"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServicePredecessor",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/Predecessor"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServicePredecessorResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServicePutValueInner invokes InternalService_PutValueInner operation.
//
// POST /server.InternalService/PutValueInner
func (c *Client) InternalServicePutValueInner(ctx context.Context, params InternalServicePutValueInnerParams) error {
	_, err := c.sendInternalServicePutValueInner(ctx, params)
	return err
}

func (c *Client) sendInternalServicePutValueInner(ctx context.Context, params InternalServicePutValueInnerParams) (res *InternalServicePutValueInnerOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_PutValueInner"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/PutValueInner"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServicePutValueInner",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/PutValueInner"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "key" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "key",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Key.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "value" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "value",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Value.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServicePutValueInnerResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InternalServiceSuccessors invokes InternalService_Successors operation.
//
// POST /server.InternalService/Successors
func (c *Client) InternalServiceSuccessors(ctx context.Context) error {
	_, err := c.sendInternalServiceSuccessors(ctx)
	return err
}

func (c *Client) sendInternalServiceSuccessors(ctx context.Context) (res *InternalServiceSuccessorsOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Successors"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Successors"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "InternalServiceSuccessors",
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/server.InternalService/Successors"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeInternalServiceSuccessorsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
