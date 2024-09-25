package schema

import "time"

type MeterData struct {
	ID            string            `json:"_id"`
	Date          string            `json:"date"`
	PropertyID    string            `json:"propertyId"`
	DeviceID      int64             `json:"deviceId"`
	UnitID        string            `json:"unitId"`
	UtilityTypeID int               `json:"utilityTypeId"`
	Data          []*MeterDataEntry `json:"data"`
	LatestRead    MeterLatestRead   `json:"latestRead"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

type MeterDataEntry struct {
	Type                                int       `json:"type"`
	DeviceID                            int64     `json:"deviceId"`
	LastCheckIn                         time.Time `json:"lastCheckIn"`
	UtilityTypeID                       int       `json:"utilityTypeId"`
	UomTypeID                           int       `json:"uomTypeId"`
	PulseCount                          int       `json:"pulseCount"`
	Multiplier                          int       `json:"multiplier"`
	ResetCount                          int       `json:"resetCount"`
	FirstHopID                          int64     `json:"firstHopId"`
	LinkQuality                         int       `json:"linkQuality"`
	Temperature                         int       `json:"temperature"`
	BatteryLevel                        int       `json:"batteryLevel"`
	Hops                                int       `json:"hops"`
	CheckInInterval                     int       `json:"checkInInterval"`
	FirmwareMajor                       int       `json:"firmwareMajor"`
	FirmwareMinor                       int       `json:"firmwareMinor"`
	BatteryCapacity                     int       `json:"batteryCapacity"`
	BatteryVoltage                      int       `json:"batteryVoltage"`
	BLEInfoOrFirmwareBufferedFlags      int64     `json:"bleInfoOrFirmwareBufferedFlags"`
	BLEReservedOrFirmwareBufferedFlags2 int       `json:"bleReservedOrFirmwareBufferedFlags2"`
	Tr401ConfigFlags                    int       `json:"tr401ConfigFlags"`
	Tr401MeterFlags                     int       `json:"tr401MeterFlags"`
	Tr401Flags                          int       `json:"tr401Flags"`
	Tr401BonusFlags                     int       `json:"tr401BonusFlags"`
	ReadFlags                           int       `json:"readFlags"`
	RawMeterSerial                      int       `json:"rawMeterSerial"`
	Usage                               int       `json:"usage"`
}

type MeterLatestRead struct {
	Type                                int       `json:"type"`
	DeviceID                            int64     `json:"deviceId"`
	LastCheckIn                         time.Time `json:"lastCheckIn"`
	UtilityTypeID                       int       `json:"utilityTypeId"`
	UomTypeID                           int       `json:"uomTypeId"`
	PulseCount                          int       `json:"pulseCount"`
	Multiplier                          int       `json:"multiplier"`
	ResetCount                          int       `json:"resetCount"`
	FirstHopID                          int64     `json:"firstHopId"`
	LinkQuality                         int       `json:"linkQuality"`
	Temperature                         int       `json:"temperature"`
	BatteryLevel                        int       `json:"batteryLevel"`
	Hops                                int       `json:"hops"`
	CheckInInterval                     int       `json:"checkInInterval"`
	FirmwareMajor                       int       `json:"firmwareMajor"`
	FirmwareMinor                       int       `json:"firmwareMinor"`
	BatteryCapacity                     int       `json:"batteryCapacity"`
	BatteryVoltage                      int       `json:"batteryVoltage"`
	BLEInfoOrFirmwareBufferedFlags      int64     `json:"bleInfoOrFirmwareBufferedFlags"`
	BLEReservedOrFirmwareBufferedFlags2 int       `json:"bleReservedOrFirmwareBufferedFlags2"`
	Tr401ConfigFlags                    int       `json:"tr401ConfigFlags"`
	Tr401MeterFlags                     int       `json:"tr401MeterFlags"`
	Tr401Flags                          int       `json:"tr401Flags"`
	Tr401BonusFlags                     int       `json:"tr401BonusFlags"`
	ReadFlags                           int       `json:"readFlags"`
	RawMeterSerial                      int       `json:"rawMeterSerial"`
	Usage                               int       `json:"usage"`
}
