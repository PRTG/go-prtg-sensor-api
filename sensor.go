// Copyright Paessler AG.
// All Rights Reserved

/*

Package prtg implements the API for PRTG custom sensors. It provides all structs and constants needed to implements your own advanced exe sensor in go.

API reference can be found here: https://prtg.paessler.com/api.htm?username=demo&password=demodemo&tabid=7

*/
package prtg

import "encoding/json"

// Units of channel values
const (
	UnitBytesBandwidth = "BytesBandwidth"
	UnitBytesMemory    = "BytesMemory"
	UnitBytesDisk      = "BytesDisk"
	UnitBytesFile      = "BytesFile"
	UnitTemperature    = "Temperature"
	UnitPercent        = "Percent"
	UnitTimeResponse   = "TimeResponse"
	UnitTimeSeconds    = "TimeSeconds"
	UnitCustom         = "Custom"
	UnitCount          = "Count"
	UnitCPU            = "CPU"
	UnitSpeedDisk      = "SpeedDisk"
	UnitSpeedNet       = "SpeedNet"
	UnitTimeHours      = "TimeHours"
)

// Unit of VolumeSize or SpeedSize
const (
	UnitOne      = "One"
	UnitKilo     = "Kilo"
	UnitMega     = "Mega"
	UnitGiga     = "Giga"
	UnitTera     = "Tera"
	UnitByte     = "Byte"
	UnitKiloByte = "KiloByte"
	UnitMegaByte = "MegaByte"
	UnitGigaByte = "Gigabyte"
	UnitTeraByte = "TeraByte"
	UnitBit      = "Bit"
	UnitKiloBit  = "KiloBit"
	UnitMegaBit  = "MegaBit"
	UnitGigaBit  = "GigaBit"
	UnitTeraBit  = "TeraBit"
)

// Unit of speedTime
const (
	UnitSecond = "Second"
	UnitMinute = "Minute"
	UnitHour   = "Hour"
	UnitDay    = "Day"
)

// Options for DecimalMode and Mode
const (
	OptionAbsolute   = "Absolute"
	OptionDifference = "Difference"
	OptionAuto       = "Auto"
	OptionAll        = "All"
)

// SensorResponse is the struct returned by the sensor
type SensorResponse struct {
	PRTG SensorResult `json:"prtg"`
}

// SensorResult contains the sensors channels and optional error/text
type SensorResult struct {
	Result []SensorChannel `json:"result"`
	Error  string          `json:"error,omitempty"`
	Text   string          `json:"text,omitempty"`
}

// SensorChannel represents a PRTG sensor channel
type SensorChannel struct {
	// Default fields
	Channel string `json:"channel"`
	Value   string `json:"value"`
	// Options
	CustomUnit      string `json:"CustomUnit,omitempty"`
	DecimalMode     string `json:"DecimalMode,omitempty"`
	Float           int    `json:"Float,omitempty"`
	LimitErrorMsg   string `json:"LimitErrorMsg,omitempty"`
	LimitMaxError   int    `json:"LimitMaxError,omitempty"`
	LimitMaxWarning int    `json:"LimitMaxWarning,omitempty"`
	LimitMinError   int    `json:"LimitMinError,omitempty"`
	LimitMinWarning int    `json:"limitMinWarning,omitempty"`
	LimitMode       int    `json:"LimitMode,omitempty"`
	LimitWarningMsg string `json:"LimitWarningMsg,omitempty"`
	Mode            string `json:"Mode,omitempty"`
	NotifyChanged   bool   `json:"NotifyChanged,omitempty"`
	ShowChart       int    `json:"ShowChart"`
	ShowTable       int    `json:"ShowTable"`
	SpeedSize       string `json:"SpeedSize,omitempty"`
	SpeedTime       string `json:"SpeedTime,omitempty"`
	Unit            string `json:"Unit,omitempty"`
	ValueLookup     string `json:"ValueLookup,omitempty"`
	VolumeSize      string `json:"VolumeSize,omitempty"`
	Warning         int    `json:"Warning,omitempty"`
}

// AddChannel adds a new channel to a SensorResponse
func (s *SensorResponse) AddChannel(channel SensorChannel) {
	s.PRTG.Result = append(s.PRTG.Result, channel)
}

// String converts the SensorResponse to a JSON formatted string
func (s SensorResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}
