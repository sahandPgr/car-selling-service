package helper

type ResultCode int

const (
	Success             ResultCode = 200
	Created             ResultCode = 201
	BadRequest          ResultCode = 400
	InternalServerError ResultCode = 500
	Conflict            ResultCode = 409
	NotFound            ResultCode = 404
	TooManyRequests     ResultCode = 429
)
