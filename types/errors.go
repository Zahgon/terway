package types

const (
	ErrInternalError      ErrCode = "InternalError"
	ErrInvalidArgsErrCode ErrCode = "InvalidArgs"
	ErrInvalidDataType    ErrCode = "InvalidDataType"

	ErrPodIsProcessing ErrCode = "PodIsProcessing"

	ErrResourceInvalid ErrCode = "ResourceInvalid"

	ErrOpenAPIErr     ErrCode = "OpenAPIErr"
	ErrPodENINotReady ErrCode = "PodENINotReady"
	ErrIPNotAllocated ErrCode = "IPNotAllocated"

	ErrIPOutOfSyncErr ErrCode = "OutIPOfSync"
)

type ErrCode string

type Error struct {
	Code ErrCode
	Msg  string

	R error
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }
