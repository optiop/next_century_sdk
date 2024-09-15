//go:generate go run github.com/dmarkham/enumer -type=UtilityTypeId -json -transform=kebab -trimprefix=UtilityTypeId -output=utility_type_id__string.go
package enums

type UtilityTypeId int

/*
|Value 				|Description 							|
|-------------------|-------------------|
|0 					|UNDEFINED								|
|1 					|COLD_WATER								|
|2 					|HOT_WATER								|
|3 					|GAS									|
|4 					|ELECTRIC								|
|5 					|ALL_WATER								|
|6 					|RUN_TIME								|
|7 					|COMMERCIAL_WATER						|
|8 					|THERMAL_USAGE							|
*/
const (
	UtilityTypeIdUndefined       UtilityTypeId = 0
	UtilityTypeIdColdWater       UtilityTypeId = 1
	UtilityTypeIdHotWater        UtilityTypeId = 2
	UtilityTypeIdGas             UtilityTypeId = 3
	UtilityTypeIdElectric        UtilityTypeId = 4
	UtilityTypeIdAllWater        UtilityTypeId = 5
	UtilityTypeIdRunTime         UtilityTypeId = 6
	UtilityTypeIdCommercialWater UtilityTypeId = 7
	UtilityTypeIdThermalUsage    UtilityTypeId = 8
)
