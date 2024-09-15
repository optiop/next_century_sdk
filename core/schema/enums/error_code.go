//go:generate go run github.com/dmarkham/enumer -type=ErrorCode  -json -transform=kebab -trimprefix=ErrorCode -output=error_code__string.go
package enums

type ErrorCode int

/*
|Value 				|Description 							|
|-------------------|-------------------|
|1 					|COULD_NOT_READ_METER					|
*/

const (
	ErrorCodeCouldNotReadMeter ErrorCode = 1
)
