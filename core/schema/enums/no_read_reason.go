//go:generate go run github.com/dmarkham/enumer -type=NoReadReason  -json -transform=kebab -trimprefix=NoReadReason -output=no_read_reason__string.go
package enums

type NoReadReason int

/*
|Value 				|Description 							|
|-------------------|-------------------|
|0 					|OTHER									|
|1 					|ACCESS_OBSTRUCTED						|
|2 					|TENANT_REFUSED							|
|3 					|UNABLE_TO_ACCESS						|
|4 					|EQUIPMENT_ISSUE						|
|5 					|CANNOT_FIND_EQUIPMENT					|
|6 					|MISSING_EQUIPMENT						|
|7 					|UNIT_IS_CLOSED							|
|8 					|IMPORT_READ_ERROR						|
*/
const (
	NoReadReasonOther NoReadReason = iota
	NoReadReasonAccessObstructed
	NoReadReasonTenantRefused
	NoReadReasonUnableToAccess
	NoReadReasonEquipmentIssue
	NoReadReasonCannotFindEquipment
	NoReadReasonMissingEquipment
	NoReadReasonUnitIsClosed
	NoReadReasonImportReadError
)
