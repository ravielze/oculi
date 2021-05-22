package stdcode

import (
	std "github.com/ravielze/oculi/standard"
)

const (
	// Common Group

	UNKNOWN std.Code = "unknown"
	OK      std.Code = "std"

	// Authorization Group

	UNAUTHORIZED       std.Code = "unauthorized"
	ROLE_NO_PERMISSION std.Code = "role_unauthorized"

	// Controller Group

	PARAMETER_ERROR std.Code = "parameter_error"

	// Handler Group

	HANDLER_ERROR        std.Code = "handler_error"
	UNPROCESSABLE_ENTITY std.Code = "unprocessable_entity"

	// Service Group

	LOGIC_ERROR std.Code = "logic_error"

	// Repository Group

	RECORD_NOT_FOUND std.Code = "record_not_found"
	SQL_ERROR        std.Code = "sql_error"

	// Others

	TOO_MANY_REQUESTS std.Code = "too_many_requests"
	NOT_FOUND         std.Code = "not_found"
)
