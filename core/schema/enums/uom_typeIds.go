//go:generate go run github.com/dmarkham/enumer -type=UomTypeId -json -transform=kebab -trimprefix=UomTypeId -output=uom_type_id__string.go
package enums

type UomTypeId int

/*
|Value 				|Description 							|
|-------------------|-------------------|
|0 					|UNDEFINED								|
|1 					|GALLONS								|
|2 					|LITERS									|
|3 					|CUBIC_FEET								|
|4 					|CUBIC_YARDS							|
|5 					|KWHS									|
|6 					|WHS									|
|7 					|CUBIC_METERS							|
|8 					|HOURS									|
|9 					|MINUTES								|
|10 				|SECONDS								|
|11 				|BTUS									|
*/
const (
	UomTypeIdUndefined   UomTypeId = 0
	UomTypeIdGallons     UomTypeId = 1
	UomTypeIdLiters      UomTypeId = 2
	UomTypeIdCubicFeet   UomTypeId = 3
	UomTypeIdCubicYards  UomTypeId = 4
	UomTypeIdKwhs        UomTypeId = 5
	UomTypeIdWhs         UomTypeId = 6
	UomTypeIdCubicMeters UomTypeId = 7
	UomTypeIdHours       UomTypeId = 8
	UomTypeIdMinutes     UomTypeId = 9
	UomTypeIdSeconds     UomTypeId = 10
	UomTypeIdBtus        UomTypeId = 11
)
