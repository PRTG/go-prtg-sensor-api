// Copyright Paessler AG.
// All Rights Reserved

package prtg

// PRTG Sensor string constants
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
	// Required options
	ShowChart int `json:"ShowChart"`
	ShowTable int `json:"ShowTable"`
	// Additional options
	Float           int    `json:"Float,omitempty"`
	Unit            string `json:"Unit,omitempty"`
	CustomUnit      string `json:"CustomUnit,omitempty"`
	LimitMinError   int    `json:"LimitMinError,omitempty"`
	LimitMaxError   int    `json:"LimitMaxError,omitempty"`
	LimitMinWarning int    `json:"limitMinWarning,omitempty"`
	LimitMaxWarning int    `json:"LimitMaxWarning,omitempty"`
	LimitErrorMsg   string `json:"LimitErrorMsg,omitempty"`
	LimitWarningMsg string `json:"LimitWarningMsg,omitempty"`
	LimitMode       int    `json:"LimitMode,omitempty"`
	ValueLookup     string `json:"ValueLookup,omitempty"`
	NotifyChanged   bool   `json:"NotifyChanged,omitempty"`
	Warning         int    `json:"Warning,omitempty"`
}

// AddChannel adds a new channel to a SensorResponse
func (s *SensorResponse) AddChannel(channel SensorChannel) {
	s.PRTG.Result = append(s.PRTG.Result, channel)
}
