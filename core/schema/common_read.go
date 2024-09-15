package schema

import "github.com/optiop/next_century_sdk/core/schema/enums"

/*
|Property 				|Type 				|
|-----------------------|-------------------|
|utilityTypeId 			|UtilityTypeIds 	|
|uomTypeId 				|UomTypeIds 		|
|multiplier 			|Number 			|
|imr 					|Number 			|
|inputType 				|InputType 			|
|computed? 				|Number 			|
|raw? 					|Number 			|
*/
type MeterRead struct {
	UtilityTypeId enums.UtilityTypeId `json:"utilityTypeId"`
	UomTypeId     enums.UomTypeId     `json:"uomTypeId"`
	Multiplier    float64             `json:"multiplier"`
	Imr           float64             `json:"imr"`
	InputType     enums.InputType     `json:"inputType"`
	Computed      float64             `json:"computed,omitempty"`
	Raw           float64             `json:"raw,omitempty"`
}

/*
|Property 				|Type 				|
|-----------------------|-------------------|
|firstHopId 			|Number 			|
|linkQuality 			|Number 			|
|batteryLevel? 			|Number 			|
|temperature? 			|Number 			|
|tamper? 				|Number 			|
|freeze? 				|Number 			|
|leak? 				|Number 				|
*/
type Device struct {
	FirstHopId   float64 `json:"firstHopId"`
	LinkQuality  float64 `json:"linkQuality"`
	BatteryLevel float64 `json:"batteryLevel,omitempty"`
	Temperature  float64 `json:"temperature,omitempty"`
	Tamper       float64 `json:"tamper,omitempty"`
	Freeze       float64 `json:"freeze,omitempty"`
	Leak         float64 `json:"leak,omitempty"`
}

/*
|Property 				|Type 				|
|-----------------------|-------------------|
|recordedBy 			|String 			|
|imageUrl 				|String 			|
|noReadReason? 			|NoReadReason 		|
|noReadReasonOther? 	|String 			|
*/
type Manual struct {
	RecordedBy        string             `json:"recordedBy"`
	ImageUrl          string             `json:"imageUrl"`
	NoReadReason      enums.NoReadReason `json:"noReadReason,omitempty"`
	NoReadReasonOther string             `json:"noReadReasonOther,omitempty"`
}

/*
|Property 				|Type 				|
|-----------------------|-------------------|
|uploadedBy 			|String 			|
|fileUrl? 				|String 			|
|lineNumber? 			|Number 			|
|noReadReason? 			|NoReadReason 		|
*/
type Import struct {
	UploadedBy   string             `json:"uploadedBy"`
	FileUrl      string             `json:"fileUrl,omitempty"`
	LineNumber   float64            `json:"lineNumber,omitempty"`
	NoReadReason enums.NoReadReason `json:"noReadReason,omitempty"`
}

/*
|Property 				|Type 				|
|-----------------------|-------------------|
|type 					|CommonReadType 	|
|timestamp 				|Date 				|
|deviceId 				|Number 			|
|deviceTypeId 			|DeviceTypeIds 		|
|propertyId 			|String 			|
|unitId 				|String 			|
|meterRead? 			|MeterRead 			|
|device? 				|Device 			|
|manual? 				|Manual 			|
|import? 				|Import 			|
|coords? 				|[Number, Number] 	|
|customUnitId? 			|String 			|
|customPropertyId? 		|String 			|
*/
type CommonRead struct {
	Type             enums.CommonReadType `json:"type"`
	Timestamp        string               `json:"timestamp"`
	DeviceId         float64              `json:"deviceId"`
	DeviceTypeId     enums.DeviceTypeId   `json:"deviceTypeId"`
	PropertyId       string               `json:"propertyId"`
	UnitId           string               `json:"unitId"`
	MeterRead        MeterRead            `json:"meterRead,omitempty"`
	Device           Device               `json:"device,omitempty"`
	Manual           Manual               `json:"manual,omitempty"`
	Import           Import               `json:"import,omitempty"`
	Coords           []float64            `json:"coords,omitempty"`
	CustomUnitId     string               `json:"customUnitId,omitempty"`
	CustomPropertyId string               `json:"customPropertyId,omitempty"`
}
