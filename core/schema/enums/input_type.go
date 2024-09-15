//go:generate go run github.com/dmarkham/enumer -type=InputType  -json -transform=kebab -trimprefix=InputType -output=input_type__string.go
package enums

type InputType int

/*
|Value 				|Description 							|
|-------------------|-------------------|
|1 					|PULSE									|
|2 					|ENCODED								|
|3 					|MANUAL									|
|4 					|IMPORT									|
*/

const (
	InputTypePulse InputType = iota + 1
	InputTypeEncoded
	InputTypeManual
	InputTypeImport
)
