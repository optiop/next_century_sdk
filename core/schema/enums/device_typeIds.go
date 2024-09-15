//go:generate go run github.com/dmarkham/enumer -type=DeviceTypeId  -json -transform=kebab -trimprefix=DeviceTypeId -output=device_type_id__string.go
package enums

type DeviceTypeId int

/*
|Value 				|Description 								|
|-------------------|-------------------|
|1 					|TRANSCEIVER								|
|2 					|REPEATER									|
|4 					|TRANSMITTER_3RD_PARTY_FA					|
|5 					|TRANSMITTER_3RD_PARTY_ECHO					|
|6 					|REPEATER_3RD_PARTY_FA						|
|7 					|REPEATER_3RD_PARTY_ECHO					|
|12 				|REMOTE_READER								|
|13 				|REMOTE_READER_TRANSCEIVER					|
|14 				|GATEWAY_201								|
|15 				|GATEWAY_301								|
|17 				|TRANSMITTER_3RD_PARTY_UNKNOWN				|
|18 				|REPEATER_3RD_PARTY_UNKNOWN					|
|19 				|RECEIVER_3RD_PARTY_UNKNOWN					|
|20 				|MANUAL_READER								|
|23 				|TR4										|
|24 				|TR4_X										|
|25 				|TR4_I										|
|26 				|NR4										|
|31 				|RR4_TR										|
|32 				|RE4										|
|34 				|GW4										|
|35 				|GW4_LITE									|
|36 				|NM4_I										|
|37 				|TR4_E										|
|41 				|LS4										|
|99 				|UNKNOWN_3RD_PARTY							|
*/

const (
	DeviceTypeIdTransceiver                DeviceTypeId = 1
	DeviceTypeIdRepeater                   DeviceTypeId = 2
	DeviceTypeIdTransmitter3rdPartyFa      DeviceTypeId = 4
	DeviceTypeIdTransmitter3rdPartyEcho    DeviceTypeId = 5
	DeviceTypeIdRepeater3rdPartyFa         DeviceTypeId = 6
	DeviceTypeIdRepeater3rdPartyEcho       DeviceTypeId = 7
	DeviceTypeIdRemoteReader               DeviceTypeId = 12
	DeviceTypeIdRemoteReaderTransceiver    DeviceTypeId = 13
	DeviceTypeIdGateway201                 DeviceTypeId = 14
	DeviceTypeIdGateway301                 DeviceTypeId = 15
	DeviceTypeIdTransmitter3rdPartyUnknown DeviceTypeId = 17
	DeviceTypeIdRepeater3rdPartyUnknown    DeviceTypeId = 18
	DeviceTypeIdReceiver3rdPartyUnknown    DeviceTypeId = 19
	DeviceTypeIdManualReader               DeviceTypeId = 20
	DeviceTypeIdTr4                        DeviceTypeId = 23
	DeviceTypeIdTr4X                       DeviceTypeId = 24
	DeviceTypeIdTr4I                       DeviceTypeId = 25
	DeviceTypeIdNr4                        DeviceTypeId = 26
	DeviceTypeIdRr4Tr                      DeviceTypeId = 31
	DeviceTypeIdRe4                        DeviceTypeId = 32
	DeviceTypeIdGw4                        DeviceTypeId = 34
	DeviceTypeIdGw4Lite                    DeviceTypeId = 35
	DeviceTypeIdNm4I                       DeviceTypeId = 36
	DeviceTypeIdTr4E                       DeviceTypeId = 37
	DeviceTypeIdLs4                        DeviceTypeId = 41
	DeviceTypeIdUnknown3rdParty            DeviceTypeId = 99
)
