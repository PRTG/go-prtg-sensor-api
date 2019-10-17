// Copyright (c) 2019, Paessler AG.
// All Rights Reserved
package prtg

import (
	"reflect"
	"testing"
)

func TestSensorChannel_SetErrLimitMsg(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetErrLimitMsg("TestMsg")
	if channel.LimitErrorMsg != "TestMsg" {
		t.Error("Message was not set")
	}
	channel = &SensorChannel{}
	if channel.LimitErrorMsg != "" {
		t.Error("Message should be empty if not set")
	}
}
func TestSensorChannel_SetWarnLimitMsg(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetWarnLimitMsg("TestMsg")
	if channel.LimitWarningMsg != "TestMsg" {
		t.Error("Message was not set")
	}
	channel = &SensorChannel{}
	if channel.LimitWarningMsg != "" {
		t.Error("Message should be empty if not set")
	}
}

func TestSensorChannel_SetToWarning(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetToWarning()
	if channel.Warning != "1" {
		t.Error("Warning was not set to '1'")
	}
	channel = &SensorChannel{}
	if channel.Warning != "" {
		t.Error("Warning should be empty if not set")
	}
}

func TestSensorChannel_SetShowChart(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetShowChart(true)
	if channel.ShowChart != "1" {
		t.Error("Warning was not set to '1'")
	}
	channel = &SensorChannel{}
	channel = channel.SetShowChart(false)
	if channel.ShowChart != "0" {
		t.Error("Warning was not set to '0'")
	}
	channel = &SensorChannel{}
	if channel.ShowChart != "" {
		t.Error("Warning should be empty if not set")
	}
}

func TestSensorChannel_SetShowTable(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetShowTable(true)
	if channel.ShowTable != "1" {
		t.Error("ShowTable was not set to '1'")
	}
	channel = &SensorChannel{}
	channel = channel.SetShowTable(false)
	if channel.ShowTable != "0" {
		t.Error("ShowTable was not set to '0'")
	}
	channel = &SensorChannel{}
	if channel.ShowTable != "" {
		t.Error("ShowTable should be empty if not set")
	}
}

func TestSensorChannel_ValueIsDifference(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.ValueIsDifference()
	if channel.ValueMode != "Difference" {
		t.Error("ValueMode was not set to 'Difference'")
	}
	channel = &SensorChannel{}
	if channel.ValueMode != "" {
		t.Error("ValueMode should be empty if not set")
	}
}

func TestSensorChannel_SetMaxWarnLimit(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetMaxWarnLimit(100)
	if channel.LimitMode != "1" && channel.LimitMaxWarning != "100" {
		t.Error("Limit not set correctly to 100")
	}
	channel = &SensorChannel{}
	channel = channel.SetMaxWarnLimit(-2)
	if channel.LimitMode != "1" && channel.LimitMaxWarning != "-2" {
		t.Error("Limit not set correctly to -2")
	}
	channel = &SensorChannel{}
	channel = channel.SetMaxWarnLimit(0)
	if channel.LimitMode != "1" && channel.LimitMaxWarning != "0" {
		t.Error("Limit not set correctly to 0")
	}
	channel = New().AddChannel("test")
	if channel.LimitMode != "" && channel.LimitMaxWarning != "" {
		t.Error("Limits should be empty if not set")
	}
}

func TestSensorChannel_SetMinWarnLimit(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetMinWarnLimit(100)
	if channel.LimitMode != "1" && channel.LimitMinWarning != "100" {
		t.Error("Limit not set correctly to 100")
	}
	channel = &SensorChannel{}
	channel = channel.SetMinWarnLimit(-2)
	if channel.LimitMode != "1" && channel.LimitMinWarning != "-2" {
		t.Error("Limit not set correctly to -2")
	}
	channel = &SensorChannel{}
	channel = channel.SetMinWarnLimit(0)
	if channel.LimitMode != "1" && channel.LimitMinWarning != "0" {
		t.Error("Limit not set correctly to 0")
	}
	channel = New().AddChannel("test")
	if channel.LimitMode != "" && channel.LimitMinWarning != "" {
		t.Error("Limits should be empty if not set")
	}
}

