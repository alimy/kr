package xerror

import "errors"

var (
	ErrProviderNotSupported = errors.New("provider not supported")
	ErrNotReady             = errors.New("not ready")
	ErrNilConfigure         = errors.New("nil configure is not allowed")
)
