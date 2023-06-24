package skeleton

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	DataSiswaa "go-skeleton-auth/internal/entity/skeleton"
	skeletonEntity "go-skeleton-auth/internal/entity/skeleton"
	jaegerLog "go-skeleton-auth/pkg/log"
	"go-skeleton-auth/pkg/response"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

// ISkeletonSvc is an interface to Skeleton Service
// Masukkan function dari service ke dalam interface ini
type ISkeletonSvc interface {
	GetSkeleton(ctx context.Context) error
	PostSkeleton(ctx context.Context) (skeletonEntity.Skeleton, error)
	GetDataSiswa(ctx context.Context) ([]DataSiswaa.DataSiswa, error)
}

type (
	// Handler ...
	Handler struct {
		skeletonSvc ISkeletonSvc
		tracer      opentracing.Tracer
		logger      jaegerLog.Factory
	}
)

// New for bridging product handler initialization
func New(is ISkeletonSvc, tracer opentracing.Tracer, logger jaegerLog.Factory) *Handler {
	return &Handler{
		skeletonSvc: is,
		tracer:      tracer,
		logger:      logger,
	}
}

// SkeletonHandler will receive request and return response
func (h *Handler) SkeletonHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx context.Context
		//_token   string
		resp     *response.Response
		result   interface{}
		metadata interface{}
		err      error
		errRes   response.Error
	)
	spanCtx, _ := h.tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := h.tracer.StartSpan("SkeletonHandler", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	//ctx = opentracing.ContextWithSpan(ctx, span)
	//h.logger.For(ctx).Info("HTTP request received", zap.String("method", r.Method), zap.Stringer("url", r.URL))

	if err == nil {
		switch r.Method {
		// Check if request method is GET
		case http.MethodGet:
			//h.logger.For(ctx).Info("Running Service", zap.String("service", "GetSkeleton"))
			err = h.skeletonSvc.GetSkeleton(ctx)
		// Check if request method is POST
		case http.MethodPost:

			result, err = h.skeletonSvc.PostSkeleton(ctx)
		// Check if request method is PUT
		case http.MethodPut:

		// Check if request method is DELETE
		case http.MethodDelete:

		default:
			err = errors.New("400")
		}
	}

	// If anything from service or data return an error
	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   101,
			Msg:    "101 - Data Not Found",
			Status: true,
		}
		// If service returns an error
		if strings.Contains(err.Error(), "service") {
			// Replace error with server error
			errRes = response.Error{
				Code:   500,
				Msg:    "500 - Internal Server Error",
				Status: true,
			}
		}
		// If error 401
		if strings.Contains(err.Error(), "401") {
			// Replace error with server error
			errRes = response.Error{
				Code:   401,
				Msg:    "401 - Unauthorized",
				Status: true,
			}
		}
		// If error 403
		if strings.Contains(err.Error(), "403") {
			// Replace error with server error
			errRes = response.Error{
				Code:   403,
				Msg:    "403 - Forbidden",
				Status: true,
			}
		}
		// If error 400
		if strings.Contains(err.Error(), "400") {
			// Replace error with server error
			errRes = response.Error{
				Code:   400,
				Msg:    "400 - Bad Request",
				Status: true,
			}
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		h.logger.For(ctx).Error("HTTP request error", zap.String("method", r.Method), zap.Stringer("url", r.URL), zap.Error(err))
		resp.StatusCode = errRes.Code
		resp.Error = errRes
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
	//h.logger.For(ctx).Info("HTTP request done", zap.String("method", r.Method), zap.Stringer("url", r.URL))
}
