package grpc_errors

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

var (
	ErrNotFound         = errors.New("Not found")
	ErrNoCtxMetaData    = errors.New("No ctx metadata")
	ErrInvalidSessionId = errors.New("Invalid session id")
	ErrEmailExists      = errors.New("Email already exists")
)

// ParseGRPCErrStatusCode parses error and returns grpc status code
func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return codes.NotFound
	case errors.Is(err, redis.Nil):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, ErrEmailExists):
		return codes.AlreadyExists
	case errors.Is(err, ErrNoCtxMetaData):
		return codes.Unauthenticated
	case errors.Is(err, ErrInvalidSessionId):
		return codes.PermissionDenied
	case strings.Contains(err.Error(), "Validate"):
		return codes.InvalidArgument
	case strings.Contains(err.Error(), "redis"):
		return codes.NotFound
	}
	return codes.Internal
}
