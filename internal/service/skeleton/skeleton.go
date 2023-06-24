package skeleton

import (
	"context"
	"fmt"
	"go-skeleton-auth/internal/entity/auth"
	DataSiswaa "go-skeleton-auth/internal/entity/skeleton"
	skeletonEntity "go-skeleton-auth/internal/entity/skeleton"
	jaegerLog "go-skeleton-auth/pkg/log"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetDataSiswa(ctx context.Context) ([]DataSiswaa.DataSiswa, error)
}

// AuthData ...
type AuthData interface {
	CheckAuth(ctx context.Context, code string) (auth.Auth, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data     Data
	authData AuthData
	tracer   opentracing.Tracer
	logger   jaegerLog.Factory
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data, authData AuthData, tracer opentracing.Tracer, logger jaegerLog.Factory) Service {
	// Assign variable dari parameter ke object
	return Service{
		data:     data,
		authData: authData,
		tracer:   tracer,
		logger:   logger,
	}
}

// GetSkeleton ...
func (s Service) GetSkeleton(ctx context.Context) error {
	// Check if have span on context

	fmt.Println("MASUK SINI")

	// CHECK AUTH
	// auth, err := s.authData.CheckAuth(ctx, "191")
	// if err != nil {
	// 	s.logger.For(ctx).Error("Failed to check auth", zap.Error(err))
	// 	return errors.Wrap(err, "[SERVICE][GetSkeleton]")
	// }
	// if auth.Error.Status == true {
	// 	s.logger.For(ctx).Error("401 Unauthorized")
	// 	return errors.Wrap(errors.New("401 Unauthorized"), "[SERVICE][GetSkeleton]")
	// }
	// END CHECK AUTH

	return nil
}

// GetSkeleton ...
func (s Service) PostSkeleton(ctx context.Context) (skeletonEntity.Skeleton, error) {
	// Check if have span on context

	fmt.Println("MASUK SINI POST")
	var data skeletonEntity.Skeleton

	data.SkeletonID = 1
	//data.SkeletonName = "TESTING 1"
	data.SkeletonType = "BONE"

	return data, nil
}

// GetSkeleton ...
func (s Service) GetDataSiswa(ctx context.Context) ([]DataSiswaa.DataSiswa, error) {
	// Check if have span on context
	var (
		datasiswas []DataSiswaa.DataSiswa
		err        error
	)

	datasiswas, err = s.data.GetDataSiswa(ctx)
	if err != nil {
		return datasiswas, errors.Wrap(err, "[Service] [GetDataSiswa]")
	}

	return datasiswas, nil
	// fmt.Println("MASUK SINI POST")
	// var data skeletonEntity.Skeleton

	// data.SkeletonID = 1
	// //data.SkeletonName = "TESTING 1"
	// data.SkeletonType = "BONE"

	// return data, nil
}
