package ncMockSample

const DailyReadsSample = `
[
    {
        "_id": "69324e290123q1234",
        "date": "20240202",
        "propertyId": "x_1234",
        "deviceId": 3137739079,
        "unitId": "u_test",
        "utilityTypeId": 4,
        "data": [
            {
                "type": 7,
                "deviceId": 3137739079,
                "lastCheckIn": "2024-02-02T08:38:08.000Z",
                "utilityTypeId": 4,
                "uomTypeId": 5,
                "pulseCount": 0,
                "multiplier": 1,
                "resetCount": 2,
                "firstHopId": 3658489252,
                "linkQuality": 173,
                "temperature": 8,
                "batteryLevel": 4,
                "hops": 0,
                "checkInInterval": 3600,
                "firmwareMajor": 7,
                "firmwareMinor": 6,
                "batteryCapacity": 211,
                "batteryVoltage": 166,
                "bleInfoOrFirmwareBufferedFlags": 67109820,
                "bleReservedOrFirmwareBufferedFlags2": 0,
                "tr401ConfigFlags": 1,
                "tr401MeterFlags": 0,
                "tr401Flags": 25722930,
                "tr401BonusFlags": 4096,
                "readFlags": 9,
                "rawMeterSerial": 0,
                "usage": 0
            }
        ],
        "latestRead": {
            "type": 7,
            "deviceId": 3137739079,
            "lastCheckIn": "2024-02-03T04:58:19.000Z",
            "utilityTypeId": 4,
            "uomTypeId": 5,
            "pulseCount": 0,
            "multiplier": 1,
            "resetCount": 2,
            "firstHopId": 3658489252,
            "linkQuality": 169,
            "temperature": 11,
            "batteryLevel": 4,
            "hops": 0,
            "checkInInterval": 3600,
            "firmwareMajor": 7,
            "firmwareMinor": 6,
            "batteryCapacity": 211,
            "batteryVoltage": 169,
            "bleInfoOrFirmwareBufferedFlags": 67109820,
            "bleReservedOrFirmwareBufferedFlags2": 0,
            "tr401ConfigFlags": 1,
            "tr401MeterFlags": 0,
            "tr401Flags": 25722930,
            "tr401BonusFlags": 4096,
            "readFlags": 5,
            "rawMeterSerial": 0,
            "usage": 0
        },
        "updatedAt": "2024-02-03T06:54:48.243Z"
    }
]
`

const UnitsSample = `
[
    {
        "id": 1234,
        "name": "Bldg 6",
        "building_id": "Landlord",
        "property_id": 14612,
        "createdAt": "2022-07-01T18:57:41.269Z",
        "updatedAt": "2024-08-22T14:07:36.977Z"
    }
]
`
