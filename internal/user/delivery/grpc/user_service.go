package grpc

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go/log"
	"github.com/sergio-id/go-grpc-user-microservice/config"
	"github.com/sergio-id/go-grpc-user-microservice/internal/metrics"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/usecase"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/grpc_errors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/tracing"
	"github.com/sergio-id/go-grpc-user-microservice/proto"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// userService is a struct implementing the gRPC server for user-related operations.
type userService struct {
	logger                               logger.Logger
	cfg                                  config.Config
	userUseCase                          *usecase.UserUseCase
	validate                             *validator.Validate
	metrics                              *metrics.Metrics
	proto.UnimplementedUserServiceServer // This is needed to satisfy the gRPC compiler.
}

// NewUserServerGRPC creates a new instance of the gRPC user service.
func NewUserServerGRPC(
	l logger.Logger,
	cfg config.Config,
	userUseCase *usecase.UserUseCase,
	v *validator.Validate,
	m *metrics.Metrics,
) proto.UserServiceServer {
	return &userService{
		logger:      l,
		cfg:         cfg,
		userUseCase: userUseCase,
		validate:    v,
		metrics:     m,
	}
}

// Create is a gRPC method to handle user creation requests.
func (u *userService) Create(ctx context.Context, r *proto.CreateRequest) (*proto.CreateReply, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "userService.Create")
	defer span.Finish()
	span.LogFields(log.String("req", r.String()))
	u.metrics.GrpcCreateUserRequests.Inc()

	user := &domain.User{
		Email:       r.GetEmail(),
		Password:    r.GetPassword(),
		FirstName:   r.GetFirstName(),
		LastName:    r.GetLastName(),
		About:       r.GetAbout(),
		PhoneNumber: r.GetPhoneNumber(),
		Gender:      r.GetGender(),
		Status:      r.GetStatus(),
		LastIP:      r.GetLastIp(),
		LastDevice:  r.GetLastDevice(),
		AvatarURL:   r.GetAvatarUrl(),
	}

	if err := u.validate.StructCtx(ctx, user); err != nil {
		u.logger.Errorf("ValidateStruct: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ValidateStruct: %v", err)
	}

	createdUser, err := u.userUseCase.Create(ctx, user)
	if err != nil {
		u.logger.Errorf("userUseCase.Create: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUseCase.Create: %v", err)
	}

	return &proto.CreateReply{User: u.prepareProtoUser(createdUser)}, nil
}

// Update is a gRPC method to handle user creation requests.
func (u *userService) Update(ctx context.Context, r *proto.UpdateRequest) (*proto.UpdateReply, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "userService.Update")
	defer span.Finish()
	span.LogFields(log.String("req", r.String()))
	u.metrics.GrpcUpdateUserRequests.Inc()

	user := &domain.User{
		ID:          r.GetId(),
		FirstName:   r.GetFirstName(),
		LastName:    r.GetLastName(),
		About:       r.GetAbout(),
		PhoneNumber: r.GetPhoneNumber(),
		Gender:      r.GetGender(),
		Status:      r.GetStatus(),
		LastIP:      r.GetLastIp(),
		LastDevice:  r.GetLastDevice(),
		AvatarURL:   r.GetAvatarUrl(),
	}

	if err := u.validate.StructCtx(ctx, user); err != nil {
		u.logger.Errorf("ValidateStruct: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ValidateStruct: %v", err)
	}

	updatedUser, err := u.userUseCase.Update(ctx, user)
	if err != nil {
		u.logger.Errorf("userUseCase.Update: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Update: %v", err)
	}

	return &proto.UpdateReply{User: u.prepareProtoUser(updatedUser)}, nil
}

func (u *userService) Delete(ctx context.Context, r *proto.DeleteRequest) (*proto.DeleteReply, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "userService.Delete")
	defer span.Finish()
	span.LogFields(log.String("req", r.String()))
	u.metrics.GrpcDeleteUserRequests.Inc()

	err := u.userUseCase.Delete(ctx, r.GetId())
	if err != nil {
		u.logger.Errorf("userUseCase.Delete: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Delete: %v", err)
	}

	return &proto.DeleteReply{}, nil
}

// GetByID is a gRPC method to handle user get by id requests.
func (u *userService) GetByID(ctx context.Context, r *proto.GetByIDRequest) (*proto.GetByIDReply, error) {
	ctx, span := tracing.StartGrpcServerTracerSpan(ctx, "userService.GetByID")
	defer span.Finish()
	span.LogFields(log.String("req", r.String()))
	u.metrics.GrpcGetByIDUserRequests.Inc()

	user, err := u.userUseCase.GetById(ctx, r.GetId())
	if err != nil {
		u.logger.Errorf("userUseCase.GetById: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "userUseCase.GetById: %v", err)
	}

	return &proto.GetByIDReply{User: u.prepareProtoUser(user)}, nil
}

// prepareProtoUser is a helper method to prepare a gRPC User struct from a domain User struct.
func (u *userService) prepareProtoUser(user *domain.User) *proto.User {
	userProto := &proto.User{
		Id:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		About:       user.About,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		Status:      user.Status,
		LastIp:      user.LastIP,
		LastDevice:  user.LastDevice,
		AvatarUrl:   user.AvatarURL,
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
		CreatedAt:   timestamppb.New(user.CreatedAt),
	}
	return userProto
}
