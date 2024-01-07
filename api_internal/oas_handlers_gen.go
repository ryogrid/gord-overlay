// Code generated by ogen, DO NOT EDIT.

package api_internal

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.19.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleInternalServiceDeleteValueInnerRequest handles InternalService_DeleteValueInner operation.
//
// POST /server.InternalService/DeleteValueInner
func (s *Server) handleInternalServiceDeleteValueInnerRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_DeleteValueInner"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/DeleteValueInner"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceDeleteValueInner",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServiceDeleteValueInner",
			ID:   "InternalService_DeleteValueInner",
		}
	)
	params, err := decodeInternalServiceDeleteValueInnerParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServiceDeleteValueInnerOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceDeleteValueInner",
			OperationSummary: "",
			OperationID:      "InternalService_DeleteValueInner",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "key",
					In:   "query",
				}: params.Key,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServiceDeleteValueInnerParams
			Response = *InternalServiceDeleteValueInnerOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServiceDeleteValueInnerParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceDeleteValueInner(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceDeleteValueInner(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceDeleteValueInnerResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServiceFindClosestPrecedingNodeRequest handles InternalService_FindClosestPrecedingNode operation.
//
// POST /server.InternalService/FindClosestPrecedingNode
func (s *Server) handleInternalServiceFindClosestPrecedingNodeRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_FindClosestPrecedingNode"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/FindClosestPrecedingNode"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceFindClosestPrecedingNode",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServiceFindClosestPrecedingNode",
			ID:   "InternalService_FindClosestPrecedingNode",
		}
	)
	params, err := decodeInternalServiceFindClosestPrecedingNodeParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServiceFindClosestPrecedingNodeOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceFindClosestPrecedingNode",
			OperationSummary: "",
			OperationID:      "InternalService_FindClosestPrecedingNode",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "query",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServiceFindClosestPrecedingNodeParams
			Response = *InternalServiceFindClosestPrecedingNodeOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServiceFindClosestPrecedingNodeParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceFindClosestPrecedingNode(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceFindClosestPrecedingNode(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceFindClosestPrecedingNodeResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServiceFindSuccessorByListRequest handles InternalService_FindSuccessorByList operation.
//
// POST /server.InternalService/FindSuccessorByList
func (s *Server) handleInternalServiceFindSuccessorByListRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_FindSuccessorByList"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/FindSuccessorByList"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceFindSuccessorByList",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServiceFindSuccessorByList",
			ID:   "InternalService_FindSuccessorByList",
		}
	)
	params, err := decodeInternalServiceFindSuccessorByListParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServiceFindSuccessorByListOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceFindSuccessorByList",
			OperationSummary: "",
			OperationID:      "InternalService_FindSuccessorByList",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "query",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServiceFindSuccessorByListParams
			Response = *InternalServiceFindSuccessorByListOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServiceFindSuccessorByListParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceFindSuccessorByList(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceFindSuccessorByList(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceFindSuccessorByListResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServiceFindSuccessorByTableRequest handles InternalService_FindSuccessorByTable operation.
//
// POST /server.InternalService/FindSuccessorByTable
func (s *Server) handleInternalServiceFindSuccessorByTableRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_FindSuccessorByTable"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/FindSuccessorByTable"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceFindSuccessorByTable",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServiceFindSuccessorByTable",
			ID:   "InternalService_FindSuccessorByTable",
		}
	)
	params, err := decodeInternalServiceFindSuccessorByTableParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServiceFindSuccessorByTableOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceFindSuccessorByTable",
			OperationSummary: "",
			OperationID:      "InternalService_FindSuccessorByTable",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "id",
					In:   "query",
				}: params.ID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServiceFindSuccessorByTableParams
			Response = *InternalServiceFindSuccessorByTableOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServiceFindSuccessorByTableParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceFindSuccessorByTable(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceFindSuccessorByTable(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceFindSuccessorByTableResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServiceGetValueInnerRequest handles InternalService_GetValueInner operation.
//
// POST /server.InternalService/GetValueInner
func (s *Server) handleInternalServiceGetValueInnerRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_GetValueInner"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/GetValueInner"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceGetValueInner",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServiceGetValueInner",
			ID:   "InternalService_GetValueInner",
		}
	)
	params, err := decodeInternalServiceGetValueInnerParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServiceGetValueInnerOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceGetValueInner",
			OperationSummary: "",
			OperationID:      "InternalService_GetValueInner",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "key",
					In:   "query",
				}: params.Key,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServiceGetValueInnerParams
			Response = *InternalServiceGetValueInnerOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServiceGetValueInnerParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceGetValueInner(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceGetValueInner(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceGetValueInnerResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServiceNotifyRequest handles InternalService_Notify operation.
//
// POST /server.InternalService/Notify
func (s *Server) handleInternalServiceNotifyRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Notify"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Notify"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceNotify",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServiceNotify",
			ID:   "InternalService_Notify",
		}
	)
	params, err := decodeInternalServiceNotifyParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServiceNotifyOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceNotify",
			OperationSummary: "",
			OperationID:      "InternalService_Notify",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "host",
					In:   "query",
				}: params.Host,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServiceNotifyParams
			Response = *InternalServiceNotifyOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServiceNotifyParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceNotify(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceNotify(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceNotifyResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServicePingRequest handles InternalService_Ping operation.
//
// POST /server.InternalService/Ping
func (s *Server) handleInternalServicePingRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Ping"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Ping"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServicePing",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err error
	)

	var response *InternalServicePingOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServicePing",
			OperationSummary: "",
			OperationID:      "InternalService_Ping",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *InternalServicePingOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServicePing(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServicePing(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServicePingResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServicePredecessorRequest handles InternalService_Predecessor operation.
//
// POST /server.InternalService/Predecessor
func (s *Server) handleInternalServicePredecessorRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Predecessor"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Predecessor"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServicePredecessor",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err error
	)

	var response *InternalServicePredecessorOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServicePredecessor",
			OperationSummary: "",
			OperationID:      "InternalService_Predecessor",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *InternalServicePredecessorOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServicePredecessor(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServicePredecessor(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServicePredecessorResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServicePutValueInnerRequest handles InternalService_PutValueInner operation.
//
// POST /server.InternalService/PutValueInner
func (s *Server) handleInternalServicePutValueInnerRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_PutValueInner"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/PutValueInner"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServicePutValueInner",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "InternalServicePutValueInner",
			ID:   "InternalService_PutValueInner",
		}
	)
	params, err := decodeInternalServicePutValueInnerParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response *InternalServicePutValueInnerOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServicePutValueInner",
			OperationSummary: "",
			OperationID:      "InternalService_PutValueInner",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "key",
					In:   "query",
				}: params.Key,
				{
					Name: "value",
					In:   "query",
				}: params.Value,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = InternalServicePutValueInnerParams
			Response = *InternalServicePutValueInnerOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackInternalServicePutValueInnerParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServicePutValueInner(ctx, params)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServicePutValueInner(ctx, params)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServicePutValueInnerResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleInternalServiceSuccessorsRequest handles InternalService_Successors operation.
//
// POST /server.InternalService/Successors
func (s *Server) handleInternalServiceSuccessorsRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("InternalService_Successors"),
		semconv.HTTPMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/server.InternalService/Successors"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "InternalServiceSuccessors",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	s.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		err error
	)

	var response *InternalServiceSuccessorsOK
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "InternalServiceSuccessors",
			OperationSummary: "",
			OperationID:      "InternalService_Successors",
			Body:             nil,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = struct{}
			Params   = struct{}
			Response = *InternalServiceSuccessorsOK
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				err = s.h.InternalServiceSuccessors(ctx)
				return response, err
			},
		)
	} else {
		err = s.h.InternalServiceSuccessors(ctx)
	}
	if err != nil {
		recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeInternalServiceSuccessorsResponse(response, w, span); err != nil {
		recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
