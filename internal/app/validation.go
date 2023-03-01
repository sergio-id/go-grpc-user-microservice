package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/types"
)

func newValidate() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("gender", func(fl validator.FieldLevel) bool {
		for _, item := range []types.GenderType{types.Male, types.Female, types.Unknown} {
			if fl.Field().String() == string(item) {
				return true
			}
		}
		return false
	})
	_ = validate.RegisterValidation("status", func(fl validator.FieldLevel) bool {
		for _, item := range []types.StatusType{types.Active, types.Blocked, types.Deleted, types.Pending} {
			if fl.Field().String() == string(item) {
				return true
			}
		}
		return false
	})
	return validate
}
