// Copyright (c) 2019, Paessler AG.
// All Rights Reserved
package prtg

import (
	"fmt"
	"strconv"
)

// SensorChannel has the channel name and value and options
type SensorChannel struct {
	Channel string `json:"channel"`
	Value   string `json:"value"`

	// Options
	ValueMode       string `json:"mode,omitempty"`
	Unit            string `json:"unit,omitempty"`
	CustomUnit      string `json:"customunit,omitempty"`
	ValueLookup     string `json:"valuelookup,omitempty"`
	VolumeSize      string `json:"volumesize,omitempty"`
	SpeedSize       string `json:"speedsize,omitempty"`
	SpeedTime       string `json:"speedtime,omitempty"`
	Float           string `json:"float,omitempty"`
	DecimalMode     string `json:"decimalmode,omitempty"`
	ShowChart       string `json:"showchart,omitempty"`
	ShowTable       string `json:"showtable,omitempty"`
	LimitMinWarning string `json:"limitminwarning,omitempty"`
	LimitMaxWarning string `json:"limitmaxwarning,omitempty"`
	LimitWarningMsg string `json:"limitwarningmsg,omitempty"`
	LimitMinError   string `json:"limitminerror,omitempty"`
	LimitMaxError   string `json:"limitmaxerror,omitempty"`
	LimitErrorMsg   string `json:"limiterrormsg,omitempty"`
	LimitMode       string `json:"limitmode,omitempty"`
	Warning         string `json:"warning,omitempty"`
}

// The unit of the value. Default is Custom. This is useful for PRTG to be able to convert volumes and times.
func (r *SensorChannel) SetUnit(unit UnitType) *SensorChannel {
	r.Unit = string(unit)
	return r
}

// Sets a custom unit text, that is displayed behind the value.
func (r *SensorChannel) SetCustomUnit(unit string) *SensorChannel {
	r.Unit = "Custom"
	r.CustomUnit = unit

	return r
}

// Define if you want to use a lookup file (for example, to view integer values as status texts).
// Enter the ID of the lookup file you want to use, or omit this element to not use lookups.
func (r *SensorChannel) SetValueLookup(lookup string) *SensorChannel {
	r.ValueLookup = lookup
	return r
}

// Volume size used for the display value.
// For example, if you have a value of 50000 and use Kilo as size, the display is 50 kilo.
// Default is One (value used as returned).
func (r *SensorChannel) SetVolumeSize(size SizeType) *SensorChannel {
	r.VolumeSize = string(size)
	return r
}

// Speed size used for the display value.
// For example, if you have a value of 50000 and use KiloBit as size, the display is 50 kiloBits.
// Default is One (value used as returned).
func (r *SensorChannel) SetSpeedSize(size SpeedType) *SensorChannel {
	r.SpeedSize = string(size)
	return r
}

// Time size used for the display value.
// For example, if you have a value of 780 and use Minute as size, the display is 13 Minutes.
// Default is Second (value used as returned).
func (r *SensorChannel) SetSpeedTime(time TimeType) *SensorChannel {
	r.SpeedTime = string(time)
	return r
}

// Sets a value.
// All types allowed except complex and uintptr.
func (r *SensorChannel) SetValue(val interface{}) *SensorChannel {
	switch t := val.(type) {
	case float32, float64:
		r.Float = "1"
		r.Value = fmt.Sprintf("%v", t)
	case int8, uint8, int16, uint16, int32, uint32, uint64, int64, uint, int, string:
		r.Float = "0"
		r.Value = fmt.Sprintf("%v", t)
	case bool:
		r.Float = "0"
		r.Value = "0"

		if t {
			r.Value = "1"
		}
	default:
		fmt.Printf("type of %v is not supported\n", t)
	}

	return r
}

// Select if the value is an absolute value or counter. Default is "Absolute" set.
func (r *SensorChannel) ValueIsDifference() *SensorChannel {
	r.ValueMode = "Difference"
	return r
}

// If value is a float a custom number of decimal places can be set.
func (r *SensorChannel) SetDecimalMode(mode DecimalModeType) *SensorChannel {
	r.DecimalMode = string(mode)
	return r
}

// Define an upper warning limit for the channel.
// If enabled, the sensor will be set to a Warning status if this value is overrun.
func (r *SensorChannel) SetMaxWarnLimit(limit int) *SensorChannel {
	r.LimitMode = "1"
	r.LimitMaxWarning = strconv.Itoa(limit)

	return r
}

// Define an upper error limit for the channel.
// If enabled, the sensor will be set to a Down status if this value is overrun.
func (r *SensorChannel) SetMinWarnLimit(limit int) *SensorChannel {
	r.LimitMode = "1"
	r.LimitMinWarning = strconv.Itoa(limit)

	return r
}

// Define an additional message.
// It will be added to the sensor's message when entering a Warning status that is triggered by a limit.
func (r *SensorChannel) SetWarnLimitMsg(msg string) *SensorChannel {
	r.LimitWarningMsg = msg
	return r
}

// Define an upper error limit for the channel.
// If enabled, the sensor will be set to a Down status if this value is overrun.
func (r *SensorChannel) SetMaxErrLimit(limit int) *SensorChannel {
	r.LimitMode = "1"
	r.LimitMaxError = strconv.Itoa(limit)

	return r
}

// Define a lower error limit for the channel.
// If enabled, the sensor will be set to a Down status if this value is undercut.
func (r *SensorChannel) SetMinErrLimit(limit int) *SensorChannel {
	r.LimitMode = "1"
	r.LimitMinError = strconv.Itoa(limit)

	return r
}

// Define an additional message.
// It will be added to the sensor's message when entering a Down status that is triggered by a limit.
func (r *SensorChannel) SetErrLimitMsg(msg string) *SensorChannel {
	r.LimitErrorMsg = msg
	return r
}

// If enabled for at least one channel, the entire sensor is set to "Warning" status. Default is "not set".
func (r *SensorChannel) SetToWarning() *SensorChannel {
	r.Warning = "1"
	return r
}

// Init value for the Show in graphs option. Default is yes.
func (r *SensorChannel) SetShowChart(show bool) *SensorChannel {
	r.ShowChart = "0"
	if show {
		r.ShowChart = "1"
	}

	return r
}

// Init value for the Show in tables option. Default is yes.
func (r *SensorChannel) SetShowTable(show bool) *SensorChannel {
	r.ShowTable = "0"
	if show {
		r.ShowTable = "1"
	}

	return r
}