func TestSensorChannel_SetMaxErrLimit(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetMaxErrLimit(100)
	if channel.LimitMode != "1" && channel.LimitMaxError != "100" {
		t.Error("Limit not set correctly to 100")
	}
	channel = &SensorChannel{}
	channel = channel.SetMaxErrLimit(-2)
	if channel.LimitMode != "1" && channel.LimitMaxError != "-2" {
		t.Error("Limit not set correctly to -2")
	}
	channel = &SensorChannel{}
	channel = channel.SetMaxErrLimit(0)
	if channel.LimitMode != "1" && channel.LimitMaxError != "0" {
		t.Error("Limit not set correctly to 0")
	}
	channel = New().AddChannel("test")
	if channel.LimitMode != "" && channel.LimitMaxError != "" {
		t.Error("Limits should be empty if not set")
	}
}

func TestSensorChannel_SetMinErrLimit(t *testing.T) {
	channel := &SensorChannel{}
	channel = channel.SetMinErrLimit(100)
	if channel.LimitMode != "1" && channel.LimitMinError != "100" {
		t.Error("Limit not set correctly to 100")
	}
	channel = &SensorChannel{}
	channel = channel.SetMinErrLimit(-2)
	if channel.LimitMode != "1" && channel.LimitMinError != "-2" {
		t.Error("Limit not set correctly to -2")
	}
	channel = &SensorChannel{}
	channel = channel.SetMinErrLimit(0)
	if channel.LimitMode != "1" && channel.LimitMinError != "0" {
		t.Error("Limit not set correctly to 0")
	}
	channel = New().AddChannel("test")
	if channel.LimitMode != "" && channel.LimitMinError != "" {
		t.Error("Limits should be empty if not set")
	}
}

