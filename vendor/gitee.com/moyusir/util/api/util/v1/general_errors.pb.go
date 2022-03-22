// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsRegisterFormatNonstandard(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_REGISTER_FORMAT_NONSTANDARD.String() && e.Code == 400
}

func ErrorRegisterFormatNonstandard(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_REGISTER_FORMAT_NONSTANDARD.String(), fmt.Sprintf(format, args...))
}

func IsApiGatewayConnectFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_API_GATEWAY_CONNECT_FAIL.String() && e.Code == 500
}

func ErrorApiGatewayConnectFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_API_GATEWAY_CONNECT_FAIL.String(), fmt.Sprintf(format, args...))
}

func IsRequestSendFail(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_REQUEST_SEND_FAIL.String() && e.Code == 500
}

func ErrorRequestSendFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_REQUEST_SEND_FAIL.String(), fmt.Sprintf(format, args...))
}