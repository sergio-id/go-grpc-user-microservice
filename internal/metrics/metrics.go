package metrics

import (
	"fmt"
	"github.com/sergio-id/go-grpc-user-microservice/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics contains metrics for the microservice.
type Metrics struct {
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter

	GrpcCreateUserRequests  prometheus.Counter
	GrpcUpdateUserRequests  prometheus.Counter
	GrpcDeleteUserRequests  prometheus.Counter
	GrpcGetByIDUserRequests prometheus.Counter
}

// NewMetrics creates a new instance of Metrics.
func NewMetrics(cfg config.Config) *Metrics {
	return &Metrics{
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_grpc_requests_total", cfg.ServiceName),
			Help: "The total number of error grpc requests",
		}),

		GrpcCreateUserRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number create user grpc requests",
		}),
		GrpcUpdateUserRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_update_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number update user grpc requests",
		}),
		GrpcDeleteUserRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_delete_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number delete user grpc requests",
		}),
		GrpcGetByIDUserRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_by_id_user_grpc_requests_total", cfg.ServiceName),
			Help: "The total number get by id user grpc requests",
		}),
	}
}
