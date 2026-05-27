package errors

import (
	"errors"

	"github.com/alibabacloud-go/tea/tea"
	apiErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
)

const (
	ErrInternalError = "InternalError"

	ErrForbidden = "Forbidden.RAM"

	// InvalidVSwitchIDIPNotEnough AssignPrivateIpAddresses const error message
	// Reference: https://help.aliyun.com/document_detail/85917.html
	InvalidVSwitchIDIPNotEnough = "InvalidVSwitchId.IpNotEnough"

	QuotaExceededPrivateIPAddress = "QuotaExceeded.PrivateIpAddress"

	// ErrEniPerInstanceLimitExceeded CreateNetworkInterfaces const error message
	// Reference: https://help.aliyun.com/document_detail/85917.html
	ErrEniPerInstanceLimitExceeded = "EniPerInstanceLimitExceeded"

	// ErrSecurityGroupInstanceLimitExceed CreateNetworkInterfaces const error message
	// Reference: https://help.aliyun.com/document_detail/85917.html
	ErrSecurityGroupInstanceLimitExceed = "SecurityGroupInstanceLimitExceed"

	// ErrInvalidIPIPUnassigned see https://help.aliyun.com/document_detail/85919.html
	// for API UnassignPrivateIpAddresses
	ErrInvalidIPIPUnassigned = "InvalidIp.IpUnassigned"

	// ErrInvalidENINotFound ..
	// for API UnassignPrivateIpAddresses DetachNetworkInterface
	ErrInvalidENINotFound = "InvalidEniId.NotFound"

	ErrInvalidEcsIDNotFound = "InvalidEcsId.NotFound"

	ErrIPv4CountExceeded = "InvalidOperation.Ipv4CountExceeded"
	ErrIPv6CountExceeded = "InvalidOperation.Ipv6CountExceeded"

	// ErrInvalidENIState ..
	// for API DeleteNetworkInterface
	ErrInvalidENIState = "InvalidOperation.InvalidEniState"

	// ErrInvalidAllocationIDNotFound InvalidAllocationId.NotFound EIP not found
	ErrInvalidAllocationIDNotFound = "InvalidAllocationId.NotFound"

	// ErrThrottling .
	ErrThrottling = "Throttling"

	ErrOperationConflict = "Operation.Conflict"

	ErrIdempotentFailed = "IdempotentFailed"

	ErrEfloPrivateIPQuotaExecuted = 1013
	ErrEfloResourceNotFound       = 1011
)

// define well known err
var (
	ErrNotFound = errors.New("not found")
)

// ErrAssert check err is match errCode
// DEPRECATED
func ErrAssert(errCode string, err error) bool { _ = "STUB: not implemented"; return false }

func ErrorCodeIs(err error, codes ...string) bool { _ = "STUB: not implemented"; return false }

// ErrorCodeIsAny checks if error matches any of the provided error codes.
// It supports both apiErr.Error and tea.SDKError types.
func ErrorCodeIsAny(err error, codes ...string) bool {
	_ = "STUB: not implemented"
	// Try apiErr.Error first
	return false
}

// Try tea.SDKError

// ErrRequestID try to get requestID
func ErrRequestID(err error) string { _ = "STUB: not implemented"; return "" }

type E struct {
	e apiErr.Error
}

func (e *E) Error() string { _ = "STUB: not implemented"; return "" }

func (e *E) Unwrap() error { _ = "STUB: not implemented"; return nil }

func WarpError(err error) error { _ = "STUB: not implemented"; return nil }

// IsURLError if there is conn problem
func IsURLError(err error) bool { _ = "STUB: not implemented"; return false }

func WarpFn(codes ...string) CheckErr { _ = "STUB: not implemented"; return *new(CheckErr) }

type CheckErr = func(err error) bool

func ErrorIs(err error, fns ...CheckErr) bool { _ = "STUB: not implemented"; return false }

type EFLOCode struct {
	Code      int
	Message   string
	RequestID string
	Content   any
}

func (e *EFLOCode) Error() string { _ = "STUB: not implemented"; return "" }

func IsEfloCode(err error, code int) bool { _ = "STUB: not implemented"; return false }

type E2 struct {
	e     *tea.SDKError
	extra []string
}

func (e *E2) Error() string { _ = "STUB: not implemented"; return "" }

func (e *E2) Unwrap() error { _ = "STUB: not implemented"; return nil }

func WarpError2(err error, extra ...string) error { _ = "STUB: not implemented"; return nil }