func TestSensorChannel_SetValue(t *testing.T) {
	type args struct {
		val interface{}
	}
	var tests = []struct {
		name    string
		channel SensorChannel
		args    args
		want    *SensorChannel
	}{
		{
			name:    "set float (1 decimal) value to Float value",
			channel: SensorChannel{Value: "empty_test", Float: "0"},
			args:    args{val: 1.1},
			want:    &SensorChannel{Value: "1.1", Float: "1"},
		},
		{
			name:    "set float32 (0 decimal)",
			channel: SensorChannel{},
			args:    args{val: float32(2)},
			want:    &SensorChannel{Value: "2", Float: "1"},
		},
		{
			name:    "set float64 (0 decimal) value to Float value",
			channel: SensorChannel{},
			args:    args{val: float64(2)},
			want:    &SensorChannel{Value: "2", Float: "1"},
		},
		{
			name:    "set float (4 decimal) value to Float value",
			channel: SensorChannel{Value: "empty_test", Float: "0"},
			args:    args{val: 0.1999},
			want:    &SensorChannel{Value: "0.1999", Float: "1"},
		},
		{
			name:    "set uint64 value",
			channel: SensorChannel{},
			args:    args{val: uint64(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set uint32 value",
			channel: SensorChannel{},
			args:    args{val: uint32(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set uint16 value",
			channel: SensorChannel{},
			args:    args{val: uint16(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set uint8 value",
			channel: SensorChannel{},
			args:    args{val: uint8(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set uint value",
			channel: SensorChannel{},
			args:    args{val: uint(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set int64 value",
			channel: SensorChannel{},
			args:    args{val: int64(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set int32 value",
			channel: SensorChannel{},
			args:    args{val: int32(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set int16 value",
			channel: SensorChannel{},
			args:    args{val: int16(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set int8 value",
			channel: SensorChannel{},
			args:    args{val: int8(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set int value",
			channel: SensorChannel{},
			args:    args{val: int(2)},
			want:    &SensorChannel{Value: "2", Float: "0"},
		},
		{
			name:    "set bool value, true",
			channel: SensorChannel{},
			args:    args{val: true},
			want:    &SensorChannel{Value: "1", Float: "0"},
		},
		{
			name:    "set bool value, false",
			channel: SensorChannel{},
			args:    args{val: false},
			want:    &SensorChannel{Value: "0", Float: "0"},
		},
		{
			name:    "set string value",
			channel: SensorChannel{},
			args:    args{val: "testString123"},
			want:    &SensorChannel{Value: "testString123", Float: "0"},
		},
		{
			name:    "set complex value, shall fail",
			channel: SensorChannel{},
			args:    args{val: 6 + 7i},
			want:    &SensorChannel{},
		},
		{
			name:    "set uintptr value, shall fail",
			channel: SensorChannel{},
			args:    args{val: new(uintptr)},
			want:    &SensorChannel{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetValue(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetUnit(t *testing.T) {
	tests := []struct {
		name     string
		channel  SensorChannel
		unitType UnitType
		want     *SensorChannel
	}{
		{
			name:     "returns custom string in Unit field",
			channel:  SensorChannel{},
			unitType: "freeText",
			want:     &SensorChannel{Unit: "freeText"},
		},
		{
			name:     "returns exact constant string as called with, BytesBandwidth",
			channel:  SensorChannel{},
			unitType: BytesBandwidth,
			want:     &SensorChannel{Unit: string(BytesBandwidth)},
		},
		{
			name:     "returns exact constant string as called with, BytesDisk",
			channel:  SensorChannel{},
			unitType: BytesDisk,
			want:     &SensorChannel{Unit: string(BytesDisk)},
		},
		{
			name:     "returns exact constant string as called with, Temperature",
			channel:  SensorChannel{},
			unitType: Temperature,
			want:     &SensorChannel{Unit: string(Temperature)},
		},
		{
			name:     "returns exact constant string as called with, Percent",
			channel:  SensorChannel{},
			unitType: Percent,
			want:     &SensorChannel{Unit: string(Percent)},
		},
		{
			name:     "returns exact constant string as called with, TimeResponse",
			channel:  SensorChannel{},
			unitType: TimeResponse,
			want:     &SensorChannel{Unit: string(TimeResponse)},
		},
		{
			name:     "returns exact constant string as called with, TimeSeconds",
			channel:  SensorChannel{},
			unitType: TimeSeconds,
			want:     &SensorChannel{Unit: string(TimeSeconds)},
		},
		{
			name:     "returns exact constant string as called with, Custom",
			channel:  SensorChannel{},
			unitType: Custom,
			want:     &SensorChannel{Unit: string(Custom)},
		},
		{
			name:     "returns exact constant string as called with, Count",
			channel:  SensorChannel{},
			unitType: Count,
			want:     &SensorChannel{Unit: string(Count)},
		},
		{
			name:     "returns exact constant string as called with, CPU",
			channel:  SensorChannel{},
			unitType: CPU,
			want:     &SensorChannel{Unit: string(CPU)},
		},
		{
			name:     "returns exact constant string as called with, BytesFile",
			channel:  SensorChannel{},
			unitType: BytesFile,
			want:     &SensorChannel{Unit: string(BytesFile)},
		},
		{
			name:     "returns exact constant string as called with, SpeedDisk",
			channel:  SensorChannel{},
			unitType: SpeedDisk,
			want:     &SensorChannel{Unit: string(SpeedDisk)},
		},
		{
			name:     "returns exact constant string as called with, SpeedNet",
			channel:  SensorChannel{},
			unitType: SpeedNet,
			want:     &SensorChannel{Unit: string(SpeedNet)},
		},
		{
			name:     "returns exact constant string as called with, TimeHours",
			channel:  SensorChannel{},
			unitType: TimeHours,
			want:     &SensorChannel{Unit: string(TimeHours)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetUnit(tt.unitType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetCustomUnit(t *testing.T) {
	var tests = []struct {
		name        string
		channel     SensorChannel
		unitTypeStr string
		want        *SensorChannel
	}{
		{
			name:        "returns custom string to CustomUnit and 'Custom' to Unit",
			channel:     SensorChannel{},
			unitTypeStr: "mySpecialUnitType asdf 123 🍌",
			want:        &SensorChannel{Unit: "Custom", CustomUnit: "mySpecialUnitType asdf 123 🍌"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetCustomUnit(tt.unitTypeStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCustomUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetValueLookup(t *testing.T) {
	var tests = []struct {
		name    string
		channel SensorChannel
		lookup  string
		want    *SensorChannel
	}{
		{
			name:    "returns ValueLookup as given",
			channel: SensorChannel{},
			lookup:  "mySpecialValueLookup 123 🍌",
			want:    &SensorChannel{ValueLookup: "mySpecialValueLookup 123 🍌"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetValueLookup(tt.lookup); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetValueLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetVolumeSize(t *testing.T) {
	var tests = []struct {
		name       string
		channel    SensorChannel
		volumeSize SizeType
		want       *SensorChannel
	}{
		{
			name:       "returns VolumeSize as string from given SizeType, One",
			channel:    SensorChannel{},
			volumeSize: One,
			want:       &SensorChannel{VolumeSize: string(One)},
		},
		{
			name:       "returns VolumeSize as string from given SizeType, Kilo",
			channel:    SensorChannel{},
			volumeSize: Kilo,
			want:       &SensorChannel{VolumeSize: string(Kilo)},
		},
		{
			name:       "returns VolumeSize as string from given SizeType, KiloByte",
			channel:    SensorChannel{},
			volumeSize: KiloByte,
			want:       &SensorChannel{VolumeSize: string(KiloByte)},
		},
		{
			name:       "returns VolumeSize as string from given SizeType, Tera",
			channel:    SensorChannel{},
			volumeSize: Tera,
			want:       &SensorChannel{VolumeSize: string(Tera)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetVolumeSize(tt.volumeSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVolumeSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetSpeedSize(t *testing.T) {
	var tests = []struct {
		name      string
		channel   SensorChannel
		speedType SpeedType
		want      *SensorChannel
	}{
		{
			name:      "returns given speed as string, Bit",
			channel:   SensorChannel{},
			speedType: Bit,
			want:      &SensorChannel{SpeedSize: string(Bit)},
		},
		{
			name:      "returns given speed as string, KiloBit",
			channel:   SensorChannel{},
			speedType: KiloBit,
			want:      &SensorChannel{SpeedSize: string(KiloBit)},
		},
		{
			name:      "returns given speed as string, TeraBit",
			channel:   SensorChannel{},
			speedType: TeraBit,
			want:      &SensorChannel{SpeedSize: string(TeraBit)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetSpeedSize(tt.speedType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSpeedSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetSpeedTime(t *testing.T) {
	var tests = []struct {
		name      string
		channel   SensorChannel
		speedTime TimeType
		want      *SensorChannel
	}{
		{
			name:      "returns speedTime as string, Second",
			channel:   SensorChannel{},
			speedTime: Second,
			want:      &SensorChannel{SpeedTime: string(Second)},
		},
		{
			name:      "returns speedTime as string, Minute",
			channel:   SensorChannel{},
			speedTime: Minute,
			want:      &SensorChannel{SpeedTime: string(Minute)},
		},
		{
			name:      "returns speedTime as string, Hour",
			channel:   SensorChannel{},
			speedTime: Hour,
			want:      &SensorChannel{SpeedTime: string(Hour)},
		},
		{
			name:      "returns speedTime as string, Day",
			channel:   SensorChannel{},
			speedTime: Day,
			want:      &SensorChannel{SpeedTime: string(Day)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetSpeedTime(tt.speedTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSpeedTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSensorChannel_SetDecimalMode(t *testing.T) {
	var tests = []struct {
		name     string
		channel  SensorChannel
		modeType DecimalModeType
		want     *SensorChannel
	}{
		{
			name:     "returns modeType as string, OptionAuto",
			channel:  SensorChannel{},
			modeType: OptionAuto,
			want:     &SensorChannel{DecimalMode: string(OptionAuto)},
		},
		{
			name:     "returns modeType as string, OptionAll",
			channel:  SensorChannel{},
			modeType: OptionAll,
			want:     &SensorChannel{DecimalMode: string(OptionAll)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorChannel{
				Channel:         tt.channel.Channel,
				Value:           tt.channel.Value,
				ValueMode:       tt.channel.ValueMode,
				Unit:            tt.channel.Unit,
				CustomUnit:      tt.channel.CustomUnit,
				ValueLookup:     tt.channel.ValueLookup,
				VolumeSize:      tt.channel.VolumeSize,
				SpeedSize:       tt.channel.SpeedSize,
				SpeedTime:       tt.channel.SpeedTime,
				Float:           tt.channel.Float,
				DecimalMode:     tt.channel.DecimalMode,
				ShowChart:       tt.channel.ShowChart,
				ShowTable:       tt.channel.ShowTable,
				LimitMinWarning: tt.channel.LimitMinWarning,
				LimitMaxWarning: tt.channel.LimitMaxWarning,
				LimitWarningMsg: tt.channel.LimitWarningMsg,
				LimitMinError:   tt.channel.LimitMinError,
				LimitMaxError:   tt.channel.LimitMaxError,
				LimitErrorMsg:   tt.channel.LimitErrorMsg,
				LimitMode:       tt.channel.LimitMode,
				Warning:         tt.channel.Warning,
			}
			if got := r.SetDecimalMode(tt.modeType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDecimalMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
