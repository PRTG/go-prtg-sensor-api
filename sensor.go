// Copyright Paessler AG.
// All Rights Reserved

// Package prtg implements the API for PRTG custom sensors. It provides all structs and constants needed to implement your own advanced exe sensor in go.
//
// API reference can be found here: https://prtg.paessler.com/api.htm?username=demo&password=demodemo&tabid=7
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
	SensorResult `json:"prtg"`
}

// SensorResult contains the sensors channels and optional error/text
type SensorResult struct {
	Channels []SensorChannel `json:"result"`
	Error    int             `json:"error,omitempty"`
	Text     string          `json:"text,omitempty"`
}

// SensorChannel represents a PRTG sensor channel
type SensorChannel struct {
	Name      string  `json:"channel,omitempty"`
	ChannelID *int    `json:"channelid,omitempty"`
	Value     float64 ``

	// Options
	Unit            string   `json:",omitempty"`
	CustomUnit      string   `json:",omitempty"`
	SpeedSize       string   `json:",omitempty"`
	VolumeSize      string   `json:",omitempty"`
	SpeedTime       string   `json:",omitempty"`
	Mode            string   `json:",omitempty"`
	Float           int      `json:",omitempty"`
	DecimalMode     string   `json:",omitempty"`
	Warning         int      `json:",omitempty"`
	ShowChart       *int     `json:",omitempty"`
	ShowTable       *int     `json:",omitempty"`
	LimitMaxError   *float64 `json:",omitempty"`
	LimitMaxWarning *float64 `json:",omitempty"`
	LimitMinWarning *float64 `json:",omitempty"`
	LimitMinError   *float64 `json:",omitempty"`
	LimitErrorMsg   string   `json:",omitempty"`
	LimitWarningMsg string   `json:",omitempty"`
	LimitMode       int      `json:",omitempty"`
	ValueLookup     string   `json:",omitempty"`
	NotifyChanged   bool     `json:",omitempty"`
}

// AddChannel adds a new channel to a SensorResponse
func (s *SensorResponse) AddChannel(channel SensorChannel) {
	s.Channels = append(s.Channels, channel)
}

// String converts the SensorResponse to a JSON formatted string
func (s SensorResponse) String() string {
	b, _ := json.Marshal(s)
	return string(b)
}
