package main

import (
	"errors"
	"strings"

	v1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
)

func validateCreateTableRequest(req *v1.CreateTableRequest) error {
	if req == nil {
		return errors.New("req is required")
	}
	if strings.TrimSpace(req.GetName()) == "" {
		return errors.New("рname is required")
	}
	if req.GetOwnerUserId() == "" {
		return errors.New("owner user is required")
	}
	if req.GetRequestId() == "" {
		return errors.New("request id is required")
	}

	switch req.GetAccessMode() {
	case v1.AccessMode_ACCESS_MODE_UNSPECIFIED,
		v1.AccessMode_ACCESS_MODE_PUBLIC,
		v1.AccessMode_ACCESS_MODE_INVITATION:
		if req.GetPassword() != "" {
			return errors.New("password is only allowed for password-protected tables")
		}

	case v1.AccessMode_ACCESS_MODE_PASSWORD:
		if req.GetPassword() == "" {
			return errors.New("password is required")
		}

	default:
		return errors.New("invalid access_mode")
	}
	return nil
}
