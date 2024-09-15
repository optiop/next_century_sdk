//go:generate go run github.com/dmarkham/enumer -type=CommonReadType -json -transform=kebab -trimprefix=CommonReadType -output=common_read_type__string.go
package enums

type CommonReadType int

/*
|Value 				|Description 							|
|-------------------|-------------------|
|1 					|METER									|
|2 					|NETWORK									|
|3 					|ACCESSORY									|
*/
const (
	CommonReadTypeMeter     CommonReadType = 1
	CommonReadTypeNetwork   CommonReadType = 2
	CommonReadTypeAccessory CommonReadType = 3
)
