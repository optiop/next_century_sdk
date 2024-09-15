package schema

import "github.com/optiop/next_century_sdk/core/schema/enums"

/*
|Property 				|Type 				|
|-----------------------|-------------------|
|state 					|ReadDownloadState 	|
|progress 				|Number 			|
|url? 					|String 			|
|expiresAt? 			|Date 				|
|totalProperties? 		|Number 			|
|date? 					|String 			|
|propertyId? 			|String 			|
|totalNumberOfDates? 	|Number 			|
*/

type ReadDownloadStatus struct {
	State              enums.ReadDownloadState `json:"state"`
	Progress           float64                 `json:"progress"`
	Url                string                  `json:"url,omitempty"`
	ExpiresAt          string                  `json:"expiresAt,omitempty"`
	TotalProperties    float64                 `json:"totalProperties,omitempty"`
	Date               string                  `json:"date,omitempty"`
	PropertyId         string                  `json:"propertyId,omitempty"`
	TotalNumberOfDates float64                 `json:"totalNumberOfDates,omitempty"`
}
