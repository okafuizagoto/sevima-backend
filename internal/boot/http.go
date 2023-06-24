package boot

import (
	"go-skeleton-auth/internal/data/auth"
	"go-skeleton-auth/pkg/httpclient"
	"go-skeleton-auth/pkg/tracing"
	"log"
	"net/http"

	"go-skeleton-auth/internal/config"
	jaegerLog "go-skeleton-auth/pkg/log"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	skeletonData "go-skeleton-auth/internal/data/skeleton"
	skeletonServer "go-skeleton-auth/internal/delivery/http"
	skeletonHandler "go-skeleton-auth/internal/delivery/http/skeleton"
	skeletonService "go-skeleton-auth/internal/service/skeleton"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s     skeletonServer.Server    // HTTP Server Object
		sd    skeletonData.Data        // Domain data layer
		ss    skeletonService.Service  // Domain service layer
		sh    *skeletonHandler.Handler // Domain handler
		cfg   *config.Config           // Configuration object
		httpc *httpclient.Client
		ad    auth.Data

		logger *zap.Logger
	)

	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()
	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Set logger used for jaeger
	logger, _ = zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "skeleton"))
	zlogger := jaegerLog.NewFactory(zapLogger)

	// Set tracer for service
	tracer, closer := tracing.Init("skeleton", zlogger)
	defer closer.Close()

	httpc = httpclient.NewClient(tracer)
	ad = auth.New(httpc, cfg.API.Auth)

	// Diganti dengan domain yang anda buat
	sd = skeletonData.New(db, tracer, zlogger)
	ss = skeletonService.New(sd, ad, tracer, zlogger)
	sh = skeletonHandler.New(ss, tracer, zlogger)

	s = skeletonServer.Server{
		Skeleton: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
