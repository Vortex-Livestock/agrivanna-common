package common

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var (
	ErrEmptyID       = errors.New("ID cannot be empty")
	ErrMissingFields = errors.New("missing required fields")
)

func MapServiceErrorToStatus(err error, action string) error {
	switch {
	case errors.Is(err, ErrEmptyID):
		return status.Errorf(codes.InvalidArgument, "failed to %s. %v", action, err)
	case errors.Is(err, ErrMissingFields):
		return status.Errorf(codes.InvalidArgument, "failed to %s. %v", action, err)
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return status.Errorf(codes.AlreadyExists, "failed to %s. duplicated key", action)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return status.Errorf(codes.NotFound, "failed to %s. not found", action)
	default:
		return status.Errorf(codes.Internal, "failed to %s. %v", action, err)
	}
}

func MapGrpcErrorToHTTP(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}
	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError, "unknown error. " + err.Error()
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest, st.Message()
	case codes.NotFound:
		return http.StatusNotFound, st.Message()
	case codes.AlreadyExists:
		return http.StatusConflict, st.Message()
	case codes.Internal:
		return http.StatusInternalServerError, st.Message()
	default:
		return http.StatusInternalServerError, st.Message()
	}
}
