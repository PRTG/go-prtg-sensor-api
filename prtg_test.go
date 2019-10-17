// Copyright (c) 2019, Paessler AG.
// All Rights Reserved
package prtg

import (
	"math/rand"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	os.Exit(m.Run())
}

func TestAddChannel(t *testing.T) {
	// Initiate PRTG instance
	sensor := New()
	// Create channel with random data
	channel1 := sensor.AddChannel("channel1").SetValue(1000.25).SetCustomUnit("someUnit")
	channel2 := sensor.AddChannel("channel2").SetValue(50).SetUnit(Percent)
	// Create empty SensorResults
	sc1 := &SensorChannel{Channel: "channel1", Value: "1000.25", Unit: "Custom", CustomUnit: "someUnit", Float: "1"}
	sc2 := &SensorChannel{Channel: "channel2", Value: "50", Unit: "Percent", Float: "0"}

	if !reflect.DeepEqual(channel1, sc1) {
		t.Errorf("sensor.AddChannel() => failed to add channel")
	}

	if !reflect.DeepEqual(channel2, sc2) {
		t.Errorf("sensor.AddChannel() => failed to add channel")
	}

	// Create output
	json, _ := sensor.MarshalToString()
	wantedJson := "{\"prtg\":{\"result\":[{\"channel\":\"channel1\",\"value\":\"1000.25\",\"unit\":\"Custom\",\"customunit\":\"someUnit\",\"float\":\"1\"},{\"channel\":\"channel2\",\"value\":\"50\",\"unit\":\"Percent\",\"float\":\"0\"}]}}"
	// Check if it safely arrived
	if json != wantedJson {
		t.Error("AddChannel() => failed to add channel")
	}
}

func TestMarshalToString(t *testing.T) {
	// Initiate PRTG instance
	sensor := New()
	// Create dummy sensor
	sensor.AddChannel("").SetValue("")
	// Create output
	_, err := sensor.MarshalToString()
	if err != nil {
		t.Error("MarshalToString() => failed to create json string")
	}
}

func TestSetError(t *testing.T) {
	var tests = []struct {
		name    string
		results SensorResults
		arg     bool
		want    *SensorResults
	}{
		{
			name:    "set error true with not set before",
			results: SensorResults{},
			arg:     true,
			want:    &SensorResults{Error: "1"},
		},
		{
			name:    "set error true with set false before",
			results: SensorResults{Error: "0"},
			arg:     true,
			want:    &SensorResults{Error: "1"},
		},
		{
			name:    "set error false with not set before",
			results: SensorResults{},
			arg:     false,
			want:    &SensorResults{Error: "0"},
		},
		{
			name:    "set error false with set true before",
			results: SensorResults{Error: "1"},
			arg:     false,
			want:    &SensorResults{Error: "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorResults{
				Error: tt.results.Error,
			}
			if got := r.SetError(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetSensorText(t *testing.T) {
	var tests = []struct {
		name    string
		results SensorResults
		arg     string
		want    *SensorResults
	}{
		{
			name:    "set text with no text set before",
			results: SensorResults{},
			arg:     "Some Text!",
			want:    &SensorResults{Text: "Some Text!"},
		},
		{
			name:    "overwrite text",
			results: SensorResults{Text: "previous text"},
			arg:     "Some other Text!",
			want:    &SensorResults{Text: "Some other Text!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SensorResults{
				Error: tt.results.Error,
			}
			if got := r.SetSensorText(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetSensorText() = %v, want %v", got, tt.want)
			}
		})
	}
}
