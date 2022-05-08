package reporting

import "errors"

var (
	ErrNotHavePermission   = errors.New("request is not permitted")
	ErrUnknownReportStatus = errors.New("report status is unknown")
)
